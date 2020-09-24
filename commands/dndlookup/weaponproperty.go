package dndlookup

import (
	// "strings"
	"fmt"
	"log"

	go5e "github.com/elliotcubit/go-5e-srd-api"
)

func doWeaponProperty(query string) string {
	searchResults, err := go5e.SearchWeaponPropertyByName(query)
	if err != nil || searchResults.Count < 1 {
		log.Println(err)
		return ""
	}
	spellIndex := getBestMatch(query, searchResults)
	spell, err := go5e.GetWeaponProperty(spellIndex)
	if err != nil {
		log.Println(err)
		return ""
	}
	return formatWeaponProperty(spell)
}

func formatWeaponProperty(res go5e.WeaponProperty) string {
	formatString := "%s\n%s"

	return fmt.Sprintf(formatString,
		res.Name,
		res.Desc[0],
	)
}
