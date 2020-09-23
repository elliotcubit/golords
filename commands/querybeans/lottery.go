package querybeans

import (
  "fmt"
  "math/rand"
  "time"

  "github.com/bwmarrin/discordgo"
)

const ticketPrice = 500

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
    return
  }
  newLottery := &Lottery{
    ServerID: serverID,
    ChannelID: m.ChannelID,
    Tickets: []*LotteryTicket{},
  }
  // Mark ourselves as existing then wait
  runningLotteries[serverID] = newLottery
  go newLottery.AwaitBeanLottery()
}

// Adds a new ticket to the bean lottery
// after docking the entrant the price of the lottery
func (h Bean) EnterBeanLottery(s *discordgo.Session, m *discordgo.MessageCreate) string {
  serverID := m.GuildID
  if lottery, exists := runningLotteries[serverID]; !exists {
    return "There is no lottery running in this server currently.\n Start one with !startbeanlottery"
  }
  ticket := &LotteryTicket{
    UserID: m.Author,
  }
  // Do we want to limit people to one ticket?
  if lottery.InLottery(ticket) {
    return "You are already in the lottery."
  }
  balance, err := state.GetBeansForUser(serverID, ticket.UserID)
  if err != nil {
    return "There was a problem entering you into the lottery"
  }
  if balance < lotteryPrice {
    return fmt.Sprintf("Sorry, but the lottery costs %d beans", lotterPrice)
  }
  newBalance, err := state.UpdateBeans(serverID, ticket.UserID, balance-lotteryPrice)
  if err != nil {
    return "There was a problem entering you into the lottery"
  }
  l.MakeEntry(ticket)
}

func (l Lottery) InLottery(ticket *LotteryTicket){
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

func (l Lottery) AwaitBeanLottery(){
  time.Sleep(30 * time.Minute)
  l.ExecuteBeanLottery()
}

func (l Lottery) ExecuteBeanLottery(){
  // Select a random ticket
  index := rand.Intn(len(l.Tickets))
  randomTicket := l.Tickets[index]
  // Pay it whatever THE_LOTTERY's current balance is
  lotteryBalance, err := state.GetBeansForUser(l.ServerID, THE_LOTTERY)
  if err != nil {
    lotteryBalance, err = state.GetBeansForUser(l.ServerID, THE_LOTTERY)
  }

  // Set THE_LOTTERY's balance to zero

  // Remove ourself from the parent lottery map
}
