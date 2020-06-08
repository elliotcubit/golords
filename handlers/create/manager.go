package create

import (
  "github.com/bwmarrin/discordgo"
)

type CreateHandler func(*discordgo.Session, *discordgo.MessageCreate)

// Bind all commands to their handlers
func GetCreateFunctionMap() map[string]CreateHandler {
  res := make(map[string]CreateHandler)

  res["!aq"] = HandleCreateAddQuote
  res["!addquote"] = HandleCreateAddQuote
  res["!gq"] = HandleCreateGetQuote
  res["!getquote"] = HandleCreateGetQuote
  res["!ping"] = HandleCreatePing
  res["!help"] = HandleHelp
  res["!poll"] = HandleMakePoll
  res["!vote"] = HandleMakePoll
  res["!spell"] = HandleGetSpell

  return res
}
