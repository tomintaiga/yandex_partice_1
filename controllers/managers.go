package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/tomintaiga/yandex_partice_1/domain"
	"github.com/tomintaiga/yandex_partice_1/service"
	"github.com/tomintaiga/yandex_partice_1/store"
)

// getManagerLogin extrack manager login from request headers
func getManagerLogin(ctx *fiber.Ctx) (string, error) {
	l := log.With().Str("func", "controllers.getManagerLogin").Logger()

	login, ok := ctx.GetReqHeaders()["X-Manager-Login"]
	if login == "" || !ok {
		l.Error().Msg("No auth header")
		return "", fmt.Errorf("no auth header")
	}

	return login, nil
}

func getManager(ctx *fiber.Ctx, managerStore *store.ManagerStore) (*domain.Manager, error) {
	login, err := getManagerLogin(ctx)
	if err != nil {
		return nil, err
	}

	manager, err := managerStore.Get(login)
	if err != nil {
		return nil, err
	}

	return manager, nil
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

		// Get manager object
		manager, err := getManager(c, managerStore)
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
				l.Error().Str("employee", payload.Login).Str("manager", manager.Login).Msg("Employee already registered")
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

func GetManagerEmployees(managerStore *store.ManagerStore) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		manager, err := getManager(c, managerStore)
		if err != nil {
			return err
		}

		employes := make([]string, len(manager.Employees))
		for i, _ := range manager.Employees {
			employes[i] = manager.Employees[i].Login
		}

		payload := struct {
			Employees []string `json:"employees"`
		}{
			Employees: employes,
		}

		return c.JSON(payload)
	}
}

func GetEmployeeInfo(managerStore *store.ManagerStore, employeeManager *store.EmployeeStore) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		l := log.With().Str("func", "controllers.GetEmployeeInfo").Logger()

		manager, err := getManager(c, managerStore)
		if err != nil {
			return err
		}

		l = l.With().Str("manager", manager.Login).Str("employee", c.Params("login")).Logger()

		for _, cur := range manager.Employees {
			l.Debug().Msgf("Checking employee %v", cur.Login)
			if cur.Login == c.Params("login") {
				employee, err := employeeManager.GetEmployee(cur.Login)
				if err != nil {
					l.Error().Msg("Employee not found #1")
					return err
				}

				l.Debug().Msg("Employee found")
				return c.JSON(employee)
			}
		}

		l.Error().Msg("Employee not found #2")
		return fmt.Errorf("employee not found")
	}
}
