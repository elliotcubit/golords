package handlers

import (
  "golords/handlers/create/handler"
  "golords/handlers/create/addquote"
  "golords/handlers/create/ball"
  "golords/handlers/create/diceroll"
  "golords/handlers/create/dndspell"
  "golords/handlers/create/getquote"
  "golords/handlers/create/help"
  "golords/handlers/create/ping"
  "golords/handlers/create/vote"

  "github.com/bwmarrin/discordgo"
)

const (
  JON_UUID = "202565623790698498"
)

// Does this syntax even work?
var commandPrompts = [] handler.CreateHandler{
  addquote.New(),
  ball.New(),
  diceroll.New(),
  dndspell.New(),
  getquote.New(),
  help.New(),
  ping.New(),
  vote.New(),
  //youtube.New(),
}

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  // Ignore ourself no matter what
  if m.Author.ID == s.State.User.ID {
    return
  }

  // ignore jon lol
  if m.Author.ID == JON_UUID {
    return
  }

  // Run appropriate command, if there is one
  for _, handler := range commandPrompts {
    if handler.Should(m.Content) {
      handler.Do(s, m)
    }
  }
}
