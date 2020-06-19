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

  "github.com/bwmarrin/discordgo"
)

func New() handler.CreateHandler {
  return HelpHandler{}
}

type HelpHandler struct {
  handler.DefaultHandler
}

func (h HelpHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){

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


/*
var helpStr = `Help:
  !addquote/!aq: add quote to the library
  !getquote/!qg: get a random quote from the library
  !ping: Check if the bot is online
  !help: Bring up this message`

func HandleHelp(s *discordgo.Session, m *discordgo.MessageCreate){
  s.ChannelMessageSend(m.ChannelID, helpStr)
}
*/
