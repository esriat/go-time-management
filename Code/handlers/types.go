package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/datastores"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

type Env struct {
	DB datastores.IDatastore
}

type AppHandlerFunc func(http.ResponseWriter, *http.Request) *AppError

type AppError struct {
	Error   error
	Message string
	Code    int
}

type UserIntermediate struct {
	UserId               int64  `db:"user_id" json:"user_id"`
	ContractId           int64  `db:"contract_id" json:"contract_id"`
	RoleId               int64  `db:"role_id" json:"role_id"`
	Username             string `db:"username" json:"username"`
	LastName             string `db:"last_name" json:"last_name"`
	FirstName            string `db:"first_name" json:"first_name"`
	Mail                 string `db:"mail" json:"mail"`
	TheoricalHoursWorked int64  `db:"theorical_hours_worked" json:"theorical_hours_worked"`
	VacationHours        int64  `db:"vacation_hours" json:"vacation_hours"`
}

type ScheduleIntermediate struct {
	ScheduleId int64  `json:"schedule_id"`
	ProjectId  int64  `json:"project_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

func IntermediateToSchedule(SI ScheduleIntermediate) (model.Schedule, error) {
	var (
		startTime time.Time
		endTime   time.Time
		err       error
	)

	format := "2006-01-02 15:04:05"

	if startTime, err = time.Parse(format, SI.StartDate); err != nil {
		return model.Schedule{}, err
	}

	if endTime, err = time.Parse(format, SI.EndDate); err != nil {
		return model.Schedule{}, err
	}

	return model.Schedule{
		ScheduleId: SI.ScheduleId,
		ProjectId:  SI.ProjectId,
		StartDate:  sql.NullTime{Valid: true, Time: startTime},
		EndDate:    sql.NullTime{Valid: true, Time: endTime},
	}, nil
}

func ScheduleToIntermediate(S model.Schedule) ScheduleIntermediate {
	format := "2006-01-02 15:04:05"
	return ScheduleIntermediate{
		ScheduleId: S.ScheduleId,
		ProjectId:  S.ProjectId,
		StartDate:  S.StartDate.Time.Format(format),
		EndDate:    S.StartDate.Time.Format(format),
	}
}
