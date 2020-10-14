package migrations

import (
	"database/sql"
	"log"
)

type UsersTableMigration struct {
	DB *sql.DB
}

func (r *UsersTableMigration) Up() {
	query := `CREATE TABLE users (
        "email" TEXT PRIMARY KEY,
		"hash" TEXT,
		"status" INTEGER
	);`

	statement, err := r.DB.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("UsersTableMigration up")
}

func (r *UsersTableMigration) Down() {
	query := `DROP TABLE users;`
	statement, err := r.DB.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("UsersTableMigration down")
}
