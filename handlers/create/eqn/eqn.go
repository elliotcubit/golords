package eqn

import (
  "strings"
  "net/url"

  "golords/handlers/create/handler"

  "github.com/bwmarrin/discordgo"
)

var baseUrl string = `https://chart.apis.google.com/chart?cht=tx&chf=bg,s,FFFFFF00&chl=%0D%0A`

func New() handler.CreateHandler {
  return EqnHandler{}
}

type EqnHandler struct {
  handler.DefaultHandler
}

func (h EqnHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 2)
  if len(data) == 1 {
    return
  }
  s.ChannelMessageSend(m.ChannelID, baseUrl+url.QueryEscape(data[1]))
}

func (h EqnHandler) GetPrompts() []string {
  return []string{"!eqn", "!latex", "!equation", "!math"}
}

func (h EqnHandler) Help() string {
  return "Turn your message into LaTeX"
}

func (h EqnHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}
