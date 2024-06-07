package controllers

import (
	"gym-manager/initializers"
	"gym-manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Func to create a class
func CreateAClass(c *gin.Context) {
	var body struct {
		Name        string
		ImageURL    string
		Description string
		StartTime   string
		EndTime     string
		TrainerIDs  []uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read the request body...",
		})
		return
	}

	var trainers []*models.Trainer
	for _, id := range body.TrainerIDs {
		var trainer models.Trainer
		if initializers.DB.First(&trainer, id).Error == nil {
			trainers = append(trainers, &trainer)
		}
	}

	class := models.Class{
		Name:        body.Name,
		ImageURL:    body.ImageURL,
		Description: body.Description,
		StartTime:   body.StartTime,
		EndTime:     body.EndTime,
		Trainers:    trainers,
	}

	result := initializers.DB.Create(&class)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to create class...",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Class created successfully",
		"class":   class,
	})
}

// FUnc to enroll member in class
func EnrollMemberInClass(c *gin.Context) {
	var body struct {
		UserID  uint
		ClassID uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read request body...",
		})
		return
	}

	var user models.User
	userResult := initializers.DB.First(&user, body.UserID)
	if userResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var class models.Class
	classResult := initializers.DB.First(&class, body.ClassID)

	if classResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Class not found",
		})
		return
	}

	initializers.DB.Model(&class).Association("Attendees").Append(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Enrolled in class successfully"})
}

// Func to get all classes
func GetAllClasses(c *gin.Context) {
	var classes []models.Class
	result := initializers.DB.Preload("Trainers").Preload("Attendees").Find(&classes)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Classes not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"classes": classes})
}

// Func to get a single class
func GetClassById(c *gin.Context) {
	var class models.Class
	classID := c.Param("id")

	result := initializers.DB.First(&class, classID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Class not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"class": class})
}
