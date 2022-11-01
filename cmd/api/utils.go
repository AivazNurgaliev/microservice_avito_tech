package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/url"
	"strconv"
)

type envelope map[string]interface{}

func (app *application) readIdFromPathParam(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) readReceiverIdFromPathParam(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	ToId, err := strconv.ParseInt(params.ByName("ToId"), 10, 64)

	if err != nil || ToId < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return ToId, nil
}

func (app *application) readDepositReserveParam(r *http.Request) (int64, int64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	id_account, err := strconv.ParseInt(params.ByName("id_account"), 10, 64)
	id_service, err := strconv.ParseInt(params.ByName("id_service"), 10, 64)

	if err != nil || id_account < 1 || id_service < 1 {
		return 0, 0, errors.New("invalid parameters")
	}

	return id_account, id_service, nil
}

func (app *application) readYearFromPathParam(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	year, err := strconv.ParseInt(params.ByName("year"), 10, 64)

	if err != nil || year < 1 {
		return 0, errors.New("invalid year parameter")
	}

	return year, nil
}

func (app *application) readMonthFromPathParam(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	month, err := strconv.ParseInt(params.ByName("month"), 10, 64)

	if err != nil || month < 1 || month > 12 {
		return 0, errors.New("invalid month parameter")
	}

	return month, nil
}

func (app *application) readCashFromPathParam(r *http.Request) (float64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	accountCash, err := strconv.ParseFloat(params.ByName("account_cash"), 64)
	if err != nil || accountCash < 0 {
		return 0, errors.New("invalid id parameter")
	}

	return accountCash, nil
}

func (app *application) readFileNameFromPathParam(r *http.Request) (string, error) {
	params := httprouter.ParamsFromContext(r.Context())
	filename := params.ByName("filename")
	return filename, nil
}
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {

	js, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	return nil
}

func (app *application) readIntFromQueryParam(qs url.Values, key string, defaultValue int) int {

	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}

	if i < 1 {
		return defaultValue
	}

	return i
}

func (app *application) readFloatFromQuery(qs url.Values, key string, defaultValue float64) float64 {

	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultValue
	}

	return i
}

func (app *application) readStringFromQuery(qs url.Values, key string, defaultValue string) (string, error) {
	s := qs.Get(key)

	if s == "" {
		return defaultValue, nil
	}

	if s != "date" && s != "sum" && s != "-date" && s != "-sum" {
		return "", errors.New("Invalid sort parametrs")
	}

	return s, nil
}

func (app *application) readServiceFromPathParam(r *http.Request) (string, float64, error) {

	params := httprouter.ParamsFromContext(r.Context())
	name := params.ByName("name")
	price, err := strconv.ParseFloat(params.ByName("price"), 64)

	if err != nil || price < 0 {
		return "", 0, errors.New("invalid parameter")
	}

	return name, price, nil
}
