package main

import (
	"log"
	"net/http"
	"os"

	"./dbmanager"
	_ "github.com/mattn/go-sqlite3"
)

func handlerInit(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handlerSaveData(w http.ResponseWriter, r *http.Request) {
	url, email := r.FormValue("url"), r.FormValue("email")

	db := dbmanager.OpenDbConnection("sqlite3.db")
	defer dbmanager.CloseDbConnection(db)

	dbmanager.AddUrl(db, url, 0)
	urlID := dbmanager.GetUrlID(db, url)
	dbmanager.AddSubscription(db, email, urlID)
}

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "--reset" {
			dbmanager.DbReset()
		}
	}

	// resp, err := http.Get("https://m.avito.ru/moskva/tovary_dlya_zhivotnyh/kofta_dlya_koshki_ili_horka_2008986079?s=22")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer resp.Body.Close()
	// for true {

	// 	bs := make([]byte, 1014)
	// 	n, err := resp.Body.Read(bs)
	// 	fmt.Println(string(bs[:n]))

	// 	if n == 0 || err != nil {
	// 		break
	// 	}
	// }

	// url := "https://m.avito.ru/moskva/tovary_dlya_zhivotnyh/kofta_dlya_koshki_ili_horka_2008986079"
	// client := &http.Client{}
	// req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	// req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 OPR/71.0.3770.228")
	// req.Header.Set("referer", "https://www.avito.ru/rossiya")
	// req.Header.Set("connection", "keep-alive")
	// req.Header.Set("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,lt;q=0.6")
	// // req.Header.Set("server", "server")
	// // req.Header.Set("server", "server")
	// // req.Header.Set("server", "server")
	// // req.Header.Set("server", "server")
	// // req.Header.Set("server", "server")
	// // req.Header.Set("server", "server")
	// res, _ := client.Do(req)
	// for true {

	// 	bs := make([]byte, 1014)
	// 	n, err := res.Body.Read(bs)
	// 	fmt.Println(string(bs[:n]))

	// 	if n == 0 || err != nil {
	// 		break
	// 	}
	// }

	// resp, err := http.Get("https://m.avito.ru/moskva/tovary_dlya_zhivotnyh/kofta_dlya_koshki_ili_horka_2008986079")
	// if err != nil {
	// 	panic(err)
	// }

	// if res.StatusCode == 403 {
	// 	log.Fatal("FATAL!!! - ", res.StatusCode)
	// }

	http.HandleFunc("/", handlerInit)
	http.HandleFunc("/save/", handlerSaveData)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
