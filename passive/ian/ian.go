package ian

import (
  "strings"

  "golords/handlers"

  "github.com/bwmarrin/discordgo"
)

/*
  TODO only enable this in my server.
*/

func init(){
  handlers.RegisterPassiveModule(
    Ian{},
  )
}

type Ian struct{}

func (h Ian) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  content := strings.ToLower(m.Content)

  buyWords := []string{
    "buy",
    "bought",
    "purchase",
    "get a new",
    "house",
    "pay",
    "spend",
    "buying",
    "irl money",
    "for me",
    "for us",
    "for eric",
    "gil",
    "csgo crate",
    "get me",
    "give me",
    "cash",
  }

  shouldTrigger := false

  for _, word := range buyWords {
    shouldTrigger = shouldTrigger || strings.Contains(content, word)
  }

  shouldTrigger = shouldTrigger && (strings.Contains(content, "ian") || strings.Contains(content, "ina"))

  if !shouldTrigger || strings.Contains(content, "don't") || strings.Contains(content, "dont") {
    return
  }

  s.ChannelMessageSend(m.ChannelID, "<@208773246009475072>, don't buy that thing!!!!!")
}

func (h Ian) Help() string {
  return "Stop Ian from spending all his money"
}
