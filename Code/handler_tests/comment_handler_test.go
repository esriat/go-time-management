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
	TESTED : GET /comments
	TESTED : GET /comments/{id}
	TESTED : GET /users/{id}/comments
	TESTED : GET /projects/{id}/comments
	TESTED : GET /schedules/{id}/comments
	TESTED : POST /comments
	TESTED : PATCH /comments/{id}
	TESTED : DELETE /comments/{id}
*/
func TestCommentHandler(t *testing.T) {
	var (
		err     error
		request *http.Request

		jsonObject []byte
	)

	// Creating a temporary struct to decode the json.
	var tmp struct {
		CommentId int `json:"comment_id"`
	}

	// Preparing for HTTP requests
	rr := httptest.NewRecorder()

	//
	//	GET /comments
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/comments", nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing the request
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbAllComments model.Comments
	if err = json.NewDecoder(rr.Body).Decode(&dbAllComments); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(fakeComments, dbAllComments) {
		t.Error("Comments are not the same")
	}

	globals.Log.Debug("GET /comments - PASSED")

	//
	//	POST /comments
	//

	comment1 := model.Comment{
		ScheduleId:  1,
		Comment:     "New test comment",
		IsImportant: false,
	}
	// Turning the object into JSON
	if jsonObject, err = json.Marshal(comment1); err != nil {
		t.Error(err)
	}

	if request, err = http.NewRequest(http.MethodPost, "/comments", bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)

	// Sending the request
	r.ServeHTTP(rr, request)

	// Checking the result
	if err = json.NewDecoder(rr.Body).Decode(&tmp); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("POST /comments - PASSED")

	comment1.CommentId = int64(tmp.CommentId)

	//
	//	GET /comments/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/comments/"+strconv.FormatInt(comment1.CommentId, 10), nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbComment1 model.Comment
	if err = json.NewDecoder(rr.Body).Decode(&dbComment1); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(comment1, dbComment1) {
		t.Error("Comments are not the same")
	}
	globals.Log.Debug("GET /comments/{id} - PASSED")

	//
	//	GET /users/{id}/comments
	//

	// Fetching the comments of company 1
	commentsOfUser, _ := env.DB.GetCommentsOfUser(1)

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/users/1/comments", nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbCommentsOfUser model.Comments
	if err = json.NewDecoder(rr.Body).Decode(&dbCommentsOfUser); err != nil {
		t.Error(err)
	}

	// Verifying the result
	if !cmp.Equal(commentsOfUser, dbCommentsOfUser) {
		t.Error("Comments are not the same")
	}
	globals.Log.Debug("GET /companies/{id}/comments - PASSED")

	//
	//	GET /projects/{id}/comments
	//
	// Fetching the comments of project 1
	commentsOfProject, _ := env.DB.GetCommentsOfProject(1)

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/projects/1/comments", nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbCommentsOfProject model.Comments
	if err = json.NewDecoder(rr.Body).Decode(&dbCommentsOfProject); err != nil {
		t.Error(err)
	}

	// Verifying the result
	if !cmp.Equal(commentsOfProject, dbCommentsOfProject) {
		t.Error("Comments are not the same")
	}
	globals.Log.Debug("GET /projects/{id}/comments - PASSED")

	//
	//	GET /schedules/{id}/comments
	//
	// Fetching the comments of schedule 1
	commentsOfSchedule, _ := env.DB.GetCommentsOfSchedule(1)

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/schedules/1/comments", nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbCommentsOfSchedule model.Comments
	if err = json.NewDecoder(rr.Body).Decode(&dbCommentsOfSchedule); err != nil {
		t.Error(err)
	}

	// Verifying the result
	if !cmp.Equal(commentsOfSchedule, dbCommentsOfSchedule) {
		t.Error("Comments are not the same")
	}
	globals.Log.Debug("GET /schedules/{id}/comments - PASSED")

	//
	//	PATCH /comments/{id}
	//

	// Fetching a comment
	comment := fakeComments[1]
	// Modifying it
	comment.Comment = "Modified comment"
	// JSON-ing the object
	if jsonObject, err = json.Marshal(comment); err != nil {
		t.Error(err)
	}

	// Creating the request
	if request, err = http.NewRequest(http.MethodPatch, "/comments/"+strconv.FormatInt(comment.CommentId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new comment so we can see if he got changed : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/comments/"+strconv.FormatInt(comment.CommentId, 10), nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbComment model.Comment
	if err = json.NewDecoder(rr.Body).Decode(&dbComment); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(comment, dbComment) {
		t.Error("Comments are not the same")
	}
	globals.Log.Debug("PATCH /comments/{id} - PASSED")

	//
	//	DELETE /comments/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodDelete, "/comments/"+strconv.FormatInt(comment.CommentId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new comment so we can see if it got deleted : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/comments/"+strconv.FormatInt(comment.CommentId, 10), nil); err != nil {
		t.Error(err)
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	if err = json.NewDecoder(rr.Body).Decode(&dbComment); err == nil {
		t.Error(err)
	}

	globals.Log.Debug("DELETE /comments/{id} - PASSED")
}
