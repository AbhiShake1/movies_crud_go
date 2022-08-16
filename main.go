package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"movies_crud_go/model"
	"movies_crud_go/views"
)

var movies []model.Movie

func main() {
	req := gin.Default()
	getMovieRoutes(req)
	err := req.Run()
	if err != nil {
		return
	}
}

func getMovieRoutes(req *gin.Engine) {
	req.GET(MOVIES, views.GetMovies)
	req.POST(MOVIES, views.CreateMovie)
	req.PUT(fmt.Sprintf("%v:id", MOVIES), func(c *gin.Context) {
		id := c.Param("id")
		views.UpdateMovie(c, id)
	})
	req.DELETE(fmt.Sprintf("%v:id", MOVIES), func(c *gin.Context) {
		id := c.Param("id")
		views.DeleteMovie(c, id)
	})
	req.GET(fmt.Sprintf("%v:id", MOVIES), func(c *gin.Context) {
		id := c.Param("id")
		views.GetMovie(c, id)
	})
}
