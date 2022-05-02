package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/models"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payload struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name,omitempty"`
	Price string             `json:"price" bson:"price,omitempty"`
	Date  time.Time          `json:"time" bson:"time,omitempty"`
}

type resp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (app *application) addRecord(w http.ResponseWriter, r *http.Request) {

	var payload Payload

	_ = json.NewDecoder(r.Body).Decode(&payload)

	var rec models.Record

	rec.Name = payload.Name
	rec.Price, _ = strconv.ParseFloat(payload.Price, 64)
	// location, _ := time.LoadLocation("Aisa/Bangkok")
	rec.Date = time.Now()

	err := app.models.DB.AddRecord(rec)

	if err != nil {
		app.errorJson(w, err)
		return
	}

	recTime := rec.Date.Format(time.UnixDate)
	resp := resp{
		Status:  "OK",
		Message: fmt.Sprintf("Added at %s", recTime),
	}

	err = app.writeJSON(w, http.StatusOK, resp)

	if err != nil {
		app.errorJson(w, err)
		return
	}

}

type body struct {
	Status string           `json:"status"`
	Data   []*models.Record `json:"data"`
}

func (app *application) getAllRecord(w http.ResponseWriter, r *http.Request) {
	records, err := app.models.DB.All()

	if err != nil {
		app.errorJson(w, err)
		return
	}

	resp := body{
		Status: "OK",
		Data:   records,
	}

	err = app.writeJSON(w, http.StatusOK, resp)

	if err != nil {
		app.errorJson(w, err)
		return
	}

}

func (app *application) getRecordInMonth(w http.ResponseWriter, r *http.Request) {
	records, err := app.models.DB.WithinMonth()

	if err != nil {
		app.errorJson(w, err)
		return
	}

	resp := body{
		Status: "OK",
		Data:   records,
	}

	err = app.writeJSON(w, http.StatusOK, resp)

	if err != nil {
		app.errorJson(w, err)
		return
	}
}
