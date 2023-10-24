package main

import (
	"api/auth/db"
	"api/auth/models"
	"api/auth/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {

	//variable globales
	if err := godotenv.Load("./env"); err != nil {
		log.Fatal(err)
	}

	//tomamos variables de entorno
	IP_DB := os.Getenv("IP_DB")
	PORT_DB := os.Getenv("PORT_DB")
	USER_DB := os.Getenv("USER_DB")
	PASS_DB := os.Getenv("PASS_DB")
	NAME_DB := os.Getenv("NAME_DB")

	//vamos a conectar con la base de datos
	db.DBConnection(IP_DB, PORT_DB, USER_DB, PASS_DB, NAME_DB)

	//vamos a crear las tablas
	db.DB.AutoMigrate(models.UserAuth{})

	//acá se cre un objeto ruta del modulo mux
	router := mux.NewRouter()

	// se crean las primeras rutas
	//la funcion handlefunc lo que hace es recibir dos parametros
	//el primero es la ruta a la cual se va a dirigir
	//el segundo recibe la funcion de lo que va a responder
	//responde con una funcion
	router.HandleFunc("/", routes.Test).Methods("POST")
	router.HandleFunc("/Docs", routes.Docs).Methods("POST")
	router.HandleFunc("/Register", routes.Register).Methods("POST")
	router.HandleFunc("/Login", routes.Loggin).Methods("POST")
	router.HandleFunc("/Auth", routes.Auth).Methods("POST")
	router.HandleFunc("/User/{id_user}", routes.UserById).Methods("POST")

	//inicializamos el servidor
	//recibe el puerto y el router inicializador
	http.ListenAndServe(":3000", router)
}
