package querybeans

import (
  "fmt"
  "log"
  "strconv"
  "strings"

  "golords/state"
  "github.com/bwmarrin/discordgo"
)

func sendBeanLeaderboard(s *discordgo.Session, m *discordgo.MessageCreate, ascending bool){
  data := strings.SplitN(m.Content, " ", 2)
  var err error
  amount := 5
  if len(data) > 1 {
    amount, err = strconv.Atoi(data[1])
    if err != nil {
      amount = 5
    }
  }
  var results []*state.BeanData
  if ascending {
    results, err = state.GetTopNBeans(m.GuildID, amount)
    if err != nil {
      log.Println(err)
      return
    }
  } else {
    results, err = state.GetBottomNBeans(m.GuildID, amount)
    if err != nil {
      log.Println(err)
      return
    }
  }
  out := ""
  for ranking, data := range results {
    out += fmt.Sprintf("%d | %-32s %8d beans\n", ranking, data.User, data.Amount)
  }

  embed := &discordgo.MessageEmbed{Color:0x3498DB}
  embed.Title = "Most Beans Leaderboard"
  if !ascending {
    embed.Title = "Least Beans Leaderboard"
  }
  embed.Description = out
  s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
