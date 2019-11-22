package database

// import (
// 	"database/sql"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/spf13/viper"
// )

// /*
// SaveResponseTime ...
// */
// func SaveResponseTime(responseTime int64) {
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath(".")
// 	viper.AutomaticEnv()
// 	viper.SetConfigType("yml")

// 	dbname := viper.GetString("database.dbname")
// 	dbuser := viper.GetString("database.dbuser")
// 	dbpassword := viper.GetString("database.dbpassword")
// 	db, err := sql.Open("mysql", dbuser+":"+dbpassword+"@tcp(127.0.0.1:3306)/"+dbname)

// 	// if there is an error opening the connection, handle it
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// defer the close till after the main function has finished
// 	defer db.Close()

// 	// perform a db.Query insert
// 	insert, err := db.Prepare("INSERT INTO responsetimes(requesttime) VALUES(?)")

// 	// if there is an error inserting, handle it
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	insert.Exec(responseTime)

// 	// be careful deferring Queries if you are using transactions
// 	defer insert.Close()
// }
