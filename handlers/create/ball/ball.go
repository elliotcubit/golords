package ball

import (
  "math/rand"

  "golords/handlers/create/handler"
  
  "github.com/bwmarrin/discordgo"
)

func getResponses() []string {
    return []string{
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
}

func New() handler.CreateHandler {
  return BallHandler{}
}

type BallHandler struct {
  handler.DefaultHandler
}

func (h BallHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  responses := getResponses()
  msg := responses[rand.Intn(len(responses))]
  s.ChannelMessageSend(m.ChannelID, msg)
}

func (h BallHandler) GetPrompts() []string {
  return []string{"!8ball"}
}

func (h BallHandler) Help() string {
  return "Ask magic 8 ball a question"
}

func (h BallHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}
