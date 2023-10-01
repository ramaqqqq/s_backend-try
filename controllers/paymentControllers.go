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

func C_AddPayment(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(jwt.MapClaims)
	userId := user["user_id"].(string)
	purchaseId := mux.Vars(r)["purchase_id"]
	purchaseIdint, _ := strconv.Atoi(purchaseId)

	result, err := handlers.H_PaymentByTopUps(purchaseIdint, userId)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(201, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "Created payment, response: "+string(logger))
	helpers.Response(w, http.StatusCreated, rMsg)

}

func C_GetAllPayment(w http.ResponseWriter, r *http.Request) {
	result, err := handlers.H_GetAllPayment()
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	helpers.Logger("info", "view all payment")
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_DeletePayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paymentId := params["payment_id"]
	paymentIdInt, _ := strconv.Atoi(paymentId)

	result, err := handlers.H_DeletePayment(paymentIdInt)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	helpers.Logger("info", "deleted payment")
	helpers.Response(w, http.StatusOK, rMsg)
}
