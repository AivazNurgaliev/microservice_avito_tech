package main

import (
	"encoding/csv"
	"fmt"
	"microservice-balance_avito_internship/internal/data"
	"net/http"
	"os"
)

//todo переименовать
// функция получения временного платежа

func (app *application) getReportTempHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdFromPathParam(r)

	if err != nil || id < 1 {
		app.logger.Println(err)
		return
	}

	report, err := app.models.Report.GetTemp(id)
	if err != nil {
		app.logger.Println(err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"temp_report": report}, nil)

	if err != nil {
		app.logger.Println(err)
	}
}

// функция подтверждения пэймента
func (app *application) compensationRecognitionHandler(w http.ResponseWriter, r *http.Request) {
	tempCompensationID, err := app.readIdFromPathParam(r)
	if err != nil || tempCompensationID < 1 {
		app.logger.Println(err)
		return
	}

	tempCompensation, err := app.models.Report.GetTemp(tempCompensationID)
	if tempCompensation == nil {
		app.logger.Println("Incorrect id")
		return
	}
	if err != nil {
		app.logger.Println(err)
	}

	service, err := app.models.Service.Get(tempCompensation.ServiceId)
	if err != nil || tempCompensationID < 1 {
		app.logger.Println(err)
		return
	}
	account, err := app.models.Account.GetAccountById(tempCompensation.AccountId)
	if err != nil || tempCompensationID < 1 {
		app.logger.Println(err)
		return
	}
	app.models.Report.CreatePayment(tempCompensation.AccountId, tempCompensation.ServiceId, tempCompensationID, account.AccountReservedCash-service.ServicePrice)
	if err != nil {
		app.logger.Println(err)
	}
}

// функция получения платежа
func (app *application) getCompensationReportHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdFromPathParam(r)

	if err != nil || id < 1 {
		app.logger.Println(err)
		return
	}

	report, err := app.models.Report.Get(id)
	if err != nil {
		app.logger.Println("report by this id not found")
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"report": report}, nil)

	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) getMonthlyReportLinkHandler(w http.ResponseWriter, r *http.Request) {

	year, err := app.readYearFromPathParam(r)
	if err != nil {
		app.logger.Println(err)
		return
	}

	month, err := app.readMonthFromPathParam(r)
	if err != nil {
		app.logger.Println(err)
		return
	}

	reports, err := app.models.Report.GetMonthlyReport(year, month)

	if err != nil {
		app.logger.Println(err)
		return
	}

	fileName := fmt.Sprintf("report%d-%d.csv", year, month)

	filePath := "reports/" + fileName
	csvFile, err := os.Create(filePath)

	if err != nil {
		app.logger.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	err = writer.WriteAll(reports)
	if err != nil {
		app.logger.Println(err)
		return
	}

	writer.Flush()

	link := fmt.Sprintf("http://localhost:4000/file/%s", fileName)
	fmt.Println(link)
	err = app.writeJSON(w, http.StatusOK, envelope{"link": link}, nil)

	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) getFileOfReportHandler(w http.ResponseWriter, r *http.Request) {

	fileName, _ := app.readFileNameFromPathParam(r)
	filePath := "reports/" + fileName
	http.ServeFile(w, r, filePath)
}

// История покупок пользователя различных услуг
func (app *application) getUserReportHistoryHandler(w http.ResponseWriter, r *http.Request) {
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

	userReportHistory, metadata, err := app.models.Report.GetUserHistory(id, input.Filters)

	if userReportHistory == nil {
		app.logger.Println("Incorrect account id")
		return
	}

	if err != nil {
		app.logger.Println(err)
		return
	}
	err = app.writeJSON(w, http.StatusOK,
		envelope{"metadata": metadata}, nil)
	err = app.writeJSON(w, http.StatusOK,
		envelope{"User Buying Services History": userReportHistory}, nil)

	if err != nil {
		app.logger.Println(err)
	}
}
