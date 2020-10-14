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
        "url" TEXT,
        FOREIGN KEY(email) REFERENCES users(email),
        FOREIGN KEY(url) REFERENCES urls(url),
        CONSTRAINT unic_rows UNIQUE(email, url)
    );`

    statement, err := r.DB.Prepare(query)
    if err != nil {
        log.Fatal(err.Error())
    }
    statement.Exec()
    log.Println("SubscriptionsTableMigration up")
}

func (r *SubscriptionsTableMigration) Down() {
    query := `DROP TABLE subscriptions;`
    statement, err := r.DB.Prepare(query)
    if err != nil {
        log.Fatal(err.Error())
    }
    statement.Exec()
    log.Println("SubscriptionsTableMigration down")
}
