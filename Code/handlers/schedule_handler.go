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

//	GetScheduleHandler
/*	The handler called by the following endpoint : GET /schedules/{id}
	This method is used to get an existing schedule.
*/
func (env *Env) GetScheduleHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err      error
		schedule model.Schedule
		id       int
	)

	globals.Log.Debug("Calling GetScheduleHandler")

	vars := mux.Vars(r)

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if schedule, err = env.DB.GetSchedule(int64(id)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching schedule",
			Code:    http.StatusInternalServerError,
		}
	}

	format := "2006-01-02 15:04:05"
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

//	GetSchedulesOfUserHandler
/*	The handler called by the following endpoint : GET /users/{user_id}/schedules
	This method is used to get the list of schedules of a specific user
*/
func (env *Env) GetSchedulesOfUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err           error
		schedules     model.Schedules
		intermediates []ScheduleIntermediate
		userId        int
	)

	globals.Log.Debug("Calling GetSchedulesOfUserHandler")

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if schedules, err = env.DB.GetSchedulesOfUser(int64(userId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching schedules",
			Code:    http.StatusInternalServerError,
		}
	}

	format := "2006-01-02 15:04:05"

	for _, schedule := range schedules {
		startDate := schedule.StartDate.Time.Format(format)
		endDate := schedule.EndDate.Time.Format(format)

		intermediates = append(intermediates, ScheduleIntermediate{
			ScheduleId: schedule.ScheduleId,
			ProjectId:  schedule.ProjectId,
			StartDate:  startDate,
			EndDate:    endDate,
		})
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(intermediates)

	return nil
}

//	GetSchedulesOfProjectHandler
/*	The handler called by the following endpoint : GET /projects/{project_id}/schedules
	This method is used to get the list of schedules of a specific project
*/
func (env *Env) GetSchedulesOfProjectHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err           error
		schedules     model.Schedules
		intermediates []ScheduleIntermediate
		projectId     int
	)

	globals.Log.Debug("Calling GetSchedulesOfProjectHandler")

	vars := mux.Vars(r)

	if projectId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if schedules, err = env.DB.GetSchedulesOfProject(int64(projectId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching schedules",
			Code:    http.StatusInternalServerError,
		}
	}

	format := "2006-01-02 15:04:05"

	for _, schedule := range schedules {
		startDate := schedule.StartDate.Time.Format(format)
		endDate := schedule.EndDate.Time.Format(format)

		intermediates = append(intermediates, ScheduleIntermediate{
			ScheduleId: schedule.ScheduleId,
			ProjectId:  schedule.ProjectId,
			StartDate:  startDate,
			EndDate:    endDate,
		})
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(intermediates)

	return nil
}

//	CreateScheduleHandler
/*	The handler called by the following endpoint : POST /schedules
	This method is used to create a new schedule.
*/
func (env *Env) CreateScheduleHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		schedule   model.Schedule
		scheduleId int64
		startTime  time.Time
		endTime    time.Time
	)

	globals.Log.Debug("CreateScheduleHandler called")

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

	globals.Log.Debug("Calling CreateSchedule method")

	if scheduleId, err = env.DB.CreateSchedule(schedule); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error creating the schedule",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Schedule created")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(struct {
		ScheduleId int64 `json:"schedule_id"`
	}{
		ScheduleId: scheduleId,
	}); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the schedule id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	UpdateScheduleHandler
/*	The handler called by the following endpoint : PATCH /schedules/{id}
	This method is used to update an existing schedule.
*/
func (env *Env) UpdateScheduleHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		schedule   model.Schedule
		scheduleId int
		startTime  time.Time
		endTime    time.Time
	)

	globals.Log.Debug("CreateScheduleHandler called")

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

	globals.Log.Debug("Calling UpdateSchedule method")

	if schedule, err = env.DB.UpdateSchedule(schedule); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when updating the schedule",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Schedule updated")

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

//	DeleteScheduleHandler
/*	The handler called by the following endpoint : DELETE /schedules/{id}
	This method is used to delete a schedule.
*/
func (env *Env) DeleteScheduleHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		scheduleId int
	)

	globals.Log.Debug("DeleteScheduleHandler called")

	vars := mux.Vars(r)

	if scheduleId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Calling DeleteSchedule method")

	if err = env.DB.DeleteSchedule(int64(scheduleId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when deleting the schedule",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Schedule deleted")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}
