package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := s.Repo.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *UserService) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := s.Repo.CreateUser(c.Request.Context(), user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch the created user from the database.
	createdUser, err := s.Repo.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve created user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"responseCode": http.StatusCreated,
		"message":      "Request Processed",
		"responseData": createdUser,
	}) // Return the whole user
}

func (s *UserService) GetAllUsers(c *gin.Context) {
	users, err := s.Repo.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": http.StatusOK,
		"message":      "Request Processed",
		"responseData": users,
	})
}
