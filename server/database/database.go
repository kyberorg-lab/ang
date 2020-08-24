package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kyberorg/ang/database/model"
	"github.com/kyberorg/ang/utils/defaults"
	"github.com/kyberorg/ang/utils/env"
	"log"
	"os"
)

var DBConn *gorm.DB

func InitDatabase() {
	var err error

	dbFile, dbLocExist := os.LookupEnv(env.DatabaseFile)
	if !dbLocExist {
		dbFile = defaults.DatabaseFile
	}

	DBConn, err = gorm.Open(defaults.DatabaseDialect, dbFile)
	if err != nil {
		fmt.Println("Error " + err.Error() + " DB File: " + dbFile)
		panic("Failed to connect to database")
	}
	fmt.Println("Database successfully connected. Database location:", dbFile)

	//auto migrate
	DBConn.AutoMigrate(&model.Hero{})
	fmt.Println("Database migrations are executed")
}

func CloseDatabase() {
	log.Fatal(DBConn.Close())
}
