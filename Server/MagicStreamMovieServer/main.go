package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	controllers "github.com/ayechanhan/GoNextMovieApp/Server/MagicStreamMovieServer/controllers"
)

func main(){
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context){
		c.String(200, "Hello Magic Stream Movies")
	})

	router.GET("/movies", controllers.GetMovies())
	
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Fail to start server!", err)
	}
}