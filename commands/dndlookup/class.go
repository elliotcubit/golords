package dndlookup

import (
	// "strings"
	"fmt"
	"log"

	go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doClass(query string) string {
	searchResults, err := go5e.SearchClassByName(query)
	if err != nil || searchResults.Count < 1 {
		log.Println(err)
		return ""
	}
	spellIndex := getBestMatch(query, searchResults)
	spell, err := go5e.GetClass(spellIndex)
	if err != nil {
		log.Println(err)
		return ""
	}
	return formatClass(spell)
}

func formatClass(res go5e.Class) string {
	formatString := "%s\nHit Die: %d\n"

	return fmt.Sprintf(formatString,
		res.Name,
		res.HitDie,
	)
}
