package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mefellows/home/config"
)

// Health struct
type Health struct {
	Status bool
	DB     bool
}

// HealthController wraps up student related routes.
type HealthController struct {
	config *config.Config
}

// NewHealthController returns a new controller.
func NewHealthController(config *config.Config) *HealthController {
	return &HealthController{
		config: config,
	}
}

// Get Health Route.
func (s *HealthController) Get(c *gin.Context) {
	health := Health{true, true}
	var err error
	if err == nil {
		c.JSON(200, health)
	} else {
		log.Println("[DEBUG] unable to render health endpoint")
		c.JSON(http.StatusNotFound, gin.H{"status": "file not found"})
	}
}
