package controllers

import (
	"encoding/json"
	"strconv"
	"synapsis-go-try/handlers"
	"synapsis-go-try/helpers"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func C_AddItem(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(jwt.MapClaims)
	userId := user["user_id"].(string)

	datum := handlers.Item{}
	err := json.NewDecoder(r.Body).Decode(&datum)
	if err != nil {
		helpers.Logger("error", "In Server: Oopss server someting wrong"+err.Error())
		msg := helpers.MsgErr(http.StatusInternalServerError, "internal server error", err.Error())
		helpers.Response(w, http.StatusInternalServerError, msg)
		return
	}

	result, err := datum.H_AddItem(userId)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(201, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "created item, response: "+string(logger))
	helpers.Response(w, http.StatusCreated, rMsg)
}

func C_GetAllItem(w http.ResponseWriter, r *http.Request) {
	result := handlers.H_GetAllItem()
	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	helpers.Logger("info", "View all item")
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_UpdateItem(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(jwt.MapClaims)
	userId := user["user_id"].(string)
	userIdInt, _ := strconv.Atoi(userId)

	itemId := mux.Vars(r)["item_id"]
	itemIdInt, _ := strconv.Atoi(itemId)

	datum := handlers.Item{}
	err := json.NewDecoder(r.Body).Decode(&datum)
	if err != nil {
		helpers.Logger("error", "In Server: Oopss server someting wrong"+err.Error())
		msg := helpers.MsgErr(http.StatusInternalServerError, "internal server error", err.Error())
		helpers.Response(w, http.StatusInternalServerError, msg)
		return
	}

	result, err := datum.H_UpdateSingleItem(itemIdInt, userIdInt)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}
	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "Updated payment, response: "+string(logger))
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_DeleteItem(w http.ResponseWriter, r *http.Request) {
	parse := mux.Vars(r)
	itemId := parse["item_id"]

	result, err := handlers.H_DeleteItem(itemId)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Succesfully")
	rMsg["deleted"] = result
	helpers.Logger("info", "Deleted item")
	helpers.Response(w, http.StatusOK, rMsg)
}
