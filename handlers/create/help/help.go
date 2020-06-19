package help

import (
  "golords/handlers/create/addquote"
  "golords/handlers/create/ball"
  "golords/handlers/create/diceroll"
  "golords/handlers/create/dndspell"
  "golords/handlers/create/getquote"
  "golords/handlers/create/ping"
  "golords/handlers/create/vote"

  "golords/handlers/create/handler"

  "strings"

  "github.com/bwmarrin/discordgo"
)

func New() handler.CreateHandler {
  return HelpHandler{
    // TODO these should really just be pointers,
    // And create.go should populate this list...
    Handlers: []handler.CreateHandler{
      addquote.New(),
      ball.New(),
      diceroll.New(),
      dndspell.New(),
      getquote.New(),
      ping.New(),
      vote.New(),
    },
  }
}

type HelpHandler struct {
  handler.DefaultHandler
  Handlers []handler.CreateHandler
}

func (h HelpHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  str := "\nHelp:\n"
  for _, module := range h.Handlers {
    str = appendHelp(str, module)
  }
  // Put ourselves in there for shits and giggles
  str = appendHelp(str, h)

  s.ChannelMessageSend(m.ChannelID, str)
}

func (h HelpHandler) GetPrompts() []string {
  return []string{"!help"}
}

func (h HelpHandler) Help() string {
  return "Display this message"
}

func (h HelpHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}

func appendHelp(base string, module handler.CreateHandler) string {
  base += strings.Join(module.GetPrompts(), ", ")
  base += ": "
  base += module.Help()
  base += "\n"
  return base
}
