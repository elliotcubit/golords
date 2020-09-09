package ian

import (
  "string"
  "fmt"

  "golords/handlers/create/handler"

  "github.com/bwmarrin/discordgo"
)

/*
  TODO only enable this in my server.
*/

func New() handler.IanHandler {
  return IanHandler{}
}

type IanHandler struct {
  handler.DefaultHandler
}

func (h IanHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  content := strings.toLower(m.Content)

  buyWords := []string{"buy", "bought", "purchase", "get a new", "house"}

  shouldTrigger := false

  for word := range buyWords {
    shouldTrigger = shouldTrigger || strings.Contains(content, word)
  }

  shouldTrigger = shouldTrigger && strings.contains("ian")

  if !shouldTrigger {
    return
  }

  // s.ChannelMessageSend(m.ChannelID, "@'Ina#2077, don't buy that thing!!!!!")
  s.ChannelMessageSend(m.ChannelID, "@jesvschrist#1548, don't buy that thing!!!!!")
}

func (h IanHandler) GetPrompts() []string {
  return []string{}
}

func (h IanHandler) Help() string {
  return ""
}

func (h IanHandler) Should(hint string) bool {
  // Always call Do() from the handler
  return true
}
