package main

import (
	"chatbot/model"
	"chatbot/routers"
	"os"
)

func main() {
	router := routers.NewRouter()
	model.Database(os.Getenv("MYSQL_DSN"))
	router.Run(":8080")
	// router.RunTLS(":8443", "cert/certificate.crt", "cert/private.key")
}
