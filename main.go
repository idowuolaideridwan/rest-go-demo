package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-go-demo/controllers"
	"rest-go-demo/database"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletPersonByID).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "ls-eab037bb7a56f39111861e2bec9ac27d321bf964.c3bcoovup1j1.eu-west-2.rds.amazonaws.com:3306",
			User:       "dbmasteruser",
			Password:   "{w{prh30dk=V(A$N(up[jwNAF1n3%T",
			DB:         "ewms_db",
		}

	fmt.Println("Connecting to database...")
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	//database.MigrateTables(&entity.Payment{}, &entity.History{})
}
