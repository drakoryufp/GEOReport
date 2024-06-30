package controller

import (
	"georeportapi/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorityHomepage(c *gin.Context) {
	// Retrieve user ID from the context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Type assertion to convert the user value to the User struct
	user, ok := userID.(entity.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type in context"})
		return
	}

	// Return user-specific information
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the authority homepage!",
		"user":    user,
	})
}