package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//	GetContractsHandler
/*	The handler called by the following endpoint : GET /contracts
	This method simply fetches the list of contracts.
*/
func (env *Env) GetContractsHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		contracts model.Contracts
	)

	globals.Log.Debug("Calling GetContractsHandler")

	if contracts, err = env.DB.GetContracts(); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error retrieving all the contracts",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(contracts)
	return nil
}

//	GetContractsHandler
/*	The handler called by the following endpoint : GET /contracts/{id}
	This method fetches the wanted contract.
*/
func (env *Env) GetContractHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err      error
		contract model.Contract
		id       int
	)

	globals.Log.Debug("Calling GetContractHandler")

	vars := mux.Vars(r)

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if contract, err = env.DB.GetContract(int64(id)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching contract",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(contract)
	return nil
}

//	GetContractOfUserHandler
/*	The handler called by the following endpoint : GET /users/{user_id}/contract
	This method fetches the contract of a user.
*/
func (env *Env) GetContractOfUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err      error
		contract model.Contract
		userId   int
	)

	globals.Log.Debug("Calling GetContractOfUserHandler")

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if contract, err = env.DB.GetContractOfUser(int64(userId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching contract",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(contract)
	return nil
}

//	CreateContractHandler
/*	The handler called by the following endpoint : POST /contracts
	This method is used to create a new contract.
*/
func (env *Env) CreateContractHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		contract   model.Contract
		contractId int64
	)

	globals.Log.Debug("CreateContractHandler called")

	if err = json.NewDecoder(r.Body).Decode(&contract); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	globals.Log.Debug("Decoded contract : " + contract.String())
	globals.Log.Debug("Calling CreateContract method")

	if contractId, err = env.DB.CreateContract(contract); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error creating the contract",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Contract created")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(struct {
		ContractId int64 `json:"contract_id"`
	}{
		ContractId: contractId,
	}); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the contract id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	CreateContractHandler
/*	The handler called by the following endpoint : PATCH /contracts/{id}
	This method is used to update an existing contract.
*/
func (env *Env) UpdateContractHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		contract   model.Contract
		contractId int
	)

	globals.Log.Debug("CreateContractHandler called")

	if err = json.NewDecoder(r.Body).Decode(&contract); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	vars := mux.Vars(r)

	if contractId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	contract.ContractId = int64(contractId)

	globals.Log.Debug("Decoded contract : " + contract.String())
	globals.Log.Debug("Calling CreateContract method")

	if contract, err = env.DB.UpdateContract(contract); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when updating the contract",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Contract updated")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(contract); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the contract id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	DeleteContractHandler
/*	The handler called by the following endpoint : DELETE /contracts/{id}
	This method is used to delete an existing contract.
*/
func (env *Env) DeleteContractHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		contractId int
	)

	globals.Log.Debug("DeleteContractHandler called")

	vars := mux.Vars(r)

	if contractId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Calling CreateContract method")

	if err = env.DB.DeleteContract(int64(contractId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when deleting the contract",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Contract deleted")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}
