package main

import "net/http"

func (app *application) createServiceHandler(w http.ResponseWriter, r *http.Request) {
	name, price, err := app.readServiceFromPathParam(r)

	if err != nil {
		app.logger.Println(err)
		return
	}

	app.models.Service.Create(name, price)

	if err != nil {
		app.logger.Println(err)
	}

}

func (app *application) getServiceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdFromPathParam(r)

	if err != nil {
		app.logger.Println(err)
		return
	}

	service, err := app.models.Service.Get(id)
	if err != nil {
		app.logger.Println(err)
		return
	}

	service.ServiceId = id

	err = app.writeJSON(w, http.StatusOK, envelope{"service": service}, nil)
	if err != nil {
		app.logger.Println(err)
	}
}
