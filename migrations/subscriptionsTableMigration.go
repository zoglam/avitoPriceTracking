package migrations

import (
    "database/sql"
    "log"
)

type SubscriptionsTableMigration struct {
    DB *sql.DB
}

func (r *SubscriptionsTableMigration) Up() {
    query := `CREATE TABLE subscriptions (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "email" TEXT,
        "id_url" INTEGER,
        FOREIGN KEY(id_url) REFERENCES urls(id),
        CONSTRAINT unic_rows UNIQUE(email, id_url));`

    statement, err := r.DB.Prepare(query)
    if err != nil {
        log.Fatal(err.Error())
    }
    statement.Exec()
    log.Println("UserTableMigration up")
}

func (r *SubscriptionsTableMigration) Down() {
    query := `DROP TABLE subscriptions;`
    statement, err := r.DB.Prepare(query)
    if err != nil {
        log.Fatal(err.Error())
    }
    statement.Exec()
    log.Println("UserTableMigration down")
}
