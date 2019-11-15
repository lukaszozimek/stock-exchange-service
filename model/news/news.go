package news

import (
	"github.com/jinzhu/gorm"
	"github.com/lukaszozimek/stock-exchange-service/model"
	u "github.com/lukaszozimek/stock-exchange-service/util"
)

type News struct {
	gorm.Model
	HeadContent string `json:"head_content"`
	PublisherId uint   `json:"publisher_id"`
}

func (news *News) Validate() (map[string]interface{}, bool) {

	if news.HeadContent == "" {
		return u.Message(false, "Stock Exchange name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (news *News) Create() map[string]interface{} {

	if resp, ok := news.Validate(); !ok {
		return resp
	}

	model.GetDB().Create(news)

	resp := u.Message(true, "success")
	resp["news"] = news
	return resp
}

func GetNews(id uint) *News {

	stockExchange := &News{}
	err := model.GetDB().Table("news").Where("id = ?", id).First(stockExchange).Error
	if err != nil {
		return nil
	}
	return stockExchange
}
