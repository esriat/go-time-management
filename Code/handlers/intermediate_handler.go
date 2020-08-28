package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//	CreateCompanyUserHandler
/*	The handler called by the following endpoint : POST /companies/{company_id}/users/{user_id}
	This method is used to create a new link between a company and a user.
*/
func (env *Env) CreateCompanyUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		companyId int
		userId    int
	)

	globals.Log.Debug("Calling CreateCompanyUserHandler")

	vars := mux.Vars(r)

	if companyId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if userId, err = strconv.Atoi(vars["other_id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if err = env.DB.CreateCompanyUser(model.CompanyUser{
		CompanyId: int64(companyId),
		UserId:    int64(userId),
	}); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error creating a Company-User link",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}

//	DeleteCompanyUserHandler
/*	The handler called by the following endpoint : DELETE /companies/{company_id}/users/{user_id}
	This method is used to delete a link between a company and a user.
*/
func (env *Env) DeleteCompanyUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		companyId int
		userId    int
	)

	globals.Log.Debug("Calling CreateCompanyUserHandler")

	vars := mux.Vars(r)

	if companyId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if userId, err = strconv.Atoi(vars["other_id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if err = env.DB.DeleteCompanyUser(model.CompanyUser{
		CompanyId: int64(companyId),
		UserId:    int64(userId),
	}); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error creating a Company-User link",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}

//	CreateUserScheduleHandler
/*	The handler called by the following endpoint : POST /users/{user_id}/schedules/{schedule_id}
	This method is used to create a new link between a user and a schedule.
*/
func (env *Env) CreateUserScheduleHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		userId     int
		scheduleId int
	)

	globals.Log.Debug("Calling CreateUserScheduleHandler")

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if scheduleId, err = strconv.Atoi(vars["other_id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if err = env.DB.CreateUserSchedule(model.UserSchedule{
		UserId:     int64(userId),
		ScheduleId: int64(scheduleId),
	}); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error creating a User-Schedule link",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}

//	DeleteUserScheduleHandler
/*	The handler called by the following endpoint : DELETE /users/{user_id}/schedules/{schedule_id}
	This method is used to delete a link between a user and a schedule.
*/
func (env *Env) DeleteUserScheduleHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		userId     int
		scheduleId int
	)

	globals.Log.Debug("Calling DeleteUserScheduleHandler")

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if scheduleId, err = strconv.Atoi(vars["other_id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if err = env.DB.DeleteUserSchedule(model.UserSchedule{
		UserId:     int64(userId),
		ScheduleId: int64(scheduleId),
	}); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error deleting a Company-User link",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}

//	CreateCompanyProjectHandler
/*	The handler called by the following endpoint : POST /companies/{company_id}/projects/{project_id}
	This method is used to create a new link between a company and a project.
*/
func (env *Env) CreateCompanyProjectHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		companyId int
		projectId int
	)

	globals.Log.Debug("Calling CreateCompanyProjectHandler")

	vars := mux.Vars(r)

	if companyId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if projectId, err = strconv.Atoi(vars["other_id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if err = env.DB.CreateCompanyProject(model.CompanyProject{
		CompanyId: int64(companyId),
		ProjectId: int64(projectId),
	}); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error creating a Company-Project link",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}

//	DeleteCompanyProjectHandler
/*	The handler called by the following endpoint : DELETE /companies/{company_id}/projects/{project_id}
	This method is used to delete a link between a company and a project.
*/
func (env *Env) DeleteCompanyProjectHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		companyId int
		projectId int
	)

	globals.Log.Debug("Calling DeleteUserScheduleHandler")

	vars := mux.Vars(r)

	if companyId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if projectId, err = strconv.Atoi(vars["other_id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if err = env.DB.DeleteCompanyProject(model.CompanyProject{
		CompanyId: int64(companyId),
		ProjectId: int64(projectId),
	}); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error deleting a Company-Project link",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}

//	CreateUserFunctionHandler
/*	The handler called by the following endpoint : POST /users/{user_id}/functions/{function_id}
	This method is used to create a new link between a user and a function.
*/
func (env *Env) CreateUserFunctionHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		userId     int
		functionId int
	)

	globals.Log.Debug("Calling CreateUserFunctionHandler")

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if functionId, err = strconv.Atoi(vars["other_id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if err = env.DB.CreateUserFunction(model.UserFunction{
		UserId:     int64(userId),
		FunctionId: int64(functionId),
	}); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error creating a User-Function link",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}

//	DeleteUserFunctionHandler
/*	The handler called by the following endpoint : DELETE /users/{user_id}/functions/{function_id}
	This method is used to delete a link between a user and a function.
*/
func (env *Env) DeleteUserFunctionHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		userId     int
		functionId int
	)

	globals.Log.Debug("Calling DeleteUserFunctionHandler")

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if functionId, err = strconv.Atoi(vars["other_id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if err = env.DB.DeleteUserFunction(model.UserFunction{
		UserId:     int64(userId),
		FunctionId: int64(functionId),
	}); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error deleting a User-Function link",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}
