package diceroll

import (
  "golords/handlers"
  "github.com/bwmarrin/discordgo"
  "strings"
  "golords/diceroll"
)

// TODO move diceroll from root to this package

func init() {
  handlers.RegisterActiveModule(
    DiceRoll{},
  )
}

type DiceRoll struct{}

func (h DiceRoll) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 2)
  if len(data) == 1 {
    return
  }
  query := data[1]
  strings.Replace(query, " ", "", -1)
  msg := diceroll.Do(query)
  if msg != "" {
    s.ChannelMessageSend(m.ChannelID, msg)
  }
}

func (h DiceRoll) Help() string {
  return "Roll a die. {sides}d{amount}+{extra rolls / constants}"
}

func (h DiceRoll) Prefixes() []string {
  return []string{"roll", "r"}
}
