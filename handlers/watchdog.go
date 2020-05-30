package handlers

import (
  "strings"
  "fmt"
  "bufio"
  "os"
  "time"

  "github.com/bwmarrin/discordgo"
)

var banWords []string

func LoadBanWords(fname string) error{
  file, err := os.Open(fname)
  if err != nil {
    return err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    banWords = append(banWords, scanner.Text())
  }
  return nil
}

// Returns a banned word in the string or 0 if none
func containsBanWord(query string) string {
  for _, s := range banWords {
    if strings.Contains(strings.ToLower(query), s) {
      return s
    }
  }
  return ""
}

func OnMessageUpdate(s *discordgo.Session, mup *discordgo.MessageUpdate) {
  // Ignore ourself, and do nothing if previous message wasn't cached
  // MessageUpdate.Author is nil when requests made via webhook.
  if mup.Author == nil || mup.Author.ID == s.State.User.ID || mup.BeforeUpdate == nil {
    return
  }

  // Then message does not have a banned word anymore
  wordHidden := containsBanWord(mup.BeforeUpdate.Content)
  if wordHidden != "" && containsBanWord(mup.Content) == "" {
    s.GuildMemberNickname(mup.GuildID, "@me", "CancelBot")

    punish := fmt.Sprintf(
      "Nice try, but you can't hide that from me. User %v said \"%v\" in a message and tried to edit it out... tsk tsk\n",
      mup.Author.Mention(),
      wordHidden,
    )

    punish2 := fmt.Sprintf(
      "Their original message said \"%v\"\n",
      mup.BeforeUpdate.Content,
    )

    s.ChannelMessageSend(mup.ChannelID, punish+punish2)

    go func(){
      fmt.Println("Waiting to fix our nickname")
      time.Sleep(10 * time.Second)
      s.GuildMemberNickname(mup.GuildID, "@me", "GolordBot")
      fmt.Println("Should be safe to kill the bot. :)")
    }()
  }
}

// TODO MessageDelete only sends an ID of the message seemingly
func OnMessageDelete(s *discordgo.Session, m *discordgo.MessageDelete){
  // Ignore ourself
  if m.Author == nil || m.Author.ID == s.State.User.ID {
    return
  }

  wordHidden := containsBanWord(m.Content)

  if wordHidden != "" {
    s.GuildMemberNickname(m.GuildID, "@me", "CancelBot")

    punish := fmt.Sprintf(
      "Nice try, but you can't hide that from me. User %v said \"%v\" in a message and tried to delete it... tsk tsk\n",
      m.Author.Mention(),
      wordHidden,
    )

    punish2 := fmt.Sprintf(
      "Their original message said \"%v\"\n",
      m.Content,
    )

    s.ChannelMessageSend(m.ChannelID, punish+punish2)

    go func(){
      fmt.Println("Waiting to fix our nickname")
      time.Sleep(10 * time.Second)
      s.GuildMemberNickname(m.GuildID, "@me", "GolordBot")
      fmt.Println("Should be safe to kill the bot. :)")
    }()
  }
}
