package dndlookup

import (
	// "strings"
	"fmt"
	"log"

	go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doMagicSchool(query string) string {
	searchResults, err := go5e.SearchMagicSchoolByName(query)
	if err != nil || searchResults.Count < 1 {
		log.Println(err)
		return ""
	}
	spellIndex := getBestMatch(query, searchResults)
	spell, err := go5e.GetMagicSchool(spellIndex)
	if err != nil {
		log.Println(err)
		return ""
	}
	return formatMagicSchool(spell)
}

func formatMagicSchool(res go5e.MagicSchool) string {
	formatString := "%s\n%s"

	return fmt.Sprintf(formatString,
		res.Name,
		res.Desc,
	)
}
