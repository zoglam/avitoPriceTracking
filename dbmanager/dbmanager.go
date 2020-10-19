package dbmanager

import (
    "crypto/md5"
    "database/sql"
    "fmt"
    "log"
    "os"
    "strconv"

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

func AddUser(db *sql.DB, email string) int {
    query := `INSERT or IGNORE INTO users(email, hash, status) VALUES (?, ?, 0)`
    hash := fmt.Sprintf("%x", md5.Sum([]byte(email)))
    statement, err := db.Prepare(query)

    if err != nil {
        log.Println(err.Error())
        return 1
    }
    _, err = statement.Exec(email, hash)
    if err != nil {
        log.Println(err.Error())
        return 1
    }
    return 0
}

func GetUserStatus(db *sql.DB, email string) int {
    query := `SELECT status FROM users WHERE email = ?`
    rows, err := db.Query(query, email)
    if err != nil {
        log.Fatalln(err.Error())
    }
    defer rows.Close()
    for rows.Next() {
        var status int
        rows.Scan(&status)
        return status
    }
    return 0
}

func GetUserHash(db *sql.DB, email string) string {
    query := `SELECT hash FROM users WHERE email = ?`
    rows, err := db.Query(query, email)
    if err != nil {
        log.Fatalln(err.Error())
    }
    defer rows.Close()
    for rows.Next() {
        var hash string
        rows.Scan(&hash)
        return hash
    }
    return ""
}

func UserActivateEmail(db *sql.DB, hash string) {
    query := `UPDATE users SET status = 1 WHERE hash = ?`
    statement, err := db.Prepare(query)

    if err != nil {
        log.Fatalln(err.Error())
    }
    _, err = statement.Exec(hash)
    if err != nil {
        log.Fatalln(err.Error())
    }
}

func GetEmailsByUrl(db *sql.DB, url string) []string {
    var result []string
    query := `SELECT email FROM subscriptions WHERE url = ?`
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

func AddSubscription(db *sql.DB, email string, url string) {
    query := `INSERT or IGNORE INTO subscriptions(email, url) VALUES (?, ?)`
    statement, err := db.Prepare(query)

    if err != nil {
        log.Fatalln(err.Error())
    }
    _, err = statement.Exec(email, url)
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

    usersTable := migrations.UsersTableMigration{DB: sqliteDatabase}
    usersTable.Up()

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
