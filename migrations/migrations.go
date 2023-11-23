package migrations

import (
	"example.com/goproject9/entities"
	"gorm.io/gorm"
)

type MigrationsPackageModel struct {
}

type INewMigrationsPackage interface {
	DropAllTablesAndMigrateAllEntities(dbContext *gorm.DB)
	MigrateAllEntities(dbContext *gorm.DB)
	MigrateRoleEntity(dbContext *gorm.DB)
	MigrateScreenEntity(dbContext *gorm.DB)
	MigrateCustomerEntity(dbContext *gorm.DB)
	MigrateProfileEntity(dbContext *gorm.DB)
}

func NewMigrationsPackage() INewMigrationsPackage {
	return &MigrationsPackageModel{}
}

func (migrationsPackageModel *MigrationsPackageModel) DropAllTablesAndMigrateAllEntities(dbContext *gorm.DB) {
	if dbContext.Migrator().HasTable(&entities.CustomerEntity{}) {
		dbContext.Migrator().DropTable(&entities.CustomerEntity{})
	}
	if dbContext.Migrator().HasTable(&entities.RoleEntity{}) {
		dbContext.Migrator().DropTable(&entities.RoleEntity{})
	}
	if dbContext.Migrator().HasTable(&entities.ScreenEntity{}) {
		dbContext.Migrator().DropTable(&entities.ScreenEntity{})
	}
	if dbContext.Migrator().HasTable(&entities.ProfileEntity{}) {
		dbContext.Migrator().DropTable(&entities.ProfileEntity{})
	}
	dbContext.AutoMigrate(&entities.RoleEntity{}, &entities.CustomerEntity{}, &entities.ScreenEntity{}, entities.ProfileEntity{})
}

func (migrationsPackageModel *MigrationsPackageModel) MigrateAllEntities(dbContext *gorm.DB) {
	if !dbContext.Migrator().HasTable(&entities.CustomerEntity{}) && !dbContext.Migrator().HasTable(&entities.RoleEntity{}) && !dbContext.Migrator().HasTable(&entities.ScreenEntity{}) && !dbContext.Migrator().HasTable(&entities.ProfileEntity{}) {
		if err := dbContext.AutoMigrate(&entities.RoleEntity{}, &entities.CustomerEntity{}, &entities.ScreenEntity{}, entities.ProfileEntity{}); err != nil {
			panic(err)
		}
	}
}

func (migrationsPackageModel *MigrationsPackageModel) MigrateRoleEntity(dbContext *gorm.DB) {
	if !dbContext.Migrator().HasTable(&entities.RoleEntity{}) {
		err := dbContext.AutoMigrate(&entities.RoleEntity{})
		if err != nil {
			panic(err)
		}
	}
}
func (migrationsPackageModel *MigrationsPackageModel) MigrateScreenEntity(dbContext *gorm.DB) {
	if !dbContext.Migrator().HasTable(&entities.ScreenEntity{}) {
		err := dbContext.AutoMigrate(&entities.ScreenEntity{})
		if err != nil {
			panic(err)
		}
	}
}
func (migrationsPackageModel *MigrationsPackageModel) MigrateCustomerEntity(dbContext *gorm.DB) {
	if !dbContext.Migrator().HasTable(&entities.CustomerEntity{}) {
		err := dbContext.AutoMigrate(&entities.CustomerEntity{})
		if err != nil {
			panic(err)
		}
	}
}
func (migrationsPackageModel *MigrationsPackageModel) MigrateProfileEntity(dbContext *gorm.DB) {
	if !dbContext.Migrator().HasTable(&entities.ProfileEntity{}) {
		err := dbContext.AutoMigrate(&entities.ProfileEntity{})
		if err != nil {
			panic(err)
		}
	}
}
