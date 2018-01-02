/*

Directory: arbitrage
File: main.go
Compile and run: go build && ./arbitrage

Author: Tim Siwula

*/


package main

import (
  "fmt"
  "time"
)


func main() {

  start := time.Now()

  go loop()
  elapsed := time.Since(start)
  fmt.Println("Loop took ", elapsed)
  time.Sleep(time.Second * 5)

}

// loop logic
func loop() {
  for range time.Tick(time.Second *1){
      fmt.Println("Foo %d", time.Second *1)
  }
 }
