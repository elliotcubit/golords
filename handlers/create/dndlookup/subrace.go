package dndlookup

import (
  // "strings"
  "log"
  "fmt"

  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doSubrace(query string) string {
  searchResults, err := go5e.SearchSubraceByName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return ""
  }
  spellIndex := getBestMatch(query, searchResults)
  spell, err := go5e.GetSubrace(spellIndex)
  if err != nil {
    log.Println(err)
    return ""
  }
  return formatSubrace(spell)
}

// TODO more info is available from api
func formatSubrace(res go5e.Subrace) string {
  formatString := "%s\nRace: %s\nLanguages: %s\n%s"

  return fmt.Sprintf(formatString,
    res.Name,
    res.Race.Name,
    res.LanguageDesc,
    res.Desc,
  )
}
