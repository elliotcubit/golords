package diceroll

import (
  "github.com/bwmarrin/discordgo"
  "strings"
  "golords/diceroll"
)

func HandleDiceRoll(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 2)
  if len(data) == 1 {
    return
  }
  query := data[1]
  strings.Replace(query, " ", "", -1)
  msg := diceroll.Do(query)
  if msg != "" {
    s.ChannelMessageSend(m.ChannelID, msg)
  }
}
