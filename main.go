package main

import (
	"fmt"
)

func main() {
	var Gary = character{
		name:       "Gary",
		speed:      1,
		cost:       [3]int{200, 500, 1000},
		currentLvl: 0,
	}
	var Bob = character{
		name:       "Bob",
		speed:      1,
		cost:       [3]int{200, 500, 1000},
		currentLvl: 0,
	}
	var Nick = character{
		name:       "Nick",
		speed:      1,
		cost:       [3]int{200, 500, 1000},
		currentLvl: 0,
	}
	var Lary = character{
		name:       "Lary",
		speed:      1,
		cost:       [3]int{200, 500, 1000},
		currentLvl: 0,
	}
	var Karl = character{
		name:       "Karl",
		speed:      1,
		cost:       [3]int{200, 500, 1000},
		currentLvl: 0,
	}
	var Jim = character{
		name:       "Jim",
		speed:      1,
		cost:       [3]int{200, 500, 1000},
		currentLvl: 0,
	}
	var n string
	var name string
	b := 0
	Channel := make(chan string)
	done := make(chan string)
	Channel2 := make(chan string)
	sum := 0
	check := 0
	time := 0
	total := 0
	info()
	for true {
		n = Scan2()
		switch n {
		case "info":
			info()
		case "start":
			if check == 0 {
				fmt.Println("Enter the name")
				fmt.Scan(&name)
				check = 1
				go Timer(&time)
				go Finder(Channel, done, Gary)
				go Miner(Channel, done, Channel2, Bob)
				go Smelter(Channel2, done, Nick, &b, &total)
			}
		default:
			switch n {
			case "buy Finder":
				switch sum {
				case 0:
					if possibleBuy(Karl, b) {
						fmt.Println("Congratulations you have hired: Lary")
						go Finder(Channel, done, Lary)
						sum++
						b -= Lary.cost[0]
					} else {
						fmt.Println("Not enough gold")
					}
				case 1:
					if possibleBuy(Karl, b) {
						fmt.Println("Congratulations you have hired: Karl")
						go Finder(Channel, done, Karl)
						sum++
						b -= Karl.cost[0]
					} else {
						fmt.Println("Not enough gold")
					}
				case 2:
					if possibleBuy(Jim, b) {
						fmt.Println("Congratulations you have hired: Jim")
						go Finder(Channel, done, Jim)
						sum++
						b -= Jim.cost[0]
					} else {
						fmt.Println("Not enough gold")
					}
				default:
					fmt.Println("Already the maximum number of miners")
				}
			case "buy Miner":
				switch sum {
				case 0:
					if possibleBuy(Lary, b) {
						fmt.Println("Congratulations you have hired: Lary")
						go Miner(Channel, done, Channel2, Lary)
						sum++
						b -= Lary.cost[0]
					} else {
						fmt.Println("Not enough gold")
					}
				case 1:
					if possibleBuy(Karl, b) {
						fmt.Println("Congratulations you have hired: Karl")
						go Miner(Channel, done, Channel2, Karl)
						sum++
						b -= Karl.cost[0]
					} else {
						fmt.Println("Not enough gold")
					}
				case 2:
					if possibleBuy(Jim, b) {
						fmt.Println("Congratulations you have hired: Jim")
						go Miner(Channel, done, Channel2, Jim)
						sum++
						b -= Jim.cost[0]
					} else {
						fmt.Println("Not enough gold")
					}
				default:
					fmt.Println("Already the maximum number of miners")
				}
			case "buy Smelter":
				switch sum {
				case 0:
					if possibleBuy(Karl, b) {
						fmt.Println("Congratulations you have hired: Lary")
						go Smelter(Channel2, done, Lary, &b, &total)
						sum++
						b -= Lary.cost[0]
					} else {
						fmt.Println("Not enough gold")
					}
				case 1:
					if possibleBuy(Karl, b) {
						fmt.Println("Congratulations you have hired: Karl")
						go Smelter(Channel2, done, Karl, &b, &total)
						sum++
						b -= Karl.cost[0]
					} else {
						fmt.Println("Not enough gold")
					}
				case 2:
					if possibleBuy(Jim, b) {
						fmt.Println("Congratulations you have hired: Jim")
						go Smelter(Channel2, done, Jim, &b, &total)
						sum++
						b -= Jim.cost[0]
					} else {
						fmt.Println("Not enough gold")
					}
				default:
					fmt.Println("Already the maximum  miners")
				}
			case "lvl Gary":
				if possibleUP(Gary, b) {
					Lvlup(&Gary, &b)
					fmt.Println("Congratulations, lvl Gary was upgrade,current Gary's lvl is: ", Gary.currentLvl)
				}
			case "lvl Bob":
				if possibleUP(Bob, b) {
					Lvlup(&Bob, &b)
					fmt.Println("Congratulations, lvl Bob was upgrade,current Bob's lvl is: ", Bob.currentLvl)
				}
			case "lvl Nick":
				if possibleUP(Nick, b) {
					Lvlup(&Nick, &b)
					fmt.Println("Congratulations, lvl Nick was upgrade,current Nick's lvl is: ", Nick.currentLvl)
				}
			case "lvl Lary":
				if possibleUP(Lary, b) && sum == 1 {
					Lvlup(&Lary, &b)
					fmt.Println("Congratulations, lvl Lary was upgrade,current Lary's lvl is: ", Lary.currentLvl)
				}
			case "lvl Karl":
				if possibleUP(Karl, b) && sum == 2 {
					Lvlup(&Karl, &b)
					fmt.Println("Congratulations, lvl Karl was upgrade,current Karl's lvl is: ", Karl.currentLvl)
				}
			case "lvl Jim":
				if possibleUP(Jim, b) && sum == 3 {
					Lvlup(&Jim, &b)
					fmt.Println("Congratulations, lvl Jim was upgrade,current Jim's lvl is: ", Jim.currentLvl)
				}
			case "end":
				fmt.Print("\n", name, "\nTotal earned: ", total, "\nTime: ", conv(time))
				return
			}
		}
	}
}
