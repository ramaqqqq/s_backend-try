package controllers

import (
	"encoding/json"
	"net/http"
	"synapsis-go-try/handlers"
	"synapsis-go-try/helpers"
)

func C_Login(w http.ResponseWriter, r *http.Request) {
	datum := &handlers.User{}
	err := json.NewDecoder(r.Body).Decode(datum)
	if err != nil {
		helpers.Logger("error", "In Server: Oopss server someting wrong"+err.Error())
		msg := helpers.MsgErr(http.StatusInternalServerError, "internal server error", err.Error())
		helpers.Response(w, http.StatusInternalServerError, msg)
		return
	}

	result, err := datum.H_Login()
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	Logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "Login, Response: "+string(Logger))
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_Register(w http.ResponseWriter, r *http.Request) {
	datum := &handlers.User{}
	err := json.NewDecoder(r.Body).Decode(datum)
	if err != nil {
		helpers.Logger("error", "In Server: Oopss server someting wrong"+err.Error())
		msg := helpers.MsgErr(http.StatusInternalServerError, "internal server error", err.Error())
		helpers.Response(w, http.StatusInternalServerError, msg)
		return
	}

	result, err := datum.H_Register()
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	Logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "Register, response: "+string(Logger))
	helpers.Response(w, http.StatusCreated, rMsg)
}
