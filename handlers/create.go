package handlers

import (
  "strings"

  "golords/handlers/create"
  "github.com/bwmarrin/discordgo"
)

var commandPrompts = create.GetCreateFunctionMap()

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  // Ignore ourself no matter what
  if m.Author.ID == s.State.User.ID {
    return
  }
  // Check to see if a command was requested
  for trigger, fun := range commandPrompts {
    if strings.HasPrefix(m.Content, trigger){
      fun(s, m)
    }
  }
}
