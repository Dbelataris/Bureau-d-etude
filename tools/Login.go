package tools

import (
	"math/rand"
	"time"
)

type Login struct {
	Username       string
	Password       string
	StartTime      string
	ExpirationTime int //En minute
}

func (log *Login) GenerateLogin(nbreCara int, Expiration int) {
	const base string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, nbreCara)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = base[rand.Intn(len(base))]
	}
	log.Username = string(b) //Conversion du tableau de bytes en string
	log.Password = string(b)
	log.StartTime = time.Now().Format("2006-01-02 15:04:05")
	log.ExpirationTime = Expiration

}
