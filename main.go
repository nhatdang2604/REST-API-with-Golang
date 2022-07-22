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

	albums = []album {
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	PATH = "localhost:8080"
	albums_PATH = "/albums"
)

//Respond with the list of all albums as JSON
func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

//Add an album from JSON recieved in the request body
func postAlbum(context *gin.Context) {
	var newAlbum album
	
	//Call BindJSON to bind the recievedJSON to newAlbum
	if err := context.BindJSON(&newAlbum); nil != err {
		return
	}

	//Add the new album to the slice
	albums = append(albums, newAlbum)
	context.IntendedJSON(http.StatusCreated, newAlbum)


}

func main() {
	router := gin.Default()
	router.GET(albums_PATH, getAlbums)
	router.Run(PATH)
}
