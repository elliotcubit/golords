package create

import (
  "github.com/bwmarrin/discordgo"
  "log"
  "strings"
)

func HandleMakePoll(s *discordgo.Session, m *discordgo.MessageCreate){
  log.Println("Creating poll")
  data := strings.SplitN(m.Content, " ", 2)
  if len(data) == 1 {
    return
  }

  p, err := s.ChannelMessageSend(m.ChannelID, data[1])
  if err != nil {
    log.Println("Couldn't send message for the poll")
    return
  }
  s.MessageReactionAdd(p.ChannelID, p.ID, "<:gamertime:533131012121821204>")
  s.MessageReactionAdd(p.ChannelID, p.ID, "<:robertscalls:533131034024214548>")
}
