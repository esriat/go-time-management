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
	TESTED : GetSchedule(ScheduleId int64) (model.Schedule, error)
	TESTED : GetSchedulesOfUser(UserId int64) (model.Schedules, error)
	TESTED : GetSchedulesOfProject(ProjectId int64) (model.Schedules, error)
	TESTED : CreateSchedule(Schedule model.Schedule) (int64, error)
	TESTED : DeleteSchedule(ScheduleId int64) error
	TESTED : UpdateSchedule(Schedule model.Schedule) (model.Schedule, error)
*/
func TestSchedule(t *testing.T) {
	// Initializing variables
	var err error
	var projectId int64

	testDatastore, err := datastores.NewDatabase("myTestDatabase.db")
	if err != nil {
		t.Error(err)
	}

	// Creating some data
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

	projectId, err = testDatastore.CreateProject(model.Project{
		ProjectId:   0,
		ProjectName: "Test project"})
	if err != nil {
		t.Error(err)
	}

	schedule := model.Schedule{
		ScheduleId: 0,
		ProjectId:  projectId,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	//
	// Test CreateSchedule
	//

	// Creating the schedule
	if schedule.ScheduleId, err = testDatastore.CreateSchedule(schedule); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("CreateSchedule test - PASSED")

	//
	// Test GetSchedule
	//

	// Trying to retrieve the schedule we just created
	var databaseSchedule model.Schedule
	if databaseSchedule, err = testDatastore.GetSchedule(schedule.ScheduleId); err != nil {
		t.Error(err)
	}

	// Comparing the data
	if !cmp.Equal(schedule, databaseSchedule) {
		t.Error(err)
	}

	globals.Log.Debug("GetSchedule test - PASSED")

	// Doing data stuff
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

	user2 := model.User{
		UserId:               0,
		ContractId:           contract.ContractId,
		RoleId:               role.RoleId,
		Username:             "Username 2",
		Password:             "This is a password",
		LastName:             "LastName 2",
		FirstName:            "FirstName 2",
		Mail:                 "thisistheseconduser@user.com",
		TheoricalHoursWorked: 50,
		VacationHours:        50,
	}

	if user2.UserId, err = testDatastore.CreateUser(user2); err != nil {
		t.Error(err)
	}

	schedule2 := model.Schedule{
		ScheduleId: 0,
		ProjectId:  projectId,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	if schedule2.ScheduleId, err = testDatastore.CreateSchedule(schedule2); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user1.UserId,
		ScheduleId: schedule.ScheduleId}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user2.UserId,
		ScheduleId: schedule2.ScheduleId}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserSchedule(model.UserSchedule{
		UserId:     user1.UserId,
		ScheduleId: schedule2.ScheduleId}); err != nil {
		t.Error(err)
	}

	schedulesOfUser1 := model.Schedules{}
	schedulesOfUser1 = append(schedulesOfUser1, schedule)
	schedulesOfUser1 = append(schedulesOfUser1, schedule2)

	schedulesOfUser2 := model.Schedules{}
	schedulesOfUser2 = append(schedulesOfUser2, schedule2)

	shedulesOfProject := model.Schedules{}
	shedulesOfProject = append(shedulesOfProject, schedule)
	shedulesOfProject = append(shedulesOfProject, schedule2)

	//
	// Test GetScheduleOfProject
	//

	// Getting the schedules of a project
	var dbScheduleOfProject model.Schedules
	if dbScheduleOfProject, err = testDatastore.GetSchedulesOfProject(projectId); err != nil {
		t.Error(err)
	}

	// Comparing..
	if !cmp.Equal(shedulesOfProject, dbScheduleOfProject) {
		t.Error(err)
	}

	// Creating data
	projectId2, err := testDatastore.CreateProject(model.Project{
		ProjectId:   0,
		ProjectName: "Test project 2"})
	if err != nil {
		t.Error(err)
	}

	// Getting data
	var (
		nothing       model.Schedules
		shouldBeEmpty model.Schedules
	)
	if shouldBeEmpty, err = testDatastore.GetSchedulesOfProject(projectId2); err != nil {
		t.Error(err)
	}

	// Comparing

	// I had to comment these two lines because they fail anyways.
	// We want to get 2 empty arrays, and we have two empty arrays
	// But it looks like cmp.Equal does not work on empty arrays...
	// Also, in Go, arrays can only be compared to nil.
	// Guess I have to work with the length of the objects then..
	// That's ugly, but I don't have the choice

	/*if !cmp.Equal(nothing, shouldBeEmpty) {
		t.Error()
	}*/

	if len(nothing) != 0 || len(shouldBeEmpty) != 0 {
		t.Error()
	}

	globals.Log.Debug("GetSchedulesOfProject test - PASSED")

	//
	// Test GetSchedulesOfUser
	//
	var (
		dbScheduleOfUser1 model.Schedules
		dbScheduleOfUser2 model.Schedules
	)

	// Getting the data
	if dbScheduleOfUser1, err = testDatastore.GetSchedulesOfUser(user1.UserId); err != nil {
		t.Error(err)
	}

	if dbScheduleOfUser2, err = testDatastore.GetSchedulesOfUser(user2.UserId); err != nil {
		t.Error(err)
	}

	// Comparing the data
	if !cmp.Equal(dbScheduleOfUser1, schedulesOfUser1) {
		t.Error(err)
	}

	if !cmp.Equal(dbScheduleOfUser2, schedulesOfUser2) {
		t.Error(err)
	}

	globals.Log.Debug("GetSchedulesOfUser test - PASSED")

	//
	// Test UpdateSchedule
	//

	// Updating the schedule
	format := "2006-01-02T15:04:05.000Z"
	testDate := "2015-06-15T10:19:30.000Z"
	time, err := time.Parse(format, testDate)
	schedule.StartDate = sql.NullTime{Valid: true, Time: time}

	// Saving the changes
	if _, err = testDatastore.UpdateSchedule(schedule); err != nil {
		t.Error(err)
	}

	// Getting the new data
	var updatedSchedule model.Schedule
	if updatedSchedule, err = testDatastore.GetSchedule(schedule.ScheduleId); err != nil {
		t.Error(err)
	}

	// To see if it got saved properly
	if !cmp.Equal(schedule, updatedSchedule) {
		t.Error(err)
	}

	globals.Log.Debug("UpdateSchedule test - PASSED")

	//
	// Test DeleteSchedule
	//

	// Creating a schedule so we can delete it
	var ind int64
	if ind, err = testDatastore.CreateSchedule(schedule2); err != nil {
		t.Error(err)
	}

	// Deleting the schedule
	if err = testDatastore.DeleteSchedule(ind); err != nil {
		t.Error(err)
	}

	// Trying to get the schedule
	if _, err = testDatastore.GetSchedule(ind); err == nil {
		t.Error()
	}

	globals.Log.Debug("DeleteSchedule test - PASSED")

	testDatastore.CloseDatabase()
}
