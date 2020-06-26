package dndlookup

import (
  // "strings"
  "log"
  "fmt"

  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doEquipmentCategory(query string) string {
  searchResults, err := go5e.SearchEquipmentCategoryByName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return ""
  }
  spellIndex := getBestMatch(query, searchResults)
  spell, err := go5e.GetEquipmentCategory(spellIndex)
  if err != nil {
    log.Println(err)
    return ""
  }
  return formatEquipmentCategory(spell)
}

func formatEquipmentCategory(res go5e.EquipmentCategory) string {
  formatString := "%s"

  return fmt.Sprintf(formatString,
    res.Name,
  )
}
