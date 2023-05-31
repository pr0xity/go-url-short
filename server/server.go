package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pr0xity/go-url-short/model"
	"github.com/pr0xity/go-url-short/utils"
)

func redirect(ctx *fiber.Ctx) error {
	shrtUrl := ctx.Params("redirect_url")
	shrt, err := model.FindByShrtUrl(shrtUrl)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not find shrt in database " + err.Error(),
		})
	}
	shrt.Clicked += 1
	err = model.UpdateShrt(shrt)
	if err != nil {
		fmt.Printf("error updating shrt: %v\n", err)
	}

	return ctx.Redirect(shrt.Original_url, fiber.StatusTemporaryRedirect)
}

func getAllShrts(ctx *fiber.Ctx) error {
	shrts, err := model.GetAllShrts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retreive all shrt links from database " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(shrts)
}

func getShrt(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "could not parse id " + err.Error(),
		})
	}

	shrt, err := model.GetShrtById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retreive shrt from database " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(shrt)
}

func createShrt(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var shrt model.Shrt
	err := ctx.BodyParser(&shrt)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error parsing JSON " + err.Error(),
		})
	}
	if shrt.Random {
		shrt.Shrt = utils.RandomUrl(10)
	}

	err = model.CreateShrt(shrt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create shrt in database " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(shrt)
}

func updateShrt(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var shrt model.Shrt
	err := ctx.BodyParser(&shrt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not parse JSON " + err.Error(),
		})
	}

	err = model.UpdateShrt(shrt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not update shrt in database " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(shrt)
}

func deleteShrt(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not parse id " + err.Error(),
		})
	}
	err = model.DeleteShrt(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not delete shrt from database " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "shrt deleted.",
	})
}

func SetupAndListen(app_host, app_port string) {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect_url", redirect)

	router.Get("/shrts", getAllShrts)
	router.Get("/shrts/:id", getShrt)
	router.Post("/shrts", createShrt)
	router.Patch("/shrts/:id", updateShrt)
	router.Delete("/shrts/:id", deleteShrt)

	router.Listen(fmt.Sprintf("%s:%s", app_host, app_port))
}
