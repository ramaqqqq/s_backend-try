package handlers

import (
	"errors"
	"strconv"
	"synapsis-go-try/config"
	"synapsis-go-try/helpers"
	"synapsis-go-try/models"
	"time"
)

type Shoppingcart models.Shoppingcart

func H_AddShoppCart(userId string, itemId int, quantity int) (H, error) {
	tx := config.GetDB().Begin()

	var user User
	if err := tx.Debug().First(&user, "user_id = ?", userId).Error; err != nil {
		tx.Rollback()
		helpers.Logger("error", "In Server: [ShoppCartHandler.GetUser] - id is not exist "+err.Error())
		return nil, err
	}

	var item Item
	if err := tx.Debug().First(&item, "item_id = ?", itemId).Error; err != nil {
		tx.Rollback()
		helpers.Logger("error", "In Server: [ShoppCartHandler.GetItem] - id is not exist "+err.Error())
		return nil, err
	}

	if item.Stock < quantity {
		tx.Debug().Rollback()
		helpers.Logger("error", "In Server: [ShoppCartHandler.CheckStock] - Insufficient stock")
		return nil, errors.New("In Server: Insufficient stock")
	}

	totalPrice := item.Price * float64(quantity)
	userIdint, _ := strconv.Atoi(userId)

	purchase := Shoppingcart{}
	purchase.UserID = userIdint
	purchase.ItemID = itemId
	purchase.Quantity = quantity
	purchase.TotalPrice = totalPrice
	purchase.PurchaseDate = time.Now()

	err := tx.Debug().Create(&purchase).Error
	if err != nil {
		tx.Rollback()
		helpers.Logger("error", "In Server: [ShoppCartHandler.Saved] - failed insert: "+err.Error())
		return nil, err
	}

	tx.Debug().Commit()

	msg := H{}
	msg["purhase_id"] = purchase.PurchaseID
	msg["user_id"] = purchase.UserID
	msg["item_id"] = purchase.ItemID
	msg["quantity"] = purchase.Quantity
	msg["total_price"] = purchase.TotalPrice
	msg["purchase_date"] = purchase.PurchaseDate

	return msg, nil
}

func H_GetAllShoppCart() (*[]Shoppingcart, error) {
	var datum []Shoppingcart
	err := config.GetDB().Debug().Find(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: [ShoppCartHandlers.GetAll] - failed view all data "+err.Error())
		return nil, err
	}
	return &datum, nil
}

func H_UpdateShoppCart(purchaseId int, userId string, itemId int, quantity int) (H, error) {
	tx := config.GetDB().Begin()

	var purchase Shoppingcart
	if err := tx.Debug().First(&purchase, "purchase_id = ?", purchaseId).Error; err != nil {
		tx.Rollback()
		helpers.Logger("error", "In Server: [ShoppCartHandler.GetPurchase] - id is not exist "+err.Error())
		return nil, err
	}

	var item Item
	if err := tx.Debug().First(&item, "item_id = ?", itemId).Error; err != nil {
		tx.Rollback()
		helpers.Logger("error", "In Server: [ShoppCartHandler.GetItem] - id is not exist "+err.Error())
		return nil, err
	}

	if item.Stock < quantity {
		tx.Rollback()
		helpers.Logger("error", "In Server: [ShoppCartHandler.CheckStock] - Insufficient stock")
		return nil, errors.New("In Server: Insufficient stock")
	}

	newTotalPrice := item.Price * float64(quantity)
	userIdInt, _ := strconv.Atoi(userId)

	purchase.UserID = userIdInt
	purchase.ItemID = itemId
	purchase.Quantity = quantity
	purchase.TotalPrice = newTotalPrice

	err := tx.Debug().Save(&purchase).Error
	if err != nil {
		tx.Rollback()
		helpers.Logger("error", "In Server: [ShoppCartHandler.Updated] - failed update: "+err.Error())
		return nil, err
	}

	tx.Debug().Commit()

	msg := H{}
	msg["purchase_id"] = purchase.PurchaseID
	msg["user_id"] = purchase.UserID
	msg["item_id"] = purchase.ItemID
	msg["quantity"] = purchase.Quantity
	msg["total_price"] = purchase.TotalPrice

	return msg, nil
}

func H_DeleteShoppCart(purchaseId int) (string, error) {
	tx := config.GetDB().Begin()
	rowsAffected := tx.Debug().Model(Shoppingcart{}).Where("purchase_id = ?", purchaseId).Delete(Shoppingcart{}).RowsAffected
	if rowsAffected == 0 {
		helpers.Logger("error", "In Server: [BlogHandler.Deleted] - id is not exist")
		return "", errors.New("id is not exist")
	}
	tx.Commit()
	return "success to deleted purchase", nil
}
