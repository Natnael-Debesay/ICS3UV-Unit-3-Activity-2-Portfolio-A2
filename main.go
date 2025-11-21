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
	var provinceTerritory1 string
	var Capital1 string
	var provinceTerritory2 string
	var Capital2 string
	var provinceTerritory3 string
	var Capital3 string
	var provinceTerritory4 string
	var Capital4 string
	var provinceTerritory5 string
	var Capital5 string
	var provinceTerritory6 string
	var Capital6 string
	var provinceTerritory7 string
	var Capital7 string
	var provinceTerritory8 string
	var Capital8 string
	var provinceTerritory9 string
	var Capital9 string
	var provinceTerritory10 string
	var Capital10 string
	var provinceTerritory11 string
	var Capital11 string
	var provinceTerritory12 string
	var Capital12 string
	var provinceTerritory13 string
	var Capital13 string




	reader:= bufio.NewReader(os.Stdin)
	//input
	fmt.Print("Province / Territory 1: ")
	inputProvinceTerritory1, _ := reader.ReadString('\n')
	provinceTerritory1 = strings.TrimSpace(inputProvinceTerritory1)

	fmt.Print("Capital 1: ")
	inputCapital1, _ := reader.ReadString('\n')
	Capital1 = strings.TrimSpace(inputCapital1)


	fmt.Print("Province / Territory 2: ")
	inputProvinceTerritory2, _ := reader.ReadString('\n')
	provinceTerritory2 = strings.TrimSpace(inputProvinceTerritory2)

	fmt.Print("Capital 2: ")
	inputCapital2, _ := reader.ReadString('\n')
	Capital2 = strings.TrimSpace(inputCapital2)


	fmt.Print("Province / Territory 3: ")
	inputProvinceTerritory3, _ := reader.ReadString('\n')
	provinceTerritory3 = strings.TrimSpace(inputProvinceTerritory3)

	fmt.Print("Capitalv 3: ")
	inputCapital3, _ := reader.ReadString('\n')
	Capital3 = strings.TrimSpace(inputCapital3)

	
	fmt.Print("Province / Territory 4: ")
	inputProvinceTerritory4, _ := reader.ReadString('\n')
	provinceTerritory4 = strings.TrimSpace(inputProvinceTerritory4)

	fmt.Print("Capital 4: ")
	inputCapital4, _ := reader.ReadString('\n')
	Capital4 = strings.TrimSpace(inputCapital4)

	
	fmt.Print("Province / Territory 5: ")
	inputProvinceTerritory5, _ := reader.ReadString('\n')
	provinceTerritory5 = strings.TrimSpace(inputProvinceTerritory5)

	fmt.Print("Capital 5: ")
	inputCapital5, _ := reader.ReadString('\n')
	Capital5 = strings.TrimSpace(inputCapital5)

	fmt.Print("Province / Territory 6: ")
	inputProvinceTerritory6, _ := reader.ReadString('\n')
	provinceTerritory6 = strings.TrimSpace(inputProvinceTerritory6)

	fmt.Print("Capital 6: ")
	inputCapital6, _ := reader.ReadString('\n')
	Capital6 = strings.TrimSpace(inputCapital6)

	fmt.Print("Province / Territory 7: ")
	inputProvinceTerritory7, _ := reader.ReadString('\n')
	provinceTerritory7 = strings.TrimSpace(inputProvinceTerritory7)

	fmt.Print("Capital 7: ")
	inputCapital7, _ := reader.ReadString('\n')
	Capital7 = strings.TrimSpace(inputCapital7)

	fmt.Print("Province / Territory 8: ")
	inputProvinceTerritory8, _ := reader.ReadString('\n')
	provinceTerritory8 = strings.TrimSpace(inputProvinceTerritory8)

	fmt.Print("Capital 8: ")
	inputCapital8, _ := reader.ReadString('\n')
	Capital8 = strings.TrimSpace(inputCapital8)

	fmt.Print("Province / Territory 9: ")
	inputProvinceTerritory9, _ := reader.ReadString('\n')
	provinceTerritory9 = strings.TrimSpace(inputProvinceTerritory9)

	fmt.Print("Capital 9: ")
	inputCapital9, _ := reader.ReadString('\n')
	Capital9 = strings.TrimSpace(inputCapital9)

	fmt.Print("Province / Territory 10: ")
	inputProvinceTerritory10, _ := reader.ReadString('\n')
	provinceTerritory10 = strings.TrimSpace(inputProvinceTerritory10)

	fmt.Print("Capital 10: ")
	inputCapital10, _ := reader.ReadString('\n')
	Capital10 = strings.TrimSpace(inputCapital10)

	fmt.Print("Province / Territory 11: ")
	inputProvinceTerritory11, _ := reader.ReadString('\n')
	provinceTerritory11 = strings.TrimSpace(inputProvinceTerritory11)

	fmt.Print("Capital 11: ")
	inputCapital11, _ := reader.ReadString('\n')
	Capital11 = strings.TrimSpace(inputCapital11)
	
	fmt.Print("Province / Territory 12: ")
	inputProvinceTerritory12, _ := reader.ReadString('\n')
	provinceTerritory12 = strings.TrimSpace(inputProvinceTerritory12)

	fmt.Print("Capital 12: ")
	inputCapital12, _ := reader.ReadString('\n')
	Capital12 = strings.TrimSpace(inputCapital12)

	fmt.Print("Province / Territory 13: ")
	inputProvinceTerritory13, _ := reader.ReadString('\n')
	provinceTerritory13 = strings.TrimSpace(inputProvinceTerritory13)

	fmt.Print("Capital 13: ")
	inputCapital13, _ := reader.ReadString('\n')
	Capital13 = strings.TrimSpace(inputCapital13)

	// output
	fmt.Println()
	fmt.Printf("%s", provinceTerritory1 + ": " + Capital1 );
	fmt.Printf("%s", provinceTerritory2 + ": " + Capital2 );
	fmt.Printf("%s", provinceTerritory3 + ": " + Capital3 );
	fmt.Printf("%s", provinceTerritory4 + ": " + Capital4 );
	fmt.Printf("%s", provinceTerritory5 + ": " + Capital5 );
	fmt.Printf("%s", provinceTerritory6 + ": " + Capital6 );
	fmt.Printf("%s", provinceTerritory7 + ": " + Capital7 );
	fmt.Printf("%s", provinceTerritory8 + ": " + Capital8 );
	fmt.Printf("%s", provinceTerritory9 + ": " + Capital9 );
	fmt.Printf("%s", provinceTerritory10 + ": " + Capital10 );
	fmt.Printf("%s", provinceTerritory11 + ": " + Capital11 );
	fmt.Printf("%s", provinceTerritory12 + ": " + Capital12 );
	fmt.Printf("%s", provinceTerritory13 + ": " + Capital13 );

	fmt.Println("\nDone.")
}