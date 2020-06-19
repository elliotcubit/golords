package getquote

import (
  "golords/handlers/create/handler"
  "github.com/bwmarrin/discordgo"
  "golords/quotemanager"
  "fmt"
  "time"
  "strings"
)

func New() handler.CreateHandler {
  return GetQuoteHandler{}
}

type GetQuoteHandler struct {
  handler.DefaultHandler
}

func (h GetQuoteHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  q := quotemanager.GetRandomQuote()
  t, _ := time.Parse(time.RFC3339, q.Timestamp) // TODO error checking
  ts := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
  msg := fmt.Sprintf("\"%v\" added by %v on %v", q.Text, q.AddedBy, ts)
  s.ChannelMessageSend(m.ChannelID, msg)
}

func (h GetQuoteHandler) GetPrompts() []string {
  return []string{"!getquote", "!gq"}
}

func (h GetQuoteHandler) Help() string {
  return "Get a quote from the library"
}

func (h GetQuoteHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}
