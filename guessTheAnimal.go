package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"strings"
	"math/rand"
)

func main() {
	lives := 3
	animalList := []string{
		"Horse",
		"Snake",
		"Butterfly",
		"Spider",
		"Zebra",
	}

	newSource := rand.NewSource(time.Now().UnixNano())
	x := rand.New(newSource)
	randomAnimal := animalList[x.Intn(len(animalList))]

	for lives > 0 {
		scanGuess := bufio.NewScanner(os.Stdin)
		fmt.Println("Guess the animal: ")
		scanGuess.Scan()
		guessedAnimal := scanGuess.Text()

		if strings.Title(guessedAnimal) == randomAnimal {
			fmt.Printf("You've correctly guessed the animal: %v\nYou had %v Lives remaining.", randomAnimal, lives)
			break
		}else {
			lives--
			fmt.Printf("Wrong Guess. \nYou have %v Lives remaining.\n\n", lives)
			
			if lives <= 1 {
				fmt.Println("Hint:")
				switch randomAnimal {
				case "Horse":
					fmt.Println("Neigh.")
				case "Snake":
					fmt.Println("Slither.")
				case "Butterfly":
					fmt.Println("Metamorphosis")
				case "Spider":
					fmt.Println("2x4")
				case "Zebra":
					fmt.Println("Yin Yang")
				}
			}
		}
	}
}