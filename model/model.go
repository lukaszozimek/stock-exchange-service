package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/lukaszozimek/stock-exchange-service/model/company"
	"github.com/lukaszozimek/stock-exchange-service/model/financial_instruments"
	"github.com/lukaszozimek/stock-exchange-service/model/news"
	"github.com/lukaszozimek/stock-exchange-service/model/stock_exchange"
	"github.com/lukaszozimek/stock-exchange-service/model/user"
	"os"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(
		&user.Account{},
		&user.Contact{},
		&stock_exchange.StockExchange{},
		&news.News{},
		&news.Publisher{},
		&company.Company{},
		&financial_instruments.Bonds{},
		&financial_instruments.ExchangeRate{},
		&financial_instruments.Resource{},
		&financial_instruments.Stock{},
		&financial_instruments.StockIndex{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
