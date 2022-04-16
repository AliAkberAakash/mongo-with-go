package main

import (
	"github.com/AliAkberAakash/mongo-with-go/config"
	routes "github.com/AliAkberAakash/mongo-with-go/route"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDB()

	routes.MoviesRoute(router)

	router.Run("localhost:8080")
}
