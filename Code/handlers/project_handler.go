package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//	GetProjectsHandler
/*	The handler called by the following endpoint : GET /projects
	This method is used to get the list of all existing projects.
*/
func (env *Env) GetProjectsHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err      error
		projects model.Projects
	)

	globals.Log.Debug("Calling GetProjectsHandler")

	if projects, err = env.DB.GetProjects(); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error retrieving all the projects",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(projects)
	return nil
}

//	GetProjectHandler
/*	The handler called by the following endpoint : GET /projects/{id}
	This method is used to get a project.
*/
func (env *Env) GetProjectHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err     error
		project model.Project
		id      int
	)

	globals.Log.Debug("Calling GetProjectHandler")

	vars := mux.Vars(r)

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if project, err = env.DB.GetProject(int64(id)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching project",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(project)
	return nil
}

//	GetProjectsOfCompanyHandler
/*	The handler called by the following endpoint : GET /companies/{company_id}/projects
	This method is used the list of projects of a company
*/
func (env *Env) GetProjectsOfCompanyHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		projects  model.Projects
		companyId int
	)

	globals.Log.Debug("Calling GetProjectsOfCompanyHandler")

	vars := mux.Vars(r)

	if companyId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if projects, err = env.DB.GetProjectsOfCompany(int64(companyId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching projects",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(projects)
	return nil
}

//	GetProjectsOfCompanyHandler
/*	The handler called by the following endpoint : GET /companies/{company_id}/projects
	This method is used the list of projects of a user
*/
func (env *Env) GetProjectsOfUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err      error
		projects model.Projects
		userId   int
	)

	globals.Log.Debug("Calling GetProjectsOfUserHandler")

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if projects, err = env.DB.GetProjectsOfUser(int64(userId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching projects",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(projects)
	return nil
}

//	CreateProjectHandler
/*	The handler called by the following endpoint : POST /projects
	This method is used to create a new project.
*/
func (env *Env) CreateProjectHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		project   model.Project
		projectId int64
	)

	globals.Log.Debug("CreateProjectHandler called")

	if err = json.NewDecoder(r.Body).Decode(&project); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	globals.Log.Debug("Calling CreateProject method")

	if projectId, err = env.DB.CreateProject(project); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error creating the project",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Project created")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(struct {
		ProjectId int64 `json:"project_id"`
	}{
		ProjectId: projectId,
	}); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the project id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	UpdateProjectHandler
/*	The handler called by the following endpoint : PATCH /projects
	This method is used to update an existing project.
	Trying to update the Vacation project will result in an error.
*/
func (env *Env) UpdateProjectHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err             error
		project         model.Project
		vacationProject model.Project
		projectId       int
	)

	globals.Log.Debug("CreateProjectHandler called")

	if vacationProject, err = env.DB.GetVacationProject(); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when checking projects",
			Code:    http.StatusInternalServerError,
		}
	}

	if err = json.NewDecoder(r.Body).Decode(&project); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	vars := mux.Vars(r)

	if projectId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	project.ProjectId = int64(projectId)

	if project.ProjectId == vacationProject.ProjectId || vacationProject.ProjectName == project.ProjectName {
		return &AppError{
			Message: "Updating the vacation error is unauthorized",
			Code:    http.StatusUnauthorized,
		}
	}

	globals.Log.Debug("Calling CreateProject method")

	if project, err = env.DB.UpdateProject(project); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when updating the project",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Project updated")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(project); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the project id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	DeleteProjectHandler
/*	The handler called by the following endpoint : DELETE /projects/{id}
	This method is used to delete an existing project.
	Trying to delete the Vacation project will result in an error.
*/
func (env *Env) DeleteProjectHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		projectId int
	)

	globals.Log.Debug("DeleteProjectHandler called")

	vars := mux.Vars(r)

	if projectId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Calling CreateProject method")

	if err = env.DB.DeleteProject(int64(projectId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when deleting the project",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Project deleted")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}
