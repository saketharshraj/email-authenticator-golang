package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateEmailPayload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload struct {
			Email   string `json:"email"`
			Project string `json:"project"`
		}

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			c.Abort()
			return
		}

		if payload.Email == "" || payload.Project == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and Project Name must be present in the payload"})
			c.Abort()
			return
		}

		c.Set("email", payload.Email)
		c.Set("project_key", payload.Project)
		c.Next()
	}
}
