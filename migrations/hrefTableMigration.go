package migrations

import (
	"database/sql"
	"log"
)

type UrlsTableMigration struct {
	Db *sql.DB
}

func (r *UrlsTableMigration) Up() {
	query := `CREATE TABLE urls (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"url" TEXT UNIQUE,
		"price" NUMERIC);`

	statement, err := r.Db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("UrlsTableMigration up")
}

func (r *UrlsTableMigration) Down() {
	query := `DROP TABLE urls;`
	statement, err := r.Db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("UrlsTableMigration down")

}
