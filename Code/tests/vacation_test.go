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
	TESTED : GetVacationsOfUser(UserId int64) (model.Schedules, error)
	TESTED : GetVacation(VacationId int64) (model.Schedule, error)
	TESTED : CreateVacation(Schedule model.Schedule) (int64, error)
	TESTED : DeleteVacation(VacationId int64) error
	TESTED : UpdateVacation(Vacation model.Schedule) (model.Schedule, error)
*/
func TestVacation(t *testing.T) {
	// Initializing variables
	var (
		testDatastore *datastores.ConcreteDatastore
		err           error
	)

	if testDatastore, err = datastores.NewDatabase("myTestDatabase.db"); err != nil {
		t.Error(err)
	}

	// Setting up some data
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
		Username:             "Username",
		Password:             "This is a password",
		LastName:             "LastName",
		FirstName:            "FirstName",
		Mail:                 "user@user.com",
		TheoricalHoursWorked: 50,
		VacationHours:        50,
	}

	if user1.UserId, err = testDatastore.CreateUser(user1); err != nil {
		t.Error(err)
	}

	//
	// Test CreateVacation()
	//
	schedule := model.Schedule{
		ScheduleId: 0,
		ProjectId:  0,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	// Creating a new vacation
	if schedule.ScheduleId, err = testDatastore.CreateVacation(schedule); err != nil {
		t.Error(err)
	}

	// And creating a link between a user and a schedule
	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user1.UserId,
		ScheduleId: schedule.ScheduleId}); err != nil {
		t.Error(err)
	}

	var vacationProject model.Project
	if vacationProject, err = testDatastore.GetVacationProject(); err != nil {
		t.Error(err)
	}
	schedule.ProjectId = vacationProject.ProjectId

	//
	// Test GetVacation(VacationId)
	//
	var vacation model.Schedule

	// Fetching a vacation
	if vacation, err = testDatastore.GetVacation(schedule.ScheduleId); err != nil {
		t.Error(err)
	}

	// Verifying the data
	if !cmp.Equal(vacation, schedule) {
		t.Error(err)
	}

	globals.Log.Debug("GetVacation test - PASSED")

	// Setting up some data
	user2 := model.User{
		UserId:               0,
		ContractId:           contract.ContractId,
		RoleId:               role.RoleId,
		Username:             "Username",
		Password:             "This is a password",
		LastName:             "LastName",
		FirstName:            "FirstName",
		Mail:                 "user2@user.com",
		TheoricalHoursWorked: 50,
		VacationHours:        50,
	}

	user1Vacation2 := model.Schedule{
		ScheduleId: 0,
		ProjectId:  vacationProject.ProjectId,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	user2Vacation1 := model.Schedule{
		ScheduleId: 0,
		ProjectId:  vacationProject.ProjectId,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	user2Vacation2 := model.Schedule{
		ScheduleId: 0,
		ProjectId:  vacationProject.ProjectId,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	if user2.UserId, err = testDatastore.CreateUser(user2); err != nil {
		t.Error(err)
	}

	if user1Vacation2.ScheduleId, err = testDatastore.CreateVacation(user1Vacation2); err != nil {
		t.Error(err)
	}

	if user2Vacation1.ScheduleId, err = testDatastore.CreateVacation(user2Vacation1); err != nil {
		t.Error(err)
	}

	if user2Vacation2.ScheduleId, err = testDatastore.CreateVacation(user2Vacation2); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user1.UserId,
		ScheduleId: user1Vacation2.ScheduleId}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user2.UserId,
		ScheduleId: user2Vacation1.ScheduleId}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user2.UserId,
		ScheduleId: user2Vacation2.ScheduleId}); err != nil {
		t.Error(err)
	}

	vacationsOfUser1 := model.Schedules{}
	vacationsOfUser1 = append(vacationsOfUser1, schedule)
	vacationsOfUser1 = append(vacationsOfUser1, user1Vacation2)

	vacationsOfUser2 := model.Schedules{}
	vacationsOfUser2 = append(vacationsOfUser2, user2Vacation1)
	vacationsOfUser2 = append(vacationsOfUser2, user2Vacation2)

	//
	// Test GetVacationOfUser(UserId)
	//

	var (
		databaseVacationsOfUser1 model.Schedules
		databaseVacationsOfUser2 model.Schedules
	)

	// Getting vacation of users data
	if databaseVacationsOfUser1, err = testDatastore.GetVacationsOfUser(user1.UserId); err != nil {
		t.Error(err)
	}

	if databaseVacationsOfUser2, err = testDatastore.GetVacationsOfUser(user2.UserId); err != nil {
		t.Error(err)
	}

	// Comparing the data
	if !cmp.Equal(databaseVacationsOfUser1, vacationsOfUser1) {
		t.Error(err)
	}

	if !cmp.Equal(databaseVacationsOfUser2, vacationsOfUser2) {
		t.Error(err)
	}

	globals.Log.Debug("GetVacationsOfUser test - PASSED")

	//
	// Test UpdateVacation(Vacation)
	//

	// Updating the data
	format := "2006-01-02T15:04:05.000Z"
	testDate := "2015-06-15T10:19:30.000Z"
	time, err := time.Parse(format, testDate)
	schedule.StartDate = sql.NullTime{Valid: true, Time: time}

	// Saving the changes
	if _, err = testDatastore.UpdateVacation(schedule); err != nil {
		t.Error(err)
	}

	// Retrieving the new data
	var updatedVacation model.Schedule
	if updatedVacation, err = testDatastore.GetVacation(schedule.ScheduleId); err != nil {
		t.Error(err)
	}

	// Checking if it changed
	if !cmp.Equal(schedule, updatedVacation) {
		t.Error(err)
	}

	globals.Log.Debug("UpdateVacation test - PASSED")

	//
	// Test DeleteVacation(VacationId)
	//

	// Creating a vacation so we can delete it
	var ind int64
	if ind, err = testDatastore.CreateVacation(schedule); err != nil {
		t.Error(err)
	}

	// Deleting the vacation
	if err = testDatastore.DeleteVacation(ind); err != nil {
		t.Error(err)
	}

	// Trying to get the deleted vacation
	if _, err = testDatastore.GetVacation(ind); err == nil {
		t.Error()
	}

	globals.Log.Debug("DeleteVacation test - PASSED")

	testDatastore.CloseDatabase()
}
