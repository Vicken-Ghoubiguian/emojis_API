package bdd

//
import (
	"database/sql"
)

//
func getEmojiFromItsId(emoji_id int64) {

	db, err := sql.Open("sqlite3", "./unicode_emojis.db")
	checkErr(err)

	stmt, err := db.Prepare("SELECT * FROM unicode_emojis WHERE number = ?")
	checkErr(err)

	rows, err := stmt.Exec("astaxieupdate", emoji_id)
	checkErr(err)

	print(rows)
	checkErr(err)

	db.Close()
}

func checkErr(err error) {

	if err != nil {

		panic(err)
	}
}
