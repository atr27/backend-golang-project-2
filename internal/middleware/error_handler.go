package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ispu-monitoring/backend/internal/model"
)

// ErrorHandler handles errors globally
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		
		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			
			// Log the error
			log.Printf("Error occurred: %v", err.Err)
			
			// Return error response
			c.JSON(http.StatusInternalServerError, model.APIResponse{
				Success: false,
				Error: &model.APIError{
					Code:    "INTERNAL_ERROR",
					Message: "An internal error occurred",
					Details: err.Err.Error(),
				},
				Meta: &model.MetaData{
					Timestamp: time.Now(),
					Version:   "1.0.0",
				},
			})
			
			// Abort to prevent further processing
			c.Abort()
		}
	}
}

// Recovery middleware recovers from panics
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				
				c.JSON(http.StatusInternalServerError, model.APIResponse{
					Success: false,
					Error: &model.APIError{
						Code:    "PANIC",
						Message: "Server panic occurred",
						Details: err,
					},
					Meta: &model.MetaData{
						Timestamp: time.Now(),
						Version:   "1.0.0",
					},
				})
				
				c.Abort()
			}
		}()
		
		c.Next()
	}
}
