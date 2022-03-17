package main

import (
	UserControllers "Exercirr-api/Controllers/UserControllers"
	Database "Exercirr-api/Database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	Database.ConnectDb()
	router := gin.Default()
	router.POST("login/", UserControllers.Login)
	router.POST("sign-up/", UserControllers.SignUp)

	router.Run("localhost:8080")

}
