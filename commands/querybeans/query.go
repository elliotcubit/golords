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

var bannedRecipients []string = []string{
  "235088799074484224", // Rythmbot
}

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
  case "givebeans":
    data = strings.SplitN(m.Content, " ", 3)
    // !givebeans [amount] [user] ...
    if len(data) < 3 {
      return
    }
    // ping a person to give them bean
    if len(m.Mentions) < 1 {
      return
    }
    amount := 0
    amount, err = strconv.Atoi(data[1])
    if err != nil {
      return
    }
    if amount <= 0 {
      out += "That's kind of scummy"
      break
    }
    recipient := m.Mentions[0]
    recipientID := recipient.String()
    donatorID := m.Author.String()

    // Can we name loops/switches in go to prevent this?
    stp := false
    for _, bannedID := range bannedRecipients {
      if recipient.ID == bannedID {
        out += "You can't do that"
        stp = true
      }
    }
    if stp {
      break
    }

    // Verify they have the necessary funds
    donatorBalance, err := state.GetBeansForUser(m.GuildID, donatorID)
    if err != nil {
      out += "Bean transfer failed...!\n"
      break
    }
    if donatorBalance < amount {
      out += fmt.Sprintf("Everyone point and laugh - <@%s> doesn't have enough money!", m.Author.ID)
      break
    }
    // This should not have been split into it's own function but it's too late to go back
    out += h.TransferBeans(donatorID, recipientID, m.GuildID, donatorBalance, amount)
  case "mybeans":
    user := m.Author.String()
    amount, err := state.GetBeansForUser(m.GuildID, user)
    if err != nil {
      log.Println(err)
      return
    }
    out += fmt.Sprintf("%s: %d beans\n", user, amount)
  case "topbeans":
    amount := 5
    if len(data) > 1 {
      amount, err = strconv.Atoi(data[1])
    }
    results, err := state.GetTopNBeans(m.GuildID, amount)
    if err != nil {
      log.Println(err)
      return
    }

    for _, data := range results {
      out += fmt.Sprintf("%v: %d beans\n", data.User, data.Amount)
    }
  case "bottombeans":
    amount := 5
    if len(data) > 1 {
      amount, err = strconv.Atoi(data[1])
    }
    results, err := state.GetBottomNBeans(m.GuildID, amount)
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

func (h Bean) TransferBeans(don, rec, gid string, donb, amount int) string {
  // Take money away from donator
  donb, err1 := state.UpdateBeans(gid, don, -amount)
  // If this fails just stop
  if err1 != nil {
    return "Bean transfer failed...!\n"
  }
  // Give money to recipient
  recb, err2 := state.UpdateBeans(gid, rec, amount)
  // This isn't good... try again and hope for the best?
  if err2 != nil {
    recb, err2 = state.UpdateBeans(gid, rec, amount)
    // if we failed a second time try to give back the money
    // but don't make any promises
    if err2 != nil {
      donb, err1 = state.UpdateBeans(gid, don, amount)
      return "Bean transfer failed...!\n"
    }
  }
  s := ""
  s += "Bean transfer successfuly.\n"
  s += fmt.Sprintf("%s now has %d beans, while\n", don, donb)
  s += fmt.Sprintf("%s now has %d beans.\n", rec, recb)
  return s
}

func (h Bean) Help() string {
  return "Find out how many beans someone has.\n!topbeans: top 5\n!beans @someone: someone's beans (can be you)"
}

func (h Bean) Prefixes() []string {
  return []string{
    "topbeans",
    "bottombeans",
    "beans",
    "mybeans",
    "givebeans",
  }
}
