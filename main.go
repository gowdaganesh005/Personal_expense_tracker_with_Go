package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type expense struct {
	Amt    uint
	Reason string
}

var exp_map = make(map[string][]expense)
var total uint

func main() {

	loadfile(exp_map, "expenses1.json")

	user := "Ganesh"

	var response string

	fmt.Printf("###########################################################################")
	fmt.Println()
	total = findtotal()

	fmt.Printf("Welcome %v your expenses are %v \n", user, total)

	for {
		fmt.Print("Do you would like to add an expense \n #####     IF want to add an expense Enter ADD      #####  \n #####    Type q to quit and exit      #####\n")

		fmt.Scan(&response)
		switch response {
		case "ADD":
			addexp(total)
			Savefile(exp_map, "expenses1.json")
			total = 0
			total = findtotal()
			fmt.Printf("Your all time Total is  %v \n \n", total)

		case "q":
			fmt.Printf("########################   Your Grand Total is %v #########################\n", total)
			return

		default:
			fmt.Println("Enter Valid Response")
			fmt.Println()
			continue

		}

	}
}
func addexp(total uint) {

	var date string
	fmt.Println("Enter the date in format DD/MM/YYYY")
	fmt.Scan(&date)

	var amt uint
	fmt.Println("Enter the amount to be added:")
	fmt.Scan(&amt)

	var reason string
	fmt.Println("Enter the reason of expense:")
	fmt.Scan(&reason)

	newexp := expense{
		Amt:    amt,
		Reason: reason,
	}
	exp_map[date] = append(exp_map[date], newexp)

	fmt.Printf("###########################################################################")
	fmt.Println()

	//to just crooss check wheter my struct is working properly
	//for date, newexp := range exp_map {
	//	fmt.Printf("date : %v\n ", date)
	//	for _, newexps := range newexp {
	//		fmt.Printf("amt :%v\n", newexps.Amt)
	//		fmt.Printf("reason :%v\n", newexps.Reason)
	//	}
	//}

}
func Savefile(exp_map map[string][]expense, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("error in creating a file")
		return err

	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(exp_map); err != nil {
		return err
	}

	return nil
}
func loadfile(exp_map map[string][]expense, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR in opening the file")
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&exp_map); err != nil {
		fmt.Println("Error in decoding")
		return err

	}
	return nil
}
func findtotal() uint {
	for _, exp := range exp_map { ///calcullating all total expenses till now

		for _, exps := range exp {
			total += exps.Amt

		}
	}
	return total

}
