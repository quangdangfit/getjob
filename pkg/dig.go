package pkg

import (
	"github.com/quangdangfit/gocommon/jwt"
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	if err := container.Provide(jwt.New); err != nil {
		return err
	}

	return nil
}
