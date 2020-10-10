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

func ParseAvitoPrice(url string) (int, error) {
	client := &http.Client{
		Timeout: 100 * time.Second,
	}

	request, err := http.NewRequest("GET", url, nil)
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
		return 0, errors.New("Price not found")
	}
	price, err := strconv.Atoi(match[1])
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return price, nil
}
