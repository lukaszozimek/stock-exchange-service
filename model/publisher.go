package model

import (
	"github.com/jinzhu/gorm"
	u "github.com/lukaszozimek/stock-exchange-service/util"
)

type Publisher struct {
	gorm.Model
	Name string
}

func (publisher *Publisher) Validate() (map[string]interface{}, bool) {

	if publisher.Name == "" {
		return u.Message(false, "Publisher name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (publisher *Publisher) Create() map[string]interface{} {

	if resp, ok := publisher.Validate(); !ok {
		return resp
	}

	GetDB().Create(publisher)

	resp := u.Message(true, "success")
	resp["publisher"] = publisher
	return resp
}

func GetPublisher(id uint) *Publisher {

	publisher := &Publisher{}
	err := GetDB().Table("publisher").Where("id = ?", id).First(publisher).Error
	if err != nil {
		return nil
	}
	return publisher
}
