package routes

import (
	"github.com/AliAkberAakash/mongo-with-go/controller"
	"github.com/gin-gonic/gin"
)

func MoviesRoute(router *gin.Engine) {
	router.GET("/movies", controller.GetAllMovies)
	router.POST("/movie", controller.InsertMovie)
	router.PUT("/movie/:id", controller.UpdateMovie)
	router.DELETE("/movie/:id", controller.DeleteMovie)
}
