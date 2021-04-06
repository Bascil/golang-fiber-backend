package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"os"
)

func GenerateJwt(issuer string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: issuer, //convert integer to string
	    ExpiresAt: time.Now().Add(time.Hour*24).Unix(), // Convert 24 hours to unit time
	})

	return claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseJwt(cookie string)(string,error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error){
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims) //cast as claims

	return claims.Issuer, nil
}