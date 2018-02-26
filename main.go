package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Login struct {
	Username string
	Password string
}

func (log *Login) generateLogin(nbreCara int) {
	const base string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, nbreCara)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = base[rand.Intn(len(base))]
	}
	log.Username = string(b)
	log.Password = string(b)
}
func main() {
	for i := 0; i < 10; i++ {
		var login1 Login
		login1.generateLogin(8)
		fmt.Printf("%v", login1)
	}
}
