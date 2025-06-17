package main

import (
	"github.com/gin-gonic/gin"
    "authentication-app/config"
    "authentication-app/routes"
)

func main(){
	config.ConnectDB()

    r := gin.Default()
    routes.SetupRoutes(r)

    r.Run(":8080")

}

