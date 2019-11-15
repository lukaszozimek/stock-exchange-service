package stock_exchange

import (
	"github.com/jinzhu/gorm"
	"github.com/lukaszozimek/stock-exchange-service/model"
	u "github.com/lukaszozimek/stock-exchange-service/util"
)

type StockExchange struct {
	gorm.Model
	Name string `json:"name"`
}

func (stockExchange *StockExchange) Validate() (map[string]interface{}, bool) {

	if stockExchange.Name == "" {
		return u.Message(false, "Stock Exchange name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (stockExchange *StockExchange) Create() map[string]interface{} {

	if resp, ok := stockExchange.Validate(); !ok {
		return resp
	}

	model.GetDB().Create(stockExchange)

	resp := u.Message(true, "success")
	resp["stockExchange"] = stockExchange
	return resp
}

func GetStockExchange(id uint) *StockExchange {

	stockExchange := &StockExchange{}
	err := model.GetDB().Table("stock_exchange").Where("id = ?", id).First(stockExchange).Error
	if err != nil {
		return nil
	}
	return stockExchange
}
