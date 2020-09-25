package contribute

import (
	"github.com/bwmarrin/discordgo"
	"golords/handlers"
)

func init() {
	handlers.RegisterActiveModule(
		Contribute{},
	)
}

type Contribute struct{}

func (h Contribute) Do(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{Color: 0x3498DB}
	embed.Title = "Contribute to golordsbot on GIthub"
	embed.URL = "https://github.com/elliotcubit/golords"
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func (h Contribute) Help() string {
	return "Learn to contribute to golordsbot"
}

func (h Contribute) Prefixes() []string {
	return []string{"contribute"}
}
