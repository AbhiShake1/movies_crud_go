package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "getMovies",
	})
}

func GetMovie(c *gin.Context, id string) {
	c.JSON(http.StatusOK, gin.H{
		"message": "getMovie",
		"id":      id,
	})
}

func CreateMovie(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "createMovie",
	})
}

func UpdateMovie(c *gin.Context, id string) {
	c.JSON(http.StatusOK, gin.H{
		"message": "updateMovie",
		"id":      id,
	})
}

func DeleteMovie(c *gin.Context, id string) {
	c.JSON(http.StatusOK, gin.H{
		"message": "deleteMovie",
		"id":      id,
	})
}
