package UserServices

import (
	"Exercirr-api/Database"
	"Exercirr-api/Helpers"
	Entity "Exercirr-api/Types/Entities"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(user Entity.User, c *gin.Context) {
	db := Database.DbConnection
	var getUserById string = fmt.Sprintf(`select * from user_entity ue WHERE id LIKE '%s'`, user.ID)
	dbResponse, err := db.Query(getUserById)
	if err != nil {
		log.Default().Println(err)
		panic(err)
	}
	foundUser := Entity.User{}
	for dbResponse.Next() {
		err := dbResponse.Scan(
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
	var passwordsMatch bool = Helpers.CompareHashAndPassword(user.Password, foundUser.Password)
	if user.Username != foundUser.Username || !passwordsMatch || foundUser.ID == "" {
		c.JSON(http.StatusUnauthorized, "Invalid Credentials")
		return
	} else {
		token, err := Helpers.CreateToken(user)
		if err != nil {
			log.Default().Println(err)
			panic(err)
		}
		c.JSON(http.StatusOK, token)
		return
	}
}
