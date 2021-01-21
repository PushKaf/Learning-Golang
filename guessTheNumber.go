package main 

import (
	"fmt"
	"strconv"
	"bufio"
	"math/rand"
	"time"
	"os"
)

func main() {
	lives := 3

	newSource := rand.NewSource(time.Now().UnixNano())
	x := rand.New(newSource)
	secretNumber := x.Intn(20)

	for lives > 0{
		scanGuess := bufio.NewScanner(os.Stdin)
		fmt.Print("Guess the number [0-20]: ")
		scanGuess.Scan()
		guessedNumber, _ := strconv.Atoi(scanGuess.Text())

		if guessedNumber == secretNumber {
			fmt.Printf("You've guessed the secret number: %v. \nYou had %v Lives remaining...", secretNumber, lives)
			break
		}else {
			lives--
			fmt.Println("Wrong number... Try again. \nLives remaining:", lives)
		}
	}
	fmt.Println("The secret number was:", secretNumber)
}