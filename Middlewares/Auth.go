package Middlewares

import (
	"SP/Helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct{}

func (s *Auth) AuthForDoctors(c *gin.Context) {
	clientToken := c.Request.Header.Get("token")
	if clientToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
		c.Abort()
		return
	}

	claims, err := Helper.ValidateToken(clientToken)
	if claims.Role_Type != "Doctor" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unauthorized")})
		c.Abort()
		return
	}
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

}

func (s *Auth) AuthForStudents(c *gin.Context) {
	clientToken := c.Request.Header.Get("token")
	if clientToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
		c.Abort()
		return
	}

	claims, err := Helper.ValidateToken(clientToken)
	if claims.Role_Type != "Student" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Unauthorized")})
		c.Abort()
		return
	}
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

}
