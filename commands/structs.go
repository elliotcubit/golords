package commands

import (
  "github.com/bwmarrin/discordgo"
)

type ActiveModule interface {
  // Do what the module Do
  Do(s *discordgo.Session, m *discordgo.MessageCreate)
  // Return a description of what the module does
  Help() string
  // An array of each !<prefix> that should execute this module
  Prefixes() []string
}
