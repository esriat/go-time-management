package handler_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

/*
	TESTED : GET /contracts
	TESTED : GET /contracts/{id}
	TESTED : GET /users/{id}/contracts
	TESTED : POST /contracts
	TESTED : PATCH /contracts/{id}
	TESTED : DELETE /contracts/{id}
*/
func TestContractHandler(t *testing.T) {
	var (
		err     error
		request *http.Request

		jsonObject []byte
	)

	// Creating a temporary struct to decode the json.
	var tmp struct {
		ContractId int `json:"contract_id"`
	}

	// Preparing for HTTP requests
	rr := httptest.NewRecorder()

	//
	//	GET /contracts
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/contracts", nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing the request
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbAllContracts model.Contracts
	if err = json.NewDecoder(rr.Body).Decode(&dbAllContracts); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(fakeContracts, dbAllContracts) {
		t.Error("Contracts are not the same")
	}

	globals.Log.Debug("GET /contracts - PASSED")

	//
	//	POST /contracts
	//

	contract1 := model.Contract{
		ContractName: "New test contract",
	}
	// Turning the object into JSON
	if jsonObject, err = json.Marshal(contract1); err != nil {
		t.Error(err.Error())
	}

	if request, err = http.NewRequest(http.MethodPost, "/contracts", bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)

	// Sending the request
	r.ServeHTTP(rr, request)

	// Checking the result
	if err = json.NewDecoder(rr.Body).Decode(&tmp); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("POST /contracts - PASSED")

	contract1.ContractId = int64(tmp.ContractId)

	//
	//	GET /contracts/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/contracts/"+strconv.FormatInt(contract1.ContractId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbContract1 model.Contract
	if err = json.NewDecoder(rr.Body).Decode(&dbContract1); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(contract1, dbContract1) {
		t.Error("Contracts are not the same")
	}
	globals.Log.Debug("GET /contracts/{id} - PASSED")

	//
	//	GET /users/{id}/contracts
	//

	// Fetching the contracts of user 1
	contractOfUser, _ := env.DB.GetContractOfUser(1)

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/users/1/contract", nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbContractOfUser model.Contract
	if err = json.NewDecoder(rr.Body).Decode(&dbContractOfUser); err != nil {
		t.Error(err)
	}

	// Verifying the result
	if !cmp.Equal(contractOfUser, dbContractOfUser) {
		t.Error("Contracts are not the same")
	}
	globals.Log.Debug("GET /users/{id}/contracts - PASSED")

	//
	//	PATCH /contracts/{id}
	//

	// Fetching a contract
	contract := fakeContracts[1]
	// Modifying it
	contract.ContractName = "Modified contract"
	// JSON-ing the object
	if jsonObject, err = json.Marshal(contract); err != nil {
		t.Error(err.Error())
	}

	// Creating the request
	if request, err = http.NewRequest(http.MethodPatch, "/contracts/"+strconv.FormatInt(contract.ContractId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new contract so we can see if he got changed : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/contracts/"+strconv.FormatInt(contract.ContractId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbContract model.Contract
	if err = json.NewDecoder(rr.Body).Decode(&dbContract); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(contract, dbContract) {
		t.Error("Contracts are not the same")
	}
	globals.Log.Debug("PATCH /contracts/{id} - PASSED")

	//
	//	DELETE /contracts/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodDelete, "/contracts/"+strconv.FormatInt(contract.ContractId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new contract so we can see if it got deleted : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/contracts/"+strconv.FormatInt(contract.ContractId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	if err = json.NewDecoder(rr.Body).Decode(&dbContract); err == nil {
		t.Error(err)
	}

	globals.Log.Debug("DELETE /contracts/{id} - PASSED")
}
