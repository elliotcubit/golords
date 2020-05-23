package create

import (
  "github.com/bwmarrin/discordgo"
)

func HandleCreatePing(s *discordgo.Session, m *discordgo.MessageCreate){
  s.ChannelMessageSend(m.ChannelID, "pong!")
}
