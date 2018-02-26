package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Login struct {
	username string
	password string
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", generateLogin(8))
	}
}

func generateLogin(nbreCara int) string {
	const base string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, nbreCara)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = base[rand.Intn(len(base))]
	}
	return string(b)

}
