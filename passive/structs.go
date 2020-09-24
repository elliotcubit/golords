package passive

import (
	"github.com/bwmarrin/discordgo"
)

type PassiveModule interface {
	Do(s *discordgo.Session, m *discordgo.MessageCreate)
	Help() string
}
