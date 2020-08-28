package datastores

import (
	"database/sql"

	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//  GetUsers() (model.Users, error)
/*  This method is used to get the list of all the users
    Returns the list of all users or an error
*/
func (db *ConcreteDatastore) GetUsers() (model.Users, error) {
	// Preparing the request and executing it
	request := "SELECT * FROM User"
	rows, err := db.Queryx(request)
	if err != nil {
		return nil, err
	}

	// Formatting the data
	usersList := model.Users{}

	for rows.Next() {
		user := model.User{}
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}

	return usersList, nil
}

//  GetUser(RoleId int64) (model.Role, error)
/*	This method is used to get a specific user from the database.
 */
func (db *ConcreteDatastore) GetUser(UserId int64) (model.User, error) {
	var (
		err  error
		user model.User
	)

	// Setting up and executing the request
	request := `SELECT * FROM User WHERE user_id=?`
	if err = db.Get(&user, request, UserId); err != nil {
		return model.User{}, err
	}

	return user, nil
}

//  GetUsersOfCompany(CompanyId int64) (model.Users, error)
/*	This method is used to get the users of a company.
 */
func (db *ConcreteDatastore) GetUsersOfCompany(CompanyId int64) (model.Users, error) {
	// Executing the request
	request := `SELECT *
				FROM User, CompanyUser
				WHERE User.user_id = CompanyUser.user_id
				AND CompanyUser.company_id=?`
	rows, err := db.Queryx(request, CompanyId)
	if err != nil {
		return nil, err
	}

	// Formatting the data
	usersList := model.Users{}
	for rows.Next() {
		user := model.User{}
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}

	return usersList, nil
}

// 	GetUserFromEmail(Email string) (model.User, error)
/*	This method is used to fetch a user from his email
	Returns the user or an error
*/
func (db *ConcreteDatastore) GetUserFromEmail(Email string) (model.User, error) {
	// Setting up and executing the request
	var (
		err  error
		user model.User
	)

	// Setting up and executing the request
	request := `SELECT * FROM User WHERE mail=?`
	if err = db.Get(&user, request, Email); err != nil {
		return model.User{}, err
	}

	return user, nil
}

//  GetUsersOfProject(CompanyId int64) (model.Users, error)
/*	This method is used to get the users of a project.
 */
func (db *ConcreteDatastore) GetUsersOfProject(ProjectId int64) (model.Users, error) {
	// Executing the request
	request := `SELECT *
				FROM User, UserSchedule, Schedule 
				WHERE User.user_id = UserSchedule.user_id
				AND UserSchedule.schedule_id=Schedule.schedule_id
				AND Schedule.project_id=?`

	rows, err := db.Queryx(request, ProjectId)
	if err != nil {
		return nil, err
	}

	usersList := model.Users{}

	// Formatting the data
	for rows.Next() {
		user := model.User{}
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}

	return usersList, nil
}

//  GetUsersOfSchedule(ScheduleId int64) (model.Users, error)
/*	This method is used to get the users of a schedule.
 */
func (db *ConcreteDatastore) GetUsersOfSchedule(ScheduleId int64) (model.Users, error) {
	// Executing the request
	request := `SELECT *
				FROM User, UserSchedule
				WHERE User.user_id = UserSchedule.user_id
				AND UserSchedule.schedule_id=?`
	rows, err := db.Queryx(request, ScheduleId)
	if err != nil {
		return nil, err
	}

	// Formatting the data
	usersList := model.Users{}

	for rows.Next() {
		user := model.User{}
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}

	return usersList, nil
}

//  CreateUser(User model.User) (int64, error)
/*	This method is used to create a new user.
 */
func (db *ConcreteDatastore) CreateUser(User model.User) (int64, error) {
	var (
		tx     *sql.Tx
		err    error
		res    sql.Result
		userId int64
	)

	// Starting
	if tx, err = db.Begin(); err != nil {
		return -1, err
	}

	// Executing the request
	request := `INSERT INTO User(contract_id, role_id, username, password, last_name, first_name, mail, theorical_hours_worked, vacation_hours) 
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if res, err = tx.Exec(request, User.ContractId, User.RoleId, User.Username, User.Password, User.LastName, User.FirstName, User.Mail, User.TheoricalHoursWorked, User.VacationHours); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Getting the id of the new user
	if userId, err = res.LastInsertId(); err != nil {
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

	return userId, nil
}

//  DeleteUser(UserId int64) error
/*	This method is used to delete a user
 */
func (db *ConcreteDatastore) DeleteUser(UserId int64) error {
	request := `DELETE FROM User 
	WHERE user_id=?`
	if _, err := db.Exec(request, UserId); err != nil {
		return err
	}
	return nil
}

//  UpdateUser(User model.User) (model.User, error)
/*	This method is used to update an existing user
 */
func (db *ConcreteDatastore) UpdateUser(User model.User) (model.User, error) {
	var (
		tx  *sql.Tx
		err error
	)

	// Starting
	if tx, err = db.Begin(); err != nil {
		return model.User{}, err
	}

	// Executing the request
	request := `UPDATE User
	SET contract_id=?, role_id=?, username=?, password=?, last_name=?, first_name=?, mail=?, theorical_hours_worked=?, vacation_hours=? 
	WHERE user_id =?`
	if _, err = tx.Exec(request, User.ContractId, User.RoleId, User.Username, User.Password, User.LastName, User.FirstName, User.Mail, User.TheoricalHoursWorked, User.VacationHours, User.UserId); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.User{}, errr
		}
		return model.User{}, err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.User{}, errr
		}
		return model.User{}, err
	}

	return User, nil
}
