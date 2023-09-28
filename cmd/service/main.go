package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tomintaiga/yandex_partice_1/controllers"
	"github.com/tomintaiga/yandex_partice_1/service"
	"github.com/tomintaiga/yandex_partice_1/store"
)

const (
	secret = "AAAAbbbbb"
)

func main() {
	app := fiber.New()

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	managerRepo := store.ManagerStore{}
	managerService, err := service.NewManagerService(secret, &managerRepo)
	if err != nil {
		log.Panic().Err(err).Msg("Can't init manager service")
	}

	app.Post("/api/managers/register", controllers.ManagerRegister(managerService))

	app.Listen(":8080")
}
