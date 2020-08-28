package datastores

import (
	"database/sql"

	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//  CreateCompanyProject(CP model.CompanyProject) error
/*  Creates a link between a Company and a Project.
    Can return an error
*/
func (db *ConcreteDatastore) CreateCompanyProject(CP model.CompanyProject) error {
	var (
		tx  *sql.Tx
		err error
	)

	// Preparing the request
	if tx, err = db.Begin(); err != nil {
		return err
	}

	// Setting up and executing the request
	request := `INSERT INTO CompanyProject(company_id, project_id) VALUES (?, ?)`
	if _, err = tx.Exec(request, CP.CompanyId, CP.ProjectId); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return errr
		}
		return err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return errr
		}
		return err
	}

	return nil
}

//  CreateCompanyUser(CU model.CompanyUser) error
/*  Creates a link between a Company and a User.
    Can return an error
*/
func (db *ConcreteDatastore) CreateCompanyUser(CU model.CompanyUser) error {
	var (
		tx  *sql.Tx
		err error
	)

	// Preparing the request
	if tx, err = db.Begin(); err != nil {
		return err
	}

	// Setting up and executing the request
	request := `INSERT INTO CompanyUser(company_id, user_id) VALUES (?, ?)`
	if _, err = tx.Exec(request, CU.CompanyId, CU.UserId); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return errr
		}
		return err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return errr
		}
		return err
	}

	return nil
}

//  CreateUserSchedule(US model.UserSchedule) error
/*  Creates a link between a User and a Schedule.
    Can return an error
*/
func (db *ConcreteDatastore) CreateUserSchedule(US model.UserSchedule) error {
	var (
		tx  *sql.Tx
		err error
	)

	// Preparing the request
	if tx, err = db.Begin(); err != nil {
		return err
	}

	// Setting up the request and executing it
	request := `INSERT INTO UserSchedule(user_id, schedule_id) VALUES (?, ?)`
	if _, err = tx.Exec(request, US.UserId, US.ScheduleId); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return errr
		}
		return err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return errr
		}
		return err
	}

	return nil
}

//  CreateUserFunction(UF model.UserFunction) error
/*  Creates a link between a User and a Function.
    Can return an error
*/
func (db *ConcreteDatastore) CreateUserFunction(UF model.UserFunction) error {
	var (
		tx  *sql.Tx
		err error
	)

	// Starting a request
	if tx, err = db.Begin(); err != nil {
		return err
	}

	// Setting up and executing the request
	request := `INSERT INTO UserFunction(user_id, function_id) VALUES (?, ?)`
	if _, err = tx.Exec(request, UF.UserId, UF.FunctionId); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return errr
		}
		return err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return errr
		}
		return err
	}

	return nil
}

//  DeleteCompanyProject(CP model.CompanyProject) error
/*  Deletes a link between a Company and a Project
    Can return an error
*/
func (db *ConcreteDatastore) DeleteCompanyProject(CP model.CompanyProject) error {
	// Setting up and executing a request
	request := `DELETE FROM CompanyProject 
	WHERE company_id=?
	AND project_id=?`
	if _, err := db.Exec(request, CP.CompanyId, CP.ProjectId); err != nil {
		return err
	}
	return nil
}

//  DeleteCompanyUser(CU model.CompanyUser) error
/*  Deletes a link between a Company and a User
    Can return an error
*/
func (db *ConcreteDatastore) DeleteCompanyUser(CU model.CompanyUser) error {
	// Setting up and executing a request
	request := `DELETE FROM CompanyUser 
	WHERE company_id=?
	AND user_id=?`
	if _, err := db.Exec(request, CU.CompanyId, CU.UserId); err != nil {
		return err
	}
	return nil
}

//  DeleteUserSchedule(US model.UserSchedule) error
/*  Deletes a link between a User and a Schedule
    Can return an error
*/
func (db *ConcreteDatastore) DeleteUserSchedule(US model.UserSchedule) error {
	// Setting up and executing a request
	request := `DELETE FROM UserSchedule 
	WHERE user_id=?
	AND schedule_id=?`
	if _, err := db.Exec(request, US.UserId, US.ScheduleId); err != nil {
		return err
	}
	return nil
}

//  DeleteUserFunction(UF model.UserFunction) error
/*  Deletes a link between a User and a Function
    Can return an error
*/
func (db *ConcreteDatastore) DeleteUserFunction(UF model.UserFunction) error {
	// Setting up and executing a request
	request := `DELETE FROM UserFunction 
	WHERE user_id=?
	AND function_id=?`
	if _, err := db.Exec(request, UF.UserId, UF.FunctionId); err != nil {
		return err
	}
	return nil
}
