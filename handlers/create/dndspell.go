package create

import (
	"fmt"
	"strings"
  "log"

	"github.com/bwmarrin/discordgo"
	go5e "github.com/elliotcubit/go-5e-srd-api"
)

/*
Acid Arrow
2nd level evocation
Casting Time: 1 action
Range: 90 feet
Components: V S M (Powdered rhubarb leaf and an adder's stomach)
Duration: Instantaneous
Classes: Wizard
A shimmering green arrow streaks toward a target within range and bursts in a spray of acid. Make a ranged spell attack against the target. On a hit, the target takes 4d4 acid damage immediately and 2d4 acid damage at the end of its next turn. On a miss, the arrow splashes the target with acid for half as much of the initial damage and no damage at the end of its next turn.
At Higher Levels: When you cast this spell using a spell slot of 3rd level or higher, the damage (both initial and later) increases by 1d4 for each slot level above 2nd.
*/

func formatSpell(spell go5e.Spell) string {

	formatString := "%s\nLevel %d %s\nCasting Time: %s\nRange: %s\nComponents:%s\nDuration: %s\nClasses: %s\n%s\nAt Higher Levels: %s"

	componentsPPrint := strings.Join(spell.Components, " ") + " " + spell.Material

	classesPPrint := ""

  for _, val := range spell.Classes {
    classesPPrint = classesPPrint + val.Name + " "
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
		spell.Desc[0],
		spell.HigherLevel[0],
	)
}

func HandleGetSpell(s *discordgo.Session, m *discordgo.MessageCreate) {
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

  spellIndex := searchResults.Results[0].Index

  spell, err := go5e.GetSpell(spellIndex)
  if err != nil {
    log.Println(err)
    return
  }

  s.ChannelMessageSend(m.ChannelID, formatSpell(spell))
}
