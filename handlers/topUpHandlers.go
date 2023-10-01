package handlers

import (
	"synapsis-go-try/config"
	"synapsis-go-try/helpers"
)

func H_TopUpCustomers(userId string, amount int) (H, error) {
	tx := config.GetDB().Begin()
	datum := User{}
	if err := tx.Debug().Where("user_id = ?", userId).First(&datum).Error; err != nil {
		tx.Debug().Rollback()
		helpers.Logger("error", "In Server: "+err.Error())
		return nil, err
	}
	datum.Balance += float64(amount)
	if err := tx.Debug().Model(&datum).Where("user_id = ?", userId).Update("balance", datum.Balance).Error; err != nil {
		tx.Debug().Rollback()
		helpers.Logger("error", "In Server: [TopupHandler.Saved] - failed amount: "+err.Error())
		return nil, err
	}
	tx.Debug().Commit()
	msg := H{}
	msg["id"] = userId
	msg["amount"] = datum.Balance
	return msg, nil
}
