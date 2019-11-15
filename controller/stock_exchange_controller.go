package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lukaszozimek/stock-exchange-service/model/financial_instruments"
	"github.com/lukaszozimek/stock-exchange-service/model/stock_exchange"
	u "github.com/lukaszozimek/stock-exchange-service/util"
	"net/http"
	"strconv"
)

var CreateStockExchange = func(w http.ResponseWriter, r *http.Request) {
	stockExchange := &stock_exchange.StockExchange{}
	err := json.NewDecoder(r.Body).Decode(stockExchange)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := stockExchange.Create()
	u.Respond(w, resp)
}

var UpdateStockExchange = func(w http.ResponseWriter, r *http.Request) {

}
var FindAllStockExchange = func(w http.ResponseWriter, r *http.Request) {

}
var FindOneStockExchange = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := financial_instruments.GetStockIndex(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
var DeleteStockExchange = func(w http.ResponseWriter, r *http.Request) {

}
