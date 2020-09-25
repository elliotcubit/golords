package querybeans

import (
  "fmt"
  "log"
  "math/rand"
  "strconv"
  "strings"
  "time"

  "github.com/bwmarrin/discordgo"
  "golords/state"
)

var coinGames []*CoinGame

type CoinGame struct {
  ServerID string
  Challenger string
  Challengee string
  Amount int
}

func BetBeanHandler(s *discordgo.Session, m *discordgo.MessageCreate) string {
  data := strings.SplitN(m.Content, " ", 3)
  if len(m.Mentions) < 1{
    return "You must @ someone"
  }
  serverID := m.GuildID
  amount, _ := strconv.Atoi(data[1])
  challenger := m.Author.String()
  challengee := m.Mentions[0].String()

  // Check if any challenge applies to us already
  // No specified amount means accept any offer
  for thisGameIndex, game := range coinGames {
    if game.ServerID == serverID &&
    game.Challenger == challengee &&
    game.Challengee == challenger {
      if amount == 0 || game.Amount == amount {
        return executeCoinGame(thisGameIndex)
      }
    }
  }

  // Otherwise, create a new challenge.
  // Do not create a challenge of amount 0
  if amount <= 0 {
    return "You cannot challenge for <=0 beans"
  }

  // Verify challenger has enough beans
  challengersBalance, err := state.GetBeansForUser(serverID, challenger)
  if err != nil {
    return "There was a problem creating your challenge"
  }
  if challengersBalance < amount {
    return "You do not have enough beans to make that bet"
  }

  // Verify challengee has enough beans
  challengeesBalance, err := state.GetBeansForUser(serverID, challengee)
  if err != nil {
    return "There was a problem creating your challenge"
  }
  if challengeesBalance < amount {
    return "The person you're challenging doesn't have enough beans to make that bet"
  }

  // Verify the game doesn't already exist
  for _, game := range coinGames {
    if game.ServerID == serverID &&
    game.Challenger == challengee &&
    game.Challengee == challenger &&
    game.Amount == amount {
      return "You have already made that challenge and it was not unaccepted"
    }
  }

  game := &CoinGame{
    ServerID: serverID,
    Challenger: challenger,
    Challengee: challengee,
    Amount: amount,
  }
  coinGames = append(coinGames, game)
  return fmt.Sprintf("Challenge created for %d beans. Accept by challenging back", amount)
}

func executeCoinGame(ind int) string {
  game := coinGames[ind]

  // Select winner
  choice := rand.Intn(2)
  var winner string
  var loser string
  if choice == 0 {
    winner = game.Challenger
    loser = game.Challengee
  } else {
    winner = game.Challengee
    loser = game.Challenger
  }

  // Transfer funds
  _, err := state.AddBeans(game.ServerID, winner, game.Amount)
  if err != nil {
    return "There was a problem finishing the challenge"
  }
  // We have to finish if this happens
  // TODO move to goroutine to prevent bot hanging on this if something bad happens
  _, err = state.AddBeans(game.ServerID, loser, -game.Amount)
  for ; err != nil ; {
    log.Println("Problem resolving coin game, retrying in 30s")
    time.Sleep(30 * time.Second)
    _, err = state.AddBeans(game.ServerID, loser, -game.Amount)
  }

  // Remove the challenge from the list
  coinGames[ind] = coinGames[len(coinGames)-1]
  coinGames[len(coinGames)-1] = nil
  coinGames = coinGames[:len(coinGames)-1]

  return fmt.Sprintf("%s won the bet between %s and %s for %d beans", winner, winner, loser, game.Amount)
}
