package anim

import (
  "strings"
  "fmt"
  "time"

  "golords/handlers/create/handler"
  "github.com/bwmarrin/discordgo"
)

const chars = 5

func New() handler.CreateHandler {
  return AnimHandler{}
}

type AnimHandler struct {
  handler.DefaultHandler
}

func (h AnimHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 2)
  if len(data) < 2 {
    return
  }


  // Do everything in a goroutine so we can move on and
  // only pause that thread
  go func() {
    chanID := m.ChannelID
    msg, _ := s.ChannelMessageSend(chanID, data[1])
    msgID := msg.ID

    for iter := range getCycle(data[1]) {
        s.ChannelMessageEdit(chanID, msgID, iter)
        time.Sleep(500 * time.Millisecond)
    }

    s.ChannelMessageEdit(chanID, msgID, data[1])
  }()

}

func (h AnimHandler) GetPrompts() []string {
  return []string{"!anim"}
}

func (h AnimHandler) Help() string {
  return "Pretty text :)"
}

func (h AnimHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}

func cycle(channel chan string, seed string) {
  for i := 0; i<len(seed); i++ {
    numLeft := 0
    numRight := chars
    if len(seed) < i + chars {
      numRight = len(seed) - i
      numLeft = chars - numRight
    }

    left := seed[0:numLeft]
    right := seed[i:i+numRight]

    formatString := fmt.Sprintf("`%%s%%%ds`", i - numLeft + numRight)

    channel <- fmt.Sprintf(formatString, left, right)

  }

}

func getCycle(seed string) chan string {
  channel := make(chan string)
  go func() {
    cycle(channel, seed)
    close(channel)
  }()
  return channel
}
