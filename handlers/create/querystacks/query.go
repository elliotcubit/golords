package querystacks

import (
  "strings"
  "fmt"
  "log"

  pp "golords/plusplus"
  "golords/handlers/create/handler"

  "github.com/bwmarrin/discordgo"
)

func New() handler.CreateHandler {
  return QueryHandler{}
}

type QueryHandler struct {
  handler.DefaultHandler
}

func (h QueryHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 2)

  var out string
  var err error
  switch data[0] {
  case "!topstacks":
    out, err = pp.TopQuery()
  case "!mystacks":
    fuckYou := make([]*discordgo.User, 1)
    fuckYou = append(fuckYou, m.Author)
    out, err = pp.PeopleQuery(fuckYou)
  case "!stacks":
    out, err = pp.PeopleQuery(m.Mentions)
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

func (h QueryHandler) GetPrompts() []string {
  return []string{"!topstacks", "!mystacks", "!stacks"}
}

func (h QueryHandler) Help() string {
  return "Find out how many stacks someone has.\n!topstacks: top 5\n!mystacks: your stacks\n!stacks @someone: someone elses stacks"
}

func (h QueryHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}
