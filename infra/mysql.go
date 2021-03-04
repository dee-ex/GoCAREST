package infra

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMySQLSession creates a session to our database
func NewMySQLSession() (*gorm.DB, error) {
	var (
		usrnm, paswd, proto, host, port, dbname string = os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")
		mxopen, mxlife                          int    = 10, 15
		partm                                   bool   = true
	)

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=%t", usrnm, paswd, proto, host, port, dbname, partm)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	mysqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	mysqlDB.SetMaxOpenConns(mxopen)
	mysqlDB.SetConnMaxLifetime(time.Duration(mxlife) * time.Second)

	return db, nil
}
