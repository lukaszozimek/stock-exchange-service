package model

import (
	"github.com/jinzhu/gorm"
	u "github.com/lukaszozimek/stock-exchange-service/util"
)

type StockIndex struct {
	gorm.Model
	Name            string `json:"name"`
	Shortcut        string `json:"shortcut"`
	StockExchangeId uint   `json:"stock_exchange_id"`
}

func (stockIndex *StockIndex) Validate() (map[string]interface{}, bool) {

	if stockIndex.Name == "" {
		return u.Message(false, "Stock Index name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (stockIndex *StockIndex) Create() map[string]interface{} {

	if resp, ok := stockIndex.Validate(); !ok {
		return resp
	}

	GetDB().Create(stockIndex)
	resp := u.Message(true, "success")
	resp["stockIndex"] = stockIndex
	return resp
}

func GetStockIndex(id uint) *StockIndex {

	stockIndex := &StockIndex{}
	err := GetDB().Table("stock_index").Where("id = ?", id).First(stockIndex).Error
	if err != nil {
		return nil
	}
	return stockIndex
}
