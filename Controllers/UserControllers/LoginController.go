package UserControllers

import (
	UserServices "Exercirr-api/Services/UserServices"
	Entity "Exercirr-api/Types/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user Entity.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Request")
		return
	}
	UserServices.Login(user, c)
}
