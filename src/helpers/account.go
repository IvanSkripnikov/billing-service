package helpers

import (
	"encoding/json"
	"net/http"
	"strings"

	"billing-service/models"

	logger "github.com/IvanSkripnikov/go-logger"
)

func GetAccountsList(w http.ResponseWriter, _ *http.Request) {
	category := "/v1/account/list"
	var accounts []models.Account

	query := "SELECT id, user_id, balance, created, updated, active FROM accounts WHERE active = 1"
	rows, err := DB.Query(query)
	if err != nil {
		logger.Error(err.Error())
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	for rows.Next() {
		account := models.Account{}
		if err = rows.Scan(&account.ID, &account.UserID, &account.Balance, &account.Created, &account.Updated, &account.Active); err != nil {
			logger.Error(err.Error())
			continue
		}
		accounts = append(accounts, account)
	}

	data := ResponseData{
		"data": accounts,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	category := "/v1/account/get-balance"
	var account models.Account

	account.UserID, _ = getIDFromRequestString(strings.TrimSpace(r.URL.Path))
	if account.UserID == 0 {
		FormatResponse(w, http.StatusUnprocessableEntity, category)
		return
	}

	if !isExists("SELECT * FROM accounts WHERE user_id = ?", account.UserID) {
		FormatResponse(w, http.StatusNotFound, category)
		return
	}

	query := "SELECT id, user_id, balance, created, updated, active FROM accounts WHERE user_id = ?"
	rows, err := DB.Prepare(query)

	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
	}()

	err = rows.QueryRow(account.UserID).Scan(&account.ID, &account.UserID, &account.Balance, &account.Created, &account.Updated, &account.Active)
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"data": account.Balance,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	category := "/v1/account/create"
	var account models.Account

	err := json.NewDecoder(r.Body).Decode(&account)
	if checkError(w, err, category) {
		return
	}

	query := "INSERT INTO accounts (user_id, balance, created, updated) VALUES (?, ?, ?, ?)"
	currentTimestamp := GetCurrentTimestamp()
	rows, err := DB.Query(query, account.UserID, account.Balance, currentTimestamp, currentTimestamp)

	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
		_ = rows.Err()
	}()

	data := ResponseData{
		"result": "Account successfully created!",
	}
	SendResponse(w, data, category, http.StatusOK)
}

func ChangeBalance(w http.ResponseWriter, r *http.Request, operation string) {
	category := "/v1/account/buy"
	var account models.Account

	err := json.NewDecoder(r.Body).Decode(&account)
	if checkError(w, err, category) {
		return
	}

	if !isExists("SELECT * FROM accounts WHERE user_id = ?", account.UserID) {
		FormatResponse(w, http.StatusNotFound, category)
		return
	}

	Balance := account.Balance

	query := "SELECT id, user_id, balance, created, updated, active FROM accounts WHERE user_id = ?"
	rows, err := DB.Prepare(query)

	if checkError(w, err, category) {
		return
	}

	defer func() {
		_ = rows.Close()
	}()

	err = rows.QueryRow(account.UserID).Scan(&account.ID, &account.UserID, &account.Balance, &account.Created, &account.Updated, &account.Active)
	if checkError(w, err, category) {
		return
	}

	var NewBalance float32
	if operation == "buy" {
		NewBalance = account.Balance - Balance
	} else {
		NewBalance = account.Balance + Balance
	}

	currentTimestamp := GetCurrentTimestamp()
	query = "UPDATE accounts SET balance = ?, updated = ? WHERE user_id = ?"
	_, err = DB.Query(query, NewBalance, currentTimestamp, account.UserID)

	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"message": "success",
	}
	SendResponse(w, data, category, http.StatusOK)
}
