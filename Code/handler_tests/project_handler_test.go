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
	TESTED : GET /projects
	TESTED : GET /projects/{id}
	TESTED : GET /companies/{id}/projects
	TESTED : GET /users/{id}/projects
	TESTED : POST /projects
	TESTED : PATCH /projects/{id}
	TESTED : DELETE /projects/{id}
*/
func TestProjectHandler(t *testing.T) {
	var (
		err     error
		request *http.Request

		jsonObject []byte
	)

	// Creating a temporary struct to decode the json.
	var tmp struct {
		ProjectId int `json:"project_id"`
	}

	// Preparing for HTTP requests
	rr := httptest.NewRecorder()

	//
	//	GET /projects
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/projects", nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing the request
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbAllProjects model.Projects
	if err = json.NewDecoder(rr.Body).Decode(&dbAllProjects); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(fakeProjects, dbAllProjects) {
		t.Error("Projects are not the same")
	}

	globals.Log.Debug("GET /projects - PASSED")

	//
	//	POST /projects
	//

	project1 := model.Project{
		ProjectName: "New test project",
	}
	// Turning the object into JSON
	if jsonObject, err = json.Marshal(project1); err != nil {
		t.Error(err.Error())
	}

	if request, err = http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)

	// Sending the request
	r.ServeHTTP(rr, request)

	// Checking the result
	if err = json.NewDecoder(rr.Body).Decode(&tmp); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("POST /projects - PASSED")

	project1.ProjectId = int64(tmp.ProjectId)

	//
	//	GET /projects/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/projects/"+strconv.FormatInt(project1.ProjectId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbProject1 model.Project
	if err = json.NewDecoder(rr.Body).Decode(&dbProject1); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(project1, dbProject1) {
		t.Error("Projects are not the same")
	}
	globals.Log.Debug("GET /projects/{id} - PASSED")

	//
	//	GET /companies/{id}/projects
	//

	// Fetching the projects of company 1
	projectsOfcompany1, _ := env.DB.GetProjectsOfCompany(1)

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/companies/1/projects", nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbProjectsOfCompany1 model.Projects
	if err = json.NewDecoder(rr.Body).Decode(&dbProjectsOfCompany1); err != nil {
		t.Error(err)
	}

	// Verifying the result
	if !cmp.Equal(projectsOfcompany1, dbProjectsOfCompany1) {
		t.Error("Projects are not the same")
	}
	globals.Log.Debug("GET /companies/{id}/projects - PASSED")

	//
	//	GET /users/{id}/projects
	//

	// Fetching the projects of user 1
	projectsOfUser1, _ := env.DB.GetProjectsOfUser(1)

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/users/1/projects", nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbProjectsOfUser1 model.Projects
	if err = json.NewDecoder(rr.Body).Decode(&dbProjectsOfUser1); err != nil {
		t.Error(err)
	}

	// Verifying the result
	if !cmp.Equal(projectsOfUser1, dbProjectsOfUser1) {
		t.Error("Projects are not the same")
	}
	globals.Log.Debug("GET /users/{id}/projects - PASSED")

	//
	//	PATCH /projects/{id}
	//

	// Fetching a project
	project := fakeProjects[2]
	// Modifying it
	project.ProjectName = "Modified project"
	// JSON-ing the object
	if jsonObject, err = json.Marshal(project); err != nil {
		t.Error(err.Error())
	}

	// Creating the request
	if request, err = http.NewRequest(http.MethodPatch, "/projects/"+strconv.FormatInt(project.ProjectId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new project so we can see if he got changed : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/projects/"+strconv.FormatInt(project.ProjectId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbProject model.Project
	if err = json.NewDecoder(rr.Body).Decode(&dbProject); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(project, dbProject) {
		t.Error("Projects are not the same")
	}
	globals.Log.Debug("PATCH /projects/{id} - PASSED")

	//
	//	DELETE /projects/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodDelete, "/projects/"+strconv.FormatInt(project.ProjectId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new projects so we can see if it got deleted : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/projects/"+strconv.FormatInt(project.ProjectId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	if err = json.NewDecoder(rr.Body).Decode(&dbProject); err == nil {
		t.Error(err)
	}

	globals.Log.Debug("DELETE /projects/{id} - PASSED")
}
