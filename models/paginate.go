package models

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	"math"
)

func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map {
	limit := 15
	offset := (page - 1) * limit; //if page is 2, start from the 5th record

	data := entity.Take(db, limit, offset)
	total := entity.Count(db)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":total,
			"page": page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	}

}

