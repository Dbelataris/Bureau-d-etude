package tools

import (
	"database/sql"
	"fmt"
	"log"
)

func CheckDbCon(db *sql.DB) bool {
	err := db.Ping()
	var dbStatus bool

	if err == nil {
		dbStatus = true
	} else {
		dbStatus = false
	}
	return dbStatus
}

func StoreLogin(db *sql.DB, user Login) {
	stmt, err := db.Prepare("INSERT INTO radcheck (username, attribute, op, value, time_Gen,TempsExpiration, useTime) VALUES (?, 'Cleartext-Password', ':=', ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Username, user.Password, user.StartTime, user.ExpirationTime, user.UseTime)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

}

func DeleteLogin(db *sql.DB, user string) {
	stmt, err := db.Prepare("DELETE FROM radcheck WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}
