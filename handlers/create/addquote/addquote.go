package addquote

import (
  "github.com/bwmarrin/discordgo"
  "strings"

  "golords/handlers/create/handler"
  "golords/quotemanager"
)

func New() handler.CreateHandler {
  return AddQuoteHandler{}
}

type AddQuoteHandler struct {
  handler.DefaultHandler
}

func (h AddQuoteHandler) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 2)
  if len(data) == 1 {
    return
  }
  quotemanager.AddQuote(m.Author.String(),
                        data[1],
                        string(m.Timestamp))
}

func (h AddQuoteHandler) GetPrompts() []string {
  return []string{"!addquote", "!aq"}
}

func (h AddQuoteHandler) Help() string {
  return "Adds a quote to the library"
}

func (h AddQuoteHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}
