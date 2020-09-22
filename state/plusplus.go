package state

import (
  "fmt"
)

var createStackStatement string = `INSERT INTO stacks(serverID, userID, amount) VALUES ('%s', '%s', %d)`
var getStackRowStatement string = `SELECT amount FROM stacks WHERE serverID='%s' AND userID='%s'`
var updateStackRowStatement string = `UPDATE stacks SET amount=%d WHERE serverID='%s' AND userID='%s'`
var getTopStackRowStatement string = `SELECT userID, amount FROM stacks WHERE serverID='%s' ORDER BY amount LIMIT %d`

func GetStacksForUser(server, user string) (int, error) {
    var amount int;
    rows, err := database.Query(fmt.Sprintf(getStackRowStatement, server, user))
    if err != nil {
      return 0, err
    }
    defer rows.Close()
    for rows.Next() {
      err := rows.Scan(&amount)
      if err != nil {
        return 0, err
      }
    }
    // If we didn't get a result, amount will be 0 so we're gravy
    return amount, nil
}

func GetTopNStacks(server string, n int) (map[string]int, error) {
  var user string
  var amount int
  result := make(map[string]int, 0)
  rows, err := database.Query(fmt.Sprintf(getTopStackRowStatement, server, n))
  if err != nil {
    return result, err
  }
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&user, &amount)
    if err != nil {
      return result, err
    }
    result[user] = amount
  }
  return result, nil
}

func UpdateStacks(server, user string, amount int) (int, error) {
  var currentScore int;
  var updatedScore int;
  rows, err := database.Query(fmt.Sprintf(getStackRowStatement, server, user))
  if err != nil {
    return 0, err
  }
  defer rows.Close()
  didGetResult := false
  for rows.Next() {
    err := rows.Scan(&currentScore)
    if err != nil {
      return 0, err
    }
    didGetResult = true
  }
  // Create user if there wasn't a row
  if !didGetResult {
    updatedScore = amount
    err := ppCreateUser(server, user, updatedScore)
    if err != nil {
      return 0, err
    }
    // Otherwise, update the row
  } else {
    updatedScore = currentScore + amount
    // Update the row
    _, err := database.Exec(fmt.Sprintf(updateStackRowStatement, updatedScore, server, user))
    if err != nil {
      return 0, err
    }
  }
  return updatedScore, nil
}

func ppCreateUser(server, user string, amount int) error {
  _, err := database.Exec(fmt.Sprintf(createStackStatement, server, user, amount))
  return err
}
