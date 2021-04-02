package migration

import (
	"go.uber.org/dig"

	"github.com/quangdangfit/getjob/app/interfaces"
	"github.com/quangdangfit/getjob/app/models"
)

// Migrate migrate to database
func Migrate(container *dig.Container) error {
	return container.Invoke(func(
		db interfaces.IDatabase,
	) error {
		User := models.User{}

		db.GetInstance().AutoMigrate(&User)
		//db.GetInstance().Model(&User).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")

		return nil
	})
}
