package services

import (
	"go.uber.org/dig"
)

// Inject services
func Inject(container *dig.Container) error {
	if err := container.Provide(NewUserService); err != nil {
		return err
	}
	if err := container.Provide(NewCompanyService); err != nil {
		return err
	}
	if err := container.Provide(NewExperienceService); err != nil {
		return err
	}

	return nil
}
