package controllers

import (
	"encoding/json"
	"synapsis-go-try/handlers"
	"synapsis-go-try/helpers"

	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func C_TopUpCustomer(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(jwt.MapClaims)
	userId := user["user_id"].(string)

	amount := mux.Vars(r)["amount"]
	amountint, _ := strconv.Atoi(amount)

	result, err := handlers.H_TopUpCustomers(userId, amountint)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(201, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "response: "+string(logger))
	helpers.Response(w, http.StatusCreated, rMsg)
}
