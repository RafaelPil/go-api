package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type RouteResponse struct {
	Message string `json:"message"`
	ID      string `json:"id,omitempty"`
}

func main() {
	log.Println("Starting server...")

	router := mux.NewRouter()

	log.Println("Setting up routes...")

	router.Handle("/register", alice.New(loggingMiddleware).ThenFunc(register)).Methods("POST")
	
	router.Handle("/login", alice.New(loggingMiddleware).ThenFunc(login)).Methods("POST")
	
	router.Handle("/projects", alice.New(loggingMiddleware).ThenFunc(createProject)).Methods("POST")
	
	router.Handle("/projects/{id}", alice.New(loggingMiddleware).ThenFunc(updateProject)).Methods("PUT")
	
	router.Handle("/projects", alice.New(loggingMiddleware).ThenFunc(getProjects)).Methods("GET")
	
	router.Handle("/projects/{id}", alice.New(loggingMiddleware).ThenFunc(getProject)).Methods("GET")
	
	router.Handle("/projects/{id}", alice.New(loggingMiddleware).ThenFunc(deleteProject)).Methods("DELETE")

	log.Println("Listeinig on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}

// register
func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "Register"})
}

// login
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "login"})
}

// createProject
func createProject(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "createProject"})
}

// updateProject
func updateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "updateProject", ID: id})
}

// getProjects
func getProjects(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "getProjects"})
}

// getProject
func getProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "getProject", ID: id})
}

// deleteProject
func deleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(RouteResponse{Message: "deleteProject", ID: id})
}
