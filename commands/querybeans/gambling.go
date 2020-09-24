package querybeans

import (
  "fmt"
  "math/rand"
  "strconv"

  "github.com/bwmarrin/discordgo"
  "golords/state"
)

var coinGames map[string]*CoinGame

type CoinGame struct {
  Offerer string
  Amount int
}

func init(){
  coinGames = make(map[string]string, 0)
}

func BetBeanHandler(s *discordgo.Session, m *discordgo.MessageCreate) string {
  data := strings.SplitN(m.Content, " ", 3)
  if len(m.Mentions < 1){
    return ""
  }
  amount, _ := strconv.Atoi(data[1])
  offerer := m.Author.String()
  recipient := m.Mentions[0].String()
  // If we're betting zero or the second arg wasn't a number
  // Then look for an existing game
  if amount == 0 {
    _, ok := coinGames[offerer]
    if !ok {
      return "That person did not offer you a game"
    }
    return executeCoinGame(s, m, offerer)
  } else {
    game := &CoinGame{
      Offerer: offerer,
      Amount: amount,
    }
    coinGames[recipient] = game
    return fmt.Sprintf("Coin game betting %d created", amount)
  }
}

// Only called if key is guaranteed to be in map
func executeCoinGame(s *discordgo.Session, m *discordgo.MessageCreate, key string) string {

}
