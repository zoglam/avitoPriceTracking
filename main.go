package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "time"

    "./dbmanager"
    "./mailmanager"
    "./parse"

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
            newPrice, err := parse.GetAdsPrice(url)
            if err != nil {
                log.Println(url, err)
            } else {
                if newPrice != price {
                    log.Printf("New price for %s is %d -> %d\n", url, price, newPrice)
                    dbmanager.UpdateUrl(db, url, newPrice)
                    emails := dbmanager.GetEmailsByUrl(db, url)
                    for _, email := range emails {
                        status := dbmanager.GetUserStatus(db, email)
                        if status == 1 {
                            u.OutChan <- []string{
                                url,
                                strconv.Itoa(price),
                                strconv.Itoa(newPrice),
                                email,
                            }
                        }
                    }
                }
            }
            time.Sleep(2 * time.Second)
        }
        time.Sleep(1 * time.Second)
    }
}

func (u *UrlQueue) sendNotifications() {
    log.Printf("sendNotifications started")
    for item := range u.OutChan {
        url, price, newPrice, email := item[0], item[1], item[2], item[3]
        body := fmt.Sprintf("New price for %s is %s -> %s", url, price, newPrice)
        mailmanager.SendMessage(body, email)
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

    price, err := parse.GetAdsPrice(url)
    if err != nil {
        log.Println(err)
    } else {
        dbmanager.AddUrl(db, url, price)
        s := dbmanager.AddUser(db, email)
        if s == 1 {
            hash := dbmanager.GetUserHash(db, email)
            body := fmt.Sprintf("http://%s:%s/activate_email/?hash=%s", ipAddress, portNumber, hash)
            mailmanager.SendMessage(body, email)
        }
        dbmanager.AddSubscription(db, email, url)
        log.Printf("New value for %s - %s - %d", url, email, price)
    }

    http.Redirect(w, r, "/", 302)
}

func handlerActivateEmail(w http.ResponseWriter, r *http.Request) {
    db := dbmanager.OpenDbConnection("sqlite3.db")
    defer dbmanager.CloseDbConnection(db)

    query := r.URL.Query()
    hash, _ := query["hash"]
    dbmanager.UserActivateEmail(db, hash[0])
    fmt.Fprintln(w, "email activated")
}

var ipAddress = os.Getenv("IP")
var portNumber = os.Getenv("PORT")

func main() {
    if _, err := os.Stat("sqlite3.db"); err != nil {
        log.Println("sqlite3 reset")
        dbmanager.DbReset()
    }

    if ipAddress == "" {
        ipAddress = "95.165.148.222"
    }
    if portNumber == "" {
        portNumber = "8081"
    }

    u := UrlQueue{
        OutChan: make(chan []string, 10),
    }
    go u.urlsChecking()
    go u.sendNotifications()

    log.Println("Server started on port:", portNumber)
    r := mux.NewRouter()
    r.HandleFunc("/", handlerInit).Methods("GET")
    r.HandleFunc("/save/", handlerSaveData).Methods("POST")
    r.HandleFunc("/activateemail/", handlerActivateEmail).Methods("GET")
    log.Fatal(http.ListenAndServe(":"+portNumber, r))
}
