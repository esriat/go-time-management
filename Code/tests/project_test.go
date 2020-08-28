package tests

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/datastores"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

/*
	TESTED : GetProjects() (model.Projects, error)
	TESTED : GetProject(ProjectId int64) (model.Project, error)
	TESTED : GetProjectsOfCompany(CompanyId int64) (model.Projects, error)
	TESTED : GetProjectsOfUser(UserId int64) (model.Projects, error)
	TESTED : CreateProject(Project model.Project) (int64, error)
	TESTED : DeleteProject(ProjectId int64) error
	TESTED : UpdateProject(Project model.Project) (model.Project, error)
*/
func TestProject(t *testing.T) {
	// Initializing variables
	testDatastore, err := datastores.NewDatabase("myTestDatabase.db")
	if err != nil {
		t.Error(err)
	}

	project1 := model.Project{
		ProjectId:   0,
		ProjectName: "Project 1",
	}

	project2 := model.Project{
		ProjectId:   0,
		ProjectName: "Project 2",
	}

	project3 := model.Project{
		ProjectId:   0,
		ProjectName: "Project 3",
	}

	//
	//	CreateProject
	//

	project1.ProjectId, err = testDatastore.CreateProject(project1)
	if err != nil {
		t.Error(err)
	}

	project2.ProjectId, err = testDatastore.CreateProject(project2)
	if err != nil {
		t.Error(err)
	}

	project3.ProjectId, err = testDatastore.CreateProject(project3)
	if err != nil {
		t.Error(err)
	}

	globals.Log.Debug("CreateProject test - PASSED")

	//
	//	GetProjects
	//
	var (
		vacationProject model.Project
		allProjects     model.Projects
	)

	// Retrieving the Vacation project
	if vacationProject, err = testDatastore.GetVacationProject(); err != nil {
		t.Error(err)
	}

	// Formatting the data
	projectList := model.Projects{}
	projectList = append(projectList, vacationProject)
	projectList = append(projectList, project1)
	projectList = append(projectList, project2)
	projectList = append(projectList, project3)

	// Getting all projects
	if allProjects, err = testDatastore.GetProjects(); err != nil {
		t.Error(err)
	}

	// Comparing
	if !cmp.Equal(allProjects, projectList) {
		t.Error(err)
	}

	globals.Log.Debug("GetProjects test - PASSED")

	//
	// Test GetProject(ProjectId)
	//

	// Fetching a project
	var project model.Project
	if project, err = testDatastore.GetProject(project1.ProjectId); err != nil {
		t.Error(err)
	}

	// Comparing it
	if !cmp.Equal(project1, project) {
		t.Error(err)
	}

	globals.Log.Debug("GetProject test - PASSED")

	//
	// Test GetProjectsOfCompany
	//

	// Formatting the data
	projectsOfCompany1 := model.Projects{}
	projectsOfCompany1 = append(projectsOfCompany1, project1)
	projectsOfCompany1 = append(projectsOfCompany1, project2)

	projectsOfCompany2 := model.Projects{}
	projectsOfCompany2 = append(projectsOfCompany2, project2)
	projectsOfCompany2 = append(projectsOfCompany2, project3)

	// Create some data stuff
	company1 := model.Company{
		CompanyId:   0,
		CompanyName: "Company 1",
	}

	company2 := model.Company{
		CompanyId:   0,
		CompanyName: "Company 2",
	}

	if company1.CompanyId, err = testDatastore.CreateCompany(company1); err != nil {
		t.Error(err)
	}

	if company2.CompanyId, err = testDatastore.CreateCompany(company2); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateCompanyProject(model.CompanyProject{
		CompanyId: company1.CompanyId,
		ProjectId: project1.ProjectId,
	}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateCompanyProject(model.CompanyProject{
		CompanyId: company1.CompanyId,
		ProjectId: project2.ProjectId,
	}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateCompanyProject(model.CompanyProject{
		CompanyId: company2.CompanyId,
		ProjectId: project2.ProjectId,
	}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateCompanyProject(model.CompanyProject{
		CompanyId: company2.CompanyId,
		ProjectId: project3.ProjectId,
	}); err != nil {
		t.Error(err)
	}

	// Now testing
	var (
		test1 model.Projects
		test2 model.Projects
	)

	// Getting the projects of the companies
	if test1, err = testDatastore.GetProjectsOfCompany(company1.CompanyId); err != nil {
		t.Error(err)
	}

	if test2, err = testDatastore.GetProjectsOfCompany(company2.CompanyId); err != nil {
		t.Error(err)
	}

	// Verifying the results
	if !cmp.Equal(test1, projectsOfCompany1) {
		t.Error(err)
	}

	if !cmp.Equal(test2, projectsOfCompany2) {
		t.Error(err)
	}

	globals.Log.Debug("GetProjectsOfCompany test - PASSED")

	//
	// Test GetProjectsOfUser
	//

	// Creating some data..
	contract := model.Contract{
		ContractId:   0,
		ContractName: "Test contract",
	}

	if contract.ContractId, err = testDatastore.CreateContract(contract); err != nil {
		t.Error(err)
	}

	role := model.Role{
		RoleId:   0,
		RoleName: "Test role"}

	if role.RoleId, err = testDatastore.CreateRole(role); err != nil {
		t.Error(err)
	}

	user1 := model.User{
		UserId:               0,
		ContractId:           contract.ContractId,
		RoleId:               role.RoleId,
		Username:             "First User",
		Password:             "Password",
		LastName:             "User",
		FirstName:            "First",
		Mail:                 "FirstUser@user.com",
		TheoricalHoursWorked: 50,
		VacationHours:        50,
	}

	user2 := model.User{
		UserId:               0,
		ContractId:           contract.ContractId,
		RoleId:               role.RoleId,
		Username:             "Second User",
		Password:             "Password",
		LastName:             "User",
		FirstName:            "Seconds",
		Mail:                 "SecondUser@user.com",
		TheoricalHoursWorked: 50,
		VacationHours:        50,
	}

	if user1.UserId, err = testDatastore.CreateUser(user1); err != nil {
		t.Error(err)
	}

	if user2.UserId, err = testDatastore.CreateUser(user2); err != nil {
		t.Error(err)
	}

	projectsOfUser1 := model.Projects{}
	projectsOfUser1 = append(projectsOfUser1, project1)
	projectsOfUser1 = append(projectsOfUser1, project2)

	projectsOfUser2 := model.Projects{}
	projectsOfUser2 = append(projectsOfUser2, project2)
	projectsOfUser2 = append(projectsOfUser2, project3)

	sp1 := model.Schedule{
		ScheduleId: 0,
		ProjectId:  project1.ProjectId,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	sp2 := model.Schedule{
		ScheduleId: 0,
		ProjectId:  project2.ProjectId,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	sp3 := model.Schedule{
		ScheduleId: 0,
		ProjectId:  project3.ProjectId,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	if sp1.ScheduleId, err = testDatastore.CreateSchedule(sp1); err != nil {
		log.Fatal(err)
	}

	if sp2.ScheduleId, err = testDatastore.CreateSchedule(sp2); err != nil {
		log.Fatal(err)
	}

	if sp3.ScheduleId, err = testDatastore.CreateSchedule(sp3); err != nil {
		log.Fatal(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user1.UserId,
		ScheduleId: sp1.ScheduleId}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user1.UserId,
		ScheduleId: sp2.ScheduleId}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user2.UserId,
		ScheduleId: sp2.ScheduleId}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user2.UserId,
		ScheduleId: sp3.ScheduleId}); err != nil {
		t.Error(err)
	}

	// Now testing
	dbProjectsOfUser1 := model.Projects{}
	dbProjectsOfUser2 := model.Projects{}

	// Getting the projects of users
	if dbProjectsOfUser1, err = testDatastore.GetProjectsOfUser(user1.UserId); err != nil {
		t.Error(err)
	}

	if dbProjectsOfUser2, err = testDatastore.GetProjectsOfUser(user2.UserId); err != nil {
		t.Error(err)
	}

	// Verifying the results
	if !cmp.Equal(dbProjectsOfUser1, projectsOfUser1) {
		t.Error(err)
	}

	if !cmp.Equal(dbProjectsOfUser2, dbProjectsOfUser2) {
		t.Error(err)
	}

	globals.Log.Debug("GetProjectsOfUser test - PASSED")

	//
	//	UpdateProject(Project)
	//

	var updatedProject model.Project
	// Modifying the project
	project1.ProjectName = "New project name"
	if _, err = testDatastore.UpdateProject(project1); err != nil {
		t.Error(err)
	}

	// Saving the changes
	if updatedProject, err = testDatastore.GetProject(project1.ProjectId); err != nil {
		t.Error(err)
	}

	// Comparing
	if !cmp.Equal(project1, updatedProject) {
		t.Error(err)
	}

	globals.Log.Debug("UpdateProject test - PASSED")

	//
	//	DeleteProject(ProjectId)

	// Creating a new project so we can delete it
	var ind int64
	if ind, err = testDatastore.CreateProject(model.Project{
		ProjectId:   0,
		ProjectName: "Project",
	}); err != nil {
		t.Error(err)
	}

	// Deleting it
	if err = testDatastore.DeleteProject(ind); err != nil {
		t.Error(err)
	}

	// Trying to get the deleted project. No error -> test fails
	if _, err = testDatastore.GetProject(ind); err == nil {
		t.Error()
	}

	globals.Log.Debug("DeleteProject test - PASSED")

	testDatastore.CloseDatabase()
}
