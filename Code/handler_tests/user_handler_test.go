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
	TESTED : GET /users
	TESTED : GET /users/{id}
	TESTED : GET /companies/{id}/users
	TESTED : GET /projects/{id}/users
	TESTED : GET /schedules/{id}/users
	TESTED : POST /users
	TESTED : PATCH /users/{id}
	TESTED : DELETE /users/{id}
*/
func TestUserHandler(t *testing.T) {
	var (
		err     error
		request *http.Request

		jsonObject []byte
	)

	// Creating a temporary struct to decode the json.
	var tmp struct {
		UserId int `json:"user_id"`
	}

	// Preparing for HTTP requests
	rr := httptest.NewRecorder()

	//
	//	GET /users
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/users", nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing the request
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbAllUsers model.Users
	if err = json.NewDecoder(rr.Body).Decode(&dbAllUsers); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(fakeUsers, dbAllUsers) {
		t.Error("Users are not the same")
	}

	globals.Log.Debug("GET /users - PASSED")

	//
	//	POST /users
	//

	user1 := model.User{
		RoleId:     1,
		ContractId: 1,
		Username:   "New test user",
		Mail:       "Newtestuser@mail",
	}
	// Turning the object into JSON
	if jsonObject, err = json.Marshal(user1); err != nil {
		t.Error(err)
	}

	if request, err = http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)

	// Sending the request
	r.ServeHTTP(rr, request)

	// Checking the result
	if err = json.NewDecoder(rr.Body).Decode(&tmp); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("POST /users - PASSED")

	user1.UserId = int64(tmp.UserId)

	//
	//	GET /users/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/users/"+strconv.FormatInt(user1.UserId, 10), nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbUser1 model.User
	if err = json.NewDecoder(rr.Body).Decode(&dbUser1); err != nil {
		t.Error(err)
	}

	// Doing that because the password got crypted, so we need to update it here
	user1.Password = dbUser1.Password

	if !cmp.Equal(user1, dbUser1) {
		t.Error("Users are not the same")
	}
	globals.Log.Debug("GET /users/{id} - PASSED")

	//
	//	GET /companies/{id}/users
	//

	// Fetching the users of company 1
	usersOfCompany, _ := env.DB.GetUsersOfCompany(1)

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/companies/1/users", nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbUsersOfCompany model.Users
	if err = json.NewDecoder(rr.Body).Decode(&dbUsersOfCompany); err != nil {
		t.Error(err)
	}

	// Verifying the result
	if !cmp.Equal(usersOfCompany, dbUsersOfCompany) {
		t.Error("Users are not the same")
	}
	globals.Log.Debug("GET /companies/{id}/users - PASSED")

	//
	//	GET /projects/{id}/users
	//
	// Fetching the users of project 1
	usersOfProject, _ := env.DB.GetUsersOfProject(1)

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/projects/1/users", nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbUsersOfProject model.Users
	if err = json.NewDecoder(rr.Body).Decode(&dbUsersOfProject); err != nil {
		t.Error(err)
	}

	// Verifying the result
	if !cmp.Equal(usersOfProject, dbUsersOfProject) {
		t.Error("Users are not the same")
	}
	globals.Log.Debug("GET /projects/{id}/users - PASSED")

	//
	//	GET /schedules/{id}/users
	//
	// Fetching the users of schedule 1
	usersOfSchedule, _ := env.DB.GetUsersOfSchedule(1)

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/schedules/1/users", nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbUsersOfSchedule model.Users
	if err = json.NewDecoder(rr.Body).Decode(&dbUsersOfSchedule); err != nil {
		t.Error(err)
	}

	// Verifying the result
	if !cmp.Equal(usersOfSchedule, dbUsersOfSchedule) {
		t.Error("Users are not the same")
	}
	globals.Log.Debug("GET /schedules/{id}/users - PASSED")

	//
	//	PATCH /users/{id}
	//

	// Fetching a user
	user := fakeUsers[1]
	// Modifying it
	user.Username = "Modified user"
	// JSON-ing the object
	if jsonObject, err = json.Marshal(user); err != nil {
		t.Error(err)
	}

	// Creating the request
	if request, err = http.NewRequest(http.MethodPatch, "/users/"+strconv.FormatInt(user.UserId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new user so we can see if he got changed : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/users/"+strconv.FormatInt(user.UserId, 10), nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbUser model.User
	if err = json.NewDecoder(rr.Body).Decode(&dbUser); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(user, dbUser) {
		t.Error("Users are not the same")
	}
	globals.Log.Debug("PATCH /users/{id} - PASSED")

	//
	//	DELETE /users/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodDelete, "/users/"+strconv.FormatInt(user.UserId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new user so we can see if it got deleted : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/users/"+strconv.FormatInt(user.UserId, 10), nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	if err = json.NewDecoder(rr.Body).Decode(&dbUser); err == nil {
		t.Error(err)
	}

	globals.Log.Debug("DELETE /users/{id} - PASSED")
}
