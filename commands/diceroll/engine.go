package diceroll

import (
  "math/rand"
  "fmt"
  "strings"
  "errors"
  "sort"
)

const (
  MINUS = -1
  PLUS  = 1
)

type Roll struct {
  amount   int
  sides    int
  operator int
  highest  int
  lowest   int
}

type DiceQuery struct {
    rolls []Roll
    results []int
}

// Return as a string the results of the query
func executeQuery(query string) string {
  if query == "stats" {
      old := doStats()
      stats := make([]interface{}, len(old))
      for i, v := range old {
        stats[i] = v
      }
      return fmt.Sprintf("Stats: %d, %d, %d, %d, %d, %d", stats...)
  }

  dq, err := parseQuery(query)
  if err != nil {
    return ""
  }

  err = executeDq(dq) //
  if err != nil {
    return ""
  }

  sum := 0
  for _, v := range dq.results {
    sum += v
  }

  s := fmt.Sprintf("Rolls: %v\n", dq.results)

  return fmt.Sprintf("%s... = %d", s, sum)
}

func parseQuery(query string) (*DiceQuery, error) {
  dq := &DiceQuery{}

  roll, q, err := parseOne(query, true)
  if err != nil {
    return dq, err
  }
  dq.rolls = append(dq.rolls, roll)

  for ; len(q) > 0 ; {
    roll, q, err = parseOne(q, false)
    if err != nil {
      return dq, err
    }
    dq.rolls = append(dq.rolls, roll)
  }

  return dq, nil
}

func parseOne(query string, first bool) (Roll, string, error) {
  var roll Roll

  // Get modifier
  switch {
  case query[0] == '-':
    roll.operator = MINUS
    query = query[1:]
  case query[0] == '+':
    roll.operator = PLUS
    query = query[1:]
  case first:
    roll.operator = PLUS
  default:
    return roll, "", errors.New("Invalid query")
  }

  // Find end of the query
  next := strings.IndexAny(query, "+-")
  curr := query
  if next != -1 {
    curr = query[:next]
  }

  // Try all three formats
  // There is probably a cleaner way to do this lol
  _, err := fmt.Sscanf(curr, "%dd%dkh%d", &roll.amount, &roll.sides, &roll.highest)
  if err != nil {
    _, err = fmt.Sscanf(curr, "%dd%dkl%d", &roll.amount, &roll.sides, &roll.lowest)
    if err != nil {
      _, err = fmt.Sscanf(curr, "%dd%d", &roll.amount, &roll.sides)
      if err != nil {
        _, err = fmt.Sscanf(curr, "%d", &roll.amount)
        if err != nil {
          return roll, "", err
        } else {
          // Constant values. Do it a stupid way.
          roll.sides = 1
          // This will be something like 20d1 for a constant, lol
        }
      }
    }
  }

  // Empty string if there isn't another operator
  if next == -1 {
    return roll, "", nil
  } else {
    return roll, query[next:], nil
  }
}

func executeDq(dq *DiceQuery) error {

  for _, roll := range dq.rolls {
    res, err := executeRoll(roll)
    if err != nil {
      return err
    }
    dq.results = append(dq.results, res)
  }

  return nil
}

// Return the sum result of this roll as an int
func executeRoll(roll Roll) (int, error) {
  // Filter out things that won't work,
  // Or might take a very long time.
  if roll.amount < 1   ||
     roll.amount > 100 ||
     roll.sides > 1000000 || // arbitrary, watching for overflows
     roll.highest > roll.sides ||
     roll.lowest > roll.sides ||
     roll.sides < 1 {
    return 0, errors.New("Invalid arguments for dice roll.")
  }

  var res []int
  for i := 0; i < roll.amount; i++ {
    res = append(res, (rand.Intn(roll.sides)+1)*roll.operator)
  }

  // Reverse order of negative results.
  if roll.operator == MINUS {
    sort.Sort(sort.Reverse(sort.IntSlice(res)))
  } else {
    sort.Ints(res)
  }

  if roll.highest > 0 {
    res = res[roll.amount-roll.highest:]
  } else if roll.lowest > 0 {
    res = res[:roll.lowest]
  }

  sum := 0
  for _, v := range res {
    sum = sum + v
  }

  return sum, nil
}

func doStats() []int {
  r := Roll{amount:4,sides:6,operator:PLUS,highest:3}
  var ret []int
  for i := 0; i<6; i++ {
    sum, _ := executeRoll(r)
    ret = append(ret, sum)
  }

  return ret
}
