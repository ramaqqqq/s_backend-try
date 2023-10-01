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

func C_AddShoppCart(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(jwt.MapClaims)
	userId := user["user_id"].(string)

	itemId := mux.Vars(r)["item_id"]
	quantity := mux.Vars(r)["quantity"]
	itemIdInt, _ := strconv.Atoi(itemId)
	quantityInt, _ := strconv.Atoi(quantity)

	result, err := handlers.H_AddShoppCart(userId, itemIdInt, quantityInt)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(201, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "created shopping-cart, response: "+string(logger))
	helpers.Response(w, http.StatusCreated, rMsg)
}

func C_GetAllShoppCart(w http.ResponseWriter, r *http.Request) {
	result, err := handlers.H_GetAllShoppCart()
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request | services purchase", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	helpers.Logger("info", "view all shopping-cart")
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_UpdateShoppCart(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(jwt.MapClaims)
	userId := user["user_id"].(string)

	itemId := mux.Vars(r)["item_id"]
	quantity := mux.Vars(r)["quantity"]
	purchaseId := mux.Vars(r)["purchase_id"]
	itemIdInt, _ := strconv.Atoi(itemId)
	quantityInt, _ := strconv.Atoi(quantity)
	purchaseidInt, _ := strconv.Atoi(purchaseId)

	result, err := handlers.H_UpdateShoppCart(purchaseidInt, userId, itemIdInt, quantityInt)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "status bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "Updated shopping-cart, response: "+string(logger))
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_DeleteShoppCart(w http.ResponseWriter, r *http.Request) {
	purchaseId := mux.Vars(r)["purchase_id"]
	purchaseidInt, _ := strconv.Atoi(purchaseId)

	result, err := handlers.H_DeleteShoppCart(purchaseidInt)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Deleted")
	rMsg["body"] = result
	helpers.Logger("info", "deleted shopping-cart")
	helpers.Response(w, http.StatusOK, rMsg)
}
