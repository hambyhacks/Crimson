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
	"time"

	kitlogger "github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"

	api "github.com/hambyhacks/CrimsonIMS/api/routes"
	prodsrv "github.com/hambyhacks/CrimsonIMS/service/products"
	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

func main() {
	var cfg config
	var svc prodsrv.ProductService

	flag.IntVar(&cfg.port, "port", 9000, "API Server port.")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|staging|prod)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("IMS_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")
	flag.Parse()

	// Set logger to write msg to stdout.
	logger := log.New(os.Stdout, "User-API ", log.LstdFlags|log.Ldate|log.Ltime)
	klogger := kitlogger.NewLogfmtLogger(os.Stderr)

	// Open DB connection
	db := openDB(cfg)
	defer db.Close()
	logger.Println("DB connection pool established")

	svc = &prodsrv.ProdServ{}
	{
		repo, err := prodsrv.NewProdRepo(db, klogger)
		if err != nil {
			level.Error(klogger).Log("exit", err)
			os.Exit(1)
		}
		defer db.Close()
		svc = prodsrv.NewProdServ(repo, klogger)
	}

	// Declare http.Server struct
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      api.NewHTTPHandler(svc),
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
	signal.Notify(c, os.Kill)

	sig := <-c
	logger.Println("Signal: ", sig)

	// Shutdown context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)

}

func openDB(cfg config) *sql.DB {
	// Open up PostgreSQL using DSN specified in the environment variable
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil
	}

	// Set configuration of DB
	// Connection pool
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	// Idle Connections
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	// Set duration and check for errors
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil
	}
	return db
}
