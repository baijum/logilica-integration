package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"test": "ok}`))
}

func route() (n *negroni.Negroni, rt *mux.Router) {
	router := mux.NewRouter()

	router.HandleFunc("/", HandleWebhook).Methods("GET", "POST")
	n = negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	return n, router
}
func main() {

	n, router := route()
	n.UseHandler(router)
	n.Run(":8080")

}
