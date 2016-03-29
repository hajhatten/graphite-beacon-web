package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite" //sqlite adapter
)

// OpenDBConnection opens a db connection
func OpenDBConnection() *gorm.DB {
    db, err := gorm.Open("sqlite3", "db/development.db")
    if err != nil {
      panic("failed to connect database")
    }
    return db
}