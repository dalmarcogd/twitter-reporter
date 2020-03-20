package database

import (
	"github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/errors"
	"net/http"
)

func GetAccountById(accountId string) (AccountModel, error) {
	account := AccountModel{}
	GetConnection().Where("id = ?", accountId).First(&account)
	if account.Id != "" {
		return account, nil
	}
	return account, errors.NewError(http.StatusNotFound, "Account not found", nil)
}