package config

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GinBadRequest(c *gin.Context, message string, err error) {
	log.Printf(" ⛔ Bad Request: %s, Error: %s", message, err.Error())
	c.JSON(400, gin.H{
		"success": false,
		"message": message,
		"error":   err.Error(),
		"code":    400,
	})
}

func GinInternalServerError(c *gin.Context, message string, err error) {
	log.Printf(" 📛 Internal Server Error: %s, Error: %s", message, err.Error())
	c.JSON(500, gin.H{
		"success": false,
		"message": message,
		"error":   err.Error(),
		"code":    500,
	})
}

func GinNotFound(c *gin.Context, message string, err error) {
	log.Printf(" ❓ Not Found: %s, Error: %s", message, err.Error())
	c.JSON(404, gin.H{
		"success": false,
		"message": message,
		"error":   err.Error(),
		"code":    404,
	})
}

func GinUnauthorized(c *gin.Context, message string, err error) {
	log.Printf(" 🚷 Unauthorized: %s, Error: %s", message, err.Error())
	c.JSON(401, gin.H{
		"success": false,
		"message": message,
		"error":   err.Error(),
		"code":    401,
	})
}

func GinForbidden(c *gin.Context, message string, err error) {
	log.Printf(" 💢 Forbidden: %s, Error: %s", message, err.Error())
	c.JSON(403, gin.H{
		"success": false,
		"message": message,
		"error":   err.Error(),
		"code":    403,
	})
}

func GinStatusOK(c *gin.Context, message string, data interface{}) {
	log.Printf(" ✅ Status OK: %s", message)
	c.JSON(200, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func GinCreated(c *gin.Context, message string) {
	log.Printf(" ➕ Created: %s", message)
	c.JSON(201, gin.H{
		"success": true,
		"message": message,
	})
}
