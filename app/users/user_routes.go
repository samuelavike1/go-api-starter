package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupUserRoutes(router *gin.Engine, db *pgx.Conn) {
	repo := NewUserRepository(db)
	service := NewUserService(repo)

	router.GET("/users/:id", service.GetUser)
	router.POST("/create-user", service.CreateUser)
	router.GET("/get-all-users", service.GetAllUsers)
}
