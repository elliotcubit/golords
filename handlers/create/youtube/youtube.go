package youtube

import (
  "golords/handlers/create/handler"

  "github.com/jonas747/dca"
  "github.com/rylio/ytdl"
  "github.com/bwmarrin/discordgo"
)

func New() handler.CreateHandler {
  return YoutubeManager{}
}

func (y YoutubeManager) Do(s *discordgo.Session, m *discordgo.MessageCreate){
  data := strings.SplitN(m, " ", 3)
  switch data[1] {
  case "play":
    y.enqueue(data[2])
  case "pause":
    y.pause()
  case "start":
    y.unpause()
  case "skip":
    y.skip()
  case "stop":
    y.stop()
  case "queue":
    y.queue()
  case "now":
    y.playNow(data[2])
  }
}

// Will be like
// !yt play
// !yt skip
// ...
func (y YoutubeHandler) GetPrompts() []string {
  return []string{
    "!yt",
  }
}

func (y YoutubeHandler) Help() string {
  return "RhythmBot Clone"
}

func (y YoutubeHandler) Should(hint string) bool {
  prompts := h.GetPrompts()
  for _, v := range prompts {
    if strings.HasPrefix(hint, v) {
      return true
    }
  }
  return false
}
