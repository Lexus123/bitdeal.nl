package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"bitdeal.nl/etl"
	"bitdeal.nl/models"
)

/*
GetHomePage ...
*/
func GetHomePage(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}

	requestBody, err := json.Marshal(models.GetPricesData{
		Type:     "buy",
		Currency: "eur",
		Amount:   420,
	})

	request, err := http.NewRequest("GET", "http://localhost:8003/api/getprices", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	var exchangePrices models.GetPricesResponse
	json.Unmarshal(responseData, &exchangePrices)

	// w.Header().Set("content-type", "application/json")
	// w.Write(output)

	// Functions for the template
	var funcMap = template.FuncMap{
		"bestRateToBTC": bestRateToBTC,
		"wrap":          wrap,
	}

	templates := addTemplate("templates/pages/homepage.html")
	tmpl := template.Must(template.New("homepage.html").Funcs(funcMap).ParseFiles(templates...))
	data := models.HomepageData{
		Title:          "Homepage",
		ExchangePrices: exchangePrices,
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}

/*
GetPrices ...
*/
func GetPrices(w http.ResponseWriter, r *http.Request) {
	log.Printf("1")
	// Read JSON body of request and check for errors. For example: buy, eur, 100
	bodyAsByte, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Printf("2")

	// Unmarshal the body and check for errors
	var message models.GetPricesData
	err = json.Unmarshal(bodyAsByte, &message)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Printf("3")

	// Send message to the extractor to get data from brokers
	dataChannel := make(chan interface{}, 5)
	defer close(dataChannel)
	go etl.DoRequests(message, dataChannel)

	log.Printf("4")
	getPricesResponse := models.GetPricesResponse{
		Type:     message.Type,
		Currency: message.Currency,
	}

	for i := 0; i < 5; i++ {
		getPricesResponse.ExchangeRates = append(getPricesResponse.ExchangeRates, etl.Transform(message, <-dataChannel))
	}
	log.Printf("5")

	sort.Slice(getPricesResponse.ExchangeRates[:], func(i, j int) bool {
		return getPricesResponse.ExchangeRates[i].Rate < getPricesResponse.ExchangeRates[j].Rate
	})
	log.Printf("5")

	getPricesResponse.BestRate = getPricesResponse.ExchangeRates[0].Rate
	getPricesResponse.MostReviews = 5485

	log.Printf("7")
	output, err := json.Marshal(getPricesResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Printf("8")

	w.Header().Set("content-type", "application/json")
	w.Write(output)
	log.Printf("9")
}

/*
GetSignInPage ...
*/
// func GetSignInPage(w http.ResponseWriter, r *http.Request) {
// 	// Check whether the client already has a cookie
// 	if checkCookie(r) {
// 		// You have a cookie, then go to homepage
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 	} else {
// 		// Functions for the template
// 		var funcMap = template.FuncMap{
// 			"trimCaptcha":     trimCaptcha,
// 			"printErrorParam": printErrorParam,
// 		}

// 		// Setting up the template
// 		templates := addTemplate("templates/pages/auth/signin.gohtml", false)
// 		tmpl := template.Must(template.New("signin.gohtml").Funcs(funcMap).ParseFiles(templates...))

// 		// Configuration struct for digits captcha
// 		var config = base64Captcha.ConfigDigit{
// 			Height:     100,
// 			Width:      300,
// 			MaxSkew:    0.7,
// 			DotCount:   80,
// 			CaptchaLen: 5,
// 		}

// 		// Creation of the captcha base64 image and the answer (idKey)
// 		idKey, cap := base64Captcha.GenerateCaptcha("", config)
// 		base64string := base64Captcha.CaptchaWriteToBase64Encoding(cap)

// 		// Values that need to be stored in the cookie
// 		cookieValues := map[string]string{
// 			"captchaID": idKey,
// 		}

// 		// Creation of the actual cookie for signing in/up
// 		if encodedValues, err := secCookie.Encode("websitecaptcha", cookieValues); err == nil {
// 			cookie := &http.Cookie{
// 				Name:     "websitecaptcha",
// 				Value:    encodedValues,
// 				Path:     "/",
// 				HttpOnly: true,
// 				Expires:  time.Now().Add(5 * time.Minute),
// 			}
// 			http.SetCookie(w, cookie)
// 		}

// 		// If the URL has parameters, put them inside this var. Example "/signin?error=username,count"
// 		var params []string

// 		// Error parameter present, then do something
// 		if errorParam := r.URL.Query()["error"]; len(errorParam) > 0 {
// 			params = strings.Split(r.URL.Query()["error"][0], ",")
// 		}

// 		// Data for the template
// 		templateData := models.IndexData{
// 			Title:   "Sign in",
// 			Captcha: base64string,
// 			Params:  params,
// 		}

// 		// Execute the template
// 		tmpl.ExecuteTemplate(w, "layout", templateData)
// 	}
// }

// /*
// GetHomePage ...
// */
// func GetHomePage(w http.ResponseWriter, r *http.Request) {
// 	if checkCookie(r) {
// 		templates := addTemplate("templates/pages/homepage.gohtml", true)
// 		tmpl := template.Must(template.ParseFiles(templates...))

// 		data := models.IndexData{
// 			Title: "Homepage",
// 		}

// 		tmpl.ExecuteTemplate(w, "layout", data)
// 	} else {
// 		http.Redirect(w, r, "/signin", http.StatusSeeOther)
// 	}
// }

// /*
// GetTest ...
// */
// func GetTest(w http.ResponseWriter, r *http.Request) {
// 	if checkCookie(r) {
// 		fmt.Fprintf(w, "Salty: %v\n", security.CreateSalt())
// 		fmt.Fprintf(w, "'Salt + Hoi' SHA-3 512: %v\n", security.HashPassword(security.CreateSalt(), "hoi"))
// 	} else {
// 		http.Redirect(w, r, "/signin", http.StatusSeeOther)
// 	}
// }
