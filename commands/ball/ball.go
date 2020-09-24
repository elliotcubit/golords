package ball

import (
	"github.com/bwmarrin/discordgo"
	"golords/handlers"
	"math/rand"
)

func init() {
	handlers.RegisterActiveModule(
		Ball{},
	)
}

type Ball struct{}

func (h Ball) Do(s *discordgo.Session, m *discordgo.MessageCreate) {
	responses := []string{
		"It is certain.",
		"It is decidedly so.",
		"Without a doubt.",
		"Yes -- deifinitely.",
		"You may rely on it.",
		"As I see it, yes.",
		"Most likely.",
		"Outlook good.",
		"Yes.",
		"Signs point to yes.",
		"Reply hazy, try again.",
		"Ask again later.",
		"Better not tell you now.",
		"Cannot predict now.",
		"Concentrate and ask again.",
		"Don't count on it.",
		"My reply is no.",
		"My sources say no.",
		"Outlook not so good.",
		"Very doubtful.",
	}
	msg := responses[rand.Intn(len(responses))]
	s.ChannelMessageSend(m.ChannelID, msg)
}

func (h Ball) Help() string {
	return "Ask the magic 8 ball a question"
}

func (h Ball) Prefixes() []string {
	return []string{"8ball"}
}
