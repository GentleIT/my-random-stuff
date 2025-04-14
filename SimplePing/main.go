package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

const GoogleURL = "https://google.com"
const ExampleURL = "https://example.com"
const GoURL = "https://golang.org"

func main() {
	timeStart := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(3)

	urlSlice := []string{}
	urlSlice = append(urlSlice, GoogleURL, ExampleURL, GoURL)

	for _, r := range urlSlice {
		go GetHTTPStatus(r, &wg)
	}
	wg.Wait()
	log.Println("Program has finished in", time.Since(timeStart))
}

func GetHTTPStatus(r string, wg *sync.WaitGroup) {
	resp, err := http.Get(r)
	if err != nil {
		log.Printf("Error: %s, http: %s", err, r)
		wg.Done()
		return
	}
	log.Printf("OK - %s (%d)", r, resp.StatusCode)

	defer resp.Body.Close()
	wg.Done()
}

// То, что нужно добавить:

// Добавь таймер, чтобы проверка происходила каждые n секунд.
// for {
//     select {
//     case <-ctx.Done():
//         return
//     case <-time.After(interval):
//         // ping URL
//     }
// }

// Вывод в канал
// Пример:
// type PingResult struct {
//     URL        string
//     StatusCode int
//     Error      error
//     Timestamp  time.Time
// }
// resultCh := make(chan PingResult)
