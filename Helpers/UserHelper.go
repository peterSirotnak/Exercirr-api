package Helpers

import (
	Entity "Exercirr-api/Types/Entities"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func CreateToken(user Entity.User) (string, error) {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfokjdbutdjbotfidbjniutfdlbntfliubntdsubnfdiulbhnftibjnftlbitfnbjfksd")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	// 5e4d5s-4d8de9-48ds9d-898fd4
	return token, nil
}

func CompareHashAndPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}
