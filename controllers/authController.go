package controllers

import (
	"../database"
	"../models"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string //array with key string and value string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message":"passwords do not match",
		})
	}

	password,_ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Firstname:data["first_name"],
		Lastname:data["last_name"],
		Email:data["email"],
		Password:password,
	}

	database.DB.Create(&user); // create user by reference

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string //array with key string and value string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)), //convert integer to string
	    ExpiresAt: time.Now().Add(time.Hour*24).Unix(), // Convert 24 hours to unit time
	})

	token, err := claims.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//Store token inside a cookie
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour*24),
		HTTPOnly: true, //frontend wont be able to access this cookie
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

type Claims struct {
	jwt.StandardClaims // Get all fields on standard claim
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error){
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*Claims) //cast as claims

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)
	
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour), //set expiration time to the past
		HTTPOnly: true, //frontend wont be able to access this cookie
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})


}

