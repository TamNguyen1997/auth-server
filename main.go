package main

import (
	"github.com/gin-gonic/gin"
	"main.go/controller"
)

func main() {

	auth := controller.NewAuthenticator()

	server := gin.Default()
	auth.AddRoute(server)

	server.Run()
}
