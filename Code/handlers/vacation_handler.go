package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//	GetVacationHandler
/*	The handler called by the following endpoint : GET /vacations
	This method is used to get a vacation.
*/
func (env *Env) GetVacationHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err      error
		vacation model.Schedule
		id       int
	)

	globals.Log.Debug("Calling GetVacationHandler")

	vars := mux.Vars(r)

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if vacation, err = env.DB.GetVacation(int64(id)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching vacation",
			Code:    http.StatusInternalServerError,
		}
	}

	format := "2006-01-02 15:04:05"
	startDate := vacation.StartDate.Time.Format(format)
	endDate := vacation.EndDate.Time.Format(format)

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ScheduleIntermediate{
		ScheduleId: vacation.ScheduleId,
		ProjectId:  vacation.ProjectId,
		StartDate:  startDate,
		EndDate:    endDate,
	})

	return nil
}

//  GetVacationsOfUserHandler
/*	The handler called by the following endpoint : GET /users/{user_id}/vacations
	This method is used to get the vacations of a user.
*/
func (env *Env) GetVacationsOfUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err           error
		vacations     model.Schedules
		intermediates []ScheduleIntermediate
		userId        int
	)

	globals.Log.Debug("Calling GetVacationsOfUserHandler")

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if vacations, err = env.DB.GetVacationsOfUser(int64(userId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching vacation",
			Code:    http.StatusInternalServerError,
		}
	}

	format := "2006-01-02 15:04:05"

	for _, vacation := range vacations {
		startDate := vacation.StartDate.Time.Format(format)
		endDate := vacation.EndDate.Time.Format(format)

		intermediates = append(intermediates, ScheduleIntermediate{
			ScheduleId: vacation.ScheduleId,
			ProjectId:  vacation.ProjectId,
			StartDate:  startDate,
			EndDate:    endDate,
		})
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(intermediates)

	return nil
}

//	GetVacationHaCreateVacationHandlerndler
/*	The handler called by the following endpoint : POST /vacations
	This method is used to create a vacation.
*/
func (env *Env) CreateVacationHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		vacation   model.Schedule
		vacationId int64
		startTime  time.Time
		endTime    time.Time
	)

	globals.Log.Debug("CreateVacationHandler called")

	intermediate := ScheduleIntermediate{}

	if err = json.NewDecoder(r.Body).Decode(&intermediate); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	format := "2006-01-02 15:04:05"

	if startTime, err = time.Parse(format, intermediate.StartDate); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error with the date",
			Code:    http.StatusInternalServerError,
		}
	}

	if endTime, err = time.Parse(format, intermediate.EndDate); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error with the date",
			Code:    http.StatusInternalServerError,
		}
	}

	vacation = model.Schedule{
		ProjectId: intermediate.ProjectId,
		StartDate: sql.NullTime{Valid: true, Time: startTime},
		EndDate:   sql.NullTime{Valid: true, Time: endTime},
	}

	globals.Log.Debug("Calling CreateVacation method")

	if vacationId, err = env.DB.CreateVacation(vacation); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error creating the vacation",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Vacation created")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(struct {
		VacationId int64 `json:"vacation_id"`
	}{
		VacationId: vacationId,
	}); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the vacation id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	UpdateVacationHandler
/*	The handler called by the following endpoint : PATCH /vacations/{id}
	This method is used to update an existing vacation.
*/
func (env *Env) UpdateVacationHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		schedule   model.Schedule
		scheduleId int
		startTime  time.Time
		endTime    time.Time
	)

	globals.Log.Debug("UpdateVacationHandler called")

	intermediate := ScheduleIntermediate{}

	if err = json.NewDecoder(r.Body).Decode(&intermediate); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	format := "2006-01-02 15:04:05"

	if startTime, err = time.Parse(format, intermediate.StartDate); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error with the date",
			Code:    http.StatusInternalServerError,
		}
	}

	if endTime, err = time.Parse(format, intermediate.EndDate); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error with the date",
			Code:    http.StatusInternalServerError,
		}
	}

	schedule = model.Schedule{
		ProjectId: intermediate.ProjectId,
		StartDate: sql.NullTime{Valid: true, Time: startTime},
		EndDate:   sql.NullTime{Valid: true, Time: endTime},
	}

	vars := mux.Vars(r)

	if scheduleId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	schedule.ScheduleId = int64(scheduleId)

	globals.Log.Debug("Calling UpdateVacation method")

	if schedule, err = env.DB.UpdateVacation(schedule); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when updating the schedule",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Vacation updated")

	startDate := schedule.StartDate.Time.Format(format)
	endDate := schedule.EndDate.Time.Format(format)

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ScheduleIntermediate{
		ScheduleId: schedule.ScheduleId,
		ProjectId:  schedule.ProjectId,
		StartDate:  startDate,
		EndDate:    endDate,
	})

	return nil
}

//	DeleteVacationHandler
/*	The handler called by the following endpoint : DELETE /vacations/{id}
	This method is used to delete a schedule.
*/
func (env *Env) DeleteVacationHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		scheduleId int
	)

	globals.Log.Debug("DeleteVacationHandler called")

	vars := mux.Vars(r)

	if scheduleId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Calling DeleteVacation method")

	if err = env.DB.DeleteVacation(int64(scheduleId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when deleting the schedule",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Vacation deleted")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}
