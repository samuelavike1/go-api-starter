package album

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type response struct {
	ResponseCode string  `json:"response_code"`
	Message      string  `json:"message"`
	Data         []album `json:"data"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, response{
		ResponseCode: "SUCCESS",
		Message:      "Request Successful",
		Data:         albums,
	})
}

func Test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"responseCode": "SUCCESS",
		"message":      "success",
		"data":         albums,
		//"data": gin.H{
		//	"albums": albums,
		//},
	})
}

func TestRequest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"title": "Album Test",
		},
	})
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, response{
			Message: "failed",
			Data:    nil,
		})
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, response{
		Message: "success",
		Data:    albums,
	})
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
