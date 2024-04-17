package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/orololuwa/db-driver-conn-and-repository/driver"
)
const portNumber = ":8080"

func main(){
	db, err := run()
	if (err != nil){
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))


	srv := &http.Server{
		Addr: portNumber,
		Handler: nil,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run()(*driver.DB, error){
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "bookings"
	dbUser := "orololuwa"
	dbPassword := ""
	dbSSL := "disable"

	// Connecto to DB
	log.Println("Connecting to dabase")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUser, dbPassword, dbSSL)

	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal("Cannot conect to database: Dying!", err)
	}
	log.Println("Connected to database")

	return db, nil
}