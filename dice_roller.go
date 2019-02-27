package main

import (
  "fmt"
  "strings"
  "strconv"
  "math/rand"
  "time"
)

func main(){
  var input string
  fmt.Scanln(&input)
  split := strings.Split(input, "d")
  first, second := split[0], split[1]
  first_int, _ := strconv.Atoi(first)
  second_int, _ := strconv.Atoi(second)
  var sum int
  a := make([]string, first_int)
  for i := 1; i <= first_int; i++ {
    s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)
    num := r.Intn(second_int + 1) - 1
    sum += num
    a[i - 1] = strconv.Itoa(num)
  }
  fmt.Println(sum)
  fmt.Println(strings.Join(a, ","))
}