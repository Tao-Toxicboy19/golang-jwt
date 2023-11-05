package orm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConfigDb() {
	dsn := "host=localhost user=postgres password=testpass123 dbname=nestjs port=4500 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the schema
	DB.AutoMigrate(&User{})
}
