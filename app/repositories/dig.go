package repositories

import (
	"go.uber.org/dig"
)

// Inject repositories
func Inject(container *dig.Container) error {
	if err := container.Provide(NewUserRepository); err != nil {
		return err
	}
	if err := container.Provide(NewCompanyRepository); err != nil {
		return err
	}

	return nil
}
