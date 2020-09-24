package dndlookup

import (
	"fmt"
	"log"
	"strings"

	go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doSpell(query string) string {
	searchResults, err := go5e.SearchSpellByName(query)
	if err != nil || searchResults.Count < 1 {
		log.Println(err)
		return ""
	}
	spellIndex := getBestMatch(query, searchResults)
	spell, err := go5e.GetSpell(spellIndex)
	if err != nil {
		log.Println(err)
		return ""
	}
	return formatSpell(spell)
}

func formatSpell(spell go5e.Spell) string {

	formatString := "%s\nLevel %d %s\nCasting Time: %s\nRange: %s\nComponents: %s\nDuration: %s\nClasses: %s\n%s\n%s"

	componentsPPrint := strings.Join(spell.Components, " ") + " (" + spell.Material + ")"

	classesPPrint := ""

	for _, val := range spell.Classes {
		classesPPrint = classesPPrint + val.Name + " "
	}

	descStr := ""
	higherStr := ""

	if len(spell.Desc) > 0 {
		descStr = spell.Desc[0]
	}
	if len(spell.HigherLevel) > 0 {
		higherStr = "At higher levels: " + spell.HigherLevel[0]
	}
	return fmt.Sprintf(formatString,
		spell.Name,
		spell.Level,
		spell.School.Name,
		spell.CastingTime,
		spell.Range,
		componentsPPrint,
		spell.Duration,
		classesPPrint,
		descStr,
		higherStr,
	)
}
