package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"

  "golords/credentials"
  "golords/handlers"
  "github.com/bwmarrin/discordgo"
)

func main(){
  creds, err := credentials.LoadCreds("cred.json")
  if err != nil {
    panic("Problem reading json")
  }

  dg, err := discordgo.New("Bot " + creds.Token)
  if err != nil {
    fmt.Println("error creating Discord session,", err)
    return
  }

  // Keep the last 20 messages cached.
  dg.State.MaxMessageCount = 20

  // Register handlers
  dg.AddHandler(handlers.OnMessageCreate)

  // dg.AddHandler(mainhandlers.OnMessageUpdate)

  // Open connection
  err = dg.Open()
  if err != nil {
    fmt.Println("error opening connection,", err)
    return
  }

  // Listen for a ^C with syscall
  fmt.Println("Golords bot is alive. ^C exits.")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <- sc

  // Shut down nicely
  dg.Close()
}
