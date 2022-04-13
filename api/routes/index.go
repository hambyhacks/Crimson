package api

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Crimson IMS\n"))
}
