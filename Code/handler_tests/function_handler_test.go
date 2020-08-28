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
	TESTED : GET /functions
	TESTED : GET /functions/{id}
	TESTED : GET /users/{id}/functions
	TESTED : POST /functions
	TESTED : PATCH /functions/{id}
	TESTED : DELETE /functions/{id}
*/
func TestFunctionHandler(t *testing.T) {
	var (
		err     error
		request *http.Request

		jsonObject []byte
	)

	// Creating a temporary struct to decode the json.
	var tmp struct {
		FunctionId int `json:"function_id"`
	}

	// Preparing for HTTP requests
	rr := httptest.NewRecorder()

	//
	//	GET /functions
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/functions", nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing the request
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbAllFunctions model.Functions
	if err = json.NewDecoder(rr.Body).Decode(&dbAllFunctions); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(fakeFunctions, dbAllFunctions) {
		t.Error("Functions are not the same")
	}

	globals.Log.Debug("GET /functions - PASSED")

	//
	//	POST /functions
	//

	function1 := model.Function{
		FunctionName: "New test function",
	}
	// Turning the object into JSON
	if jsonObject, err = json.Marshal(function1); err != nil {
		t.Error(err.Error())
	}

	if request, err = http.NewRequest(http.MethodPost, "/functions", bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)

	// Sending the request
	r.ServeHTTP(rr, request)

	// Checking the result
	if err = json.NewDecoder(rr.Body).Decode(&tmp); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("POST /functions - PASSED")

	function1.FunctionId = int64(tmp.FunctionId)

	//
	//	GET /functions/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/functions/"+strconv.FormatInt(function1.FunctionId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbFunction1 model.Function
	if err = json.NewDecoder(rr.Body).Decode(&dbFunction1); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(function1, dbFunction1) {
		t.Error("Functions are not the same")
	}
	globals.Log.Debug("GET /functions/{id} - PASSED")

	//
	//	GET /users/{id}/functions
	//

	// Fetching the functions of user 1
	functionsOfUser, _ := env.DB.GetFunctionsOfUser(1)

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/users/1/functions", nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbFunctionsOfUser model.Functions
	if err = json.NewDecoder(rr.Body).Decode(&dbFunctionsOfUser); err != nil {
		t.Error(err)
	}

	// Verifying the result
	if !cmp.Equal(functionsOfUser, dbFunctionsOfUser) {
		t.Error("Functions are not the same")
	}
	globals.Log.Debug("GET /users/{id}/functions - PASSED")

	//
	//	PATCH /functions/{id}
	//

	// Fetching a function
	function := fakeFunctions[1]
	// Modifying it
	function.FunctionName = "Modified function"
	// JSON-ing the object
	if jsonObject, err = json.Marshal(function); err != nil {
		t.Error(err.Error())
	}

	// Creating the request
	if request, err = http.NewRequest(http.MethodPatch, "/functions/"+strconv.FormatInt(function.FunctionId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new function so we can see if he got changed : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/functions/"+strconv.FormatInt(function.FunctionId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbFunction model.Function
	if err = json.NewDecoder(rr.Body).Decode(&dbFunction); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(function, dbFunction) {
		t.Error("Functions are not the same")
	}
	globals.Log.Debug("PATCH /functions/{id} - PASSED")

	//
	//	DELETE /functions/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodDelete, "/functions/"+strconv.FormatInt(function.FunctionId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new function so we can see if it got deleted : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/functions/"+strconv.FormatInt(function.FunctionId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	if err = json.NewDecoder(rr.Body).Decode(&dbFunction); err == nil {
		t.Error(err)
	}

	globals.Log.Debug("DELETE /functions/{id} - PASSED")
}
