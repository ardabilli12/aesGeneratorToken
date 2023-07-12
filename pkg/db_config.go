package pkg

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnectionGorm() *gorm.DB {
	gormDb, err := gorm.Open(mysql.Open(dbConfiguration()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return gormDb
}

func dbConfiguration() string {
	user := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
}
