package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: issuer, //convert integer to string
	    ExpiresAt: time.Now().Add(time.Hour*24).Unix(), // Convert 24 hours to unit time
	})

	return claims.SignedString([]byte(SecretKey))
}

func ParseJwt(cookie string)(string,error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error){
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims) //cast as claims

	return claims.Issuer, nil
}