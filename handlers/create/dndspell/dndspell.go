package dndspell

import (
	"fmt"
	"strings"
  "log"
	"sort"

	"golords/handlers/create/handler"

	"github.com/bwmarrin/discordgo"
	go5e "github.com/elliotcubit/go-5e-srd-api"
	"github.com/toldjuuso/go-jaro-winkler-distance"
)

func New() handler.CreateHandler {
  return SpellHandler{}
}

type SpellHandler struct {
  handler.DefaultHandler
}

func (h SpellHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
	data := strings.SplitN(m.Content, " ", 2)
	if len(data) == 1 {
		return
	}

  query := strings.ReplaceAll(data[1], " ", "+")

  searchResults, err := go5e.SearchSpellName(query)
  if err != nil || searchResults.Count < 1 {
    log.Println(err)
    return
  }

	// Sort with highest-first similarity based on Jaro-Winkler string distance
	sort.SliceStable(searchResults.Results, func(i, j int) bool {
		return jwd.Calculate(data[1], searchResults.Results[i].Name) > jwd.Calculate(data[1], searchResults.Results[j].Name)
	})

  spellIndex := searchResults.Results[0].Index

  spell, err := go5e.GetSpell(spellIndex)
  if err != nil {
    log.Println(err)
    return
  }

  s.ChannelMessageSend(m.ChannelID, formatSpell(spell))
}

func (h SpellHandler) GetPrompts() []string {
  return []string{"!spell"}
}

func (h SpellHandler) Help() string {
  return "Get information about a 5e spell"
}

func (h SpellHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
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

  if len(spell.Desc) > 0{
    descStr = spell.Desc[0]
  }
  if len(spell.HigherLevel) > 0{
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
