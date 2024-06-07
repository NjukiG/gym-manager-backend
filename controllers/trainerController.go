package controllers

import (
	"gym-manager/initializers"
	"gym-manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTrainer(c *gin.Context) {
	var body struct {
		FirstName string
		LastName  string
		ImageURL  string
		Email     string
		Bio       string
		Phone     string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request data...",
		})
		return
	}

	trainer := models.Trainer{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		ImageURL:  body.ImageURL,
		Email:     body.Email,
		Bio:       body.Bio,
		Phone:     body.Phone,
	}

	result := initializers.DB.Create(&trainer)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to  create new Trainer...",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Trainer created successfully",
		"Trainer": trainer,
	})
}

// Fun to fetch all trianers available
func GetAllTrainers(c *gin.Context) {
	var trainers []models.Trainer
	result := initializers.DB.Find(&trainers)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Trianers not found...",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Trainers": trainers,
	})
}

func GetTrainerById(c *gin.Context) {
	var trainer models.Trainer
	trainerID := c.Param("id")

	result := initializers.DB.First(&trainer, trainerID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Trainer not found...",
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"Trainer": trainer,
	})
}

// Assign a trainer to a class
func AssignTrainerToClass(c *gin.Context) {
	var body struct {
		TrainerID uint
		ClassID   uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read request body",
		})
		return
	}

	var trainer models.Trainer
	trainerResult := initializers.DB.First(&trainer, body.TrainerID)

	if trainerResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Trainer with that ID not found..",
		})
		return
	}

	var class models.Class
	classResult := initializers.DB.First(&class, body.ClassID)

	if classResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Class with that ID not found",
		})
		return
	}
	initializers.DB.Model(&class).Association("Trainers").Append(&trainer)
	c.JSON(http.StatusOK, gin.H{"message": "Trainer assigned to class successfully"})

}
