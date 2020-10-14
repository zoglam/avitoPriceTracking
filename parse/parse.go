package parse

import (
    "errors"
    "io/ioutil"
    "log"
    "net/http"
    "regexp"
    "strconv"
    "time"
)

// GetAdsPrice returns price from ad
func GetAdsPrice(urlName string) (int, error) {

    // transport := &http.Transport{
    //     Proxy: http.ProxyURL(&url.URL{
    //         Scheme: "http",
    //         User:   url.UserPassword("user", "password"),
    //         Host:   "host:port",
    //     }),
    // }

    client := &http.Client{
        Timeout: 20 * time.Second,
        // Transport: transport,
    }

    request, err := http.NewRequest("GET", urlName, nil)
    if err != nil {
        log.Println(err.Error())
        return 0, err
    }

    resp, err := client.Do(request)
    if err != nil {
        log.Println(err.Error())
        return 0, err
    }
    defer resp.Body.Close()

    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println(err.Error())
        return 0, err
    }

    var patt = regexp.MustCompile(`itemprop="price" content="(\d+)"`)
    match := patt.FindStringSubmatch(string(responseData))
    if len(match) < 2 {
        // file, _ := os.Create("dwa.html")
        // defer file.Close()
        // file.WriteString(string(responseData))
        return 0, errors.New("Price not found")
    }

    price, err := strconv.Atoi(match[1])
    if err != nil {
        log.Println(err.Error())
        return 0, err
    }
    return price, nil
}
