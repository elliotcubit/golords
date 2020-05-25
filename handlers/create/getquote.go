package create

import (
  "github.com/bwmarrin/discordgo"
  "golords/quotemanager"
  "fmt"
  "time"
)

func HandleCreateGetQuote(s *discordgo.Session, m *discordgo.MessageCreate){
  q := quotemanager.GetRandomQuote()
  t, _ := time.Parse(time.RFC3339, q.Timestamp) // TODO error checking
  ts := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
  msg := fmt.Sprintf("\"%v\" added by %v on %v", q.Text, q.AddedBy, ts)
  s.ChannelMessageSend(m.ChannelID, msg)
}
