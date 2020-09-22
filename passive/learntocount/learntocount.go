package learntocount

import (
  "fmt"
  "log"
  "strconv"
  "strings"

  "github.com/bwmarrin/discordgo"

  "golords/state"
  "golords/handlers"
)

func init(){
  handlers.RegisterPassiveModule(
    LearnToCount{},
  )
}

const ltcChan = "755266387014058145"
const NOT_SET =   -99999999
const NEVER_SET = -99999998
var previous int = NEVER_SET

type LearnToCount struct{}

func (h LearnToCount) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  // Only operate in the learn to count channel
  if m.ChannelID !=  ltcChan {
    return
  }

  // Attempt to get a number from the current message
  i, err := strconv.Atoi(strings.TrimSpace(m.Content))

  if err != nil {
    // Nothing wrong with this, but no change in previous
    if previous == NOT_SET || previous == NEVER_SET {
      return
      // A mistake was made if previous is set, but the message isn't a number
    } else {
      punish(s, m, -(previous*(previous+1)/2))
      previous = NOT_SET
    }
    // The message IS a number
  } else {
    // Blindly trust the first message we see
    // No points for blind trust message
    if previous == NEVER_SET {
      previous = i
      return
    }
    // If we just reset, the new number must be 1
    if previous == NOT_SET {
      if i != 1 {
        return
      } else {
        previous = i
        reward(s, m, previous)
      }
      // If we didn't it must be previous + 1
    } else {
      // A mistake
      if i != (previous + 1){
        punish(s, m, -(previous*(previous+1)/2))
        previous = NOT_SET
      } else {
        previous = i
        reward(s, m, previous)
      }
    }
  }
}

func (h LearnToCount) Help() string {
  return "Why is it so hard to learn to count?"
}

func reward(s *discordgo.Session, m *discordgo.MessageCreate, amount int){
  if m.Author == nil {
    log.Println("Something bad happened rewarded beancounters")
    return
  }
  // No message for rewards
  state.UpdateBeans(m.GuildID, m.Author.String(), amount)
}

func punish(s *discordgo.Session, m *discordgo.MessageCreate, amount int){
  if m.Author == nil {
    log.Println("An error was made in counting, but the message doesn't seem to have an author to punish")
    return
  }
  state.UpdateBeans(m.GuildID, m.Author.String(), amount)
  pingString := fmt.Sprintf("<@%s> fucked up. -1 stacks.", m.Author.ID)
  s.ChannelMessageSend(m.ChannelID, pingString)
}
