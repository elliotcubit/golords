package quotemanager

import (
  "io/ioutil"
  "os"
  "encoding/json"
  "math/rand"
  "time"
)

type Quote struct {
  AddedBy   string `json:"addedBy"`
  Text      string `json:"text"`
  Timestamp string `json:"timestamp"`
}

type QuoteList []Quote

var loadedQuotes QuoteList

func LoadQuoteList(fname string) error {
  rand.Seed(time.Now().Unix())
  jsonFile, err := os.Open(fname)
  if err != nil {
    return err
  }
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)
  err = json.Unmarshal(byteValue, &loadedQuotes)
  if err != nil {
    return err
  }

  return nil
}

func WriteQuoteList() error {
  file, err := json.Marshal(loadedQuotes)
  if err != nil {
    return err
  }
  err = ioutil.WriteFile("quotelist.json", file, 0666)
  return err
}

func AddQuote(by string, text string, timestamp string){
  quote := Quote{AddedBy: by, Text: text, Timestamp: timestamp}
  loadedQuotes = append(loadedQuotes, quote)
  WriteQuoteList()
}

func GetRandomQuote() Quote {
  ind := rand.Intn(len(loadedQuotes))
  return loadedQuotes[ind]
}
