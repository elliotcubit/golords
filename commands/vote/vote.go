package vote

import (
  "golords/handlers"

  "github.com/bwmarrin/discordgo"
  "log"
  "strings"
)

func init(){
  handlers.RegisterActiveModule(
    Vote{},
  )
}

type Vote struct{}

func (h Vote) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 2)
  if len(data) == 1 {
    return
  }

  p, err := s.ChannelMessageSend(m.ChannelID, "POLL: " + data[1])
  if err != nil {
    log.Println("Couldn't send message for the poll")
    return
  }
  s.MessageReactionAdd(p.ChannelID, p.ID, "<:gamertime:533131012121821204")
  s.MessageReactionAdd(p.ChannelID, p.ID, "<:robertscalls:533131034024214548")
}

func (h Vote) Help() string {
  return "Start a vote in the server. Place a vote with the reactions!"
}

func (h Vote) Prefixes() []string {
  return []string{"poll", "vote", "v"}
}
