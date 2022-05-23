package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	kitlogger "github.com/go-kit/log"
	"github.com/go-kit/log/level"

	api "github.com/hambyhacks/CrimsonIMS/api/routes"
	prodsrv "github.com/hambyhacks/CrimsonIMS/internal/service/products"
	_ "github.com/lib/pq"
)

type config struct {
	port   int
	env    string
	proddb struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

func main() {
	// Configuration
	var cfg config
	var prodsvc prodsrv.ProductService

	// General Environment flags
	flag.IntVar(&cfg.port, "port", 9000, "API Server port.")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|staging|prod)")

	// Products database flags
	flag.StringVar(&cfg.proddb.dsn, "prod-db-dsn", os.Getenv("PRODDB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.proddb.maxOpenConns, "prod-db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.proddb.maxIdleConns, "prod-db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.proddb.maxIdleTime, "prod-db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	// Set logger to write msg to stdout.
	logger := log.New(os.Stdout, "User-API ", log.LstdFlags|log.Ldate|log.Ltime)
	klogger := kitlogger.NewLogfmtLogger(os.Stderr)

	// Open DB connections
	proddb, err := openProdDB(cfg)
	if err != nil {
		logger.Println("cannot open prod_svc database")
		os.Exit(1)
	} else {
		logger.Println("prod_svc DB connection pool established")
	}
	defer proddb.Close()

	// Initialize services
	// Products service
	prodsvc = &prodsrv.ProdServ{}
	{
		prodrepo, err := prodsrv.NewProdRepo(proddb, klogger)
		if err != nil {
			level.Error(klogger).Log("exit", err)
			os.Exit(1)
		}
		defer proddb.Close()
		prodsvc = prodsrv.NewProdServ(prodrepo, klogger)
	}

	// Declare http.Server struct
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      api.NewHTTPHandler(prodsvc),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	// Start http.Server
	go func() {
		logger.Printf("Starting %s server at port %s\n", cfg.env, s.Addr)
		err := s.ListenAndServe()

		// Check errors for starting http.Server
		if err != nil {
			logger.Fatal(err)
			os.Exit(1)
		}
	}()

	// Shutdown server gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	sig := <-c
	logger.Println("Signal: ", sig)

	// Shutdown context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)

}

func openProdDB(cfg config) (*sql.DB, error) {
	// Open up PostgreSQL using DSN specified in the environment variable
	db, err := sql.Open("postgres", cfg.proddb.dsn)
	if err != nil {
		return nil, err
	}

	// Set configuration of DB
	// Connection pool
	db.SetMaxOpenConns(cfg.proddb.maxOpenConns)

	// Idle Connections
	db.SetMaxIdleConns(cfg.proddb.maxIdleConns)

	// Set duration and check for errors
	duration, err := time.ParseDuration(cfg.proddb.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
