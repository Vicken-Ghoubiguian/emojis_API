package bdd

//
import (
	"database/sql"
)

//
type sqliteConnector struct {
}

//
func (currentSqliteConnector *sqliteConnector) InitializeSqlConnector() {

}

//
func (currentSqliteConnector *sqliteConnector) getEmoji() {

	db, err := sql.Open("sqlite3", "./foo.db")

	checkErr(err)

	db.Close()
}

func checkErr(err error) {

	if err != nil {

		panic(err)
	}
}
