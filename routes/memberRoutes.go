package routes

import (
	"gym-manager/controllers"
	"gym-manager/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterMemberRoutes(router *gin.Engine) {
	protectedRoutes := router.Group("/protected")
	protectedRoutes.Use(middleware.RequireAuth)

	{
		protectedRoutes.POST("/memberships", middleware.AdminOnly(), controllers.CreateMembership)
		protectedRoutes.GET("/memberships", controllers.GetAllMemberships)
		protectedRoutes.GET("/memberships/:id", controllers.GetOneMembership)
		protectedRoutes.PUT("/memberships/:id", middleware.AdminOnly(), controllers.UpdateMembership)
		protectedRoutes.DELETE("/memberships/:id", middleware.AdminOnly(), controllers.DeleteMembership)
	}
}
