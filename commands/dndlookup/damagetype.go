package dndlookup

import (
  // "strings"
  "log"
  "fmt"

  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doDamageType(query string) string {
  searchResults, err := go5e.SearchDamageTypeByName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return ""
  }
  spellIndex := getBestMatch(query, searchResults)
  spell, err := go5e.GetDamageType(spellIndex)
  if err != nil {
    log.Println(err)
    return ""
  }
  return formatDamageType(spell)
}

func formatDamageType(res go5e.DamageType) string {
  formatString := "%s\n%s"

  return fmt.Sprintf(formatString,
    res.Name,
    res.Desc[0],
  )
}
