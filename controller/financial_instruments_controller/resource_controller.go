package financial_instruments_controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lukaszozimek/stock-exchange-service/model"
	u "github.com/lukaszozimek/stock-exchange-service/util"
	"net/http"
	"strconv"
)

var CreateResource = func(w http.ResponseWriter, r *http.Request) {
	stock := &model.Resource{}

	err := json.NewDecoder(r.Body).Decode(stock)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := stock.Create()
	u.Respond(w, resp)
}

var UpdateResource = func(w http.ResponseWriter, r *http.Request) {

}
var FindAllResource = func(w http.ResponseWriter, r *http.Request) {

}
var FindOneResource = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := model.GetResource(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
var DeleteResource = func(w http.ResponseWriter, r *http.Request) {

}
