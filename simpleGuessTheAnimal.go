package main

import (
	"fmt"
	"bufio"
	"math/rand"
	"time"
	"os"
	"strings"
)

func main() {
	animalList := []string{
		"Dog",
		"Cat",
		"Monkey",
		"Cow",
		"Elephant",
	}

	newSource := rand.NewSource(time.Now().UnixNano())
	x := rand.New(newSource)
	randomAnimal := animalList[x.Intn(len(animalList))]

	scanGuess := bufio.NewScanner(os.Stdin)
	fmt.Println("Guess the animal: ")
	scanGuess.Scan()
	guessedAnimal := scanGuess.Text()

	if strings.Title(guessedAnimal) == randomAnimal {
		fmt.Println("You've correctly guessed the animal:", randomAnimal)
	}else {
		fmt.Println("Wrong Guess.")
		fmt.Println("The animal was:", randomAnimal)
	}
}