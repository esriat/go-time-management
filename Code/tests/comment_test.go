package tests

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/datastores"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

/*
	TESTED : GetComments() (model.Comments, error)
	TESTED : GetComment(CommentId int64) (model.Comment, error)
	TESTED : GetCommentsOfUser(UserId int64) (model.Comments, error)
	TESTED : GetCommentsOfSchedule(ScheduleId int64) (model.Comments, error)
	TESTED : GetCommentsOfProject(ProjectId int64) (model.Comments, error)
	TESTED : CreateComment(Comment model.Comment) (int64, error)
	TESTED : DeleteComment(CommentId int64) error
	TESTED : UpdateComment(Comment model.Comment) (model.Comment, error)
*/
func TestComment(t *testing.T) {
	// Initializing variables
	var (
		err           error
		testDatastore *datastores.ConcreteDatastore
		allComments   model.Comments
	)

	if testDatastore, err = datastores.NewDatabase("myTestDatabase.db"); err != nil {
		t.Error(err)
	}

	// Setting up a bunch of data (as there are Foreign keys in the database)
	project1 := model.Project{
		ProjectId:   0,
		ProjectName: "Test project"}

	if project1.ProjectId, err = testDatastore.CreateProject(project1); err != nil {
		t.Error(err)
	}

	project2 := model.Project{
		ProjectId:   0,
		ProjectName: "Second test project"}

	if project2.ProjectId, err = testDatastore.CreateProject(project2); err != nil {
		t.Error(err)
	}

	schedule1 := model.Schedule{
		ScheduleId: 0,
		ProjectId:  project1.ProjectId,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	if schedule1.ScheduleId, err = testDatastore.CreateSchedule(schedule1); err != nil {
		t.Error(err)
	}

	schedule2 := model.Schedule{
		ScheduleId: 0,
		ProjectId:  project1.ProjectId,
		StartDate:  sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:    sql.NullTime{Valid: true, Time: time.Now()},
	}

	if schedule2.ScheduleId, err = testDatastore.CreateSchedule(schedule2); err != nil {
		t.Error(err)
	}

	comment1 := model.Comment{
		CommentId:   0,
		ScheduleId:  schedule1.ScheduleId,
		Comment:     "Test comment 1 to schedule 1",
		IsImportant: false,
	}

	comment2 := model.Comment{
		CommentId:   0,
		ScheduleId:  schedule1.ScheduleId,
		Comment:     "Test comment 2 to schedule 1",
		IsImportant: true,
	}

	comment3 := model.Comment{
		CommentId:   0,
		ScheduleId:  schedule2.ScheduleId,
		Comment:     "Test comment 3 to schedule 2",
		IsImportant: false,
	}

	comment4 := model.Comment{
		CommentId:   0,
		ScheduleId:  schedule2.ScheduleId,
		Comment:     "Test comment 4 to schedule 2",
		IsImportant: true,
	}

	//
	//	CreateComment(Comment)
	//

	// Trying to create all the comments
	if comment1.CommentId, err = testDatastore.CreateComment(comment1); err != nil {
		t.Error(err)
	}

	if comment2.CommentId, err = testDatastore.CreateComment(comment2); err != nil {
		t.Error(err)
	}

	if comment3.CommentId, err = testDatastore.CreateComment(comment3); err != nil {
		t.Error(err)
	}

	if comment4.CommentId, err = testDatastore.CreateComment(comment4); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("CreateComment test - PASSED")

	//
	//	GetComments()
	//

	// Formatting some data, for now or later
	commentList := model.Comments{}
	commentList = append(commentList, comment1)
	commentList = append(commentList, comment2)
	commentList = append(commentList, comment3)
	commentList = append(commentList, comment4)

	commentSchedule1List := model.Comments{}
	commentSchedule1List = append(commentSchedule1List, comment1)
	commentSchedule1List = append(commentSchedule1List, comment2)

	commentSchedule2List := model.Comments{}
	commentSchedule2List = append(commentSchedule2List, comment3)
	commentSchedule2List = append(commentSchedule2List, comment4)

	// Fetching all the comments
	if allComments, err = testDatastore.GetComments(); err != nil {
		t.Error(err)
	}

	// Comparing the data
	if !cmp.Equal(allComments, commentList) {
		t.Error(err)
	}

	globals.Log.Debug("GetComments test - PASSED")

	//
	//	GetCommentsOfUser
	//

	// Creating a contract and a role (foreign key stuff)
	contract := model.Contract{
		ContractId:   0,
		ContractName: "Test contract",
	}

	if contract.ContractId, err = testDatastore.CreateContract(contract); err != nil {
		t.Error(err)
	}

	role := model.Role{
		RoleId:   0,
		RoleName: "Test role",
	}

	if role.RoleId, err = testDatastore.CreateRole(role); err != nil {
		t.Error(err)
	}

	// Creating users
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
		VacationHours:        50}

	user2 := model.User{
		UserId:               0,
		ContractId:           contract.ContractId,
		RoleId:               role.RoleId,
		Username:             "Second User",
		Password:             "Password",
		LastName:             "User",
		FirstName:            "Second",
		Mail:                 "SecondUser@user.com",
		TheoricalHoursWorked: 50,
		VacationHours:        50}

	if user1.UserId, err = testDatastore.CreateUser(user1); err != nil {
		fmt.Print(err)
		t.Error(err)
	}

	if user2.UserId, err = testDatastore.CreateUser(user2); err != nil {
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
		UserId:     user2.UserId,
		ScheduleId: schedule2.ScheduleId,
	}); err != nil {
		t.Error(err)
	}

	// Formatting data
	commentsOfUser1 := model.Comments{}
	commentsOfUser1 = append(commentsOfUser1, comment1)
	commentsOfUser1 = append(commentsOfUser1, comment2)

	dbCommentsOfUser1 := model.Comments{}
	dbCommentsOfUser2 := model.Comments{}

	// Getting comments of the users
	if dbCommentsOfUser1, err = testDatastore.GetCommentsOfUser(user1.UserId); err != nil {
		t.Error(err)
	}

	if dbCommentsOfUser2, err = testDatastore.GetCommentsOfUser(user2.UserId); err != nil {
		t.Error(err)
	}

	// Comparing
	if !cmp.Equal(commentsOfUser1, dbCommentsOfUser1) {
		t.Error(err)
	}

	if !cmp.Equal(allComments, dbCommentsOfUser2) {
		t.Error(err)
	}

	globals.Log.Debug("GetCommentsOfUser test - PASSED")

	// Test GetComment(CommentId)
	databaseComment := model.Comment{}

	if databaseComment, err = testDatastore.GetComment(comment1.CommentId); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(databaseComment, comment1) {
		t.Error(err)
	}

	//
	//	GetCommentsOfShedule()
	//
	var (
		commentsOfSchedule1 model.Comments
		commentsOfSchedule2 model.Comments
	)

	// Getting the data
	if commentsOfSchedule1, err = testDatastore.GetCommentsOfSchedule(schedule1.ScheduleId); err != nil {
		t.Error(err)
	}

	if commentsOfSchedule2, err = testDatastore.GetCommentsOfSchedule(schedule2.ScheduleId); err != nil {
		t.Error(err)
	}

	// Comparing the data
	if !cmp.Equal(commentsOfSchedule1, commentSchedule1List) {
		t.Error(err)
	}

	if !cmp.Equal(commentsOfSchedule2, commentSchedule2List) {
		t.Error(err)
	}

	globals.Log.Debug("GetCommentsOfSchedule test - PASSED")

	//
	//	GetCommentsOfProject
	//

	var (
		commentsOfProject1     model.Comments
		commentsOfProject2     model.Comments
		realCommentsOfProject2 model.Comments
	)

	// Getting the data
	if commentsOfProject1, err = testDatastore.GetCommentsOfProject(project1.ProjectId); err != nil {
		t.Error(err)
	}

	if commentsOfProject2, err = testDatastore.GetCommentsOfProject(project2.ProjectId); err != nil {
		t.Error(err)
	}

	// Comparing it
	if !cmp.Equal(commentsOfProject1, allComments) {
		t.Error(err)
	}

	// I had to comment these two lines because they fail anyways.
	// We want to get 2 empty arrays, and we have two empty arrays
	// But it looks like cmp.Equal does not work on empty arrays...
	// Also, in Go, arrays can only be compared to nil.
	// Guess I have to work with the lenght of the objects then..
	// That's ugly, but I don't have the choice

	/*if !cmp.Equal(commentsOfProject2, realCommentsOfProject2) {
		t.Error(err)
	}*/

	if len(commentsOfProject2) != 0 || len(realCommentsOfProject2) != 0 {
		t.Error()
	}

	globals.Log.Debug("GetCommentsOfProject test - PASSED")

	//
	//	UpdateComment(Comment)
	//
	var updatedComment model.Comment

	// Update the comment
	comment1.Comment = "New comment"
	if updatedComment, err = testDatastore.UpdateComment(comment1); err != nil {
		t.Error(err)
	}

	// Get it from the database to see if the update worked
	if updatedComment, err = testDatastore.GetComment(comment1.CommentId); err != nil {
		t.Error(err)
	}

	// Compare
	if !cmp.Equal(comment1, updatedComment) {
		t.Error(err)
	}

	globals.Log.Debug("UpdateComment test - PASSED")

	//
	//	DeleteComment(CommentId)
	//

	// Deleting the comment
	if err = testDatastore.DeleteComment(comment1.CommentId); err != nil {
		t.Error(err)
	}

	// Making sure it got properly deleted
	if _, err = testDatastore.GetComment(comment1.CommentId); err == nil {
		t.Error()
	}

	globals.Log.Debug("DeleteComment test - PASSED")

	testDatastore.CloseDatabase()
}
