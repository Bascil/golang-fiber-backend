package controllers

import (
	"../database"
	"../models"
	"../util"
	"github.com/gofiber/fiber"
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

	// password,_ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Firstname:data["first_name"],
		Lastname:data["last_name"],
		Email:data["email"],
		// Password:password,
	}

	user.SetPassword(data["password"])

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

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

    token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

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

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)

	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	
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

