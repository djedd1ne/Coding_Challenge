package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Notification struct {
	Type        string `json:"Type"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

var	notifications	[]Notification

func RegisterRoutes(r *gin.Engine) {
	r.POST("/notify", handleNotification)
}

func handleNotification(c *gin.Context) {
	var notif Notification
	if err := c.BindJSON(&notif); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	// Store all notifications
	notifications = append(notifications, notif)
}