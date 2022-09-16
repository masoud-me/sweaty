package Database

import (
	"gorm.io/gorm"
	"log"
	"sweaty/Config"

	"gorm.io/driver/mysql"
)

type Db struct {
	db *gorm.DB
}

func connectToDatabase() Db {
	config := Config.AppConfig.Database
	dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/sweaty?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	return Db{db: db}
}