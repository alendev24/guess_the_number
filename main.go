// pass_fail reports wether a grade is passing or failing
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1

	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	for guess := 0; guess < 10; guess++ {
		fmt.Println("You have", 10-guess, "guess left")

		fmt.Print("Make a guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops, Your guess was LOW")
		} else if guess > target {
			fmt.Println("Oopps, your guess was HIGH")
		} else if guess == target {
			fmt.Println("That is correct number!")
		}
		break
	}
}
