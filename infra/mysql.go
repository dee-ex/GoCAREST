package infra

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMySQLSession creates a session to our database
func NewMySQLSession() (*gorm.DB, error) {
	var (
		usrnm, paswd, proto, host, dbname string = "root", "123qwe123qwe", "tcp", "127.0.0.1", "gocarest"
		port, mxopen, mxlife              int    = 3306, 10, 15
		partm                             bool   = true
	)

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=%t", usrnm, paswd, proto, host, port, dbname, partm)

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
