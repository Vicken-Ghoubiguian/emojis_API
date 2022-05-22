package bdd

//
import (
	"database/sql"
)

//
func getEmoji() {

	db, err := sql.Open("sqlite3", "./unicode_emojis.db")

	checkErr(err)

	db.Close()
}

func checkErr(err error) {

	if err != nil {

		panic(err)
	}
}
