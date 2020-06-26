package dndlookup

import (
  // "strings"
  "log"
  // "fmt"

  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doRace(query string) string {
  searchResults, err := go5e.SearchRaceByName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return ""
  }
  spellIndex := getBestMatch(query, searchResults)
  spell, err := go5e.GetRace(spellIndex)
  if err != nil {
    log.Println(err)
    return ""
  }
  return formatRace(spell)
}

func formatRace(res go5e.Race) string {
  return res.Name
}
