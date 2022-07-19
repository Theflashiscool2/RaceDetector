package main

import (
	"fmt"
	"strings"
)

var race string

func main() {
	fmt.Println("Hello? what is your race?")
	fmt.Scan(&race)
	if strings.EqualFold(race, "black") {
		for {
			fmt.Println("nigger")
		}
	}
  if strings.EqualFold(race, "mexican"){
    for {
      fmt.Println("Beaner")
    }
  }
  if strings.EqualFold(race, "white"){
    for {
      fmt.Println("Cracker")
    }
  }
  if strings.EqualFold(race, "arabic"){
    for {
      fmt.Println("terrorist")
    }
  }
}
