package model

import (
	"github.com/jinzhu/gorm"
	u "github.com/lukaszozimek/stock-exchange-service/util"
)

type Resource struct {
	gorm.Model
	Name     string `json:"name"`
	Shortcut string `json:"shortcut"`
}

func (resource *Resource) Validate() (map[string]interface{}, bool) {

	if resource.Name == "" {
		return u.Message(false, "Resource name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (resource *Resource) Create() map[string]interface{} {

	if resp, ok := resource.Validate(); !ok {
		return resp
	}

	GetDB().Create(resource)

	resp := u.Message(true, "success")
	resp["resource"] = resource
	return resp
}

func GetResource(id uint) *Resource {

	resource := &Resource{}
	err := GetDB().Table("resource").Where("id = ?", id).First(resource).Error
	if err != nil {
		return nil
	}
	return resource
}
