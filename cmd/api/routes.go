package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodPut, "/account/reserv/:id_account/:id_service", app.depositReservAccountHandler)
	router.HandlerFunc(http.MethodGet, "/static/swagger", swaggerHandler)

	router.HandlerFunc(http.MethodGet, "/service/get/:id", app.getServiceHandler)
	router.HandlerFunc(http.MethodPost, "/service/create/:name/:price", app.createServiceHandler)

	//todo service create
	router.HandlerFunc(http.MethodGet, "/tempreport/get/:id", app.getReportTempHandler)
	router.HandlerFunc(http.MethodPut, "/report/create/:id", app.compensationRecognitionHandler)
	router.HandlerFunc(http.MethodGet, "/report/get/:id", app.getCompensationReportHandler)

	router.HandlerFunc(http.MethodPost, "/account/create", app.createAccountHandler)
	router.HandlerFunc(http.MethodGet, "/account/get/:id", app.getAccountHandler)
	router.HandlerFunc(http.MethodPut, "/account/add/:id/:account_cash", app.putDepositAccountHandler)
	router.HandlerFunc(http.MethodPut, "/account/transfer/:id/:ToId/:account_cash", app.transferAccountHandler)
	router.HandlerFunc(http.MethodPut, "/account/withdrawal/:id/:account_cash", app.withdrawalepositAccountHandler)
	router.HandlerFunc(http.MethodGet, "/transaction/get/:id", app.getTransactionHandler)

	router.HandlerFunc(http.MethodGet, "/company/report/:year/:month", app.getMonthlyReportLinkHandler)
	router.HandlerFunc(http.MethodGet, "/file/:filename", app.getFileOfReportHandler)

	router.HandlerFunc(http.MethodGet, "/history/transactions/:id", app.getUserAllTransactionsHandler)
	router.HandlerFunc(http.MethodGet, "/history/user/:id", app.getUserReportHistoryHandler)

	return router
}
