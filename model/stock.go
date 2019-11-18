package model

import (
	"github.com/jinzhu/gorm"
	u "github.com/lukaszozimek/stock-exchange-service/util"
)

type Stock struct {
	gorm.Model
	Name      string `json:"name"`
	Shortcut  string `json:"shortcut"`
	CompanyId uint   `json:"company_id"`
}

func (stock *Stock) Validate() (map[string]interface{}, bool) {

	if stock.Name == "" {
		return u.Message(false, "Stock name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (stock *Stock) Create() map[string]interface{} {

	if resp, ok := stock.Validate(); !ok {
		return resp
	}

	GetDB().Create(stock)

	resp := u.Message(true, "success")
	resp["stock"] = stock
	return resp
}

func GetStock(id uint) *Stock {

	stock := &Stock{}
	err := GetDB().Table("stock").Where("id = ?", id).First(stock).Error
	if err != nil {
		return nil
	}
	return stock
}
