package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-github/v57/github"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	client := github.NewClient(nil)
	// list all organizations for user "willnorris"
	orgs, _, _ := client.Organizations.List(context.Background(), "baijum", nil)

	fmt.Println(orgs)
	var req github.Issue
	json.NewDecoder(r.Body).Decode(&req)
	fmt.Println(req)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"test": "ok}`))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"test": "ok}`))
}

func route() (n *negroni.Negroni, rt *mux.Router) {
	router := mux.NewRouter()

	router.HandleFunc("/", Handler).Methods("GET")
	router.HandleFunc("/", HandleWebhook).Methods("POST")
	n = negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	return n, router
}
func main() {

	n, router := route()
	n.UseHandler(router)
	n.Run(":8080")

}
