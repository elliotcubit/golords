package vote

import (
  "golords/handlers/create/handler"
  
  "github.com/bwmarrin/discordgo"
  "log"
  "strings"
)

func New() handler.CreateHandler {
  return VoteHandler{}
}

type VoteHandler struct {
  handler.DefaultHandler
}

func (h VoteHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
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

func (h VoteHandler) GetPrompts() []string {
  return []string{"!poll", "!vote"}
}

func (h VoteHandler) Help() string {
  return "Start a vote in the server. Place a vote with the reactions!"
}

func (h VoteHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}
