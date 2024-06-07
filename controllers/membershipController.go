package controllers

import (
	"gym-manager/initializers"
	"gym-manager/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Func to create a new membership
func CreateMembership(c *gin.Context) {
	var body struct {
		UserID    uint
		Type      models.MembershipType
		ImageURL  string
		StartDate string
		EndDate   string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read request body...",
		})
		return
	}

	membership := models.Membership{
		UserID:    body.UserID,
		Type:      body.Type,
		ImageURL:  body.ImageURL,
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
	}

	result := initializers.DB.Create(&membership)
	if result.Error != nil {
		// Log the error for debugging
		log.Printf("Error creating membership: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error":  "Membership wasn't created...",
			"Detail": result.Error.Error(), // Return the actual error message
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":    "Membership created successfully",
		"membership": membership,
	})
}

// Func to fetch all memberships
func GetAllMemberships(c *gin.Context) {
	var memberships []models.Membership

	result := initializers.DB.Find(&memberships)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Memberships not found",
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"Memberships": memberships,
	})
}

// Func to get a single membership
func GetOneMembership(c *gin.Context) {
	var membership models.Membership
	membershipID := c.Param("id")

	result := initializers.DB.First(&membership, membershipID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Mmebership not found...",
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"Membership": membership,
	})
}

// Func to update membership
func UpdateMembership(c *gin.Context) {
	userID := c.Param("id")

	var body struct {
		Type      models.MembershipType
		StartDate string
		EndDate   string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	var membership models.Membership
	memberResult := initializers.DB.First(&membership, userID)

	if memberResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Membership not found"})
		return
	}

	membership.Type = body.Type
	membership.StartDate = body.StartDate
	membership.EndDate = body.EndDate

	initializers.DB.Save(&membership)

	c.JSON(http.StatusOK, gin.H{
		"message":    "Membership updated successfully",
		"membership": membership,
	})
}

// Delete membership
func DeleteMembership(c *gin.Context) {
	userID := c.Param("id")
	result := initializers.DB.Delete(&models.Membership{}, userID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete membership",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Membership deleted successfully"})
}
