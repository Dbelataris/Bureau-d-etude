package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	. "Bureau-d-etude/tools"

	_ "Bureau-d-etude/go-sql-driver"
)

var Db *sql.DB
var Ok bool
var TEx int   //Temps Expiration
var TUtil int //Temps utilisation
var StrTimes []string

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
	readLine("./parameters")
	fmt.Println("Tex:" + StrTimes[0] + "	Tutil:" + StrTimes[1])
	TEx, _ = strconv.Atoi(StrTimes[0])

	TUtil, _ = strconv.Atoi(StrTimes[1])

	http.HandleFunc("/getLogin", LoginHandleFunc)
	http.HandleFunc("/delLogin", DeleteHandleFunc)
	http.ListenAndServe(":8080", nil)

}

func LoginHandleFunc(w http.ResponseWriter, r *http.Request) {
	var loginGen1 Login

	if Ok == true {
		loginGen1.GenerateLogin(8, TEx, TUtil)
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

func DeleteHandleFunc(w http.ResponseWriter, r *http.Request) {
	param, ok := r.URL.Query()["username"]

	if !ok || len(param) < 1 {
		log.Println("Url Param 'username' est manquant")
		return
	} else {
		DeleteLogin(Db, param[0])
	}

}

func readLine(path string) {
	var line []string
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		//fmt.Println(scanner.Text() + "***")
		line = strings.Split(scanner.Text(), " ")
		/*fmt.Println(Line[2])*/
		StrTimes = append(StrTimes, line[2])

	}
}
