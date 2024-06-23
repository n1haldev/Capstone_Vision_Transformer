package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/naren/og-auth/pkg/routes"
)

func main(){
	r:=mux.NewRouter()
	routes.RegisterRoutes(r)

	r.PathPrefix("/").Handler(http.StripPrefix("/",http.FileServer(http.Dir("D:/og-auth/"))))
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}