package handler_tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/handlers"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

/*
	TESTED : GET /schedules/{id}
	TESTED : GET /users/{id}/schedules
	TESTED : GET /projects/{id}/schedules
	TESTED : POST /schedules
	TESTED : PATCH /schedules/{id}
	TESTED : DELETE /schedules/{id}
*/
func TestScheduleHandler(t *testing.T) {
	var (
		request    *http.Request
		tmp        model.Schedule
		jsonObject []byte

		err error
	)

	// To store the ScheduleId when decoding json
	var tmpId struct {
		ScheduleId int `json:"schedule_id"`
	}

	// Preparing for HTTP requests
	rr := httptest.NewRecorder()

	//
	//	GET /users/{id}/schedules
	//

	// Getting schedules of user 1
	var schedulesOfUser1 model.Schedules
	if schedulesOfUser1, err = env.DB.GetSchedulesOfUser(1); err != nil {
		t.Error(err)
	}
	// Getting rid of the useless details
	for index, _ := range schedulesOfUser1 {
		GetRidOfScheduleDateDetails(&schedulesOfUser1[index])
	}

	// Creating a request
	if request, err = http.NewRequest(http.MethodGet, "/users/1/schedules", nil); err != nil {
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
	var dbSchedulesOfUser1 model.Schedules
	for _, SI := range SIarray1 {
		if tmp, err = handlers.IntermediateToSchedule(SI); err != nil {
			t.Error()
		}
		dbSchedulesOfUser1 = append(dbSchedulesOfUser1, tmp)
	}

	// Now comparing
	if !cmp.Equal(schedulesOfUser1, dbSchedulesOfUser1) {
		t.Error("Schedules are not the same")
	}

	globals.Log.Debug("GET /users/{id}/schedules - PASSED")

	//
	//	GET /projects/{id}/schedules
	//

	// Getting schedules of project 1
	var schedulesOfProject1 model.Schedules
	if schedulesOfProject1, err = env.DB.GetSchedulesOfProject(1); err != nil {
		t.Error(err)
	}
	// Getting rid of the useless details
	for index, _ := range schedulesOfProject1 {
		GetRidOfScheduleDateDetails(&schedulesOfProject1[index])
	}

	// Creating a request
	if request, err = http.NewRequest(http.MethodGet, "/projects/1/schedules", nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)

	// Sending the request
	r.ServeHTTP(rr, request)

	// Checking the result
	if err = json.NewDecoder(rr.Body).Decode(&SIarray1); err != nil {
		t.Error(err)
	}

	// Converting the result to model.Schedules
	var dbSchedulesOfProject1 model.Schedules
	for _, SI := range SIarray1 {
		if tmp, err = handlers.IntermediateToSchedule(SI); err != nil {
			t.Error()
		}
		dbSchedulesOfProject1 = append(dbSchedulesOfProject1, tmp)
	}

	// Now comparing
	if !cmp.Equal(schedulesOfProject1, dbSchedulesOfProject1) {
		t.Error("Schedules are not the same")
	}

	globals.Log.Debug("GET /projects/{id}/schedules - PASSED")

	//
	//	POST /schedules
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

	if request, err = http.NewRequest(http.MethodPost, "/schedules", bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)

	// Sending the request
	r.ServeHTTP(rr, request)

	// Checking the result
	if err = json.NewDecoder(rr.Body).Decode(&tmpId); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("POST /schedules - PASSED")

	SI2.ScheduleId = int64(tmpId.ScheduleId)

	//
	//	GET /schedules/{id}
	//

	// Getting the first schedule
	var schedule1 model.Schedule
	if schedule1, err = env.DB.GetSchedule(SI2.ScheduleId); err != nil {
		t.Error(err)
	}

	// Creating a request
	if request, err = http.NewRequest(http.MethodGet, "/schedules/"+strconv.FormatInt(SI2.ScheduleId, 10), nil); err != nil {
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
	var dbSchedule1 model.Schedule
	if dbSchedule1, err = handlers.IntermediateToSchedule(SI1); err != nil {
		t.Error()
	}
	GetRidOfScheduleDateDetails(&schedule1)

	// Comparing the objects
	if !cmp.Equal(schedule1, dbSchedule1) {
		t.Error("Schedules are not the same")
	}

	globals.Log.Debug("GET /schedules/{id} - PASSED")

	//
	//	PATCH /schedules/{id}
	//

	// Fetching a schedule
	ISpatch := handlers.ScheduleToIntermediate(fakeSchedules[2])
	// Modifying it
	ISpatch.StartDate = "2007-05-30 10:09:27"
	// JSON-ing the object
	if jsonObject, err = json.Marshal(ISpatch); err != nil {
		t.Error(err)
	}

	// Creating the request
	if request, err = http.NewRequest(http.MethodPatch, "/schedules/"+strconv.FormatInt(ISpatch.ScheduleId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new schedule so we can see if he got changed : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/schedules/"+strconv.FormatInt(ISpatch.ScheduleId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbSchedule handlers.ScheduleIntermediate
	if err = json.NewDecoder(rr.Body).Decode(&dbSchedule); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(ISpatch, dbSchedule) {
		t.Error("Roles are not the same")
	}
	globals.Log.Debug("PATCH /roles/{id} - PASSED")

	//
	//	DELETE /schedules/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodDelete, "/schedules/"+strconv.FormatInt(ISpatch.ScheduleId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new schedule so we can see if it got deleted : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/schedules/"+strconv.FormatInt(ISpatch.ScheduleId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	if err = json.NewDecoder(rr.Body).Decode(&dbSchedule); err == nil {
		t.Error(err)
	}

	globals.Log.Debug("DELETE /schedules/{id} - PASSED")
}

// This is used to get rid of the details of the dates
// Otherwise, we'll have like microseconds that prevent tests from passing
func GetRidOfScheduleDateDetails(S *model.Schedule) {
	var (
		format    string
		startDate string
		endDate   string

		StartDate time.Time
		EndDate   time.Time
	)

	// Setting up the format
	format = "2006-01-02 15:04:05"

	// Formatting the dates (by converting them to strings)
	startDate = S.StartDate.Time.Format(format)
	endDate = S.EndDate.Time.Format(format)

	// Converting the strings back to time.Time
	StartDate, _ = time.Parse(format, startDate)
	EndDate, _ = time.Parse(format, endDate)

	// Now setting the data in the schedule
	S.StartDate = sql.NullTime{Valid: true, Time: StartDate}
	S.EndDate = sql.NullTime{Valid: true, Time: EndDate}

	// No return cuz we're using a reference
}
