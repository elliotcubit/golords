package credentials

import (
  "os"
  "io/ioutil"
  "encoding/json"
)

// Struct storings credentials
type Credentials struct {
  Id     string `json:"client_id"`
  Secret string `json:"client_secret"`
  Token  string `json:"token"`
}

// Returns credentials populated from a JSON filepath provided.
func LoadCreds(fname string) (Credentials, error) {
  var creds Credentials

  jsonFile, err := os.Open(fname)
  if err != nil {
    return creds, err
  }
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)
  err = json.Unmarshal(byteValue, &creds)
  if err != nil{
    return creds, err
  }

  return creds, nil
}
