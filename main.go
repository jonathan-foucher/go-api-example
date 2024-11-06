package main

import (
	"fmt"
    "os"
    "net/http"
    "github.com/joho/godotenv"
    "go-api-example/handlers"
)

func main() {
	godotenv.Load()
	HTTP_PORT := os.Getenv("HTTP_PORT")
	fmt.Println("Application is starting on port", HTTP_PORT)
	api := &handlers.ApiHandler{}
	http.ListenAndServe(":" + HTTP_PORT, api)
}
