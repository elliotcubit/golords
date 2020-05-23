package create

import (
  "github.com/bwmarrin/discordgo"
  "golords/quotemanager"
  "fmt"
)

func HandleCreateGetQuote(s *discordgo.Session, m *discordgo.MessageCreate){
  q := quotemanager.GetRandomQuote()
  msg := fmt.Sprintf("\"%v\" added by %v @ %v", q.Text, q.AddedBy, q.Timestamp)
  s.ChannelMessageSend(m.ChannelID, msg)
}
