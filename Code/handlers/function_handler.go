package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//	GetFunctionsHandler
/*	The handler called by the following endpoint : GET /functions
	This method is used to retrieve the list of all existing functions.
*/
func (env *Env) GetFunctionsHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		functions model.Functions
	)

	globals.Log.Debug("Calling GetFunctionsHandler")

	if functions, err = env.DB.GetFunctions(); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error retrieving all the functions",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(functions)
	return nil
}

//	GetFunctionHandler
/*	The handler called by the following endpoint : GET /functions/{id}
	This method is used to get a specific function.
*/
func (env *Env) GetFunctionHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err      error
		function model.Function
		id       int
	)

	globals.Log.Debug("Calling GetFunctionHandler")

	vars := mux.Vars(r)

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if function, err = env.DB.GetFunction(int64(id)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching function",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(function)
	return nil
}

//	GetFunctionsOfUserHandler
/*	The handler called by the following endpoint : GET /users/{user_id}/functions
	This method is used to retrieve the list of all functions of a user.
*/
func (env *Env) GetFunctionsOfUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		functions model.Functions
		userId    int
	)

	globals.Log.Debug("Calling GetFunctionsOfUserHandler")

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if functions, err = env.DB.GetFunctionsOfUser(int64(userId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching function",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(functions)
	return nil
}

//	CreateFunctionHandler
/*	The handler called by the following endpoint : POST /functions
	This method is used to create a new function.
*/
func (env *Env) CreateFunctionHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		function   model.Function
		functionId int64
	)

	globals.Log.Debug("CreateFunctionHandler called")

	if err = json.NewDecoder(r.Body).Decode(&function); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	globals.Log.Debug("Decoded function : " + function.String())
	globals.Log.Debug("Calling CreateFunction method")

	if functionId, err = env.DB.CreateFunction(function); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error creating the function",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Function created")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(struct {
		FunctionId int64 `json:"function_id"`
	}{
		FunctionId: functionId,
	}); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the function id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	UpdateFunctionHandler
/*	The handler called by the following endpoint : PATCH /functions/{id}
	This method is used to update an existing function.
*/
func (env *Env) UpdateFunctionHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		function   model.Function
		functionId int
	)

	globals.Log.Debug("CreateFunctionHandler called")

	if err = json.NewDecoder(r.Body).Decode(&function); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	vars := mux.Vars(r)

	if functionId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	function.FunctionId = int64(functionId)

	globals.Log.Debug("Decoded function : " + function.String())
	globals.Log.Debug("Calling CreateFunction method")

	if function, err = env.DB.UpdateFunction(function); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when updating the function",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Function updated")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(function); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the function id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	DeleteFunctionHandler
/*	The handler called by the following endpoint : DELETE /functions/{id}
	This method is used to delete an existing function.
*/
func (env *Env) DeleteFunctionHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		functionId int
	)

	globals.Log.Debug("DeleteFunctionHandler called")

	vars := mux.Vars(r)

	if functionId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Calling CreateFunction method")

	if err = env.DB.DeleteFunction(int64(functionId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when deleting the function",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Function deleted")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}
