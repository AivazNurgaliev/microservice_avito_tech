package main

import (
	"net/http"
)

func (app *application) createAccountHandler(w http.ResponseWriter, r *http.Request) {

	app.models.Account.Create()

}

func (app *application) getAccountHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdFromPathParam(r)

	if err != nil || id < 1 {
		app.logger.Println(err)
		return
	}

	account, err := app.models.Account.GetAccountById(id)
	if err != nil {
		app.logger.Println(err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)

	if err != nil {
		app.logger.Println(err)
	}
}

// функция перевода денег из кошелька в резерв
func (app *application) depositReservAccountHandler(w http.ResponseWriter, r *http.Request) {
	id_account, id_service, err := app.readDepositReserveParam(r)

	if err != nil {
		app.logger.Println(err)
		return
	}

	account, err := app.models.Account.GetAccountById(id_account)
	service, err := app.models.Service.Get(id_service)

	if account == nil {
		app.logger.Println("Account not found")
		return
	}

	if service == nil {
		app.logger.Println("Service not found")
		return
	}

	if err != nil {
		app.logger.Println(err)
		return
	}

	var price = service.ServicePrice

	if account.AccountId == 0 || service.ServiceId == 0 {
		app.logger.Println("Account or Service not found")
		return
	}

	account.AccountId = id_account
	if account.AccountCash-price < 0 {
		app.logger.Println("Not enough money")
		return
	}
	account.AccountCash -= price
	account.AccountReservedCash += price

	err = app.models.Account.UpdateAccountCashData(account)
	if err != nil {
	}

	app.models.Report.CreateTemp(id_account, id_service)
	if err != nil {
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	if err != nil {
		app.logger.Println(err)
	}
}

// функция добавления денег на аккаунт
func (app *application) putDepositAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdFromPathParam(r)
	cash, err := app.readCashFromPathParam(r)

	if err != nil {
		app.logger.Println(err)
		return
	}

	account, err := app.models.Account.GetAccountById(id)
	if err != nil || account == nil {
		app.models.Account.CreateId(id)
		app.logger.Println(err)
	}
	account, err = app.models.Account.GetAccountById(id)

	account.AccountId = id
	account.AccountCash += cash

	err = app.models.Account.UpdateAccountCashById(account)
	if err != nil {
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	if err != nil {
		app.logger.Println(err)
	}
	app.models.Transaction.CreateTransaction(account.AccountId, account.AccountId, 1, cash)
}

// функция перевода денег с аккаунта на аккаунт
func (app *application) transferAccountHandler(w http.ResponseWriter, r *http.Request) {
	FromId, err := app.readIdFromPathParam(r)
	ToId, err := app.readReceiverIdFromPathParam(r)
	cash, err := app.readCashFromPathParam(r)
	if cash < 0 {
		return
	}

	accountFrom, err := app.models.Account.GetAccountById(FromId)
	if err != nil {
		app.logger.Println(err)
		return
	}

	accountTo, err := app.models.Account.GetAccountById(ToId)
	if err != nil {
		app.logger.Println(err)
		return
	}
	var newCash float64 = accountFrom.AccountCash - cash
	if newCash < 0 {
		app.logger.Println("not enough cash")
		return
	} else {

		accountFrom.AccountCash = newCash
		newCash = accountTo.AccountCash + cash
		accountTo.AccountCash = newCash
		app.models.Account.UpdateAccountCashById(accountFrom)
		app.models.Account.UpdateAccountCashById(accountTo)
		app.models.Transaction.CreateTransaction(accountFrom.AccountId, accountTo.AccountId, 3, cash)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"accountFrom": accountFrom}, nil)
	err = app.writeJSON(w, http.StatusOK, envelope{"accountTo": accountTo}, nil)
}

// функция снятия денег с аккаунта
func (app *application) withdrawalepositAccountHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdFromPathParam(r)
	cash, err := app.readCashFromPathParam(r)

	if err != nil {
		app.logger.Println(err)
		return
	}

	account, err := app.models.Account.GetAccountById(id)
	if err != nil {
		app.logger.Println(err)
		return
	}

	if account.AccountId == 0 {
		app.logger.Println("Account not found")
		return
	}

	account.AccountId = id
	account.AccountCash -= cash

	if account.AccountCash < 0 {
		app.logger.Println("negative number")
		return
	}

	err = app.models.Account.UpdateAccountCashById(account)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"account": account}, nil)
	if err != nil {
		app.logger.Println(err)
	}
	app.models.Transaction.CreateTransaction(account.AccountId, account.AccountId, 2, cash)
}
