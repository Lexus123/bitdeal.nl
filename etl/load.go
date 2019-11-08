package etl

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// type Bitvavo struct {
// 	Market string
// 	Nonce  int
// 	Bids   [][]interface{}
// 	Asks   [][]interface{}
// }

// var urls = map[string]string{
// 	// "https://api.satos.nl/v1/public/getDepthRate": "Post",
// 	"https://api.bitvavo.com/v2/BTC-EUR/book": "Get",
// 	// "https://twitter.com":                         "Get",
// }

// func getStations(body []byte) (*Bitvavo, error) {
// 	var s = new(Bitvavo)
// 	err := json.Unmarshal(body, &s)
// 	if err != nil {
// 		fmt.Println("whoops:", err)
// 	}
// 	return s, err
// }

// func fetch(url string, method string, wg *sync.WaitGroup) (string, error) {
// 	var resp *http.Response
// 	var err error

// 	if method == "Get" {
// 		resp, err = http.Get(url)
// 	} else {
// 		resp, err = http.Post(url, "application/json; charset=utf-8", nil)
// 	}

// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}

// 	defer resp.Body.Close()

// 	respData, err := ioutil.ReadAll(resp.Body)
// 	s, err := getStations([]byte(respData))

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(s.Bids[0])

// 	wg.Done()
// 	fmt.Println(resp.Status)
// 	return resp.Status, nil
// }

// func doRequests() {
// 	fmt.Println("Getting all the prices")

// 	var wg sync.WaitGroup

// 	for url, method := range urls {
// 		wg.Add(1)
// 		go fetch(url, method, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Returning Response")
// }

// /*
// StartRequesting ...
// */
// func StartRequesting() {
// 	for {
// 		<-time.After(2 * time.Second)
// 		go doRequests()
// 	}
// }
