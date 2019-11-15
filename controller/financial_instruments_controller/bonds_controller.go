package financial_instruments_controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lukaszozimek/stock-exchange-service/model/financial_instruments"
	u "github.com/lukaszozimek/stock-exchange-service/util"
	"net/http"
	"strconv"
)

var CreateBonds = func(w http.ResponseWriter, r *http.Request) {
	stockExchange := &financial_instruments.Bonds{}
	err := json.NewDecoder(r.Body).Decode(stockExchange)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := stockExchange.Create()
	u.Respond(w, resp)
}

var UpdateBonds = func(w http.ResponseWriter, r *http.Request) {

}
var FindAllBonds = func(w http.ResponseWriter, r *http.Request) {

}
var FindOneBonds = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := financial_instruments.GetBond(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
var DeleteBonds = func(w http.ResponseWriter, r *http.Request) {

}
