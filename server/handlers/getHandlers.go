package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"bitdeal.nl/database"
	"bitdeal.nl/models"

	"github.com/spf13/viper"
)

/*
GetHomePage ...
*/
func GetHomePage(w http.ResponseWriter, r *http.Request) {
	var isIphone = strings.Contains(r.Header.Get("User-Agent"), "iPhone")

	requestBody, err := json.Marshal(models.GetPricesData{
		Type:     "buy",
		Currency: "eur",
		Amount:   "420",
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response, err := http.Post("http://localhost:8003/api/getprices", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	var exchangePrices models.GetPricesResponse
	json.Unmarshal(responseData, &exchangePrices)

	// Functions for the template
	var funcMap = template.FuncMap{
		"stars": stars,
		"wrap":  wrap,
	}

	var tmpl *template.Template
	var templatesProduction []string

	templates := addTemplate("templates/pages/homepage.html")
	templates = addExtraTemplate(templates, "templates/components/homepageheaderleft.html")
	templates = addExtraTemplate(templates, "templates/components/homepageheaderright.html")
	templates = addExtraTemplate(templates, "templates/components/homepageheader.html")
	templates = addExtraTemplate(templates, "templates/components/homepagenav.html")
	templates = addExtraTemplate(templates, "templates/components/bitdeal.html")

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	if viper.GetString("environment") == "production" {
		for s := range templates {
			templatesProduction = append(templatesProduction, "/var/www/bitdeal.nl/"+templates[s])
		}
		tmpl = template.Must(template.New("homepage.html").Funcs(funcMap).ParseFiles(templatesProduction...))
	} else {
		tmpl = template.Must(template.New("homepage.html").Funcs(funcMap).ParseFiles(templates...))
	}

	data := models.HomepageData{
		Title:          "Homepage",
		ExchangePrices: exchangePrices,
		Iphone:         isIphone,
		URL:            r.RequestURI,
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

/*
GetStatsPage ...
*/
func GetStatsPage(w http.ResponseWriter, r *http.Request) {
	var funcMap = template.FuncMap{}

	var tmpl *template.Template
	var templatesProduction []string

	templates := addTemplate("templates/pages/stats.html")
	templates = addExtraTemplate(templates, "templates/components/homepagenav.html")

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	if viper.GetString("environment") == "production" {
		for s := range templates {
			templatesProduction = append(templatesProduction, "/var/www/bitdeal.nl/"+templates[s])
		}
		tmpl = template.Must(template.New("stats.html").Funcs(funcMap).ParseFiles(templatesProduction...))
	} else {
		tmpl = template.Must(template.New("stats.html").Funcs(funcMap).ParseFiles(templates...))
	}

	data := models.StatsData{
		Title: "Stats",
		URL:   r.RequestURI,
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

/*
GetInformationPage ...
*/
func GetInformationPage(w http.ResponseWriter, r *http.Request) {
	var funcMap = template.FuncMap{}

	var tmpl *template.Template
	var templatesProduction []string

	templates := addTemplate("templates/pages/information.html")
	templates = addExtraTemplate(templates, "templates/components/globalHeader.html")

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	if viper.GetString("environment") == "production" {
		for s := range templates {
			templatesProduction = append(templatesProduction, "/var/www/bitdeal.nl/"+templates[s])
		}
		tmpl = template.Must(template.New("stats.html").Funcs(funcMap).ParseFiles(templatesProduction...))
	} else {
		tmpl = template.Must(template.New("stats.html").Funcs(funcMap).ParseFiles(templates...))
	}

	data := models.StatsData{
		Title: "Stats",
		URL:   r.RequestURI,
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

/*
GetStats ...
*/
func GetStats(w http.ResponseWriter, r *http.Request) {
	output := database.GetStatistics()
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

/*
GetSunshade ...
*/
func GetSunshade(w http.ResponseWriter, r *http.Request) {
	output := database.GetSunshade()
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

/*
GetBitcoinPrices ...
*/
func GetBitcoinPrices(w http.ResponseWriter, r *http.Request) {
	database.GetBitcoinPrices()
	w.Header().Set("content-type", "application/json")
	// w.Write(output)
}
