package main

import (
	"fmt"
)

type Login struct {
	username string
	password string
}

func main() {
	fmt.Printf("Hello friend\n")
}

func generateLogin(nbreCara int) string {
	const base string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return base

}
