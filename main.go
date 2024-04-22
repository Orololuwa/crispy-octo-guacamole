package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/orololuwa/crispy-octo-guacamole/driver"
	"github.com/orololuwa/crispy-octo-guacamole/models"
	"github.com/orololuwa/crispy-octo-guacamole/repository"
)
const portNumber = ":8080"

func main(){
	db, route, err := run()
	if (err != nil){
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))


	srv := &http.Server{
		Addr: portNumber,
		Handler: route,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run()(*driver.DB, *chi.Mux, error){
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

	userRepo := repository.NewUserRepo(db.SQL)
	dbRepo := repository.NewDBRepo(db.SQL)
	router := chi.NewRouter()

	router.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		type userBody struct {
			FirstName string `json:"firstName"`
			LastName string `json:"lastName"`
			Email string `json:"email"`
			Password string `json:"password"`
		}

		var body userBody
		
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user := models.User{
			FirstName: body.FirstName,
			LastName: body.LastName,
			Email: body.Email,
			Password: body.Password,
		}

		ctx := context.Background()
		var id int

		err = dbRepo.Transaction(ctx, func(ctx context.Context, tx *sql.Tx) error {
			id, err = userRepo.CreateAUser(ctx, tx, user)
			if err != nil {
				return err
			}

			userRepo.UpdateAUsersName(ctx, tx, id, body.FirstName, "test")
			if err != nil {
				return err
			}

			return nil
		})
		
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{"message": "user created successfully", "data": id}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	return db, router, nil
}