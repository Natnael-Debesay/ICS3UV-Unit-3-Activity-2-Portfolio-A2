// Author: Natnael Debesay
// Version: 1.0.0
// Date: 2025-11-19
// Fileoverview: This program prompts the user to enter the names of the ten Canadian provinces, three territories, and their capital cities. 

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var provinceTerritory string
	var Capital string

	reader:= bufio.NewReader(os.Stdin)
	//input
	fmt.Print("Province / Territory: ")
	inputProvinceTerritory, _ := reader.ReadString('\n')
	provinceTerritory = strings.TrimSpace(inputProvinceTerritory)

	fmt.Print("Capital: ")
	inputCapital, _ := reader.ReadString('\n')
	Capital = strings.TrimSpace(inputCapital)

	// output
	fmt.Println()
	fmt.Printf(provinceTerritory + ": " + Capital);

	fmt.Println("\nDone.")
}