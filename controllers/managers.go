package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/tomintaiga/yandex_partice_1/service"
)

func ManagerRegister(srv *service.ManagerService) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		payload := struct {
			SecretCode string `json:"secret_code"`
			Login      string `json:"login"`
		}{}

		l := log.With().Str("func", "controllers.ManagerRegister").Logger()

		err := c.BodyParser(&payload)
		if err != nil {
			l.Error().Err(err)
			return err
		}

		manager, err := srv.Register(payload.SecretCode, payload.Login)
		if err != nil {
			l.Error().Err(err)
			return err
		}

		l.Info().Str("login", manager.Login).Msg("New manager")

		return nil
	}
}
