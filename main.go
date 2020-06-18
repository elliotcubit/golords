package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "log"
  "math/rand"
  "time"

  "golords/credentials"
  "golords/handlers"
  "github.com/bwmarrin/discordgo"
)

func main(){
  log.Println("Loading golordsbot")

  creds, err := credentials.LoadCreds()
  if err != nil {
    log.Fatal("Problem loading credentials")
  }

  dg, err := discordgo.New("Bot " + creds.Token)
  if err != nil {
    log.Fatalf("error creating Discord session: %v", err)
  }

  err = InitializeStoredData()
  if err != nil {
    log.Fatalf("error initializing data %v", err)
  }

  // Keep the last 20 messages cached.
  dg.State.MaxMessageCount = 20

  // Register handlers
  dg.AddHandler(handlers.OnMessageCreate)
  dg.AddHandler(handlers.OnMessageUpdate)
  dg.AddHandler(handlers.OnMessageDelete)

  // Open connection
  err = dg.Open()
  if err != nil {
    log.Fatalf("error while opening connection", err)
  }

  // Init RNG to start time.
  rand.Seed(time.Now().Unix())

  // Let 'em know
  fmt.Println("Golords bot is alive. ^C exits.")

  // Listen for a ^C with syscall
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <- sc

  fmt.Println("^C Registered. Beginning the shutdown routine.")

  // Shut down nicely
  dg.Close()
}
