package routes

import (
	"gym-manager/controllers"
	"gym-manager/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterClassRoutes(router *gin.Engine) {
	protectedRoutes := router.Group("/protected")
	protectedRoutes.Use(middleware.RequireAuth)

	{
		protectedRoutes.POST("/classes", middleware.AdminOnly(), controllers.CreateAClass)
		protectedRoutes.GET("/classes", controllers.GetAllClasses)
		protectedRoutes.GET("/classes/:id", controllers.GetClassById)
		protectedRoutes.POST("/classes/newMember", controllers.EnrollMemberInClass)
	}
}
