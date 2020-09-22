package querybeans

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
    Bean{},
  )
}

type Bean struct{}

func (h Bean) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 2)

  var out string
  var err error
  switch data[0] {
  case "topbeans":
    // TODO specify how many with command
    results, err := state.GetTopNBeans(m.GuildID, 5)
    if err != nil {
      log.Println(err)
      return
    }
    for user, amount := range results {
      out += fmt.Sprintf("%v: %d stacks\n", user, amount)
    }
  // TODO !mystacks; a health workaround with !stacks @me works at the moment
  case "beans":
    // TODO this only really needs one query
    // SELECT * FROM _ WHERE --- OR --- OR --- OR --- OR
    for _, user := range m.Mentions {
      amount, err := state.GetBeansForUser(m.GuildID, user.String())
      if err != nil {
        log.Println(err)
        return
      }
      out += fmt.Sprintf("%v: %d stacks\n", user.String(), amount)
    }
  default:
    err = fmt.Errorf("Bad command: %v", data[0])
  }

  // Mongo machine broke
  if err != nil {
    log.Printf("Error in query: %v", err)
    return
  }

  if out == "" {
    log.Println("No output for stack query")
    return
  }

  s.ChannelMessageSend(m.ChannelID, out)
}

func (h Bean) Help() string {
  return "Find out how many beans someone has.\n!topbeans: top 5\n!beans @someone: someone's beans (can be you)"
}

func (h Bean) Prefixes() []string {
  return []string{"topbeans", "beans"}
}
