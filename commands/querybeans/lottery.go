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

func init() {
	runningLotteries = make(map[string]*Lottery, 0)
}

type Lottery struct {
	ServerID  string
	ChannelID string
	Tickets   []string
}

// Start a goroutine that waits `timer` minutes
func (h Bean) StartBeanLottery(s *discordgo.Session, m *discordgo.MessageCreate, timer int) string {
	serverID := m.GuildID
	if _, exists := runningLotteries[serverID]; exists {
		return "There is already a lottery running in this server.\nBuy a ticket with !buybeanticket"
	}
	newLottery := &Lottery{
		ServerID:  serverID,
		ChannelID: m.ChannelID,
		Tickets:   []string{},
	}
	// Mark ourselves as existing then wait
	runningLotteries[serverID] = newLottery
	log.Printf("Started a lottery in server %s", serverID)
	go newLottery.AwaitBeanLottery(s, m, timer)
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
	// Do we want to limit people to one ticket?
	if lottery.InLottery(m.Author.String()) {
		return "You are already in the lottery."
	}
	// This could technically fail by producing 500 extra beans
	// But I'll take that risk B)
	_, err := state.AddBeans(serverID, THE_LOTTERY, ticketPrice)
	if err != nil {
		return "There was a problem entering you into the lottery"
	}
	_, err = state.AddBeans(serverID, m.Author.String(), -ticketPrice)
	if err != nil {
		return "There was a problem entering you into the lottery"
	}
	lottery.MakeEntry(m.Author.String())
	return "You have been entered into the bean lottery."
}

func (l *Lottery) InLottery(entrant string) bool {
	for _, entry := range l.Tickets {
		if entry == entrant {
			return true
		}
	}
	return false
}

func (l *Lottery) MakeEntry(entrant string) {
	l.Tickets = append(l.Tickets, entrant)
}

func (l *Lottery) AwaitBeanLottery(s *discordgo.Session, m *discordgo.MessageCreate, timer int) {
	time.Sleep(time.Duration(timer) * time.Minute)
	l.ExecuteBeanLottery(s, m)
}

// Execute the results of a lottery.
// If any DB call in this fails we're forced to keep retrying
// Or the lottery won't execute and the state will be unrecoverable.
// This is okay, because this function will only exist in a separate goroutine
func (l *Lottery) ExecuteBeanLottery(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Select a random ticket
	index := rand.Intn(len(l.Tickets))
	randomTicket := l.Tickets[index]

	// Get THE_LOTTERY's current balance
	lotteryBalance, err := state.GetBeansForUser(l.ServerID, THE_LOTTERY)
	for err != nil {
		time.Sleep(30 * time.Second)
		lotteryBalance, err = state.GetBeansForUser(l.ServerID, THE_LOTTERY)
	}

	// Add that balance to the winner
	_, err = state.AddBeans(l.ServerID, randomTicket, lotteryBalance)
	for err != nil {
		time.Sleep(30 * time.Second)
		_, err = state.AddBeans(l.ServerID, randomTicket, lotteryBalance)
	}

	// Set THE_LOTTERY's balance to zero
	_, err = state.AddBeans(l.ServerID, THE_LOTTERY, -lotteryBalance)
	for err != nil {
		time.Sleep(30 * time.Second)
		_, err = state.AddBeans(l.ServerID, THE_LOTTERY, -lotteryBalance)
	}

	// Remove ourself from the parent lottery map so a new lottery could start
	delete(runningLotteries, l.ServerID)

	// Send a message declaring the winner
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v has won %d beans in the lottery! Congratulations!", randomTicket, lotteryBalance))
}
