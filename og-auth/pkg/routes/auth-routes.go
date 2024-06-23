package routes

import (
	"github.com/gorilla/mux"
	
	"github.com/naren/og-auth/pkg/collections"
)

var RegisterRoutes=func(router *mux.Router){
	router.HandleFunc("/login/",collections.CheckUser).Methods("POST")
	router.HandleFunc("/register/",collections.CreateUser).Methods("POST")
	router.HandleFunc("/viewusers/",collections.GetUser).Methods("GET")
	router.HandleFunc("/deluser/{userId}", collections.DeleteUser).Methods("DELETE")
}