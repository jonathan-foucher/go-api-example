package main

import (
	"context"
	"fmt"
    "os"
    "net/http"
    db "go-api-example/database"
    "github.com/joho/godotenv"
    "go-api-example/handlers"
)

func main() {
	godotenv.Load()

	conn := db.InitDbConn()
	defer conn.Close(context.Background())

	HTTP_PORT := os.Getenv("HTTP_PORT")
	fmt.Println("Application is starting on port", HTTP_PORT)
	api := &handlers.ApiHandler{}
	http.ListenAndServe(":" + HTTP_PORT, api)
}
