package initializers

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitializeDbConnection(dbUrl string) *gorm.DB {
	newDatabaseLogger := logger.New(
		log.New(os.Stdout, "SQL\t", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second * time.Duration(5), // SQL querylerinin execute edilmelerinin ne kadar sürdükten sonra warning loglarının atılacağını belirtiriz. Eğer bir SQL querysinin execute süresi, belirtilen süreyi geçiyorsa bu query, yavaş olarak kabul edilecek ve warning logu basılacaktır.
			LogLevel:                  logger.Info,                    // Log leveli belirtiriz.
			IgnoreRecordNotFoundError: false,                          // Eğer `true` alırsa databasede alınan `ErrRecordNotFound`[databasee atılan select querylerinde eğer bir kayıt bulamıyorsa bu hata alınır] hatasını loglamaz; `false` alırsa da loglamaya devam eder.
			ParameterizedQueries:      false,                          // Eğer `false` alırsa, SQL loglarında alınan dynamic valueler de raw halinde belirtilir; `true` alırsa da bu valueler gizlenir mesela `$1$,$2$,$3$` şeklinde.
			Colorful:                  true,                           // Enable color
		},
	)
	dB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger:                 newDatabaseLogger,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",    // table name prefix, table for `User` would be `t_users`
			SingularTable: false, // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,
		},
		AllowGlobalUpdate:    false,
		FullSaveAssociations: true,
		DryRun:               false,
	})
	if err != nil {
		panic(err)
	}
	return dB
}
