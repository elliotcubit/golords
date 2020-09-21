package getquote

import (
  "golords/handlers"
  "github.com/bwmarrin/discordgo"
  "golords/quotemanager"
  "fmt"
  "time"
)

func init(){
  handlers.RegisterActiveModule(
    GetQuote{},
  )
}

type GetQuote struct{}

func (h GetQuote) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  q := quotemanager.GetRandomQuote()
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
