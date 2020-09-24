package dndlookup

import (
	// "strings"
	"fmt"
	"log"

	go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doProficiency(query string) string {
	searchResults, err := go5e.SearchProficiencyByName(query)
	if err != nil || searchResults.Count < 1 {
		log.Println(err)
		return ""
	}
	spellIndex := getBestMatch(query, searchResults)
	spell, err := go5e.GetProficiency(spellIndex)
	if err != nil {
		log.Println(err)
		return ""
	}
	return formatProficiency(spell)
}

// TODO classes / races
func formatProficiency(res go5e.Proficiency) string {
	formatString := "%s\nType: %s\n"

	return fmt.Sprintf(formatString,
		res.Name,
		res.Type,
	)
}
