package model

import (
	"database/sql"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Schedule : Represents a period of time during which something happened.
/*	ProjectId : The id of the project this schedule is linked to.
	StartDate : The start date of this Schedule.
	EndDate : The end date of this schedule.
*/
type Schedule struct {
	ScheduleId int64        `db:"schedule_id" json:"schedule_id"`
	ProjectId  int64        `db:"project_id" json:"project_id"`
	StartDate  sql.NullTime `db:"start_date" json:"start_date"`
	EndDate    sql.NullTime `db:"end_date" json:"end_date"`
}

type Schedules []Schedule
