package server

import (
	"fmt"
	"gorm/model"
	"gorm/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func redirect(c *fiber.Ctx) error {
	gormUrl := c.Params("redirect")

	gorm, err := model.FindByGormUrl(gormUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find gorm" + err.Error(),
		})
	}

	gorm.Clicked += 1

	err = model.UpdateGorm(gorm)
	if err != nil {
		fmt.Println("error updating clicks!")
	}
	return c.Redirect(gorm.Redirect, fiber.StatusTemporaryRedirect)
}

func getGorms(c *fiber.Ctx) error {
	gorms, err := model.GetAllGorms()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all gorm links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(gorms)
}

func getGorm(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}

	gorm, err := model.GetGorm(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrieve gorm from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(gorm)
}

func createGorm(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var gorm model.Gorm
	err := c.BodyParser(&gorm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if gorm.Random {
		gorm.Gorm = utils.RandomURL(8)
	}

	err = model.CreateGorm(gorm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create gorm " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(gorm)
}

func updateGorm(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var gorm model.Gorm
	err := c.BodyParser(&gorm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}

	err = model.UpdateGorm(gorm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update gorm " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(gorm)
}

func deleteGorm(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse id from url" + err.Error(),
		})
	}

	err = model.DeleteGorm(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not delete" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Gorm Deleted!",
	})
}
func SetupAndListen() {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect", redirect)
	router.Get("/gorm", getGorms)
	router.Get("/gorm/:id", getGorm)
	router.Post("/gorm", createGorm)
	router.Patch("/gorm", updateGorm)
	router.Delete("/gorm/:id", deleteGorm)

	router.Listen(":3000")
}
