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
		Company := models.Company{}
		Experience := models.Experience{}

		db.GetInstance().AutoMigrate(&User, &Company, &Experience)

		// Experience
		db.GetInstance().Model(&Experience).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
		db.GetInstance().Model(&Experience).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")

		return nil
	})
}
