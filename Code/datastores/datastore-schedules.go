package datastores

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//  GetSchedule(ScheduleId) (model.Schedule, error)
/*	This method is used to get a specific schedule from the database.
 */
func (db *ConcreteDatastore) GetSchedule(ScheduleId int64) (model.Schedule, error) {
	var (
		err      error
		schedule model.Schedule
	)

	request := `SELECT schedule_id, project_id, start_date, end_date 
	FROM Schedule 
	WHERE Schedule.schedule_id=?`
	if err = db.Get(&schedule, request, ScheduleId); err != nil {
		return model.Schedule{}, err
	}

	return schedule, nil
}

//  GetSchedulesOfUser(UserId int64) (model.Schedules, error)
/*	This method is used to get the schedules linked to a uesr.
 */
func (db *ConcreteDatastore) GetSchedulesOfUser(UserId int64) (model.Schedules, error) {
	var (
		rows *sqlx.Rows
		err  error
	)

	// Executing the request
	request := `SELECT *
	FROM Schedule S, UserSchedule US
	WHERE S.schedule_id = US.schedule_id
	AND US.user_id=?`
	if rows, err = db.Queryx(request, UserId); err != nil {
		return nil, err
	}

	// Formatting
	vacationList := model.Schedules{}
	for rows.Next() {
		vacation := model.Schedule{}
		if err = rows.StructScan(&vacation); err != nil {
			return nil, err

		}
		vacationList = append(vacationList, vacation)
	}

	return vacationList, nil
}

//  GetSchedulesOfProject(ProjectId int64) (model.Schedules, error)
/*	This method is used to get the project of a project
 */
func (db *ConcreteDatastore) GetSchedulesOfProject(ProjectId int64) (model.Schedules, error) {
	var (
		rows *sqlx.Rows
		err  error
	)

	// Executing the requet
	request := `SELECT *
	FROM Schedule S
	WHERE project_id=?`
	if rows, err = db.Queryx(request, ProjectId); err != nil {
		return nil, err
	}

	// Formatting
	vacationList := model.Schedules{}
	for rows.Next() {
		vacation := model.Schedule{}
		if err = rows.StructScan(&vacation); err != nil {
			return nil, err

		}
		vacationList = append(vacationList, vacation)
	}

	return vacationList, nil
}

//  CreateSchedule(Schedule model.Schedule) (int64, error)
/*	This method is used to create a new schedule.
 */
func (db *ConcreteDatastore) CreateSchedule(Schedule model.Schedule) (int64, error) {
	var (
		tx         *sql.Tx
		err        error
		res        sql.Result
		scheduleId int64
	)

	// Starting
	if tx, err = db.Begin(); err != nil {
		return -1, err
	}

	// Executing the request
	request := `INSERT INTO Schedule(project_id, start_date, end_date) VALUES (?, ?, ?)`
	if res, err = tx.Exec(request, Schedule.ProjectId, Schedule.StartDate.Time, Schedule.EndDate.Time); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Getting the id of the new item
	if scheduleId, err = res.LastInsertId(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	return scheduleId, nil
}

//  DeleteSchedule(ScheduleId int64) error
/*	This method is used to delete a schedule
 */
func (db *ConcreteDatastore) DeleteSchedule(ScheduleId int64) error {
	request := `DELETE FROM Schedule 
	WHERE schedule_id=?`
	if _, err := db.Exec(request, ScheduleId); err != nil {
		return err
	}
	return nil
}

//  UpdateSchedule(Schedule model.Schedule) (model.Schedule, error)
/*	This method is used to update an existing schedule
 */
func (db *ConcreteDatastore) UpdateSchedule(Schedule model.Schedule) (model.Schedule, error) {
	var (
		tx  *sql.Tx
		err error
	)

	// Starting
	if tx, err = db.Begin(); err != nil {
		return model.Schedule{}, err
	}

	// Executing the request
	request := `UPDATE Schedule
	SET project_id=?, start_date=?, end_date=?
	WHERE schedule_id=?`
	if _, err = tx.Exec(request, Schedule.ProjectId, Schedule.StartDate, Schedule.EndDate, Schedule.ScheduleId); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Schedule{}, errr
		}
		return model.Schedule{}, err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Schedule{}, errr
		}
		return model.Schedule{}, err
	}

	return Schedule, nil
}
