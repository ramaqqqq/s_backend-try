package handlers

import (
	"errors"
	"synapsis-go-try/config"
	"synapsis-go-try/helpers"
	"synapsis-go-try/models"

	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Payment models.Payment

func updateUserBalance(userID int, amount float64) error {
	var user User
	if err := config.GetDB().Debug().First(&user, userID).Error; err != nil {
		helpers.Logger("error", "In Server: [PaymentHandler.UpdateUserBalance] - id is not exist "+err.Error())
		return err
	}

	newBalance := user.Balance + amount
	if newBalance < 0 {
		helpers.Logger("error", "In Server: [PaymentHandler.UpdateUserBalance] - Saldo anda kurang")
		return errors.New("Saldo anda kurang")
	}

	user.Balance = newBalance

	return config.GetDB().Debug().Save(&user).Error
}

func deleteUserBalance(userID int, amount float64) error {
	var user User
	if err := config.GetDB().Debug().First(&user, userID).Error; err != nil {
		helpers.Logger("error", "In Server: [PaymentHandler.DeletedUserBalance] - id is not exist")
		return err
	}
	user.Balance += amount
	return config.GetDB().Debug().Save(&user).Error
}

func H_PaymentByTopUps(purchaseId int, userId string) (H, error) {
	tx := config.GetDB().Begin()

	uuid := uuid.NewV4().String()
	userIdint, _ := strconv.Atoi(userId)

	var purchase Shoppingcart
	if err := tx.Debug().First(&purchase, "purchase_id = ?", purchaseId).Error; err != nil {
		tx.Debug().Rollback()
		helpers.Logger("error", "In Server: [PaymentHandler.GetPurchase] - id is not exist "+err.Error())
		return nil, err
	}

	var user User
	if err := tx.Debug().First(&user, "user_id = ?", userId).Error; err != nil {
		tx.Debug().Rollback()
		helpers.Logger("error", "In Server: [PaymentHandler.GetUser] - id is not exist "+err.Error())
		return nil, err
	}

	if user.Balance < purchase.TotalPrice {
		tx.Debug().Rollback()
		helpers.Logger("error", "In Server: [PaymentHandler.CheckBalance] - Insufficient balance")
		return nil, errors.New("In Server: Insufficient balance")
	}

	payment := Payment{}
	payment.OrderID = uuid
	payment.UserID = userIdint
	payment.PurchaseID = purchaseId
	payment.Amount = purchase.TotalPrice
	payment.Status = "Pembayaran berhasil"
	payment.SnapUrl = "by saldo user"
	payment.PaymentDate = time.Now()

	if err := tx.Debug().Create(&payment).Error; err != nil {
		tx.Debug().Rollback()
		helpers.Logger("error", "In Server: [PaymentHandler.Saved] - failed insert: "+err.Error())
		return nil, err
	}

	user.Balance -= purchase.TotalPrice
	if err := tx.Debug().Model(&user).Where("user_id = ?", userId).Update("balance", user.Balance).Error; err != nil {
		tx.Debug().Rollback()
		helpers.Logger("error", "In Server: [PaymentHandler.UpdatedBalance] - failed update: "+err.Error())
		return nil, err
	}

	tx.Commit()

	helpers.Logger("info", "In Server: Payment created successfully")

	msg := H{}
	msg["payment_id"] = payment.PaymentID
	msg["user_id"] = payment.UserID
	msg["purchase_id"] = payment.PurchaseID
	msg["amount"] = payment.Amount
	msg["status"] = payment.Status
	msg["order_id"] = payment.OrderID
	msg["payment_date"] = payment.PaymentDate

	return msg, nil
}

func H_GetAllPayment() (*[]Payment, error) {
	var datum []Payment
	err := config.GetDB().Debug().Find(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: [PaymentHandlers.GetAll] - failed view all data "+err.Error())
		return nil, err
	}
	return &datum, nil
}

func H_DeletePayment(paymentId int) (string, error) {
	tx := config.GetDB().Begin()

	var payment Payment
	if err := tx.Debug().First(&payment, "payment_id = ?", paymentId).Error; err != nil {
		tx.Debug().Rollback()
		helpers.Logger("error", "In Server: [PaymentHandler.GetPayment] - id is not exist "+err.Error())
		return "", err
	}

	if err := tx.Debug().Delete(&payment).Error; err != nil {
		tx.Debug().Rollback()
		helpers.Logger("error", "In Server: [PaymentHandler.Deleted] - id is not exist")
		return "", errors.New("id is not exist")
	}

	if err := deleteUserBalance(payment.UserID, payment.Amount); err != nil {
		tx.Debug().Rollback()
		helpers.Logger("error", "In Server: [PaymentHandler.DeletedUserBalance] - failure delete balance")
		return "", err
	}

	tx.Debug().Commit()

	return "success deleted payment", nil
}
