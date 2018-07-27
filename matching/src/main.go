package main

import (
	"fmt"
	"time"

	"./match"
)

func main() {
	// timeFunc(func() {
	// 	fmt.Println(match.MatchV0("../match/1.jpeg"))
	// 	fmt.Println(match.MatchV0("../match/2.jpeg"))
	// })

	// timeFunc(func() {
	// 	fmt.Println(match.MatchV1("../match/1.jpeg"))
	// 	fmt.Println(match.MatchV1("../match/2.jpeg"))
	// })

	// timeFunc(func() {
	// 	fmt.Println(match.MatchV2("../match/1.jpeg"))
	// 	fmt.Println(match.MatchV2("../match/2.jpeg"))
	// })

	timeFunc(func() {
		fmt.Println(match.MatchV3("../png_match/1.PNG"))
		fmt.Println(match.MatchV3("../png_match/2.PNG"))
		fmt.Println(match.MatchV3("../png_match/3.PNG"))
	})
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}
