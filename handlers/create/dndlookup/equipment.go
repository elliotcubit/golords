package dndlookup

import (
  // "strings"
  "log"
  "fmt"

  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doEquipment(query string) string {
  searchResults, err := go5e.SearchEquipmentByName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return ""
  }
  spellIndex := getBestMatch(query, searchResults)
  spell, err := go5e.GetEquipment(spellIndex)
  if err != nil {
    log.Println(err)
    return ""
  }
  return formatEquipment(spell)
}

func formatEquipment(res go5e.Equipment) string {
  formatString := "%s\nCost: %d%s\nWeight: %s\n"

  return fmt.Sprintf(formatString,
    res.Name,
    res.Cost.Quantity,
    res.Cost.Unit,
    res.Weight,
  )
}
