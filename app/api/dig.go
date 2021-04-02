package api

import (
	"go.uber.org/dig"
)

// Inject api
func Inject(container *dig.Container) error {
	if err := container.Provide(NewUserAPI); err != nil {
		return err
	}
	return nil
}
