package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

/*
SaveResponseTime ...
*/
func SaveResponseTime(responseTime int64) {
	fmt.Println("Go MySQL Tutorial")

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

	dbname := viper.GetString("database.dbname")
	dbuser := viper.GetString("database.dbuser")
	dbpassword := viper.GetString("database.dbpassword")
	db, err := sql.Open("mysql", dbuser+":"+dbpassword+"@tcp(127.0.0.1:3306)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}
