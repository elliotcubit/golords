package dndlookup

import (
  // "strings"
  "log"
  // "fmt"

  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doCondition(query string) string {
  searchResults, err := go5e.SearchConditionByName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return ""
  }
  spellIndex := getBestMatch(query, searchResults)
  spell, err := go5e.GetCondition(spellIndex)
  if err != nil {
    log.Println(err)
    return ""
  }
  return formatCondition(spell)
}

func formatCondition(res go5e.Condition) string {
  return res.Name
}
