package dndlookup

import (
  // "strings"
  "log"
  "fmt"

  go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doSkill(query string) string {
  searchResults, err := go5e.SearchSkillByName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return ""
  }
  spellIndex := getBestMatch(query, searchResults)
  spell, err := go5e.GetSkill(spellIndex)
  if err != nil {
    log.Println(err)
    return ""
  }
  return formatSkill(spell)
}

func formatSkill(res go5e.Skill) string {
  formatString := "%s\n%s\nAbility score: %s"

  return fmt.Sprintf(formatString,
    res.Name,
    res.Desc[0],
    res.AbilityScore.Name,
  )
}
