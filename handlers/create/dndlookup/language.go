package dndlookup

import (
  // "strings"
  "log"
  // "fmt"

  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doLanguage(query string) string {
  searchResults, err := go5e.SearchLanguageByName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return ""
  }
  spellIndex := getBestMatch(query, searchResults)
  spell, err := go5e.GetLanguage(spellIndex)
  if err != nil {
    log.Println(err)
    return ""
  }
  return formatLanguage(spell)
}

func formatLanguage(res go5e.Language) string {
  return res.Name
}
