package handlers

import (
  "golords/handlers/create/handler"
  "golords/handlers/create/addquote"
  "golords/handlers/create/ball"
  "golords/handlers/create/diceroll"
  "golords/handlers/create/dndlookup"
  "golords/handlers/create/getquote"
  "golords/handlers/create/help"
  "golords/handlers/create/ping"
  "golords/handlers/create/vote"
  // "golords/handlers/create/anim"
  "golords/handlers/create/plusplus"
  "golords/handlers/create/querystacks"
  "golords/handlers/create/contribute"
  "golords/handlers/create/eqn"

  "github.com/bwmarrin/discordgo"
)

const (
  RULEBREAKER_UUID = "724433548567773235"
)

// Does this syntax even work?
var commandPrompts = [] handler.CreateHandler{
  addquote.New(),
  ball.New(),
  diceroll.New(),
  getquote.New(),
  help.New(),
  ping.New(),
  vote.New(),
  dndlookup.New(),
  plusplus.New(),
  querystacks.New(),
  contribute.New(),
  eqn.New(),
  // anim.New(),
  //youtube.New(),
}

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  // Ignore ourself no matter what
  if m.Author.ID == s.State.User.ID {
    return
  }

  // Ignore the rulebreaker role,
  // TODO set this to be configurable to
  // different roles.
  roles := m.Member.Roles
  for i := 0; i < len(roles); i++ {
    if roles[i] == RULEBREAKER_UUID {
      return
    }
  }


  // Run appropriate command, if there is one
  for _, handler := range commandPrompts {
    if handler.Should(m.Content) {
      handler.Do(s, m)
    }
  }
}
