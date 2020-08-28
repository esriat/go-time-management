package tests

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/datastores"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

/*
	TESTED : GetUsers() (model.Users, error)
	TESTED : GetUser(UserId int64) (model.User, error)
	TESTED : GetUserFromEmail(Email string) (model.User, error)
	TESTED : GetUsersOfCompany(CompanyId int64) (model.Users, error)
	TESTED : GetUsersOfProject(ProjectId int64) (model.Users, error)
	TESTED : GetUsersOfSchedule(ScheduleId int64) (model.Users, error)
	TESTED : CreateUser(User model.User) (int64, error)
	TESTED : DeleteUser(UserId int64) error
	TESTED : UpdateUser(User model.User) (model.User, error)
*/
func TestUser(t *testing.T) {
	var err error
	// Initializing variables and creating some data
	testDatastore, err := datastores.NewDatabase("myTestDatabase.db")
	if err != nil {
		t.Error(err)
	}

	contract := model.Contract{
		ContractId:   0,
		ContractName: "Test contract"}

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
		Password:             "This is a password",
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
		Password:             "This is a password",
		LastName:             "User",
		FirstName:            "Second",
		Mail:                 "SecondUser@user.com",
		TheoricalHoursWorked: 50,
		VacationHours:        50}

	user3 := model.User{
		UserId:               0,
		ContractId:           contract.ContractId,
		RoleId:               role.RoleId,
		Username:             "Third User",
		Password:             "This is a password",
		LastName:             "User",
		FirstName:            "Third",
		Mail:                 "ThirdUser@user.com",
		TheoricalHoursWorked: 50,
		VacationHours:        50}

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

	project1 := model.Project{
		ProjectId:   0,
		ProjectName: "Project 1",
	}

	project2 := model.Project{
		ProjectId:   0,
		ProjectName: "Project 2",
	}

	if project1.ProjectId, err = testDatastore.CreateProject(project1); err != nil {
		t.Error(err)
	}

	if project2.ProjectId, err = testDatastore.CreateProject(project2); err != nil {
		t.Error(err)
	}

	schedule1 := model.Schedule{
		ScheduleId: 0,
		ProjectId:  project1.ProjectId,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	schedule2 := model.Schedule{
		ScheduleId: 0,
		ProjectId:  project1.ProjectId,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	if schedule1.ScheduleId, err = testDatastore.CreateSchedule(schedule1); err != nil {
		t.Error(err)
	}

	if schedule2.ScheduleId, err = testDatastore.CreateSchedule(schedule2); err != nil {
		t.Error(err)
	}

	//
	// Test CreateUser()
	//
	if user1.UserId, err = testDatastore.CreateUser(user1); err != nil {
		t.Error(err)
	}

	if user2.UserId, err = testDatastore.CreateUser(user2); err != nil {
		t.Error(err)
	}

	if user3.UserId, err = testDatastore.CreateUser(user3); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("CreateUser test - PASSED")

	//
	// Test GetUser(UserId)
	//
	var user model.User

	// Fetching the user
	if user, err = testDatastore.GetUser(user1.UserId); err != nil {
		t.Error(err)
	}

	// Verifying the data
	if !cmp.Equal(user1, user) {
		t.Error(err)
	}

	globals.Log.Debug("GetUser test - PASSED")

	// Setting up the remaining data
	usersOfCompanyAndSchedule1 := model.Users{}
	usersOfCompanyAndSchedule1 = append(usersOfCompanyAndSchedule1, user1)
	usersOfCompanyAndSchedule1 = append(usersOfCompanyAndSchedule1, user2)

	usersOfCompanyAndSchedule2 := model.Users{}
	usersOfCompanyAndSchedule2 = append(usersOfCompanyAndSchedule2, user3)

	if err = testDatastore.CreateCompanyUser(model.CompanyUser{
		CompanyId: company1.CompanyId,
		UserId:    user1.UserId,
	}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateCompanyUser(model.CompanyUser{
		CompanyId: company1.CompanyId,
		UserId:    user2.UserId,
	}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateCompanyUser(model.CompanyUser{
		CompanyId: company2.CompanyId,
		UserId:    user3.UserId,
	}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user1.UserId,
		ScheduleId: schedule1.ScheduleId,
	}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user2.UserId,
		ScheduleId: schedule1.ScheduleId,
	}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user3.UserId,
		ScheduleId: schedule2.ScheduleId,
	}); err != nil {
		t.Error(err)
	}

	//
	// Test GetUsersOfCompany
	//
	var (
		databaseUsersOfCompany1 model.Users
		databaseUsersOfCompany2 model.Users
	)

	// Getting users of companies
	if databaseUsersOfCompany1, err = testDatastore.GetUsersOfCompany(company1.CompanyId); err != nil {
		t.Error(err)
	}

	if databaseUsersOfCompany2, err = testDatastore.GetUsersOfCompany(company2.CompanyId); err != nil {
		t.Error(err)
	}

	// checking the data
	if !cmp.Equal(usersOfCompanyAndSchedule1, databaseUsersOfCompany1) {
		t.Error(err)
	}

	if !cmp.Equal(usersOfCompanyAndSchedule2, databaseUsersOfCompany2) {
		t.Error(err)
	}

	globals.Log.Debug("GetUsersOfCompany test - PASSED")

	//
	// Test GetUsersOfSchedule
	//
	var (
		databaseUsersOfSchedule1 model.Users
		databaseUsersOfSchedule2 model.Users
	)

	// Getting users of schedules
	if databaseUsersOfSchedule1, err = testDatastore.GetUsersOfSchedule(schedule1.ScheduleId); err != nil {
		t.Error(err)
	}

	if databaseUsersOfSchedule2, err = testDatastore.GetUsersOfSchedule(schedule2.ScheduleId); err != nil {
		t.Error(err)
	}

	// verifying the data
	if !cmp.Equal(usersOfCompanyAndSchedule1, databaseUsersOfSchedule1) {
		t.Error(err)
	}

	if !cmp.Equal(usersOfCompanyAndSchedule2, databaseUsersOfSchedule2) {
		t.Error(err)
	}

	globals.Log.Debug("GetUsersOfSchedule test - PASSED")

	//
	// Test GetUsersOfProject
	//

	var (
		usersOfProject1 model.Users
		usersOfProject2 model.Users

		databaseUsersOfProject1 model.Users
		databaseUsersOfProject2 model.Users
	)
	usersOfProject1 = append(usersOfProject1, user1)
	usersOfProject1 = append(usersOfProject1, user2)
	usersOfProject1 = append(usersOfProject1, user3)

	// Getting users of projects
	if databaseUsersOfProject1, err = testDatastore.GetUsersOfProject(project1.ProjectId); err != nil {
		t.Error(err)
	}

	if databaseUsersOfProject2, err = testDatastore.GetUsersOfProject(project2.ProjectId); err != nil {
		t.Error(err)
	}

	// Verifying the data
	if !cmp.Equal(usersOfProject1, databaseUsersOfProject1) {
		t.Error(err)
	}

	// Same probleme here than in the schedule_test.go file, concerning empty arrays.

	/*if !cmp.Equal(usersOfProject2, databaseUsersOfProject2) {
		t.Error(err)
	}*/

	if len(usersOfProject2) != 0 || len(databaseUsersOfProject2) != 0 {
		t.Error()
	}

	globals.Log.Debug("GetUsersOfProject test - PASSED")

	//
	// Test GetUsers()
	//

	var (
		usersList   model.Users
		allUsers    model.Users
		defaultUser model.User
	)

	// Formatting data

	if defaultUser, err = testDatastore.GetUser(1); err != nil {
		t.Error(err)
	}

	usersList = append(usersList, defaultUser)
	usersList = append(usersList, user1)
	usersList = append(usersList, user2)
	usersList = append(usersList, user3)

	// Fetching all users
	if allUsers, err = testDatastore.GetUsers(); err != nil {
		t.Error(err)
	}

	// Comparing the data
	if !cmp.Equal(usersList, allUsers) {
		t.Error(err)
	}

	globals.Log.Debug("GetUsers test - PASSED")

	//
	// Test GetUserFromEmail
	//

	// Getting the data
	if user, err = testDatastore.GetUserFromEmail(user1.Mail); err != nil {
		t.Error(err)
	}

	// Comparing
	if !cmp.Equal(user1, user) {
		t.Error(err)
	}

	//
	// Test UpdateUser(User)
	//

	// Updating the user
	var updatedUser model.User
	user1.Username = "New user name"
	user1.FirstName = "New first name"
	user1.LastName = "New last name"

	// Saving changes
	if _, err = testDatastore.UpdateUser(user1); err != nil {
		t.Error(err)
	}

	// Checking if changes got saved
	if updatedUser, err = testDatastore.GetUser(user1.UserId); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(user1, updatedUser) {
		t.Error(err)
	}

	globals.Log.Debug("UpdateUser test - PASSED")

	//
	// Test DeleteProject(ProjectId)
	//

	// Creating a user so we can delete it
	var ind int64
	if ind, err = testDatastore.CreateUser(model.User{
		UserId:               0,
		ContractId:           contract.ContractId,
		RoleId:               role.RoleId,
		Username:             "Fourth User",
		Password:             "This is a password",
		LastName:             "User",
		FirstName:            "Fourt",
		Mail:                 "FourthUser@user.com",
		TheoricalHoursWorked: 50,
		VacationHours:        50,
	}); err != nil {
		t.Error(err)
	}

	// Deleting it
	if err = testDatastore.DeleteUser(ind); err != nil {
		t.Error(err)
	}

	// Getting it (should not work)
	if _, err = testDatastore.GetUser(ind); err == nil {
		t.Error()
	}

	globals.Log.Debug("DeleteUser test - PASSED")

	testDatastore.CloseDatabase()
}
