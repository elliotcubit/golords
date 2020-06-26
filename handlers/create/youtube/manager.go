package youtube

import (
  "time"
  "sync"
  "errors"
  "log"

  "golords/handlers/create/handler"

  "github.com/Strum355/go-queue/queue"
  "github.com/bwmarrin/discordgo"
)

type YoutubeManager interface {
  handler.DefaultHandler
  VC Voice
}

type Voice struct {
  ChannelID string

  Playing bool

  Done chan error

  *sync.RWMutex

  Queue *queue.Queue
  VoiceCon *discordgo.VoiceConnection
  StreamingSession *dca.StreamingSession
}

type YtSong struct {
  URL string `json:"url,omitempty"`
  Name string `json:"name,omitempty"`
  Image string `json:"image,omitempty"`

  Duration time.Duration `json:"duration"`
}

func (y YoutubeManager) enqueue(url string){
  
}

func (y YoutubeManager) pause(){

}

func (y YoutubeManager) skip(){

}

func (y YoutubeManager) stop(){

}

func (y YoutubeManager) queue(){

}

func (y YoutubeManager) playNow(){

}

func (y YoutubeManager) unpause(){

}


func (y YoutubeManager) createVoice(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.VoiceConnection, error){
  // For each person in a voice channel
  for _, vs := range guild.VoiceStates {
    // If they are the caller and (we are already in, or not playing)
    if vs.UserID == m.Author.ID && (vs.ChannelID == y.VC.ChannelID || !y.VC.Playing){
      vc, err := s.ChannelVoiceJoin(guild.ID, vs.ChannelID, false, true)
      if err != nil {
        log.Error(err)
        return nil, err
      }
      return vc, nil
    }
  }
  return nil, errors.New("caller not in voice channel")
}


func getGuild(s *discordgo.Session, channelID string) (g *discordgo.Guild, err error){
  var channel *discordgo.Channel
  channel, err = getChannel(s, channelID)
  if err != nil {
    return
  }
  gID := channel.GuildID
  g, err = s.State.Guild(gID)
  if err != nil {
    if err == discordgo.ErrStateNotFound {
      guildDetails, err = s.Guild(gID)
      if err != nil {
        log.Error("couldn't get guild details")
      }
    }
  }
  return
}

func getChannel(s *discordgo.Session, channelID string) (channel *discordgo.Channel, err error){
  channel, err = s.State.Channel(channelID)
  if err != nil {
    if err == discordgo.ErrStateNotFound {
      channel, err = s.Channel(channelID)
      if err != nil {
        log.Error("couldn't get channel details")
      }
    }
  }
  return
}
