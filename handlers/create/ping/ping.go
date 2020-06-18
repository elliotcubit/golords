package ping

import (
  "strings"
  
  "golords/handlers/create/handler"

  "github.com/bwmarrin/discordgo"
)

func New() handler.CreateHandler {
  return PingHandler{}
}

type PingHandler struct {
  handler.DefaultHandler
}

func (h PingHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  s.ChannelMessageSend(m.ChannelID, "pong!")
}

func (h PingHandler) GetPrompts() []string {
  return []string{"!ping"}
}

func (h PingHandler) Help() string {
  return "Check if the bot is running."
}

func (h PingHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}
