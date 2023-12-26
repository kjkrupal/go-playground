package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	PlayMenu("Welcome to Dice Roller powered by Go!\n")
}

func PlayMenu(prompt string) {

	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)

	option, _ := GetUserInput("Select number of dice (max 3): ", reader)

	switch option {
	case "1":
		fmt.Println("Hit enter after every roll, press any other key to exit")
		PlayGame(1, reader)
	case "2":
		fmt.Println("Hit enter after every roll, press any other key to exit")
		PlayGame(2, reader)
	case "3":
		fmt.Println("Hit enter after every roll, press any other key to exit")
		PlayGame(3, reader)
	default:
		PlayMenu("Invalid selection")
	}
}

func PlayGame(count int, r *bufio.Reader) {

	output := ""
	for i := 0; i < count; i++ {
		val := RollDie()
		output += fmt.Sprintf("%v", val)
		output += " "
	}
	fmt.Println(output)
	option, _ := GetUserInput("", r)

	switch option {
	case "":
		PlayGame(count, r)
	default:
		fmt.Println("Exiting game")
	}
}

func GetUserInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	val, err := r.ReadString('\n')
	return strings.TrimSpace(val), err
}

func RollDie() int {
	return rand.Intn(6) + 1
}
