package main

import (
	"fmt"
	"log"
	"os"

	dbmanager "authentication-service/DBmanager"
	"authentication-service/handlers"
	"authentication-service/repository"
	"authentication-service/server"
	"authentication-service/services"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	//fmt.Println(connection_string)

	pool, err := dbmanager.ConnectDB(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	repo := repository.NewAuthRepo(pool)

	service := services.NewAuthService(repo)

	handler := handlers.NewAuthHandler(service)

	server := server.StartServer(handler)
	fmt.Println("Server running on :8080")

	server_error := server.ListenAndServe()

	if server_error != nil {
		log.Fatal(server_error)
	}

	// var version string
	// err = pool.QueryRow(context.Background(), "select version()").Scan(&version)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(version)

}
