package handlers

import (
  "golords/handlers/create/handler"
  "golords/handlers/create/ping"
  "github.com/bwmarrin/discordgo"
)

// Does this syntax even work?
var commandPrompts = [] handler.CreateHandler{
  ping.New(),
}

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  // Ignore ourself no matter what
  if m.Author.ID == s.State.User.ID {
    return
  }

  // Run appropriate command, if there is one
  for _, handler := range commandPrompts {
    if handler.Should(m.Content) {
      handler.Do(s, m)
    }
  }
}
