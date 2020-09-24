package dndlookup

import (
	// "strings"
	"fmt"
	"log"

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
	formatString := "%s\nSpeed: %d\nAge: %s\n Size: %s\n SizeDescription: %s\n Languages: %s"

	return fmt.Sprintf(formatString,
		res.Name,
		res.Speed,
		res.Age,
		res.Size,
		res.SizeDescription,
		res.LanguageDesc,
	)
}
