package tools

import (
	"database/sql"
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
