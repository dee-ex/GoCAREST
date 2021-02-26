package infra

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// DBInitialization connects to our database and returns gorm object
func DBInitialization() *gorm.DB {
	var ts = 1
	for db, err := NewMySQLSession(); ; db, err = NewMySQLSession() {
		if err == nil {
			return db
		}

		time.Sleep(time.Duration(ts) * time.Second)

		if ts == 512 {
			log.Fatal(err)
		}
		ts *= 2
	}
}
