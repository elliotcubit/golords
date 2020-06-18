package vote

import (
  "github.com/bwmarrin/discordgo"
  "log"
  "strings"
)

func HandleMakePoll(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 2)
  if len(data) == 1 {
    return
  }

  log.Println("Creating poll for \"%v\"", data[1])
  p, err := s.ChannelMessageSend(m.ChannelID, "POLL: " + data[1])
  if err != nil {
    log.Println("Couldn't send message for the poll")
    return
  }
  s.MessageReactionAdd(p.ChannelID, p.ID, "<:gamertime:533131012121821204")
  s.MessageReactionAdd(p.ChannelID, p.ID, "<:robertscalls:533131034024214548")
}
