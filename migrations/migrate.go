package migrations

import (
	entities "example.com/goproject9/entities/customer"
	"gorm.io/gorm"
)

func MigrateEntities(dbContext *gorm.DB) {
	err := dbContext.AutoMigrate(&entities.Customer{})
	if err != nil {
		panic(err)
	}
}
