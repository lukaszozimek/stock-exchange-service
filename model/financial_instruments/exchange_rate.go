package financial_instruments

import (
	"github.com/jinzhu/gorm"
	"github.com/lukaszozimek/stock-exchange-service/model"
	u "github.com/lukaszozimek/stock-exchange-service/util"
)

type ExchangeRate struct {
	gorm.Model
	Name         string `json:"name"`
	FromCurrency string `json:"FromCurrency"`
	ToCurrency   string `json:"ToCurrency"`
}

func (exchangeRate *ExchangeRate) Validate() (map[string]interface{}, bool) {

	if exchangeRate.Name == "" {
		return u.Message(false, "Exchange Rate name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (exchangeRate *ExchangeRate) Create() map[string]interface{} {

	if resp, ok := exchangeRate.Validate(); !ok {
		return resp
	}

	model.GetDB().Create(exchangeRate)

	resp := u.Message(true, "success")
	resp["exchangeRate"] = exchangeRate
	return resp
}

func GetExchangeRate(id uint) *ExchangeRate {

	ExchangeRate := &ExchangeRate{}
	err := model.GetDB().Table("exchange_rate").Where("id = ?", id).First(ExchangeRate).Error
	if err != nil {
		return nil
	}
	return ExchangeRate
}
