package getquote

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"golords/handlers"
	"golords/state"
	"log"
	"time"
)

func init() {
	handlers.RegisterActiveModule(
		GetQuote{},
	)
}

type GetQuote struct{}

func (h GetQuote) Do(s *discordgo.Session, m *discordgo.MessageCreate) {
	q, err := state.GetRandomQuote(m.GuildID)
	if err != nil {
		log.Println("Couldn't fetch quote from DB")
		return
	}
	t, _ := time.Parse(time.RFC3339, q.Timestamp) // TODO error checking
	ts := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
	msg := fmt.Sprintf("\"%v\" added by %v on %v", q.Text, q.AddedBy, ts)
	s.ChannelMessageSend(m.ChannelID, msg)
}

func (h GetQuote) Help() string {
	return "Get a quote from the library"
}

func (h GetQuote) Prefixes() []string {
	return []string{"getquote", "gq"}
}
