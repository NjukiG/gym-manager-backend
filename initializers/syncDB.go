package initializers

import (
	"gym-manager/models"
	"log"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{}, &models.Trainer{}, &models.Class{}, &models.Membership{})

	if err != nil {
		log.Fatal("Failed to migrate model", err)
	}
}
