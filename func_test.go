package main

import (
	"testing"
)

var test = character{
	name:       "test",
	speed:      1,
	cost:       [3]int{200, 500, 1000},
	currentLvl: 0,
}

func TestPossibleUP1(t *testing.T) {
	test.currentLvl = 3
	if PossibleUP(test, 1000) {
		t.Errorf("Incorrect result, Expect False , got True")
	}
}

func TestPossibleUP2(t *testing.T) {
	test.currentLvl = 1
	if PossibleUP(test, 100) {
		t.Errorf("Incorrect result, Expect False , got True")
	}
}

func TestLvlup(t *testing.T) {
	test.currentLvl = 0
	test.speed = 1
	Lvlup(&test)
	if test.currentLvl != 1 && test.speed != 2 {
		t.Errorf("Incorrect result")
	}
}

func TestOpportunity(t *testing.T) {
	test.currentLvl = 0
	if !Opportunity(test, "Coal") {
		t.Errorf("Incorrect result, Expect True , got False")
	}
	if Opportunity(test, "Gold") {
		t.Errorf("Incorrect result, Expect False , got True")
	}
}
func TestOpportunity1(t *testing.T) {
	test.currentLvl = 1
	if !Opportunity(test, "Coal") {
		t.Errorf("Incorrect result, Expect True , got False")
	}
	if !Opportunity(test, "Iron") {
		t.Errorf("Incorrect result, Expect True , got False")
	}
	if Opportunity(test, "Gold") {
		t.Errorf("Incorrect result, Expect False , got True")
	}
}
func TestOpportunity2(t *testing.T) {
	test.currentLvl = 2
	if !Opportunity(test, "Coal") {
		t.Errorf("Incorrect result, Expect True , got False")
	}
	if !Opportunity(test, "Iron") {
		t.Errorf("Incorrect result, Expect True , got False")
	}
	if !Opportunity(test, "Gold") {
		t.Errorf("Incorrect result, Expect True , got False")
	}
}

func TestPossibleBuy(t *testing.T) {
	if !PossibleBuy(test, 2000) {
		t.Errorf("Incorrect result, Expect True , got False")
	}
	if PossibleBuy(test, 10) {
		t.Errorf("Incorrect result, Expect False , got True")
	}
}
