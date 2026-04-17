package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=postgres password=root dbname=task-manager-gorm port=5433"

func DataBaseConnection() {
	gorm.Open(postgres.Open(DSN, &gorm.Config{}))
}
