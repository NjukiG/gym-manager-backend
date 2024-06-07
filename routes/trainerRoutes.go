package routes

import (
	"gym-manager/controllers"
	"gym-manager/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTrainers(router *gin.Engine) {
	protectedRoutes := router.Group("/protected")
	protectedRoutes.Use(middleware.RequireAuth)

	{
		protectedRoutes.POST("/trainers", middleware.AdminOnly(), controllers.CreateTrainer)
		protectedRoutes.GET("/trainers", controllers.GetAllTrainers)
		protectedRoutes.GET("/trainers/:id", controllers.GetTrainerById)
		protectedRoutes.POST("/trainers/assign", controllers.AssignTrainerToClass)
	}
}
