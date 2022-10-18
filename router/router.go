package router

import (
	"middleware/controllers"
	"middleware/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.Use(middlewares.Authentication())
		userRouter.PUT("/", controllers.UserUpdate)
		userRouter.DELETE("/", controllers.UserDelete)
	}
	return r
}
