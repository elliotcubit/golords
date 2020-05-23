package create

import (
  "github.com/bwmarrin/discordgo"
  "strings"
  "golords/quotemanager"
)

func HandleCreateAddQuote(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m.Content, " ", 2)
  if len(data) == 1 {
    return
  }
  quotemanager.AddQuote(m.Author.String(),
                        data[1],
                        string(m.Timestamp))
}
