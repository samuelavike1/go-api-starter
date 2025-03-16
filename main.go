package main

import (
	"github.com/gin-gonic/gin"
	"golang/album"
)

func main() {
	router := gin.Default()
	router.GET("/albums", album.GetAlbums)
	router.POST("/albums", album.PostAlbums)
	router.GET("/albums/:id", album.GetAlbumByID)
	router.GET("/test", album.Test)

	router.Run("localhost:8089")
}
