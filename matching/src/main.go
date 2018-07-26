package main

import (
	"fmt"
	"time"

	"./match"
)

func main() {
	timeFunc(func() {
		fmt.Println(match.MatchV0("../match/1.jpeg"))
		fmt.Println(match.MatchV0("../match/2.jpeg"))
	})

	timeFunc(func() {
		fmt.Println(match.MatchV1("../match/1.jpeg"))
		fmt.Println(match.MatchV1("../match/2.jpeg"))
	})

	timeFunc(func() {
		fmt.Println(match.MatchV2("../match/1.jpeg"))
		fmt.Println(match.MatchV2("../match/2.jpeg"))
	})
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}
