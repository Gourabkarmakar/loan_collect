package routers

import (
	"github.com/gin-gonic/gin"
	"text.com/controller"
)

func UserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/register", controller.RegisterUser)
		userRoutes.PUT("/update", controller.RegisterUser) // Assuming this is the correct handler
	}
}
