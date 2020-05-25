package create

import (
  "github.com/bwmarrin/discordgo"
)

var helpStr = `Help:
  !addquote/!aq: add quote to the library
  !getquote/!qg: get a random quote from the library
  !ping: Check if the bot is online
  !help: Bring up this message`

func HandleHelp(s *discordgo.Session, m *discordgo.MessageCreate){
  s.ChannelMessageSend(m.ChannelID, helpStr)
}
