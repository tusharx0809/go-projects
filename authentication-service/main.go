package main

import (
	"fmt"
	"log"
	"os"

	dbmanager "authentication-service/DBmanager"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	var connection_string = "postgres://" + os.Getenv("USER") + ":" + os.Getenv("PASSWORD") + "@localhost:" + os.Getenv("PORT") + "/auth-server?sslmode=disable"
	//fmt.Println(connection_string)

	pool, err := dbmanager.DBManager(connection_string)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()
	fmt.Println("Connection successful")

}
