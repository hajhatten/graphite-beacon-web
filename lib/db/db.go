package db

import (
  "os"
  "log"
  "github.com/joho/godotenv"
  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite" //sqlite adapter
)

func databasePath() string {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("error loading .env file")
  }
  path := "./db/" + os.Getenv("MODE") + ".db"
  return path
}

func exists(path string) (bool, error) {
  _, err := os.Stat(path)
  if err == nil { return true, nil }
  if os.IsNotExist(err) { return false, nil }
  return true, err
}

// OpenDBConnection opens a db connection
func OpenDBConnection() *gorm.DB {
    path := databasePath()
    pathExists, err := exists(path)
    
    if err != nil || pathExists == false {
      log.Println("cannot find database folder, creating folder")
      _ = os.Mkdir("db", os.FileMode(0777))      
    }
    
    db, err := gorm.Open("sqlite3", path)
    if err != nil {
      panic("failed to connect database")
    }
    
    return db
}