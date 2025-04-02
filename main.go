package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang/album"
	"golang/app/users"
	"golang/config"
	"golang/database"
	"log"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close(context.Background())

	router := gin.Default()

	users.SetupUserRoutes(router, db)

	router.GET("/albums", album.GetAlbums)
	router.POST("/albums", album.PostAlbums)
	router.GET("/albums/:id", album.GetAlbumByID)
	router.GET("/test", album.Test)

	router.Run("localhost:8089")
}
