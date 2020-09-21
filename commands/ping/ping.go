package ping

import (
  "golords/handlers"
  "github.com/bwmarrin/discordgo"
)

func init() {
  handlers.RegisterActiveModule(
    Ping{},
  )
}

type Ping struct {}

func (h Ping) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  s.ChannelMessageSend(m.ChannelID, "pong!")
}

func (h Ping) Help() string {
  return "Check if the bot is running"
}

func (h Ping) Prefixes() []string {
  return []string{"ping"}
}
