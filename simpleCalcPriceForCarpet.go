package main 

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanFloorWidth := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Floor Width (In Feet): ")
	scanFloorWidth.Scan()
	floorWidth, _ := strconv.ParseFloat(scanFloorWidth.Text(), 10)

	scanFloorHeight := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Floor Height (In Feet): ")
	scanFloorHeight.Scan()
	floorHeight, _ := strconv.ParseFloat(scanFloorHeight.Text(), 10)

	scanCost := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Carpet Cost Per Foot (In USD): ")
	scanCost.Scan()
	cost, _ := strconv.ParseFloat(scanCost.Text(), 2)
	
	areaOfFloor := floorWidth * floorHeight
	price := cost * areaOfFloor

	fmt.Printf("The total cost for carpet will be: $%v", price)
}