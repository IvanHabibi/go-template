package datasource

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

const DB_MAX_IDLE_CONN = 2
const DB_MAX_OPEN_CONN = 500
const DB_MAX_LIFETIME_CONN = 2 * time.Minute

// OpenDB ..
func OpenDB() *gorm.DB {
	var err error
	if db == nil {
		db, err = gorm.Open(os.Getenv("DATABASE"), os.Getenv("CONNECT_DB"))
		if err != nil {
			panic(fmt.Sprintf("failed to connect database: %s", err.Error()))
		}
		db.DB().SetMaxIdleConns(DB_MAX_IDLE_CONN)
		db.DB().SetMaxOpenConns(DB_MAX_OPEN_CONN)
		db.DB().SetConnMaxLifetime(DB_MAX_LIFETIME_CONN)
		return db
	}

	if err = db.DB().Ping(); err != nil {
		db.Close()
		db, err = gorm.Open(os.Getenv("DATABASE"), os.Getenv("CONNECT_DB"))
		if err != nil {
			panic(fmt.Sprintf("failed to connect database: %s", err.Error()))
		}
		db.DB().SetMaxIdleConns(DB_MAX_IDLE_CONN)
		db.DB().SetMaxOpenConns(DB_MAX_OPEN_CONN)
		db.DB().SetConnMaxLifetime(DB_MAX_LIFETIME_CONN)
		return db
	}
	return db
}

//CloseDB ..
func CloseDB() error {
	if db == nil {
		return fmt.Errorf("Database instance nil")
	}
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
