package dndlookup

import (
  // "strings"
  "log"
  // "fmt"

  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doFeature(query string) string {
  searchResults, err := go5e.SearchFeatureByName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return ""
  }
  spellIndex := getBestMatch(query, searchResults)
  spell, err := go5e.GetFeature(spellIndex)
  if err != nil {
    log.Println(err)
    return ""
  }
  return formatFeature(spell)
}

func formatFeature(res go5e.Feature) string {
  return res.Name
}
