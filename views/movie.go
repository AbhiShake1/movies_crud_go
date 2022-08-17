package views

import (
	"github.com/gin-gonic/gin"
	"movies_crud_go/model"
	"net/http"
)

var movies []model.Movie

func GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, movies)
}

func getMovieIndex(id string) *int {
	for i, m := range movies {
		if m.ID == id {
			return &i
		}
	}
	return nil
}

func GetMovie(c *gin.Context, id string) {
	var movie = &movies[*getMovieIndex(id)]

	if movie != nil {
		c.JSON(http.StatusOK, movie)
	} else {
		c.JSON(http.StatusOK, "No movie found with this id")
	}
}

func CreateMovie(c *gin.Context) {
	id := c.PostForm("movie_id")
	isbn := c.PostForm("movie_isbn")
	title := c.PostForm("movie_title")
	firstName := c.PostForm("director_first_name")
	lastName := c.PostForm("director_last_name")
	movies = append(movies, model.Movie{
		ID:    id,
		Isbn:  isbn,
		Title: title,
		Director: &model.Director{
			FirstName: firstName,
			LastName:  lastName,
		},
	})
	c.JSON(http.StatusOK, "Movie added successfully")
}

func UpdateMovie(c *gin.Context, id string) {
	index := *getMovieIndex(id)

	movie := movies[index]

	isbn := c.DefaultPostForm("movie_isbn", movie.Isbn)
	title := c.DefaultPostForm("movie_title", movie.Title)
	firstName := c.DefaultPostForm("director_first_name", movie.Director.FirstName)
	lastName := c.DefaultPostForm("director_last_name", movie.Director.LastName)

	movies[index] = model.Movie{
		ID:    id,
		Isbn:  isbn,
		Title: title,
		Director: &model.Director{
			FirstName: firstName,
			LastName:  lastName,
		},
	}

	c.JSON(http.StatusOK, "Successfully updated")
}

func DeleteMovie(c *gin.Context, id string) {
	index := *getMovieIndex(id)

	movies = append(movies[:index], movies[index+1:]...)

	c.JSON(http.StatusOK, "Successfully deleted")
}
