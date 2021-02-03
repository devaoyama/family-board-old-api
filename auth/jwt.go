package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	PictureUrl string `json:"picture_url"`
	jwt.StandardClaims
}

func GenerateJwtToken(id int, name, pictureUrl string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["name"] = name
	claims["picture_url"] = pictureUrl
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
