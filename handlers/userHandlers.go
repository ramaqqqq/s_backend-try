package handlers

import (
	"errors"
	"strconv"
	"synapsis-go-try/config"
	"synapsis-go-try/helpers"
	"synapsis-go-try/middleware"
	"synapsis-go-try/models"

	"golang.org/x/crypto/bcrypt"
)

type H map[string]interface{}
type User models.User

func (h *User) H_Login() (H, error) {

	datum := User{}

	err := config.GetDB().Debug().Where("email = ?", h.Email).Take(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: [UserHandler.Login] - email is not exist "+err.Error())
		return nil, errors.New("email not exists")
	}

	err = helpers.VerifyPassword(datum.Password, h.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		helpers.Logger("error", "In Server: invalid verify password")
		return nil, errors.New("Password is failure")
	}

	strUserID := strconv.Itoa(datum.UserID)
	token, err := middleware.CreateToken(strUserID, datum.Username, datum.Email, datum.Role)
	if err != nil {
		helpers.Logger("error", "In Server: [UserHandler.Login] - failure create token "+err.Error())
	}

	rMsg := H{}
	rMsg["user_id"] = datum.UserID
	rMsg["username"] = datum.Username
	rMsg["email"] = datum.Email
	rMsg["role"] = datum.Role
	access := token["accessToken"]
	refresh := token["refreshToken"]

	return H{"accessToken": access, "refreshToken": refresh, "users": rMsg}, nil
}

func (h *User) H_Register() (H, error) {

	hashedPassword, err := helpers.Hash(h.Password)
	if err != nil {
		helpers.Logger("error", "In Server: [userHandler.Register] - hashed: "+err.Error())
		return nil, err
	}

	datum := User{}
	datum.UserID = h.UserID
	datum.Username = h.Username
	datum.Email = h.Email
	datum.Role = "customer"
	datum.Password = string(hashedPassword)

	err = config.GetDB().Debug().Create(&datum).Error
	if err != nil {
		helpers.Logger("error", "In Server: [userHandler.Register] - failed insert: "+err.Error())
		return nil, err
	}

	strUserID := strconv.Itoa(datum.UserID)
	token, err := middleware.CreateToken(strUserID, datum.Username, datum.Email, datum.Role)
	if err != nil {
		helpers.Logger("error", "In Server: [UserHandler.Register] - failure create token "+err.Error())
		return nil, err
	}

	rMsg := H{}
	rMsg["id"] = datum.UserID
	rMsg["username"] = datum.Username
	rMsg["email"] = datum.Email
	rMsg["role"] = datum.Role
	rMsg["created"] = datum.Created
	access := token["accessToken"]
	refresh := token["refreshToken"]

	return H{"accessToken": access, "refreshToken": refresh, "users": rMsg}, nil
}
