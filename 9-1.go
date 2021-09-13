package main

import (
	"errors"
	"fmt"
	"strings"
)

var menus = [...]string{"Crudo", "Oysters", "Antipasti", "Pasta", "Cioppino", "Granchio", "Aragosta", "Brodetto", "Cacciucco"}
var carts = []string{}

func validate(input string) (bool, error) {
	var save string
	if input == "" {
		return false, errors.New("cannot be empty")
	}
	for i, _ := range menus {
		if strings.Title(input) == menus[i] {
			save = strings.Title(input)
		}
	}
	if save == "" {
		return false, errors.New("not in the list")
	}
	return true, nil
}

func cart(menus ...string) []string {
	var menu, choice string
	for {
		fmt.Print("Please input your menu by typing (q.e : granchio) : ")
		fmt.Scan(&menu)
		if valid, err := validate(menu); valid {
			carts = append(carts, menu)
			menu = ""
		} else {
			fmt.Println(err.Error())
		}

		fmt.Print("Order another? (Y/N) : ")
		fmt.Scan(&choice)
		if strings.ToLower(choice) == "n" {
			break
		}
	}
	return carts
}

func receipt(carts []string) {
	for i, order := range carts {
		fmt.Println("Order #", i+1, " : ", order)
	}
	fmt.Println("Thanks for the order")
}

func main() {
	fmt.Println("Zinedine's Marea")
	fmt.Println("======================")
	for _, menu := range menus {
		fmt.Println(menu)
	}
	fmt.Println("======================")
	cart()
	receipt(carts)
}
