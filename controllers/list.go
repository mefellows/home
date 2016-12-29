package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mefellows/home/config"
	"github.com/mefellows/home/models"
)

// ListController wraps up student related routes.
type ListController struct {
	config *config.Config
}

// NewListController returns a new controller.
func NewListController(config *config.Config) *ListController {
	return &ListController{
		config: config,
	}
}

// List Route.
func (s *ListController) List(c *gin.Context) {
	var list models.List
	//s.config.DB.Preload("Items").Where(&models.List{Status: models.StatusNew}).Last(&list)
	//s.config.DB.Where("status = ?", models.StatusNew).Last(&list)
	s.config.DB.Preload("Items").Where("status = ?", models.StatusNew).Last(&list)

	var err error
	if err == nil {
		c.JSON(200, list)
	} else {
		log.Println("[DEBUG] unable to render health endpoint")
		c.JSON(http.StatusNotFound, gin.H{"status": "file not found"})
	}
}

// CompleteList sets the latest list to completed
func (s *ListController) CompleteList(c *gin.Context) {
	var list models.List

	// get latest list
	s.config.DB.Where("status = ?", models.StatusNew).Last(&list)

	list.Status = models.StatusDone
	if err := s.config.DB.Save(&list).Error; err != nil {
		c.JSON(500, models.Error{Code: 500, Message: err.Error()})
		return
	}
	s.config.DB.Preload("Items").Last(&list)
	c.JSON(200, &list)
}

// AppendItem adds an item to the latest incomplete list
// If there are none, create a new one and append to it.
func (s *ListController) AppendItem(c *gin.Context) {
	var list models.List
	var item models.Item

	// get latest list
	//s.config.DB.Where(&models.List{Status: models.StatusNew}).Last(&list)
	s.config.DB.Preload("Items").Where("status = ?", models.StatusNew).Last(&list)

	if c.Bind(&item) == nil {
		log.Println("[DEBUG] item", item)
		item.ListID = list.ID
		list.Items = append(list.Items, item)

		if err := s.config.DB.Save(&list).Error; err != nil {
			c.JSON(500, models.Error{Code: 500, Message: err.Error()})
			return
		}
		s.config.DB.Preload("Items").Last(&list)
		c.JSON(200, &list)
	} else {
		log.Println("[DEBUG] unable to append item to latest list")
		c.JSON(500, models.Error{Code: 500, Message: "unable to create item"})
	}
}

// GetItems route
func (s *ListController) GetItems(c *gin.Context) {
	id := c.Param("id")
	var items []models.Item
	if err := s.config.DB.First(&models.List{}, id).Related(&items).Error; err != nil {
		log.Println("[DEBUG] unable to find items for id", id)
		c.JSON(http.StatusNotFound, models.Error{Code: 400, Message: err.Error()})
	} else {
		c.JSON(200, items)
	}
}

// DeleteItem route
func (s *ListController) DeleteItem(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	s.config.DB.Delete(&models.Item{
		Model: models.Model{ID: id},
	})
	var err error
	if err == nil {
		c.JSON(200, gin.H{})
	} else {
		log.Println("[DEBUG] unable to find items for id", id)
		c.JSON(http.StatusNotFound, gin.H{"status": "file not found"})
	}
}
