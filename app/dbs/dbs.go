package dbs

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/quangdangfit/gocommon/logger"

	"github.com/quangdangfit/getjob/app/interfaces"
	"github.com/quangdangfit/getjob/config"
)

type database struct {
	db *gorm.DB
}

func NewDatabase() interfaces.IDatabase {
	dbConfig := config.Config.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	logger.Info(dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logger.Fatal("Cannot connect to database: ", err)
	}

	// Set up connection pool
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(200)

	return &database{
		db: db,
	}
}

func (d *database) GetInstance() *gorm.DB {
	return d.db
}
