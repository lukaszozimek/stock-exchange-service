package financial_instruments

import (
	"github.com/jinzhu/gorm"
	"github.com/lukaszozimek/stock-exchange-service/model"
	u "github.com/lukaszozimek/stock-exchange-service/util"
)

type Bonds struct {
	gorm.Model
	Name     string `json:"name"`
	Shortcut string `json:"shortcut"`
}

func (bonds *Bonds) Validate() (map[string]interface{}, bool) {

	if bonds.Name == "" {
		return u.Message(false, "Bonds name should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (bonds *Bonds) Create() map[string]interface{} {

	if resp, ok := bonds.Validate(); !ok {
		return resp
	}

	model.GetDB().Create(bonds)

	resp := u.Message(true, "success")
	resp["bonds"] = bonds
	return resp
}
func DeleteBond(id uint) {
	bonds := &Bonds{}
	model.GetDB().Table("bonds").Delete(&bonds, model.GetDB().Where("id = ?", id))

}

func GetBond(id uint) *Bonds {

	bonds := &Bonds{}
	err := model.GetDB().Table("bonds").Where("id = ?", id).First(bonds).Error
	if err != nil {
		return nil
	}
	return bonds
}
