package querystacks

import (
  "strings"
  "fmt"
  "log"

  "golords/state"
  "golords/handlers"

  "github.com/bwmarrin/discordgo"
)

func init(){
  handlers.RegisterActiveModule(
    Stack{},
  )
}

type Stack struct{}

func (h Stack) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 2)

  var out string
  var err error
  switch data[0] {
  case "!topstacks":
    out, err = state.TopQuery()
  // TODO !mystacks; a health workaround with !stacks @me works at the moment
  case "!stacks":
    out, err = state.PeopleQuery(m.Mentions)
  default:
    err = fmt.Errorf("Bad command")
  }

  // Mongo machine broke
  if err != nil {
    log.Printf("Error in query: %v", err)
    return
  }

  s.ChannelMessageSend(m.ChannelID, out)
}

func (h Stack) Help() string {
  return "Find out how many stacks someone has.\n!topstacks: top 5\n!stacks @someone: someone's stacks (can be you)"
}

func (h Stack) Prefixes() []string {
  return []string{"topstacks", "stacks"}
}
