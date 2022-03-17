package Helpers

import (
	Entity "Exercirr-api/Types/Entities"
	"math/rand"
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
func GenerateId() string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 27)
	for i := range b {
		b[i] = letterBytes[seededRand.Intn(len(letterBytes))]
	}
	b[6] = '-'
	b[13] = '-'
	b[20] = '-'
	return string(b)
}
