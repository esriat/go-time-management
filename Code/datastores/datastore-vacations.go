package datastores

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//  GetVacationsOfUser(UserId int64) (model.Schedules, error)
/*  This method is used to get the list of vacations of a specific user
    Returns the list of vacations of the wanted user or an error
*/
func (db *ConcreteDatastore) GetVacationsOfUser(UserId int64) (model.Schedules, error) {
	var (
		rows *sqlx.Rows
		err  error
	)

	// Executing the request
	request := `SELECT *
	FROM (SELECT *
		  FROM Schedule S, UserSchedule US
		  WHERE S.schedule_id = US.schedule_id
		  AND US.user_id=?) s
	WHERE s.project_id = (SELECT project_id
				  		  FROM Project
							WHERE project_name = "Vacation")`
	if rows, err = db.Queryx(request, UserId); err != nil {
		return nil, err
	}

	// Formatting the data
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

//  GetVacation(VacationId int64) (model.Schedule, error)
/*  This method is used to get a specific vacation
    Returns the wanted vacation or an error
*/
func (db *ConcreteDatastore) GetVacation(VacationId int64) (model.Schedule, error) {
	var (
		err             error
		schedule        model.Schedule
		vacationProject model.Project
	)

	// Fetching the "Vacation" project
	if vacationProject, err = db.GetVacationProject(); err != nil {
		return model.Schedule{}, err
	}

	// And now executing the request
	request := `SELECT schedule_id, project_id, start_date, end_date 
	FROM Schedule 
	WHERE schedule_id=?
	AND project_id=?`
	if err = db.Get(&schedule, request, VacationId, vacationProject.ProjectId); err != nil {
		return model.Schedule{}, err
	}

	return schedule, nil
}

// CreateVacation(Schedule model.Schedule) (int64, error)
/*	This method is used to create a new vacation.
 */
func (db *ConcreteDatastore) CreateVacation(Schedule model.Schedule) (int64, error) {
	var (
		tx              *sql.Tx
		err             error
		res             sql.Result
		scheduleId      int64
		vacationProject model.Project
	)

	// Starting
	if tx, err = db.Begin(); err != nil {
		return -1, err
	}

	// Fetching the vacation project
	if vacationProject, err = db.GetVacationProject(); err != nil {
		return -1, nil
	}

	// Executing the request
	request := `INSERT INTO Schedule(project_id, start_date, end_date)
	VALUES (?, ?, ?)`
	if res, err = tx.Exec(request, vacationProject.ProjectId, Schedule.StartDate.Time, Schedule.EndDate.Time); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Getting the id of the new vacation
	if scheduleId, err = res.LastInsertId(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Saving changes
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	return scheduleId, nil
}

//  DeleteVacation(VacationId int64) error
/*	This method is used to delete a vacation.
	It's the same method as DeleteSchedule, but we restrict the delete to the Vacation project only.
*/
func (db *ConcreteDatastore) DeleteVacation(VacationId int64) error {
	var (
		vacationProject model.Project
		err             error
	)
	if vacationProject, err = db.GetVacationProject(); err != nil {
		return err
	}
	request := `DELETE FROM Schedule 
	WHERE schedule_id=?
	AND project_id=?`
	if _, err = db.Exec(request, VacationId, vacationProject.ProjectId); err != nil {
		return err
	}
	return nil
}

//  UpdateVacation(Vacation model.Schedule) (model.Schedule, error)
/*	This method is used to update an existing vacation.
	It's basically the same method as UpdateSchedule, but restricts the update to the Vacation project only.
*/
func (db *ConcreteDatastore) UpdateVacation(Vacation model.Schedule) (model.Schedule, error) {
	var (
		tx              *sql.Tx
		err             error
		vacationProject model.Project
	)

	// Starting
	if tx, err = db.Begin(); err != nil {
		return model.Schedule{}, err
	}

	// Fetching the vacation project
	if vacationProject, err = db.GetVacationProject(); err != nil {
		return model.Schedule{}, nil
	}

	// Executing the request
	request := `UPDATE Schedule
	SET project_id=?, start_date=?, end_date=?
	WHERE schedule_id=?
	AND project_id=?`
	if _, err = tx.Exec(request, Vacation.ProjectId, Vacation.StartDate, Vacation.EndDate, Vacation.ScheduleId, vacationProject.ProjectId); err != nil {
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

	return Vacation, nil
}
