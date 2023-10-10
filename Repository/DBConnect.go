package Model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	connStr := "user=patoanctkz password=R7cvUjxB3uTV dbname=neondb host=ep-round-credit-135823.ap-southeast-1.aws.neon.tech sslmode=verify-full"
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
