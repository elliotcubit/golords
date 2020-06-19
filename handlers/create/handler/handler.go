package handler

import (
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
func (h DefaultHandler) Should(hint string) bool { return false }
func (h DefaultHandler) Help() string { return "" }
