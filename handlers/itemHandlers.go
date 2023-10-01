package handlers

import (
	"errors"
	"strconv"
	"synapsis-go-try/config"
	"synapsis-go-try/helpers"
	"synapsis-go-try/models"
)

type Item models.Item

func (h *Item) H_AddItem(userId string) (H, error) {
	userIdInt, _ := strconv.Atoi(userId)

	datum := Item{}
	datum.UserID = userIdInt
	datum.ItemName = h.ItemName
	datum.Price = h.Price
	datum.Stock = h.Stock

	err := config.GetDB().Debug().Create(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: [ItemHandler.Saved] - failed insert: "+err.Error())
		return nil, err
	}

	msg := H{}
	msg["item_id"] = datum.ItemID
	msg["item_name"] = datum.ItemName
	msg["item_price"] = datum.Price
	msg["item_stock"] = datum.Stock

	return msg, nil
}

func H_GetAllItem() *[]Item {
	var datum []Item
	err := config.GetDB().Debug().Find(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: [ItemHandlers.GetAll] - failed view all data "+err.Error())
		return nil
	}
	return &datum
}

func H_GetSingleItem(itemId string) *Item {
	var datum Item
	err := config.GetDB().Debug().Where("item_id = ?", itemId).Find(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: [ItemHandlers.GetOne] - view one data "+err.Error())
		return nil
	}
	return &datum
}

func (h *Item) H_UpdateSingleItem(itemId int, userId int) (*Item, error) {
	datum := Item{}
	datum.UserID = userId
	datum.ItemID = itemId
	datum.ItemName = h.ItemName
	datum.Price = h.Price
	datum.Stock = h.Stock

	err := config.GetDB().Debug().Model(datum).Where("item_id = ?", itemId).Update(&h).Error
	if err != nil {
		helpers.Logger("error", "In Server: [ItemtHandler.Updated] - failed update: "+err.Error())
		return nil, err
	}
	return &datum, err
}

func H_DeleteItem(itemId string) (string, error) {
	tx := config.GetDB().Begin()
	rowsAffected := tx.Debug().Model(Item{}).Where("item_id = ?", itemId).Delete(Item{}).RowsAffected
	if rowsAffected == 0 {
		helpers.Logger("error", "In Server: [ItemHandler.Deleted] - id is not exist")
		return "", errors.New("id is not exist")
	}
	tx.Commit()
	return "success to deleted", nil
}
