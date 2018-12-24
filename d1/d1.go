package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  count := 0

  file, err := os.Open("/Users/mattf/Documents/Code/Go/src/advent-of-code-2018/d1/input.txt")
  check(err)

  reader := bufio.NewReader(file)

  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      break
    }

    sign := line[:1]
    mag := strings.TrimSpace(line[1:])

    i, err := strconv.Atoi(mag)
    check(err)
    
    if sign == "+" {
      count += i
    } else {
      count -= i
    }
  }

  fmt.Println(count)

}
