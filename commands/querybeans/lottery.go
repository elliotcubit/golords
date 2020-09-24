package querybeans

import (
  "fmt"
  "log"
  "math/rand"
  "time"

  "golords/state"

  "github.com/bwmarrin/discordgo"
)

const ticketPrice = 500

// runningLotteries[serverID] ~> Lottery{ServerID: serverID}
var runningLotteries map[string]*Lottery

func init(){
  runningLotteries = make(map[string]*Lottery, 0)
}

type Lottery struct {
  ServerID string
  ChannelID string
  Tickets []*LotteryTicket
}

type LotteryTicket struct {
  UserID string
}

// Start a goroutine that waits 30 mins
func (h Bean) StartBeanLottery(s *discordgo.Session, m *discordgo.MessageCreate) string{
  serverID := m.GuildID
  if _, exists := runningLotteries[serverID]; exists {
    return "There is already a lottery running in this server.\nBuy a ticket with !buybeanticket"
  }
  newLottery := &Lottery{
    ServerID: serverID,
    ChannelID: m.ChannelID,
    Tickets: []*LotteryTicket{},
  }
  // Mark ourselves as existing then wait
  runningLotteries[serverID] = newLottery
  go newLottery.AwaitBeanLottery(s, m)
  return "Lottery has been started, a winner will be chosen in 30 minutes.\nYou can buy a ticket with !buybeanticket"
}

// Adds a new ticket to the bean lottery
// after docking the entrant the price of the lottery
func (h Bean) EnterBeanLottery(s *discordgo.Session, m *discordgo.MessageCreate) string {
  serverID := m.GuildID
  lottery, exists := runningLotteries[serverID]
  if !exists {
    return "There is no lottery running in this server currently.\n Start one with !startbeanlottery"
  }
  ticket := &LotteryTicket{
    UserID: m.Author.String(),
  }
  // Do we want to limit people to one ticket?
  if lottery.InLottery(ticket) {
    return "You are already in the lottery."
  }
  balance, err := state.GetBeansForUser(serverID, ticket.UserID)
  if err != nil {
    return "There was a problem entering you into the lottery"
  }
  if balance < ticketPrice {
    return fmt.Sprintf("Sorry, but the lottery costs %d beans", ticketPrice)
  }
  _, err = state.UpdateBeans(serverID, ticket.UserID, balance-ticketPrice)
  if err != nil {
    return "There was a problem entering you into the lottery"
  }
  lottery.MakeEntry(ticket)
  return "You have been entered into the bean lottery."
}

func (l Lottery) InLottery(ticket *LotteryTicket) bool {
  for _, entry := range l.Tickets {
    if entry.UserID == ticket.UserID {
      return true
    }
  }
  return false
}

func (l Lottery) MakeEntry(ticket *LotteryTicket){
  l.Tickets = append(l.Tickets, ticket)
}

func (l Lottery) AwaitBeanLottery(s *discordgo.Session, m *discordgo.MessageCreate){
  time.Sleep(30 * time.Minute)
  l.ExecuteBeanLottery(s, m)
}

// Execute the results of a lottery.
// If any DB call in this fails we're forced to keep retrying
// Or the lottery won't execute and the state will be unrecoverable.
// This is okay, because this function will only exist in a separate goroutine
func (l Lottery) ExecuteBeanLottery(s *discordgo.Session, m *discordgo.MessageCreate){
  // Select a random ticket
  index := rand.Intn(len(l.Tickets))
  randomTicket := l.Tickets[index]

  // Get THE_LOTTERY's current balance
  lotteryBalance, err := state.GetBeansForUser(l.ServerID, THE_LOTTERY)
  for ; err != nil ; {
    log.Println("Couldn't get lottery's balance. Retrying in 30sec")
    time.Sleep(30 * time.Second)
    lotteryBalance, err = state.GetBeansForUser(l.ServerID, THE_LOTTERY)
  }

  // Pay the winner
  winnerBalance, err := state.GetBeansForUser(l.ServerID, randomTicket.UserID)
  for ; err != nil ; {
    log.Println("Couldn't get lottery winner's balance. Retrying in 30sec")
    time.Sleep(30 * time.Second)
    winnerBalance, err = state.GetBeansForUser(l.ServerID, randomTicket.UserID)
  }
  winnerBalance, err = state.UpdateBeans(l.ServerID, randomTicket.UserID, winnerBalance+lotteryBalance)
  for ; err != nil ; {
    log.Println("Couldn't update lottery winner's balance. Retrying in 30sec")
    time.Sleep(30 * time.Second)
    winnerBalance, err = state.UpdateBeans(l.ServerID, randomTicket.UserID, winnerBalance+lotteryBalance)
  }

  // Set THE_LOTTERY's balance to zero
  _, err = state.UpdateBeans(l.ServerID, THE_LOTTERY, 0)
  for ; err != nil ; {
    log.Println("Couldn't set THE_LOTTERY's balance to zero. Retrying in 30sec")
    time.Sleep(30 * time.Second)
    _, err = state.UpdateBeans(l.ServerID, THE_LOTTERY, 0)
  }

  // Remove ourself from the parent lottery map so a new lottery could start
  delete(runningLotteries, l.ServerID)

  // Send a message declaring the winner
  s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v has won %d beans in the lottery! Congratulations!", randomTicket.UserID, lotteryBalance))
}
