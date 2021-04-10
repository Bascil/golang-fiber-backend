package controllers

import (
	"../database"
	"../models"
	"../middlewares"
	"github.com/gofiber/fiber"
	"strconv"
)

func GetUsers(c *fiber.Ctx) error {

	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page","1"))
	
	return c.JSON(models.Paginate(database.DB, &models.User{}, page))
}

func GetUser(c *fiber.Ctx) error {

	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id")) //id and error
	
	user := models.User{
		Id:uint(id),
	}

	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var data map[string]string //array with key string and value string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	roleId, _ := strconv.Atoi(data["role_id"]) //id and error

	user := models.User{
		Firstname:data["first_name"],
		Lastname:data["last_name"],
		Email:data["email"],
		RoleId: uint(roleId),
	}

	user.SetPassword("1234");

	database.DB.Create(&user); 
	return c.JSON(user)
}


func UpdateUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id")) //id and error
	
	user := models.User{
		Id:uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorized(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id")) //id and error
	
	user := models.User{
		Id:uint(id),
	}

	database.DB.Delete(&user)

	return nil
}