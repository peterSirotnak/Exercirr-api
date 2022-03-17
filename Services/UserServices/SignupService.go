package UserServices

import (
	Database "Exercirr-api/Database"
	"Exercirr-api/Helpers"
	Entity "Exercirr-api/Types/Entities"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	user.ID = Helpers.GenerateId()
	fmt.Println(user.ID)
	var getUserById string = fmt.Sprintf(`select * from user_entity ue WHERE id = '%s'`, user.ID)
	uniqueIdUser, err := db.Query(getUserById)
	if err != nil {
		log.Default().Println(err)
		panic(err)
	}
	foundUser := Entity.User{}
	for uniqueIdUser.Next() {
		err := uniqueIdUser.Scan(
			&foundUser.ID,
			&foundUser.Username,
			&foundUser.Password,
			&foundUser.Email,
		)
		if err != nil {
			log.Default().Println(err)
			panic(err)
		}
	}
	if foundUser.ID != "" {
		SignUp(user, c)
	}
	var createNewUser string = fmt.Sprintf(`INSERT INTO user_entity VALUES ('%s', '%s', '%s', '%s')`, user.ID, user.Username, user.Password, user.Email)
	resp, err := db.Query(createNewUser)
	if err != nil {
		log.Default().Println(err)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, "Email Already Exists")
		}
		panic(err)
	}
	fmt.Println(resp)
	c.JSON(http.StatusOK, "User Successfully Created")
}
