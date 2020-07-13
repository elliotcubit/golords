package contribute

import (
  "strings"

  "golords/handlers/create/handler"

  "github.com/bwmarrin/discordgo"
)

func New() handler.CreateHandler {
  return ContributeHandler{}
}

type ContributeHandler struct {
  handler.DefaultHandler
}

func (h ContributeHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  s.ChannelMessageSend(m.ChannelID, "You can contribute to golordbot here:\nhttps://github.com/elliotcubit/golords")
}

func (h ContributeHandler) GetPrompts() []string {
  return []string{"!contribute"}
}

func (h ContributeHandler) Help() string {
  return "Get a link to the github repo for this bot :)"
}

func (h ContributeHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}
