package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mail/config"
	"mail/models"
)

var db *gorm.DB

func InitDB() {
	c := config.GetConfig()
	dsn := c.GetString("db.user") + ":" +
		c.GetString("db.pass") + "@(" +
		c.GetString("db.host") + ":3306)/" +
		c.GetString("db.name") +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	migrationMode := c.GetBool("server.migrationMode")
	if migrationMode {
		errMigration := db.AutoMigrate(
			&models.Battalion{},
			&models.City{},
			&models.Item{},
			&models.Message{},
			&models.Parcel{},
			&models.SpecialUnit{},
			&models.Status{},
		)
		if errMigration != nil {
			return
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
