package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//	GetRolesHandler
/*	The handler called by the following endpoint : GET /roles
	This method is used to get the list of all roles.
*/
func (env *Env) GetRolesHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err   error
		roles model.Roles
	)

	globals.Log.Debug("Calling GetRolesHandler")

	if roles, err = env.DB.GetRoles(); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error retrieving all the roles",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(roles)
	return nil
}

//	GetRoleHandler
/*	The handler called by the following endpoint : GET /roles/{id}
	This method is used to get a role.
*/
func (env *Env) GetRoleHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err  error
		role model.Role
		id   int
	)

	globals.Log.Debug("Calling GetRoleHandler")

	vars := mux.Vars(r)

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if role, err = env.DB.GetRole(int64(id)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching role",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(role)
	return nil
}

//	GetRoleOfUserHandler
/*	The handler called by the following endpoint : GET /users/{user_id}/role
	This method is used to get the role of a user
*/
func (env *Env) GetRoleOfUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err    error
		role   model.Role
		userId int
	)

	globals.Log.Debug("Calling GetRoleOfUserHandler")

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if role, err = env.DB.GetRoleOfUser(int64(userId)); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error retrieving the role",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(role)
	return nil
}

//	CreateRolesHandler
/*	The handler called by the following endpoint : POST /roles
	This method is used to create a new role.
*/
func (env *Env) CreateRoleHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err    error
		role   model.Role
		roleId int64
	)

	globals.Log.Debug("CreateRoleHandler called")

	if err = json.NewDecoder(r.Body).Decode(&role); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	globals.Log.Debug("Calling CreateRole method")

	if roleId, err = env.DB.CreateRole(role); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error creating the role",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Role created")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(struct {
		RoleId int64 `json:"role_id"`
	}{
		RoleId: roleId,
	}); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the role id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	UpdateRoleHandler
/*	The handler called by the following endpoint : PATCH /roles/{id}
	This method is used to update an existing role.
	Trying to update one of the three basic roles (User, Admin and Superadmin) will result in an error.
*/
func (env *Env) UpdateRoleHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err            error
		role           model.Role
		superadminRole model.Role
		adminRole      model.Role
		userRole       model.Role
		roleId         int
	)

	globals.Log.Debug("CreateRoleHandler called")

	if superadminRole, err = env.DB.GetRoleByName("Superadmin"); err != nil {
		return &AppError{
			Error:   err,
			Message: "Internal error when checking role",
			Code:    http.StatusInternalServerError,
		}
	}

	if adminRole, err = env.DB.GetRoleByName("Admin"); err != nil {
		return &AppError{
			Error:   err,
			Message: "Internal error when checking role",
			Code:    http.StatusInternalServerError,
		}
	}

	if userRole, err = env.DB.GetRoleByName("User"); err != nil {
		return &AppError{
			Error:   err,
			Message: "Internal error when checking role",
			Code:    http.StatusInternalServerError,
		}
	}

	if err = json.NewDecoder(r.Body).Decode(&role); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	vars := mux.Vars(r)

	if roleId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	role.RoleId = int64(roleId)

	if role.RoleId == superadminRole.RoleId || role.RoleId == adminRole.RoleId || role.RoleId == userRole.RoleId {
		return &AppError{
			Message: "Modifying one of the base roles is not authorized",
			Code:    http.StatusUnauthorized,
		}
	}

	globals.Log.Debug("Calling CreateRole method")

	if role, err = env.DB.UpdateRole(role); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when updating the role",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Role updated")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(role); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the role id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	DeleteRoleHandler
/*	The handler called by the following endpoint : DELETE /roles/{id}
	This method is used to delete an existing role.
	Trying to delete one of the three basic roles (User, Admin and Superadmin) will result in an error.
*/
func (env *Env) DeleteRoleHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err    error
		roleId int
	)

	globals.Log.Debug("DeleteRoleHandler called")

	vars := mux.Vars(r)

	if roleId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Calling CreateRole method")

	if err = env.DB.DeleteRole(int64(roleId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when deleting the role",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Role deleted")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}
