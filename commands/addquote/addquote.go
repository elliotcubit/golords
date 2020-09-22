package addquote

import (
  "golords/handlers"
  "github.com/bwmarrin/discordgo"

  "strings"
  "golords/state"
)

func init(){
  handlers.RegisterActiveModule(
    AddQuote{},
  )
}

type AddQuote struct {}

func (h AddQuote) Do(s *discordgo.Session, m *discordgo.MessageCreate){
    data := strings.SplitN(m.Content, " ", 2)
    if len(data) == 1 {
      return
    }
    state.AddQuote(
      m.GuildID,
      m.Author.String(),
      data[1],
      string(m.Timestamp),
    )
}

func (h AddQuote) Help() string {
  return "Add quote to the server's library"
}

func (h AddQuote) Prefixes() []string {
  return []string{"addquote", "aq"}
}
