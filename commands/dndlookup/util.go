package dndlookup

import (
	"sort"

	go5e "github.com/elliotcubit/go-5e-srd-api"
	"github.com/toldjuuso/go-jaro-winkler-distance"
)

// Returns index field of the best match
func getBestMatch(query string, results go5e.NamedAPIResourceList) string {
	sort.SliceStable(results.Results, func(i, j int) bool {
		return jwd.Calculate(query, results.Results[i].Name) > jwd.Calculate(query, results.Results[j].Name)
	})
	return results.Results[0].Index
}
