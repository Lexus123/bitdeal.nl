package database

import (
	"database/sql"
	"encoding/json"
	"math"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

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

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
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

	responsetimes.Average = responsetimes.Average / int64(len(responsetimes.Responsetimes))

	for index, response := range responsetimes.Responsetimes {
		responsetimes.Responsetimes[index].Difference = math.Pow(float64(response.Requesttime-responsetimes.Average), 2)
		responsetimes.Variance += responsetimes.Responsetimes[index].Difference
	}

	responsetimes.Variance = responsetimes.Variance / float64(len(responsetimes.Responsetimes))
	responsetimes.Sigma = math.Sqrt(responsetimes.Variance)

	output, err := json.Marshal(responsetimes)
	if err != nil {
		panic(err.Error())
	}

	return output
}
