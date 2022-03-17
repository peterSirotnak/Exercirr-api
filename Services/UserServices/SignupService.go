package UserServices

import (
	Database "Exercirr-api/Database"
	Entity "Exercirr-api/Types/Entities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(user Entity.User, c *gin.Context) {
	db := Database.DbConnection
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashPassword)
	var createNewUser string = fmt.Sprintf(`INSERT INTO user_entity VALUES ('%s', '%s', '%s', '%s')`, user.ID, user.Username, user.Password, user.Email)
	db.Query(createNewUser)
	c.JSON(http.StatusOK, "User Successfully Created")
}
