package dndlookup

import (
  // "strings"
  "log"
  "fmt"

  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doMonster(query string) string {
  searchResults, err := go5e.SearchMonsterByName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return ""
  }
  spellIndex := getBestMatch(query, searchResults)
  spell, err := go5e.GetMonster(spellIndex)
  if err != nil {
    log.Println(err)
    return ""
  }
  return formatMonster(spell)
}

func formatMonster(res go5e.Monster) string {
  formatString := `""` +
    `"%s\n"`+
    `"HitPoints: %d\n"`+
    `"HitDice: %s\n"`+
    `"AC: %d\n"`+
    `"CR: %d\n"`+
    `"Strength: %d\n"`+
    `"Dexterity: %d\n"`+
    `"Constitution: %d\n"`+
    `"Intelligence: %d\n"`+
    `"Wisdom: %d\n"`+
    `"Charisma: %d\n"`

  return fmt.Sprintf(formatString,
    res.Name,
    res.HitPoints,
    res.HitDice,
    res.ArmorClass,
    res.ChallengeRating,
    res.Strength,
    res.Dexterity,
    res.Constitution,
    res.Intelligence,
    res.Wisdom,
    res.Charisma,
  )
}
