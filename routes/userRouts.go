package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishank838/go-jwt-implementation/controllers"
	"github.com/ishank838/go-jwt-implementation/middleware"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(middleware.GetUsers())
	incommingRoutes.GET("/users", controllers.GetUsers)
	incommingRoutes.POST("users/:user_id", controllers.GetUser())
}
