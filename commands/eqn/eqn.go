package eqn

import (
	"net/url"
	"strings"

	"golords/handlers"

	"github.com/bwmarrin/discordgo"
)

// We can plug URL-encoded LaTeX into this to get an image link. Nice!
var baseUrl string = `https://chart.apis.google.com/chart?cht=tx&chf=bg,s,FFFFFF00&chl=%0D%0A`

func init() {
	handlers.RegisterActiveModule(
		Eqn{},
	)
}

type Eqn struct{}

func (h Eqn) Do(s *discordgo.Session, m *discordgo.MessageCreate) {
	data := strings.SplitN(m.Content, " ", 2)
	if len(data) == 1 {
		return
	}
	s.ChannelMessageSend(m.ChannelID, baseUrl+url.QueryEscape(data[1]))
}

func (h Eqn) Help() string {
	return "Turn your message into LaTeX"
}

func (h Eqn) Prefixes() []string {
	return []string{"eqn", "latex", "equation", "math"}
}
