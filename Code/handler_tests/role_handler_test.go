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
	TESTED : GET /roles
	TESTED : GET /roles/{id}
	TESTED : GET /users/{id}/roles
	TESTED : POST /roles
	TESTED : PATCH /roles/{id}
	TESTED : DELETE /roles/{id}
*/
func TestRoleHandler(t *testing.T) {
	var (
		err     error
		request *http.Request

		jsonObject []byte
	)

	// Creating a temporary struct to decode the json.
	var tmp struct {
		RoleId int `json:"role_id"`
	}

	// Preparing for HTTP requests
	rr := httptest.NewRecorder()

	//
	//	GET /roles
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/roles", nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing the request
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbAllRoles model.Roles
	if err = json.NewDecoder(rr.Body).Decode(&dbAllRoles); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(allRoles, dbAllRoles) {
		t.Error("Roles are not the same")
	}

	globals.Log.Debug("GET /roles - PASSED")

	//
	//	POST /roles
	//

	role1 := model.Role{
		RoleName:             "New role",
		CanAddAndModifyUsers: false,
		CanSeeOtherSchedules: true,
		CanAddProjects:       false,
		CanSeeReports:        true,
	}

	// Turning the object into JSON
	if jsonObject, err = json.Marshal(role1); err != nil {
		t.Error(err.Error())
	}

	if request, err = http.NewRequest(http.MethodPost, "/roles", bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)

	// Sending the request
	r.ServeHTTP(rr, request)

	// Checking the result
	if err = json.NewDecoder(rr.Body).Decode(&tmp); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("POST /roles - PASSED")

	role1.RoleId = int64(tmp.RoleId)
	allRoles = append(allRoles, role1)

	//
	//	GET /roles/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/roles/"+strconv.FormatInt(role1.RoleId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbRole1 model.Role
	if err = json.NewDecoder(rr.Body).Decode(&dbRole1); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(role1, dbRole1) {
		t.Error("Roles are not the same")
	}
	globals.Log.Debug("GET /roles/{id} - PASSED")

	//
	//	GET /users/{id}/roles
	//

	// Fetching the roles of user 1
	roleOfUser, _ := env.DB.GetRoleOfUser(1)

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/users/1/role", nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbRoleOfUser model.Role
	if err = json.NewDecoder(rr.Body).Decode(&dbRoleOfUser); err != nil {
		t.Error(err)
	}

	// Verifying the result
	if !cmp.Equal(roleOfUser, dbRoleOfUser) {
		t.Error("Roles are not the same")
	}
	globals.Log.Debug("GET /users/{id}/roles - PASSED")

	//
	//	PATCH /roles/{id}
	//

	// Fetching a role
	role := allRoles[3]
	// Modifying it
	role.RoleName = "Modified role"
	// JSON-ing the object
	if jsonObject, err = json.Marshal(role); err != nil {
		t.Error(err.Error())
	}

	// Creating the request
	if request, err = http.NewRequest(http.MethodPatch, "/roles/"+strconv.FormatInt(role.RoleId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new role so we can see if he got changed : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/roles/"+strconv.FormatInt(role.RoleId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbRole model.Role
	if err = json.NewDecoder(rr.Body).Decode(&dbRole); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(role, dbRole) {
		t.Error("Roles are not the same")
	}
	globals.Log.Debug("PATCH /roles/{id} - PASSED")

	//
	//	DELETE /roles/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodDelete, "/roles/"+strconv.FormatInt(role.RoleId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new role so we can see if it got deleted : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/roles/"+strconv.FormatInt(role.RoleId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	if err = json.NewDecoder(rr.Body).Decode(&dbRole); err == nil {
		t.Error(err)
	}

	globals.Log.Debug("DELETE /roles/{id} - PASSED")
}
