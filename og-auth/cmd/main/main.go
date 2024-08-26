package main

import (
    "log"
    "net/http"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "github.com/naren/og-auth/pkg/routes"
)

func main() {
    router := mux.NewRouter()
    routes.RegisterRoutes(router)

    // Enable CORS
    headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
    origins := handlers.AllowedOrigins([]string{"*"}) // Change "*" to your frontend's domain for production

    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(router)))
}
