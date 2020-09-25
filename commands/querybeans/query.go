package querybeans

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golords/handlers"
	"golords/state"

	"github.com/bwmarrin/discordgo"
)

var bannedRecipients []string = []string{
	"235088799074484224", // Rythmbot
}

var THE_LOTTERY string = "GolordsBot#2566"

func init() {
	handlers.RegisterActiveModule(
		Bean{},
	)
}

type Bean struct{}

func (h Bean) Do(s *discordgo.Session, m *discordgo.MessageCreate) {
	data := strings.SplitN(m.Content, " ", 2)

	var out string
	var err error
	switch data[0] {
	case "beanbet":
		out += BetBeanHandler(s, m)
	case "beanrisk":
		if len(data) < 2 {
			out += "usage: !beanrisk [currentNumber]"
			break
		}
		currentNumber, err := strconv.Atoi(data[1])
		if err != nil {
			break
		}
		risk := (currentNumber*(currentNumber+1))/2
		out += fmt.Sprintf("The current amount you will lose for making a fucky wucky is: %d", risk)
	// Starts a bean lottery whose options will execute 30 minutes later
	case "startbeanlottery":
		data = strings.SplitN(m.Content, " ", 2)
		timer := 5
		if len(data) > 1 {
			// I think Atoi will set timer to 0 on fails
			timer, err = strconv.Atoi(data[1])
			if err != nil {
				timer = 5
			}
		}
		out += h.StartBeanLottery(s, m, timer)
	case "buybeanticket":
		out += h.EnterBeanLottery(s, m)
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
		sendBeanLeaderboard(s, m, true)
	case "bottombeans":
		sendBeanLeaderboard(s, m, false)
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
	donb, err1 := state.AddBeans(gid, don, -amount)
	// If this fails just stop
	if err1 != nil {
		return "Bean transfer failed...!\n"
	}
	// Give money to recipient
	recb, err2 := state.AddBeans(gid, rec, amount)
	// This isn't good... try again and hope for the best?
	if err2 != nil {
		recb, err2 = state.AddBeans(gid, rec, amount)
		// if we failed a second time try to give back the money
		// but don't make any promises
		if err2 != nil {
			donb, err1 = state.AddBeans(gid, don, amount)
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
		"startbeanlottery",
		"buybeanticket",
		"topbeans",
		"bottombeans",
		"beans",
		"mybeans",
		"givebeans",
		"beanrisk",
		"beanbet",
	}
}
