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
  formatString := "%s\nHitPoints: %d\nHitDice: %s\nAC: %d\n"CR: %d\n"Strength: %d\nDexterity: %d\nConstitution: %d\nIntelligence: %d\nWisdom: %d\nCharisma: %d\n"

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
