package create

import (
  "math/rand"

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

func HandleEightBall(s *discordgo.Session, m *discordgo.MessageCreate){
  responses := getResponses()
  msg := responses[rand.Intn(len(responses))]
  s.ChannelMessageSend(m.ChannelID, msg)
}
