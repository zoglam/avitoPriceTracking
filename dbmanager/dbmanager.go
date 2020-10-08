package dbmanager

import (
	"database/sql"
	"log"
	"os"

	"../migrations"
)

func AddUrl(db *sql.DB, url string, price int) {
	query := `INSERT or IGNORE INTO urls(url, price) VALUES (?, ?)`
	statement, err := db.Prepare(query)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(url, price)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func GetUrlID(db *sql.DB, url string) int {
	query := `SELECT id FROM urls WHERE url = ?`
	row, err := db.Query(query, url)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer row.Close()
	for row.Next() {
		var ID int
		row.Scan(&ID)
		return ID
	}
	return 0
}

func AddSubscription(db *sql.DB, email string, urlID int) {
	query := `INSERT or IGNORE INTO subscriptions(email, id_url) VALUES (?, ?)`
	statement, err := db.Prepare(query)

	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(email, urlID)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func DbReset() {
	os.Remove("sqlite3.db")
	file, err := os.Create("sqlite3.db")
	log.Println("sqlite3.db creating...")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite3.db created")

	sqliteDatabase := OpenDbConnection("sqlite3.db")
	defer CloseDbConnection(sqliteDatabase)

	urlsTable := migrations.UrlsTableMigration{Db: sqliteDatabase}
	urlsTable.Up()

	subscriptionsTable := migrations.SubscriptionsTableMigration{Db: sqliteDatabase}
	subscriptionsTable.Up()
}

func OpenDbConnection(filename string) *sql.DB {
	sqliteDatabase, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	return sqliteDatabase
}

func CloseDbConnection(db *sql.DB) {
	db.Close()
}
