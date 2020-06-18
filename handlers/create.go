package handlers

import (
  "fmt"
  
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

  fmt.Println("In message create")

  // Run appropriate command, if there is one
  for _, handler := range commandPrompts {
    fmt.Printf("Checking %v\n", handler)
    if handler.Should(m.Content) {
      fmt.Println("Should")
      handler.Do(s, m)
    }
  }
}
