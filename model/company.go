package model

import (
	"github.com/jinzhu/gorm"
	u "github.com/lukaszozimek/stock-exchange-service/util"
)

type Company struct {
	gorm.Model
	Name                string `json:"name"`
	Website             string `json:"website"`
	StockExchangeSymbol string `json:"stock_exchange_symbol"`
}

func (company *Company) Validate() (map[string]interface{}, bool) {

	if company.Name == "" {
		return u.Message(false, "Company name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (company *Company) Create() map[string]interface{} {

	if resp, ok := company.Validate(); !ok {
		return resp
	}

	GetDB().Create(company)

	resp := u.Message(true, "success")
	resp["company"] = company
	return resp
}

func GetCompany(id uint) *Company {

	company := &Company{}
	err := GetDB().Table("company").Where("id = ?", id).First(company).Error
	if err != nil {
		return nil
	}
	return company
}
