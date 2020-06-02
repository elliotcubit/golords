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
  err = s.MessageReactionAdd(p.ChannelID, p.ID, ":thinking:")
  err2 := s.MessageReactionAdd(p.ChannelID, p.ID, ":robertscalls:")

  if err != nil {
    log.Println(err)
  }
  if err2 != nil {
    log.Println(err2)
  }
}
