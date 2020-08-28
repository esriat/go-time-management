package handler_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/handlers"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

/*
	TESTED : GET /vacations/{id}
	TESTED : GET /users/{id}/vacations
	TESTED : POST /vacations
	TESTED : PATCH /vacations/{id}
	TESTED : DELETE /vacations/{id}
*/
func TestVacationHandler(t *testing.T) {
	var (
		request    *http.Request
		tmp        model.Schedule
		jsonObject []byte

		err error
	)

	// To store the ScheduleId when decoding json
	var tmpId struct {
		ScheduleId int `json:"vacation_id"`
	}

	// Preparing for HTTP requests
	rr := httptest.NewRecorder()

	//
	//	GET /users/{id}/vacations
	//

	// Getting vacations of user 1
	var vacationsOfUser1 model.Schedules
	if vacationsOfUser1, err = env.DB.GetVacationsOfUser(1); err != nil {
		t.Error(err)
	}
	// Getting rid of the useless details
	for index, _ := range vacationsOfUser1 {
		GetRidOfScheduleDateDetails(&vacationsOfUser1[index])
	}

	// Creating a request
	if request, err = http.NewRequest(http.MethodGet, "/users/1/vacations", nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)

	// Sending the request
	r.ServeHTTP(rr, request)

	// Checking the result
	var SIarray1 []handlers.ScheduleIntermediate
	if err = json.NewDecoder(rr.Body).Decode(&SIarray1); err != nil {
		t.Error(err)
	}

	// Converting the result to model.Schedule
	var dbVacationsOfUser1 model.Schedules
	for _, SI := range SIarray1 {
		if tmp, err = handlers.IntermediateToSchedule(SI); err != nil {
			t.Error()
		}
		dbVacationsOfUser1 = append(dbVacationsOfUser1, tmp)
	}

	// Now comparing
	if !cmp.Equal(vacationsOfUser1, dbVacationsOfUser1) {
		t.Error("Vacations are not the same")
	}

	globals.Log.Debug("GET /users/{id}/vacations - PASSED")

	//
	//	POST /vacations
	//
	SI2 := handlers.ScheduleIntermediate{
		ProjectId: 2,
		StartDate: "2015-06-15 10:19:30",
		EndDate:   "2020-06-15 10:19:30",
	}
	// Turning the object into JSON
	if jsonObject, err = json.Marshal(SI2); err != nil {
		t.Error(err.Error())
	}

	if request, err = http.NewRequest(http.MethodPost, "/vacations", bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)

	// Sending the request
	r.ServeHTTP(rr, request)

	// Checking the result
	if err = json.NewDecoder(rr.Body).Decode(&tmpId); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("POST /vacations - PASSED")

	SI2.ScheduleId = int64(tmpId.ScheduleId)

	//
	//	GET /vacations/{id}
	//

	// Getting the first vacation
	var vacation1 model.Schedule
	if vacation1, err = env.DB.GetVacation(SI2.ScheduleId); err != nil {
		t.Error(err)
	}

	// Creating a request
	if request, err = http.NewRequest(http.MethodGet, "/vacations/"+strconv.FormatInt(SI2.ScheduleId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)

	// Sending the request
	r.ServeHTTP(rr, request)

	// Checking the result
	var SI1 handlers.ScheduleIntermediate
	if err = json.NewDecoder(rr.Body).Decode(&SI1); err != nil {
		t.Error(err)
	}

	// Converting the result to model.Schedule
	var dbVacation1 model.Schedule
	if dbVacation1, err = handlers.IntermediateToSchedule(SI1); err != nil {
		t.Error()
	}
	GetRidOfScheduleDateDetails(&vacation1)

	// Comparing the objects
	if !cmp.Equal(vacation1, dbVacation1) {
		t.Error("Vacations are not the same")
	}

	globals.Log.Debug("GET /vacations/{id} - PASSED")

	//
	//	PATCH /vacations/{id}
	//

	// Fetching a vacation
	ISpatch := handlers.ScheduleToIntermediate(fakeSchedules[1])
	// Modifying it
	ISpatch.StartDate = "2007-05-30 10:09:27"
	// JSON-ing the object
	if jsonObject, err = json.Marshal(ISpatch); err != nil {
		t.Error(err)
	}

	// Creating the request
	if request, err = http.NewRequest(http.MethodPatch, "/vacations/"+strconv.FormatInt(ISpatch.ScheduleId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new vacation so we can see if he got changed : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/vacations/"+strconv.FormatInt(ISpatch.ScheduleId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbVacation handlers.ScheduleIntermediate
	if err = json.NewDecoder(rr.Body).Decode(&dbVacation); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(ISpatch, dbVacation) {
		t.Error("Roles are not the same")
	}
	globals.Log.Debug("PATCH /roles/{id} - PASSED")

	//
	//	DELETE /vacations/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodDelete, "/vacations/"+strconv.FormatInt(ISpatch.ScheduleId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new vacation so we can see if it got deleted : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/vacations/"+strconv.FormatInt(ISpatch.ScheduleId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	if err = json.NewDecoder(rr.Body).Decode(&dbVacation); err == nil {
		t.Error(err)
	}

	globals.Log.Debug("DELETE /vacations/{id} - PASSED")
}
