package main

import(
	"fmt"
	"math/rand"
	"bufio"
	"strconv"
	"os"
	"time"
)

func main() {
	newSource := rand.NewSource(time.Now().UnixNano())
	y := rand.New(newSource)
	secretNumber := y.Intn(10)

	scanGuess := bufio.NewScanner(os.Stdin)
	fmt.Println("Guess The Number [0-10]: ")
	scanGuess.Scan()
	gusedNumber, _ := strconv.Atoi(scanGuess.Text())

	if gusedNumber == secretNumber {
		fmt.Printf("You guessed correctly! The number was %d", secretNumber)
	} else {
		fmt.Printf("You guessed wrong...")
	}
}