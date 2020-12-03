package base

import (
  "os"
  "fmt"
  
  "Walter0697/GinBackend/models" 
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
  // Fetching environment variable for database 
  HOST := defaultIfNull(os.Getenv("DATABASE_HOST"), "", "localhost")
  PORT := defaultIfNull(os.Getenv("DATABASE_PORT"), "", "5432")
  USER := defaultIfNull(os.Getenv("DATABASE_USER"), "", "postgres")
  PWD := defaultIfNull(os.Getenv("DATABASE_PWD"), "", "postgrespwd")
  
  // Parsing into Connecting String
  dsn := fmt.Sprintf("host=%[1]s user=%[2]s password=%[3]s DB.name=postgres port=%[4]s sslmode=disable", HOST, USER, PWD, PORT)
  database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
      panic("Failed to connect to database!")
  }

  database.AutoMigrate(&models.Book{})
  database.AutoMigrate(&models.User{})

  DB = database
}

func defaultIfNull(assign_value string, check_value string, default_value string) string {
  if assign_value == check_value {
    return default_value
  }
  return assign_value
}