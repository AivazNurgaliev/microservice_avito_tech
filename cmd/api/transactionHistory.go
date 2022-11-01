package main

import (
	"microservice-balance_avito_internship/internal/data"
	"net/http"
)

// todo
func (app *application) getTransactionHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdFromPathParam(r)

	if err != nil || id < 1 {
		app.logger.Println(err)
		return
	}

	transaction, err := app.models.Transaction.GetTransactionById(id)
	if err != nil {
		app.logger.Println("Incorrect id")
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"transaction": transaction}, nil)

	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) getUserAllTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		data.Filters
	}

	qs := r.URL.Query()
	var err error
	input.Filters.Page = app.readIntFromQueryParam(qs, "page", 1)
	input.Filters.PageSize = app.readIntFromQueryParam(qs, "page_size", 5)
	input.Filters.Sort, err = app.readStringFromQuery(qs, "sort", "")

	if err != nil {
		app.logger.Println(err)
		return
	}

	id, err := app.readIdFromPathParam(r)
	if err != nil || id < 1 {
		app.logger.Println(err)
		return
	}

	userTransactionHistory, metadata, err := app.models.TransactionHistory.GetUserTransactionHistory(id, input.Filters)
	if userTransactionHistory == nil {
		app.logger.Println("Incorrect transaction id")
		return
	}

	if err != nil {
		app.logger.Println(err)
		return
	}
	err = app.writeJSON(w, http.StatusOK,
		envelope{"metadata": metadata}, nil)
	err = app.writeJSON(w, http.StatusOK,
		envelope{"User Transaction History": userTransactionHistory}, nil)

	if err != nil {
		app.logger.Println(err)
	}

}
