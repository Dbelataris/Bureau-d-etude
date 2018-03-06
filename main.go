package main

import (
	"database/sql"
	"fmt"
	"net/http"

	. "Bureau-d-etude/tools"

	_ "Bureau-d-etude/github.com/go-sql-driver/mysql"
)

func main() {
	/*again := true
	var LoginsValides = []Login{}
	var loginGen Login*/
	db, _ := sql.Open("mysql", "radius:radpass@tcp(127.0.0.1:3306)/radius")

	if CheckDbCon(db) {
		println("Connexion à la BD établit")
	} else {
		println("Connexion à la BD échouée")
	}
	/*for again {
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
	}*/
	http.HandleFunc("/getLogin", LoginHandleFunc)
	http.ListenAndServe(":8080", nil)

}

func LoginHandleFunc(w http.ResponseWriter, r *http.Request) {
	var loginGen Login
	loginGen.GenerateLogin(8, 5)
	//StoreLogin(db, loginGen)
	fmt.Println(loginGen)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(loginGen.ToJSON())
	fmt.Println("ok")
}
