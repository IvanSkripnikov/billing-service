package controllers

import (
	"net/http"

	"billing-service/helpers"
)

func GetAccountsListV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetAccountsList(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/account/list")
	}
}

func GetBalanceV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetBalance(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/account/get-balance")
	}
}

func CreateAccountV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.CreateAccount(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/account/create")
	}
}

func BuyV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		helpers.ChangeBalance(w, r, "buy")
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/account/buy")
	}
}

func DepositV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		helpers.ChangeBalance(w, r, "deposit")
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/account/deposit")
	}
}
