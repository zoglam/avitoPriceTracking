package main

import (
    "log"
    "net/http"
    "os"
    "strconv"
    "time"

    "avitopricetracking/dbmanager"
    "avitopricetracking/mailmanager"
    "avitopricetracking/parse"
    "github.com/gorilla/mux"
    _ "github.com/mattn/go-sqlite3"
)

type UrlQueue struct {
    OutChan chan []string
}

func (u *UrlQueue) urlsChecking() {
    db := dbmanager.OpenDbConnection("sqlite3.db")
    defer dbmanager.CloseDbConnection(db)
    log.Printf("urlsChecking started")

    for {
        urls := dbmanager.GetUrls(db)
        for _, item := range urls {
            url := item[0]
            price, _ := strconv.Atoi(item[1])
            newPrice, err := parse.ParseAvitoPrice(url)
            if err != nil {
                log.Println(url, err, "- Запись будет удалена мб")
            } else {
                if newPrice != price {
                    log.Println(url, "| new price:", price, "->", newPrice)
                    urlID := dbmanager.GetUrlID(db, url)
                    emails := dbmanager.GetEmailsByUrl(db, urlID)
                    for _, email := range emails {
                        dbmanager.UpdateUrl(db, url, newPrice)
                        u.OutChan <- []string{
                            url,
                            strconv.Itoa(price),
                            strconv.Itoa(newPrice),
                            email,
                        }
                    }
                }
            }
            time.Sleep(7 * time.Second)
        }
        time.Sleep(30 * time.Second)
    }
}

func (u *UrlQueue) sendNotifications() {
    for item := range u.OutChan {
        url, price, newPrice, email := item[0], item[1], item[2], item[3]
        mailmanager.SendMessage(url, price, newPrice, email)
    }
}

func handlerInit(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "templates/index.html")
}

func handlerSaveData(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    url := r.PostForm.Get("url")
    email := r.PostForm.Get("email")

    db := dbmanager.OpenDbConnection("sqlite3.db")
    defer dbmanager.CloseDbConnection(db)

    price, err := parse.ParseAvitoPrice(url)
    if err != nil {
        log.Println(err)
    } else {
        dbmanager.AddUrl(db, url, price)
        urlID := dbmanager.GetUrlID(db, url)
        dbmanager.AddSubscription(db, email, urlID)
        log.Printf("New value: %s - %s", url, email)
    }

    http.Redirect(w, r, "/", 302)
}

func main() {
    args := os.Args[1:]
    for i := 0; i < len(args); i++ {
        if args[i] == "--reset" {
            dbmanager.DbReset()
        }
    }

    portNumber := os.Getenv("PORT")
    u := UrlQueue{
        OutChan: make(chan []string, 10),
    }
    go u.urlsChecking()
    go u.sendNotifications()

    log.Println("Server started on port:", portNumber)
    r := mux.NewRouter()
    r.HandleFunc("/", handlerInit)
    r.HandleFunc("/save/", handlerSaveData).Methods("POST")
    log.Fatal(http.ListenAndServe("localhost:"+portNumber, r))
}
