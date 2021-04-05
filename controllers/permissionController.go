package controllers

import (
	"../database"
	"../models"
	"github.com/gofiber/fiber"
)

func GetPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission //slice, similar to an array
	database.DB.Find(&permissions)

	return c.JSON(permissions)
}

