package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startGame()
}

func startGame() {
	fmt.Println("Welcome to the RandGame, what's your name?")
	inputReader := bufio.NewReader(os.Stdin)
	ans, _ := inputReader.ReadString('\n')
	fmt.Printf("Hello %s, shall we start?\n", ans)
	res, _ := inputReader.ReadString('\n')
	switch strings.ToLower(res) {
	case "no\n", "n\n":
		fmt.Println("Ok, see you later")
		os.Exit(0)
	case "yes\n", "y\n":
		runGame()
		fallthrough
	default:
		startGame()
	}
}

func runGame() {
	fmt.Println("Please check number from 0 to 29, you have 5 times to guess!")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := r.Intn(30)
	for i := 0; i < 5; i++ {
		response, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		num, _ := strconv.Atoi(response[:len(response)-1])
		if result > num {
			fmt.Println("Hint: guess higher")
		} else if result < num {
			fmt.Println("Hint: guess lower")
		} else {
			fmt.Printf("bingo, answer is %d\n", result)
			fmt.Println("Thanks for playing")
			return
		}
	}
	fmt.Println("Times up, Game over")
	fmt.Printf("By the way, the answer is %d\n", result)
	os.Exit(0)

}
