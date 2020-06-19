package diceroll

import (
  "golords/handlers/create/handler"
  "github.com/bwmarrin/discordgo"
  "strings"
  "golords/diceroll"
)

func New() handler.CreateHandler {
  return DicerollHandler{}
}

type DicerollHandler struct {
  handler.DefaultHandler
}

func (h DicerollHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
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

func (h DicerollHandler) GetPrompts() []string {
  return []string{"!r", "!goroll"}
}

func (h DicerollHandler) Help() string {
  return "Roll a die. {sides}d{amount}+{extra rolls / constants}"
}

func (h DicerollHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}
