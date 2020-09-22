package state

// Sets up connection and other variables for the rest of the library

import (
  "database/sql"
  "fmt"
  "os"
  "strings"
  "log"

  _ "github.com/lib/pq"
)

var database *sql.DB

// Return a name that doesn't have FUCKING quotes in it
func sanitizeName(user string) string {
  // Replace single quotes
  s := strings.ReplaceAll(user, "'", "")
  // Replace double quotes
  s = strings.ReplaceAll(s, "\"", "")
  return s
}

func init(){
  SQL_IP := os.Getenv("YUGABYTE_IP")
  SQL_PORT := 5433
  SQL_USER := os.Getenv("YUGABYTE_USER")
  SQL_PASS := os.Getenv("YUGABYTE_PASS")
  SQL_DB_NAME := os.Getenv("YUGABYTE_DB_NAME")

  if SQL_IP == "" {
    log.Fatalf("SQL_IP not set")
  }
  if SQL_USER == "" {
    log.Fatalf("SQL_USER not set")
  }
  if SQL_PASS == "" {
    log.Fatalf("SQL_PASS not set")
  }
  if SQL_DB_NAME == "" {
    log.Fatalf("SQL_DB_NAME not set")
  }

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
  } else {
    log.Println("Successfully connected to SQL database")
  }
}
