package api

import (
	"net/http"
	"sync"
	"time"
	"github.com/gin-gonic/gin"
	"Coding_Challenge/internal/telegram"
)

type Notification struct {
	Type        string `json:"Type"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

var (
	telegramLimiter   = time.NewTicker(1 * time.Second)
	notifications     []Notification
	notificationsLock sync.Mutex
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/notify", handleNotification)
}

func handleNotification(c *gin.Context) {
	var notif Notification
	if err := c.BindJSON(&notif); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	// store all notifications
	notificationsLock.Lock()
	notifications = append(notifications, notif)
	notificationsLock.Unlock()

	// only forward warnings
	if notif.Type == "Warning" {
		<-telegramLimiter.C

		if err := telegram.Send(notif.Name, notif.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "forwarded"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ignored"})
}