package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "Bureau-d-etude/github.com/go-sql-driver/mysql"
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
	var UsersValide = []Login{}
	db, _ := sql.Open("mysql", "radius:radpass@tcp(127.0.0.1:3306)/radius")
	DbSatus := db.Ping()

	if DbSatus == nil {
		fmt.Println("Connexion à la BD réussie")
	} else {
		fmt.Println("Connexion à la BD échouée")
	}

	for i := 0; i < 5; i++ {
		var login1 Login
		login1.generateLogin(8)
		UsersValide = append(UsersValide, login1)
		//fmt.Println("%v", login1)
	}
	fmt.Println(UsersValide)
}
