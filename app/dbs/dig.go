package dbs

import (
	"go.uber.org/dig"
)

// Inject database
func Inject(container *dig.Container) error {
	if err := container.Provide(NewDatabase); err != nil {
		return err
	}
	return nil
}
