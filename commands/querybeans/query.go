package querybeans

import (
  "strings"
  "fmt"
  "log"
  "strconv"

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
  case "mybeans":
    user := m.Author.String()
    amount, err := state.GetBeansForUser(m.GuildID, user)
    if err != nil {
      log.Println(err)
      return
    }
    out += fmt.Sprintf("%s: %d beans\n", user, amount)
  case "topbeans":
    amount, err := strconv.Atoi(data[1])
    if err != nil {
      amount = 5
    }
    results, err := state.GetTopNBeans(m.GuildID, amount)
    if err != nil {
      log.Println(err)
      return
    }

    for _, data := range results {
      out += fmt.Sprintf("%v: %d beans\n", data.User, data.Amount)
    }
  case "beans":
    // TODO this only really needs one query
    // SELECT * FROM _ WHERE --- OR --- OR --- OR --- OR
    for _, user := range m.Mentions {
      amount, err := state.GetBeansForUser(m.GuildID, user.String())
      if err != nil {
        log.Println(err)
        return
      }
      out += fmt.Sprintf("%s: %d beans\n", user.String(), amount)
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
    log.Println("No output for bean query")
    return
  }

  s.ChannelMessageSend(m.ChannelID, out)
}

func (h Bean) Help() string {
  return "Find out how many beans someone has.\n!topbeans: top 5\n!beans @someone: someone's beans (can be you)"
}

func (h Bean) Prefixes() []string {
  return []string{"topbeans", "beans", "mybeans"}
}
