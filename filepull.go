package main

import (
  "golords/quotemanager"
  "golords/handlers"
)

func InitializeStoredData() error {
  err := quotemanager.LoadQuoteList("quotelist.json")
  if err != nil {
    return err
  }
  err = handlers.LoadBanWords("banwords.txt")
  if err != nil {
    return err
  }
  return nil
}
