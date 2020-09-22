package state

// Sets up connection and other variables for the rest of the library

import (
  "database/sql"
  "fmt"
  "os"
  "log"

  _ "github.com/lib/pq"
)

var database *sql.DB

func init(){
  SQL_IP := os.Getenv("YUGABYTE_IP")
  SQL_PORT := 5433
  SQL_USER := os.Getenv("YUGABYTE_USER")
  SQL_PASS := os.Getenv("YUGABYTE_PASS")
  SQL_DB_NAME := os.Getenv("YUGABYTE_DB_NAME")

  loginString := fmt.Sprintf(
    "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    SQL_IP,
    SQL_PORT,
    SQL_USER,
    SQL_PASS,
    SQL_DB_NAME,
  )

  var err error
  database, err = sql.Open("postgres", loginString)
  if err != nil {
    log.Fatal("Failed to connect to SQL database")
  }
}
