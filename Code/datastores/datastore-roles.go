package datastores

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//  GetRoles() (model.Roles, error)
/*	This method is used to get all existing roles from the database.
 */
func (db *ConcreteDatastore) GetRoles() (model.Roles, error) {
	var (
		rows *sqlx.Rows
		err  error
	)

	// Executing the request
	request := "SELECT * FROM Role;"
	if rows, err = db.Queryx(request); err != nil {
		return nil, err
	}

	// Formatting the data

	rolesList := model.Roles{}
	for rows.Next() {
		role := model.Role{}
		if err = rows.StructScan(&role); err != nil {
			return nil, err
		}
		rolesList = append(rolesList, role)
	}

	defer rows.Close()
	return rolesList, nil
}

//  GetRole(RoleId int64) (model.Role, error)
/*	This method is used to get a specific role from the database.
 */
func (db *ConcreteDatastore) GetRole(RoleId int64) (model.Role, error) {
	var (
		err  error
		role model.Role
	)

	// Executing the request
	request := `SELECT * FROM Role WHERE role_id=?`
	if err = db.Get(&role, request, RoleId); err != nil {
		return model.Role{}, err
	}

	return role, nil
}

//  GetRolesOfUser(UserId int64) (model.Roles, error)
/*	This method is used to get all the roles that a given User has.
 */
func (db *ConcreteDatastore) GetRoleOfUser(UserId int64) (model.Role, error) {
	var (
		err  error
		role model.Role
	)

	// Executing the request
	request := `SELECT * 
				FROM Role
				WHERE role_id=(SELECT role_id
							   FROM User
							   WHERE user_id=?);`
	if err = db.Get(&role, request, UserId); err != nil {
		return model.Role{}, err
	}

	return role, nil
}

//	GetRoleByName(RoleName string) (model.Role, error)
/*	This method is used to fetch a role by its name.
	Returns the wanted role or an error
*/
func (db *ConcreteDatastore) GetRoleByName(RoleName string) (model.Role, error) {
	var (
		err  error
		role model.Role
	)

	// Executing the request
	request := `SELECT * 
				FROM Role
				WHERE role_name=?`
	if err = db.Get(&role, request, RoleName); err != nil {
		return model.Role{}, err
	}

	return role, nil
}

//  CreateRole(Role model.Role) (int64, error)
/*	This method is used to create a new role.
 */
func (db *ConcreteDatastore) CreateRole(Role model.Role) (int64, error) {
	var (
		tx     *sql.Tx
		err    error
		res    sql.Result
		roleId int64
	)

	// Preparing
	if tx, err = db.Begin(); err != nil {
		return -1, err
	}

	// Setting up and executing the request
	request := `INSERT INTO Role(role_name, can_add_and_modify_users, can_see_other_schedules, can_add_projects, can_see_reports) VALUES (?, ?, ?, ?, ?)`
	if res, err = tx.Exec(request, Role.RoleName, Role.CanAddAndModifyUsers, Role.CanSeeOtherSchedules, Role.CanAddProjects, Role.CanSeeReports); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Getting the id of the new role
	if roleId, err = res.LastInsertId(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return 0, errr
		}
		return 0, err
	}

	return roleId, nil
}

//  DeleteRole(RoleId int64) error
/*	This method is used to delete a role
 */
func (db *ConcreteDatastore) DeleteRole(RoleId int64) error {
	request := `DELETE FROM Role 
	WHERE role_id=?`
	if _, err := db.Exec(request, RoleId); err != nil {
		return err
	}
	return nil
}

//  UpdateRole(Role model.Role) (model.Role, error)
/*	This method is used to update an existing role
 */
func (db *ConcreteDatastore) UpdateRole(Role model.Role) (model.Role, error) {
	var (
		tx  *sql.Tx
		err error
	)

	// Starting
	if tx, err = db.Begin(); err != nil {
		return model.Role{}, err
	}

	// Executing the request
	request := `UPDATE Role 
	SET role_name=?, can_add_and_modify_users=?, can_see_other_schedules=?, can_add_projects=?, can_see_reports=?
	WHERE role_id=?`
	if _, err = tx.Exec(request, Role.RoleName, Role.CanAddAndModifyUsers, Role.CanSeeOtherSchedules, Role.CanAddProjects, Role.CanSeeReports, Role.RoleId); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Role{}, errr
		}
		return model.Role{}, err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Role{}, errr
		}
		return model.Role{}, err
	}

	return Role, nil
}
