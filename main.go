package main

import (
	"auth/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}
	fmt.Println(e)

	port := os.Getenv("PORT")
	fmt.Println(port)

	// utils.ConnectDB()

	// Handle routes
	http.Handle("/", routes.Handlers())

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
