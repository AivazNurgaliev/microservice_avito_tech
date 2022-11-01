package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodPut, "/account/reserv/:id_account/:id_service", app.depositReservAccountHandler) //1

	router.HandlerFunc(http.MethodGet, "/service/get/:id", app.getServiceHandler)                 //1
	router.HandlerFunc(http.MethodPost, "/service/create/:name/:price", app.createServiceHandler) //1

	//todo service create
	router.HandlerFunc(http.MethodGet, "/tempreport/get/:id", app.getReportTempHandler)          //1
	router.HandlerFunc(http.MethodPut, "/report/create/:id", app.compensationRecognitionHandler) //1
	router.HandlerFunc(http.MethodGet, "/report/get/:id", app.getCompensationReportHandler)      //1

	router.HandlerFunc(http.MethodPost, "/account/create", app.createAccountHandler)                                //1
	router.HandlerFunc(http.MethodGet, "/account/get/:id", app.getAccountHandler)                                   //1
	router.HandlerFunc(http.MethodPut, "/account/add/:id/:account_cash", app.putDepositAccountHandler)              //1
	router.HandlerFunc(http.MethodPut, "/account/transfer/:id/:ToId/:account_cash", app.transferAccountHandler)     //1
	router.HandlerFunc(http.MethodPut, "/account/withdrawal/:id/:account_cash", app.withdrawalepositAccountHandler) //1
	router.HandlerFunc(http.MethodGet, "/transaction/get/:id", app.getTransactionHandler)                           //1

	router.HandlerFunc(http.MethodGet, "/company/report/:year/:month", app.getMonthlyReportLinkHandler) //1
	router.HandlerFunc(http.MethodGet, "/file/:filename", app.getFileOfReportHandler)                   //1

	router.HandlerFunc(http.MethodGet, "/history/transactions/:id", app.getUserAllTransactionsHandler) //1
	router.HandlerFunc(http.MethodGet, "/history/user/:id", app.getUserReportHistoryHandler)           //1

	return router
}
