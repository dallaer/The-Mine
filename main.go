package main

import (
	"fmt"
)

func main() {
	info()
	for true {
		n = Scan2()
		switch n {
		case "info":
			info()
		case "start": //start the timer and one character of each type
			if check == 0 {
				check = 1
				go Timer()
				go Finder(Gary)
				go Miner(Bob)
				go Smelter(Nick)
			}
		case "end": //end of the game and print game statistics
			fmt.Print("\nTotal earned: ", total, "\nTime: ", Conv(timers))
			return
		default:
			processing(n) //
		}
	}
}
