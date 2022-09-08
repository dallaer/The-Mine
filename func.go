package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//Print some information about the game
func info() {
	fmt.Println("Hi there, there is a mini game The Mine. For start the game enter \x22start\x22, for end the game enter \x22end\x22.\n" +
		"You have 3 characters: The Miner, The Finder, The Smelter, what each of them does, i think you got it. \n" +
		"You can upgrade the lvl, and buy 3 more miners of any type. For lvlup you should enter  \x22lvl Name\x22. \n" +
		"For buy someone you should enter\x22buy Finder/buy Miner/buy Smelter\x22. Good luck")
}

//Scan2 reads from the console and returns a string
func Scan2() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

//Checks the possibility of leveling up
func PossibleUP(c character, b int) bool {
	if c.currentLvl < 2 {
		if c.cost[c.currentLvl] <= b {
			return true
		}

	} else {
		fmt.Println("Already the maximum lvl")
		return false
	}
	fmt.Println("Not enough gold")
	return false
}

func Lvlup(c *character) {
	c.currentLvl++
	c.speed++
	b -= c.cost[c.currentLvl]

}

//Check out the ability to process or mine a resource
func Opportunity(c character, str string) bool {
	b := 3
	switch str {
	case "Coal":
		b = 0
	case "Iron":
		b = 1
	case "Gold":
		b = 2
	}
	if c.currentLvl >= b {
		return true
	}
	fmt.Println("Your miners can't  process this one, yet")
	return false
}

//Rdm randomly creates one of four resources
func Rdm() string {
	rand.Seed(time.Now().UTC().UnixNano())
	x := rand.Intn(100)
	switch {
	case x >= 50:
		return "Coal"
	case x >= 30 && x < 50:
		return "Iron"
	case x >= 20 && x < 30:
		return "Gold"
	}
	return "Rock"
}

//Check out the ability to buy a new character
func PossibleBuy(c character, b int) bool {
	if c.cost[0] <= b {
		return true
	}
	return false
}

//A character, launches the generator for receiving a random resource, reports it and sends the resource to a special channel
func Finder(c character) {
Loop:
	for true {
		select {
		case <-done:
			break Loop
		default:
			x := Rdm()
			if x != "Rock" {
				fmt.Println(c.name, " found: ", x)
				Channel <- x
				for i := 0; i <= 15/c.speed; i++ {
					time.Sleep(1 * time.Second)
				}
			}
		}
	}
}

//с
func Timer() {
	for true {
		timers++
		time.Sleep(1 * time.Second)
	}
}

//Сonvert seconds to required format of time
func Conv(x int) string {
	h := x / 3600
	m := x % 3600 / 60
	s := x % 3600 % 60
	str := strconv.Itoa(h) + ":" + strconv.Itoa(m) + ":" + strconv.Itoa(s)
	return str
}

//A character, takes a resource from a channel, checks if he can get it and if he can, sends it to another channel
func Miner(c character) {
Loop:
	for true {
		select {
		case foundOre, ok := <-Channel:
			if !ok {
				break Loop
			}
			if Opportunity(c, foundOre) {
				fmt.Println(c.name, " mine: ", foundOre)
				Channel2 <- foundOre
				for i := 0; i <= 15/c.speed; i++ {
					time.Sleep(1 * time.Second)
				}
			}
		case <-done:
			break Loop
		}
	}
}

//A character takes a resource from a channel checks if he can craft it and if he can adds value to balance
func Smelter(c character) {
Loop:
	for true {
		select {
		case minedOre, ok := <-Channel2:
			if !ok {
				break Loop
			}
			if Opportunity(c, minedOre) {
				fmt.Println(c.name, ": ", minedOre, " is smelted")
				switch minedOre {
				case "Coal":
					total += 10
					b += 10
				case "Iron":
					total += 20
					b += 20
				case "Gold":
					total += 30
					b += 30
				}
				fmt.Println("Coins: ", b)
				for i := 0; i <= 15/c.speed; i++ {
					time.Sleep(1 * time.Second)
				}
			}
		case <-done:
			break Loop
		}
	}
}

//Handles a request from the user and performs an action depending on the request(lvlup or buy)
func processing(n string) {
	switch n {
	case "buy Finder":
		switch sum {
		case 0:
			if PossibleBuy(Karl, b) {
				fmt.Println("Congratulations you have hired: Lary")
				go Finder(Lary)
				sum++
				b -= Lary.cost[0]
			} else {
				fmt.Println("Not enough gold")
			}
		case 1:
			if PossibleBuy(Karl, b) {
				fmt.Println("Congratulations you have hired: Karl")
				go Finder(Karl)
				sum++
				b -= Karl.cost[0]
			} else {
				fmt.Println("Not enough gold")
			}
		case 2:
			if PossibleBuy(Jim, b) {
				fmt.Println("Congratulations you have hired: Jim")
				go Finder(Jim)
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
			if PossibleBuy(Lary, b) {
				fmt.Println("Congratulations you have hired: Lary")
				go Miner(Lary)
				sum++
				b -= Lary.cost[0]
			} else {
				fmt.Println("Not enough gold")
			}
		case 1:
			if PossibleBuy(Karl, b) {
				fmt.Println("Congratulations you have hired: Karl")
				go Miner(Karl)
				sum++
				b -= Karl.cost[0]
			} else {
				fmt.Println("Not enough gold")
			}
		case 2:
			if PossibleBuy(Jim, b) {
				fmt.Println("Congratulations you have hired: Jim")
				go Miner(Jim)
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
			if PossibleBuy(Karl, b) {
				fmt.Println("Congratulations you have hired: Lary")
				go Smelter(Lary)
				sum++
				b -= Lary.cost[0]
			} else {
				fmt.Println("Not enough gold")
			}
		case 1:
			if PossibleBuy(Karl, b) {
				fmt.Println("Congratulations you have hired: Karl")
				go Smelter(Karl)
				sum++
				b -= Karl.cost[0]
			} else {
				fmt.Println("Not enough gold")
			}
		case 2:
			if PossibleBuy(Jim, b) {
				fmt.Println("Congratulations you have hired: Jim")
				go Smelter(Jim)
				sum++
				b -= Jim.cost[0]
			} else {
				fmt.Println("Not enough gold")
			}
		default:
			fmt.Println("Already the maximum  miners")
		}
	case "lvl Gary":
		if PossibleUP(Gary, b) {
			Lvlup(&Gary)
			fmt.Println("Congratulations, lvl Gary was upgrade,current Gary's lvl is: ", Gary.currentLvl)
		}
	case "lvl Bob":
		if PossibleUP(Bob, b) {
			Lvlup(&Bob)
			fmt.Println("Congratulations, lvl Bob was upgrade,current Bob's lvl is: ", Bob.currentLvl)
		}
	case "lvl Nick":
		if PossibleUP(Nick, b) {
			Lvlup(&Nick)
			fmt.Println("Congratulations, lvl Nick was upgrade,current Nick's lvl is: ", Nick.currentLvl)
		}
	case "lvl Lary":
		if PossibleUP(Lary, b) && sum == 1 {
			Lvlup(&Lary)
			fmt.Println("Congratulations, lvl Lary was upgrade,current Lary's lvl is: ", Lary.currentLvl)
		}
	case "lvl Karl":
		if PossibleUP(Karl, b) && sum == 2 {
			Lvlup(&Karl)
			fmt.Println("Congratulations, lvl Karl was upgrade,current Karl's lvl is: ", Karl.currentLvl)
		}
	case "lvl Jim":
		if PossibleUP(Jim, b) && sum == 3 {
			Lvlup(&Jim)
			fmt.Println("Congratulations, lvl Jim was upgrade,current Jim's lvl is: ", Jim.currentLvl)
		}
	}
}

type character struct {
	speed      int
	name       string
	cost       [3]int
	currentLvl int
}

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
var sum = 0
var check = 0
var timers = 0
var total = 0
var b = 0
var Channel2 = make(chan string)
var Channel = make(chan string)
var done = make(chan string)
