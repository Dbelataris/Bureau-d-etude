package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	. "Bureau-d-etude/tools"

	_ "Bureau-d-etude/github.com/go-sql-driver/mysql"
)

func main() {
	again := true
	var LoginsValides = []Login{}
	var loginGen Login
	db, _ := sql.Open("mysql", "radius:radpass@tcp(127.0.0.1:3306)/radius")

	if CheckDbCon(db) {
		println("Connexion à la BD établit")
	} else {
		println("Connexion à la BD échouée")
	}
	for again {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Générer un login? <o/n>")
		str, _ := reader.ReadString('\n')
		if strings.TrimSpace(str) == "o" {
			again = true
			loginGen.GenerateLogin(8, 5)
			StoreLogin(db, loginGen)
			LoginsValides = append(LoginsValides, loginGen)
			fmt.Println(LoginsValides)
		} else {
			again = false
		}
	}

}
