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
		db.GetInstance().Model(&User).AddForeignKey("company_id", "companies(id)", "RESTRICT", "RESTRICT")
		db.GetInstance().Model(&Experience).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.GetInstance().Model(&Experience).AddForeignKey("company_id", "companies(id)", "RESTRICT", "RESTRICT")

		return nil
	})
}
