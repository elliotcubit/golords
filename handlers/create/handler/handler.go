package handler

import (
  "strings"
  
  "github.com/bwmarrin/discordgo"
)

// Defines a handler for mesage prompts
type CreateHandler interface {
  // Do the handler's action
  Do(s *discordgo.Session, m *discordgo.MessageCreate)

  // Prompts that should call this handler.
  GetPrompts() []string

  // hint in GetPrompts() ?
  Should(hint string) bool

  // Returns a description of the command
  Help() string
}

type DefaultHandler struct {}

func (h DefaultHandler) Do() {}
func (h DefaultHandler) GetPrompts() []string { return nil }
func (h DefaultHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}
func (h DefaultHandler) Help() string { return "" }

/*
// Bind all commands to their handlers
func GetCreateHandlers() []CreateHandler {
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
  res["!8ball"] = HandleEightBall
  res["!r"] = HandleDiceRoll
  res["!goroll"] = HandleDiceRoll

  return res
}
*/
