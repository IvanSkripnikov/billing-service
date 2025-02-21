package httphandler

import (
	"net/http"
	"regexp"

	"billing-service/controllers"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	// system
	newRoute(http.MethodGet, "/health", controllers.HealthCheck),
	// orders
	newRoute(http.MethodGet, "/v1/account/list", controllers.GetAccountsListV1),
	newRoute(http.MethodGet, "/v1/account/get-balance/([0-9]+)", controllers.GetBalanceV1),
	newRoute(http.MethodPost, "/v1/account/create", controllers.CreateAccountV1),
	newRoute(http.MethodPut, "/v1/account/buy", controllers.BuyV1),
	newRoute(http.MethodPut, "/v1/account/deposit", controllers.DepositV1),
}
