package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/tomintaiga/yandex_partice_1/service"
	"github.com/tomintaiga/yandex_partice_1/store"
)

// getManagerLogin extrack manager login from request headers
func getManagerLogin(ctx *fiber.Ctx) (string, error) {
	login, ok := ctx.GetReqHeaders()["X-Manager-Login"]
	if login == "" || !ok {
		return "", fmt.Errorf("no auth header")
	}

	return login, nil
}

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

func ManagerRegisterEmployee(managerStore *store.ManagerStore, employeeStore *store.EmployeeStore) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		l := log.With().Str("func", "controllers.ManagerRegisterEmployee").Logger()

		// Get manager login
		managerLogin, err := getManagerLogin(c)
		if err != nil {
			l.Error().Err(err).Msg("No manager login found")
			return err
		}

		// Get manager object
		manager, err := managerStore.Get(managerLogin)
		if err != nil {
			l.Error().Err(err).Msg("Manager not found")
			return err
		}

		payload := struct {
			Login string `json:"login"`
		}{}

		err = c.BodyParser(&payload)
		if err != nil {
			l.Error().Err(err)
			return err
		}

		// Check if this manager already have such employee
		for _, cur := range manager.Employees {
			if cur.Login == payload.Login {
				l.Error().Str("employee", payload.Login).Str("manager", managerLogin).Msg("Employee already registered")
				return fmt.Errorf("already registered")
			}
		}

		// TODO: Start transaction
		employee, err := employeeStore.Register(payload.Login)
		if err != nil {
			l.Error().Err(err).Msg("Can't register employee")
			return err
		}

		manager, err = managerStore.AddEmployee(manager, employee)
		if err != nil {
			l.Error().Err(err).Msg("Can't add employee to manager")
			return err
		}

		l.Info().Str("manager", manager.Login).Str("employee", employee.Login).Msg("Add new employee succsessfuly")
		// TODO: End transaction

		return nil
	}
}
