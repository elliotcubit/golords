package plusplus

import (
	"fmt"
	"strings"

	"golords/handlers"
	"golords/state"

	"github.com/bwmarrin/discordgo"
)

func init() {
	handlers.RegisterPassiveModule(
		PlusPlus{},
	)
}

// The most recently added-to message
var recents []*discordgo.User
var recentIncrement int

type PlusPlus struct{}

func (h PlusPlus) Do(s *discordgo.Session, m *discordgo.MessageCreate) {
	inc := strings.Contains(m.Content, "++")
	dec := strings.Contains(m.Content, "--")

	// Do not allow increment and decrement in same operation
	// for simplicity - we will need to do a lot of parsing to
	// add syntax like @someone ++ @someonelse --

	// NOT XOR lol
	if inc == dec {
		return
	}

	// If there are no recent mentions or someone is mentioned,
	// Update the recents list and use that.

	// Change target if there are no recents or somebody is mentioned
	// Also set the value of recentIncrement in this case
	if len(recents) == 0 || len(m.Mentions) != 0 {
		recents = m.Mentions
		recentIncrement = 3
		if dec {
			recentIncrement = -3
		}
		if strings.Contains(m.Content, "++12") || strings.Contains(m.Content, "--12") {
			recentIncrement = 12
			if dec {
				recentIncrement = -12
			}
		}
	}

	// If nobody was mentioned and there are recents,
	// reuse the most recent value of increment and mentions

	// Quit if no target after that step --
	// This case is only relevant on startup before
	// anyone has called  the bot yet.
	if len(recents) == 0 {
		return
	}

	outStr := ""

	for _, user := range recents {
		if user.ID == m.Author.ID {
			// Nice try, buckwheat
			score, _ := state.UpdateStacks(m.GuildID, user.String(), -12)
			outStr = outStr + fmt.Sprintf("%v --12 for trying to edit their own stacks. They now have %d.\n", user.String(), score)
			continue
		}
		score, err := state.UpdateStacks(m.GuildID, user.String(), recentIncrement)
		if err != nil {
			return
		}
		outStr = outStr + fmt.Sprintf("%v now has %d stacks\n", user.String(), score)
	}

	// Send what is hopefully not an enormous message
	s.ChannelMessageSend(m.ChannelID, outStr)
}

func (h PlusPlus) Help() string {
	return "Karma System"
}
