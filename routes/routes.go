package routes

import (
    "github.com/gin-gonic/gin"
    "authentication-app/handlers"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/register", handlers.Register)
    r.POST("/login", handlers.Login)
}
