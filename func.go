package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type character struct {
	speed      int
	name       string
	cost       [3]int
	currentLvl int
}

func info() {
	fmt.Println("Hi there, there is a mini game The Mine. For start the game enter \x22start\x22, for end the game enter \x22end\x22.\n" +
		"You have 3 characters: The Miner, The Finder, The Smelter, what each of them does, i think you got it. \n" +
		"You can upgrade the lvl, and buy 3 more miners of any type. For lvlup you should enter  \x22lvl Name\x22. \n" +
		"For buy someone you should enter\x22buy Finder/buy Miner/buy Smelter\x22. Good luck")
}

func Scan2() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func possibleUP(c character, b int) bool {
	if c.currentLvl != 2 {
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

func Lvlup(c *character, b *int) {
	c.currentLvl++
	c.speed++
	*b -= c.cost[c.currentLvl]

}

func opportunity(c character, str string) bool {
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

func rdm() string {
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

func possibleBuy(c character, b int) bool {
	if c.cost[0] <= b {
		return true
	}
	return false
}

func Finder(Channel, done chan string, c character) {
Loop:
	for true {
		select {
		case <-done:
			break Loop
		default:
			x := rdm()
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

func Timer(t *int) {
	for true {
		*t++
		time.Sleep(1 * time.Second)
	}
}

func conv(x int) string {
	h := x / 3600
	m := x % 3600 / 60
	s := x % 3600 % 60
	str := strconv.Itoa(h) + ":" + strconv.Itoa(m) + ":" + strconv.Itoa(s)
	return str
}

func Miner(input, done, output chan string, c character) {
Loop:
	for true {
		select {
		case foundOre, ok := <-input:
			if !ok {
				break Loop
			}
			if opportunity(c, foundOre) {
				fmt.Println(c.name, " mine: ", foundOre)
				output <- foundOre
				for i := 0; i <= 15/c.speed; i++ {
					time.Sleep(1 * time.Second)
				}
			}
		case <-done:
			break Loop
		}
	}
}

func Smelter(input, done chan string, c character, b *int, total *int) {
Loop:
	for true {
		select {
		case minedOre, ok := <-input:
			if !ok {
				break Loop
			}
			if opportunity(c, minedOre) {
				fmt.Println(c.name, ": ", minedOre, " is smelted")
				switch minedOre {
				case "Coal":
					*total += 10
					*b += 10
				case "Iron":
					*total += 20
					*b += 20
				case "Gold":
					*total += 30
					*b += 30
				}
				fmt.Println("Coins: ", *b)
				for i := 0; i <= 15/c.speed; i++ {
					time.Sleep(1 * time.Second)
				}
			}
		case <-done:
			break Loop
		}
	}
}
