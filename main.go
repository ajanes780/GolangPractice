package main

import (
	"fmt"

	mascot "example.com/m/mascot"
	quote "rsc.io/quote"
)


func main () {
	fmt.Println("the best mascot is")
  fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
	
}

