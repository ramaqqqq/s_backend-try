package main

import (
	"os"
	"synapsis-go-try/controllers"

	"github.com/gorilla/mux"
)

func Factory(ctm *mux.Router) string {
	middleUrl := os.Getenv("MIDDLE_URL")

	//auth
	ctm.HandleFunc(middleUrl+"/login", controllers.C_Login).Methods("POST")
	ctm.HandleFunc(middleUrl+"/register", controllers.C_Register).Methods("POST")

	//item
	ctm.HandleFunc(middleUrl+"/item", controllers.C_AddItem).Methods("POST")
	ctm.HandleFunc(middleUrl+"/item", controllers.C_GetAllItem).Methods("GET")
	ctm.HandleFunc(middleUrl+"/item/{item_id}/edit", controllers.C_UpdateItem).Methods("PUT")
	ctm.HandleFunc(middleUrl+"/item/{item_id}/delete", controllers.C_DeleteItem).Methods("DELETE")

	//shopping-cart
	ctm.HandleFunc(middleUrl+"/shopping-cart/{item_id}/{quantity}", controllers.C_AddShoppCart).Methods("POST")
	ctm.HandleFunc(middleUrl+"/shopping-cart", controllers.C_GetAllShoppCart).Methods("GET")
	ctm.HandleFunc(middleUrl+"/shopping-cart/{purchase_id}/{item_id}/{quantity}/edit", controllers.C_UpdateShoppCart).Methods("PUT")
	ctm.HandleFunc(middleUrl+"/shopping-cart/{purchase_id}/delete", controllers.C_DeleteShoppCart).Methods("DELETE")

	//payment
	ctm.HandleFunc(middleUrl+"/payment/{purchase_id}", controllers.C_AddPayment).Methods("POST")
	ctm.HandleFunc(middleUrl+"/payment", controllers.C_GetAllPayment).Methods("GET")
	ctm.HandleFunc(middleUrl+"/payment/{payment_id}/delete", controllers.C_DeletePayment).Methods("DELETE")

	//topup
	ctm.HandleFunc(middleUrl+"/topup/{amount}", controllers.C_TopUpCustomer).Methods("POST")

	return "presenter :"
}
