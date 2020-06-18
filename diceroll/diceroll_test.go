package diceroll

import (
  "testing"
  "math/rand"
  "time"
  "fmt"
)

func init(){
  rand.Seed(time.Now().Unix())
}

func doTest(t *testing.T, exp interface{}, got interface{}){
  if exp != got {
    t.Errorf("\nExpected: '%v'\nReceived: '%v'\n", exp, got)
  }
}

func TestParseOne1(t *testing.T){
  exp := Roll{
    amount: 1,
    sides: 4,
    operator: PLUS,
  }
  got, _, _ := parseOne("1d4", true)
  doTest(t, exp, got)
}

func TestParseOne2(t *testing.T){
  exp := "+5d6"
  _, got, _ := parseOne("1d4+5d6", true)
  doTest(t, exp, got)
}

func TestParseOne3(t *testing.T){
  exp := "+5d6+1d20"
  _, got, _ := parseOne("1d4+5d6+1d20", true)
  doTest(t, exp, got)
}

func TestParseOne4(t *testing.T){
  exp := Roll{
    amount: 1,
    sides: 20,
    operator: PLUS,
  }
  got, _, _ := parseOne("+1d20", false)
  doTest(t, exp, got)
}

func TestParseOne5(t *testing.T){
  exproll := Roll{
    amount: 15,
    sides: 20,
    operator: PLUS,
  }
  expcont := ""
  gotroll, gotcont, _ := parseOne("15d20", true)
  doTest(t, exproll, gotroll)
  doTest(t, gotcont, expcont)
}

func TestParseOne6(t *testing.T){
  _, _, err := parseOne("1d20", false)
  if err == nil {
    t.Errorf("Expected error for continuation with no operator.")
  }
}

func TestExecuteRoll1(t *testing.T){
  roll := Roll{
    amount: 1,
    sides: 10000000,
    operator: PLUS,
  }
  _, err := executeRoll(roll)
  if err == nil {
    t.Errorf("Expected error for roll...")
  }
}

func TestExecuteRoll2(t *testing.T){
  roll := Roll{
    amount: 101,
    sides: 20,
    operator: MINUS,
  }
  _, err := executeRoll(roll)
  if err == nil {
    t.Errorf("Expected error for roll...")
  }
}

func TestExecuteDq1(t *testing.T){
  r1 := Roll{amount: 0, sides:20, operator: PLUS}
  r2 := Roll{amount: 1, sides:20, operator: PLUS}
  dq := &DiceQuery{rolls: []Roll{r1, r2}}
  err := executeDq(dq)
  if err == nil {
    t.Errorf("Expected error")
  }
}

func TestExecuteDq2(t *testing.T){
  r1 := Roll{amount:1,sides:20,operator:PLUS}
  dq := &DiceQuery{rolls: []Roll{r1,r1,r1,r1,r1,r1}}
  executeDq(dq)
}

func TestSpecial(t *testing.T){
  fmt.Println(Do("20"))
}
