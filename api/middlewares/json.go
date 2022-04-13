package middlewares

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hambyhacks/CrimsonIMS/app/models"
)

func MiddlewareValidateJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := &models.Products{}
		err := json.NewDecoder(r.Body).Decode(p)
		if err != nil {
			log.Println("[-] Error deserializing json", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = p.Validate()
		if err != nil {
			log.Println("[-] Error validating json", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
