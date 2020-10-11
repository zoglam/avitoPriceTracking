package dbmanager

import (
    "database/sql"
    "log"
    "os"
    "strconv"

    "avitopricetracking/migrations"
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

func GetEmailsByUrl(db *sql.DB, url int) []string {
    var result []string
    query := `SELECT email FROM subscriptions WHERE id_url = ?`
    rows, err := db.Query(query, url)
    if err != nil {
        log.Fatalln(err.Error())
    }
    defer rows.Close()
    for rows.Next() {
        var email string
        rows.Scan(&email)
        result = append(result, email)
    }
    return result
}

func GetUrls(db *sql.DB) [][]string {
    query := `SELECT url, price FROM urls`
    rows, err := db.Query(query)
    if err != nil {
        log.Fatalln(err.Error())
    }
    defer rows.Close()

    var result [][]string
    for rows.Next() {
        var url string
        var price int
        rows.Scan(&url, &price)
        result = append(result, []string{url, strconv.Itoa(price)})
    }
    return result
}

func UpdateUrl(db *sql.DB, url string, newPrice int) {
    query := `UPDATE urls SET price = ? WHERE url = ?`
    statement, err := db.Prepare(query)

    if err != nil {
        log.Fatalln(err.Error())
    }
    _, err = statement.Exec(newPrice, url)
    if err != nil {
        log.Fatalln(err.Error())
    }
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
    if err != nil {
        log.Fatal(err.Error())
    }
    file.Close()
    log.Println("sqlite3.db created")

    sqliteDatabase := OpenDbConnection("sqlite3.db")
    defer CloseDbConnection(sqliteDatabase)

    urlsTable := migrations.UrlsTableMigration{DB: sqliteDatabase}
    urlsTable.Up()

    subscriptionsTable := migrations.SubscriptionsTableMigration{DB: sqliteDatabase}
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
