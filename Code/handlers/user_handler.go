package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"

	"golang.org/x/crypto/bcrypt"
)

//	GetUsersHandler
/*	The handler called by the following endpoint : GET /users
	This method is used to retrieve the list of users.
*/
func (env *Env) GetUsersHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err   error
		users model.Users
	)

	globals.Log.Debug("Calling GetUsersHandler")

	if users, err = env.DB.GetUsers(); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error retrieving all the users",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
	return nil
}

//	GetUserHandler
/*	The handler called by the following endpoint : GET /users/{id}
	This method is used to retrieve a user.
*/
func (env *Env) GetUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err  error
		user model.User
		id   int
	)

	globals.Log.Debug("Calling GetUserHandler")

	vars := mux.Vars(r)

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if user, err = env.DB.GetUser(int64(id)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching user",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
	return nil
}

//	GetUsersOfCompanyHandler
/*	The handler called by the following endpoint : GET /companies/{company_id}/users
	This method is used to retrieve the list of users that are linked to a company
*/
func (env *Env) GetUsersOfCompanyHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		users     model.Users
		companyId int
	)

	globals.Log.Debug("Calling GetUserOfCompanyHandler")

	vars := mux.Vars(r)

	if companyId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if users, err = env.DB.GetUsersOfCompany(int64(companyId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching user",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
	return nil
}

//	GetUsersOfScheduleHandler
/*	The handler called by the following endpoint : GET /schedules/{schedule_id}/users
	This method is used to retrieve the list of users that are linked to a schedule.
*/
func (env *Env) GetUsersOfScheduleHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		users      model.Users
		scheduleId int
	)

	globals.Log.Debug("Calling GetUsersOfScheduleHandler")

	vars := mux.Vars(r)

	if scheduleId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if users, err = env.DB.GetUsersOfSchedule(int64(scheduleId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching user",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
	return nil
}

//	GetUsersOfProjectHandler
/*	The handler called by the following endpoint : GET /projects/{project_id}/users
	This method is used to retrieve the list of users that are linked to a project.
*/
func (env *Env) GetUsersOfProjectHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		users     model.Users
		projectId int
	)

	globals.Log.Debug("Calling GetUsersOfProjectHandler")

	vars := mux.Vars(r)

	if projectId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if users, err = env.DB.GetUsersOfProject(int64(projectId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching user",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
	return nil
}

//	CreateUserHandler
/*	The handler called by the following endpoint : POST /users
	This method is used to create a user.
*/
func (env *Env) CreateUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err             error
		user            model.User
		userId          int64
		cryptedPassword []byte
	)

	globals.Log.Debug("CreateUserHandler called")

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	if cryptedPassword, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when crypting the password",
			Code:    http.StatusInternalServerError,
		}
	}
	user.Password = string(cryptedPassword)

	globals.Log.Debug("Calling CreateUser method")

	if userId, err = env.DB.CreateUser(user); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error creating the user",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("User created")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(struct {
		UserId int64 `json:"user_id"`
	}{
		UserId: userId,
	}); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the user id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	UpdateUserHandler
/*	The handler called by the following endpoint : PATCH /users/{id}
	This method is used to update an existing user.
*/
func (env *Env) UpdateUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err    error
		user   model.User
		dbUser model.User
		userId int
	)

	globals.Log.Debug("CreateUserHandler called")

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	user.UserId = int64(userId)

	if dbUser, err = env.DB.GetUser(user.UserId); err != nil {
		return &AppError{
			Error:   err,
			Message: "Unexisting user",
			Code:    http.StatusInternalServerError,
		}
	}

	user.Password = dbUser.Password

	globals.Log.Debug("Calling CreateUser method")

	if user, err = env.DB.UpdateUser(user); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when updating the user",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("User updated")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(user); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the user id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	DeleteUserHandler
/*	The handler called by the following endpoint : DELETE /users/{id}
	This method is used to delete an existing user.
*/
func (env *Env) DeleteUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err    error
		userId int
	)

	globals.Log.Debug("DeleteUserHandler called")

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Calling CreateUser method")

	if err = env.DB.DeleteUser(int64(userId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when deleting the user",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("User deleted")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}
