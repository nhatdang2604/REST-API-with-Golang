package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//album represents data about a record album
type album struct {
	ID	string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64 `json:"price"`	
}

var (

	ALBUMS = []album {
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Aritst: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	PATH = "localhost:8080"
	ALBUMS_PATH = "/albums"
)

//Respond with the list of all albums as JSON
func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, ALBUMS)
}

func main() {
	router := gin.Default()
	router.GET(ALBUMS_PATH, getAlbums)
	router.Run(PATH)
}
