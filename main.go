package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	. "Bureau-d-etude/tools"

	_ "Bureau-d-etude/go-sql-driver"
)

var Db *sql.DB

func main() {
	/*var LoginsValides = []Login{}
	  var loginGen Login*/
	Db, _ = sql.Open("mysql", "radius:radpass@tcp(127.0.0.1:3306)/radius")
	if CheckDbCon(Db) {
		println("Connexion à la BD établit")
	} else {
		println("Connexion à la BD échouée")
	}

	http.HandleFunc("/getLogin", LoginHandleFunc)
	http.ListenAndServe(":8080", nil)

}

func LoginHandleFunc(w http.ResponseWriter, r *http.Request) {
	var loginGen1 Login
	/*timer1 := time.NewTimer(5 * time.Second)
	<-timer1.C*/
	time.Sleep(5 * time.Second)
	loginGen1.GenerateLogin(8, 2, 5)
	StoreLogin(Db, loginGen1)
	fmt.Println(loginGen1)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(loginGen1.ToJSON())

}
