package api

import (
	"go.uber.org/dig"
)

// Inject api
func Inject(container *dig.Container) error {
	if err := container.Provide(NewUserAPI); err != nil {
		return err
	}
	if err := container.Provide(NewCompanyAPI); err != nil {
		return err
	}
	if err := container.Provide(NewExperienceAPI); err != nil {
		return err
	}

	return nil
}
