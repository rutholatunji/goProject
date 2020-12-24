package main

import (
	"fmt"
	"math/rand"
	"time"
)

func moneyHeist() {
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true
	eludedGuards := rand.Intn(100)
	//leftSafely := rand.Intn(5)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remeber, this is the first step")
	} else if isHeistOn := false; isHeistOn {
		fmt.Println("Plan a better disguise next time?")
	}
	openedVault := rand.Intn(100)
	if isHeistOn == true && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn := false; isHeistOn {
		fmt.Println("Vault can't be opened")
	}
	fmt.Println(isHeistOn)

	leftSafely := rand.Intn(5)

	if isHeistOn == true {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Heist failed!")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside")
		case 2:
			isHeistOn = false
			fmt.Println("Better luck next time!")
		default:
			fmt.Println("Start the getaway car!")
		}
		amtStole := 10000 + rand.Intn(1000000)
		if isHeistOn == true {
			fmt.Println(amtStole)
		} else {
			fmt.Println("Vault can't be opened")
		}
	}
}
