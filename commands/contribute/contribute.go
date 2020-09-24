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
	s.ChannelMessageSend(m.ChannelID, "You can contribute to golordbot here:\nhttps://github.com/elliotcubit/golords")
}

func (h Contribute) Help() string {
	return "Learn to contribute to golordsbot"
}

func (h Contribute) Prefixes() []string {
	return []string{"contribute"}
}
