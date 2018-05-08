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
var Ok bool

func main() {
	/*var LoginsValides = []Login{}
	  var loginGen Login*/
	Db, _ = sql.Open("mysql", "root:isib@tcp(127.0.0.1:3306)/radius")
	if CheckDbCon(Db) {
		println("Connexion à la BD établit")
	} else {
		println("Connexion à la BD échouée")
	}
	defer Db.Close()
	Ok = true
	http.HandleFunc("/getLogin", LoginHandleFunc)
	http.ListenAndServe(":8080", nil)

}

func LoginHandleFunc(w http.ResponseWriter, r *http.Request) {
	var loginGen1 Login

	if Ok == true {
		loginGen1.GenerateLogin(8, 2, 5)
		StoreLogin(Db, loginGen1)
		fmt.Println(loginGen1)
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Write(loginGen1.ToJSON())
		Ok = false
		timer := time.NewTimer(3 * time.Second)
		go func() {
			<-timer.C
			/*fmt.Println("Timer  expired")*/
			Ok = true
		}()
	} /*else {
		fmt.Println("Stop")
	}*/

}
