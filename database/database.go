package database

import (
	"database/sql"
	"encoding/json"
	"math"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gonum.org/v1/gonum/stat/distuv"

	"bitdeal.nl/models"
)

func openDatabaseConnection() *sql.DB {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	dbname := viper.GetString("database.dbname")
	dbuser := viper.GetString("database.dbuser")
	dbpassword := viper.GetString("database.dbpassword")
	db, err := sql.Open("mysql", dbuser+":"+dbpassword+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func toFixed(value float64) float64 {
	return math.Round(value*100) / 100
}

/*
SaveResponseTime ...
*/
func SaveResponseTime(responseTime, created int64) {
	db := openDatabaseConnection()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO responsetimes(requesttime, created) VALUES(?,?)")

	if err != nil {
		panic(err.Error())
	}

	insert.Exec(responseTime, created)

	defer insert.Close()
}

/*
GetStatistics ...
*/
func GetStatistics() []byte {
	db := openDatabaseConnection()
	defer db.Close()

	var responsetimes models.Responsetimes

	results, err := db.Query("SELECT * FROM responsetimes")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var responsetime models.Responsetime
		err = results.Scan(&responsetime.ID, &responsetime.Requesttime, &responsetime.Created)
		if err != nil {
			panic(err.Error())
		}

		responsetimes.Responsetimes = append(responsetimes.Responsetimes, responsetime)
		responsetimes.Average += responsetime.Requesttime
	}

	responsetimes.Count = int64(len(responsetimes.Responsetimes))
	responsetimes.Average = responsetimes.Average / responsetimes.Count

	for index, response := range responsetimes.Responsetimes {
		responsetimes.Responsetimes[index].Difference = math.Pow(float64(response.Requesttime-responsetimes.Average), 2)
		responsetimes.Variance += responsetimes.Responsetimes[index].Difference
	}

	responsetimes.Variance = responsetimes.Variance / float64(responsetimes.Count)
	responsetimes.Sigma = math.Sqrt(responsetimes.Variance)
	normalDist := distuv.Normal{
		Mu:    float64(responsetimes.Average),
		Sigma: responsetimes.Sigma,
	}

	for index, response := range responsetimes.Responsetimes {
		responsetimes.Responsetimes[index].Zscore = toFixed((float64(response.Requesttime) - float64(responsetimes.Average)) / responsetimes.Sigma)
		responsetimes.Responsetimes[index].CDF = toFixed(normalDist.CDF(float64(responsetimes.Responsetimes[index].Requesttime)))
	}

	output, err := json.Marshal(responsetimes)
	if err != nil {
		panic(err.Error())
	}

	return output
}

/*
GetBitcoinPrices ...
*/
func GetBitcoinPrices() {
	//
}
