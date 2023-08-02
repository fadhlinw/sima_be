package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userID int, nama_masjid string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["nama"] = nama_masjid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}
