package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitializeDbConnection(dbUrl string) *gorm.DB {
	dB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dB.Logger.LogMode(logger.LogLevel(0))
	return dB
}
