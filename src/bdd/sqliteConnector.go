package bdd

//
import (
	"database/sql"
)

//
func getEmoji(emoji_id int64) {

	db, err := sql.Open("sqlite3", "./unicode_emojis.db")

	checkErr(err)

	rows, err := db.Query("SELECT * FROM unicode_emojis WHERE  = ?")

	checkErr(err)

	for rows.Next() {

	}

	db.Close()
}

func checkErr(err error) {

	if err != nil {

		panic(err)
	}
}
