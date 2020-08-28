package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//	GetCommentsHandler
/*	The handler called by the following endpoint : GET /comments
	This method is used to get the list of comments.
*/
func (env *Env) GetCommentsHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err      error
		comments model.Comments
	)

	globals.Log.Debug("Calling GetCommentsHandler")

	if comments, err = env.DB.GetComments(); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error retrieving all the comments",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(comments)
	return nil
}

//	GetCommentHandler
/*	The handler called by the following endpoint : GET /comments/{id}
	This method is used to get a comment.
*/
func (env *Env) GetCommentHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err     error
		comment model.Comment
		id      int
	)

	globals.Log.Debug("Calling GetCommentHandler")

	vars := mux.Vars(r)

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if comment, err = env.DB.GetComment(int64(id)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching comment",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(comment)
	return nil
}

//	GetCommentsOfUserHandler
/*	The handler called by the following endpoint : GET /users/{user_id}/comments
	This method is used to get the list of comments from a specific user.
*/
func (env *Env) GetCommentsOfUserHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err      error
		userId   int
		comments model.Comments
	)

	globals.Log.Debug("Calling GetCommentsOfUserHandler")

	vars := mux.Vars(r)

	if userId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if comments, err = env.DB.GetCommentsOfUser(int64(userId)); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error retrieving the comments",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(comments)
	return nil
}

//	GetCommentsOfScheduleHandler
/*	The handler called by the following endpoint : GET /schedule/{schedule_id}/comments
	This method is used to get the list of comments linked to a specific schedule.
*/
func (env *Env) GetCommentsOfScheduleHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err        error
		scheduleId int
		comments   model.Comments
	)

	globals.Log.Debug("Calling GetCommentsOfScheduleHandler")

	vars := mux.Vars(r)

	if scheduleId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if comments, err = env.DB.GetCommentsOfSchedule(int64(scheduleId)); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error retrieving the comments",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(comments)
	return nil
}

//	GetCommentsOfProjectHandler
/*	The handler called by the following endpoint : GET /projects/{project_id}/comments
	This method is used to get the list of comments linked to a specific project.
*/
func (env *Env) GetCommentsOfProjectHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		projectId int
		comments  model.Comments
	)

	globals.Log.Debug("Calling GetCommentsOfProjectHandler")

	vars := mux.Vars(r)

	if projectId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if comments, err = env.DB.GetCommentsOfProject(int64(projectId)); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error retrieving the comments",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(comments)
	return nil
}

//	CreateCommentHandler
/*	The handler called by the following endpoint : POST /comments
	This method is used to create a comment.
*/
func (env *Env) CreateCommentHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		comment   model.Comment
		commentId int64
	)

	globals.Log.Debug("CreateCommentHandler called")

	if err = json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	globals.Log.Debug("Calling CreateComment method")

	if commentId, err = env.DB.CreateComment(comment); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error creating the comment",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Comment created")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(struct {
		CommentId int64 `json:"comment_id"`
	}{
		CommentId: commentId,
	}); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the comment id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	UpdateCommentHandler
/*	The handler called by the following endpoint : PATCH /comments/{id}
	This method is used to update an existing comment.
*/
func (env *Env) UpdateCommentHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		comment   model.Comment
		commentId int
	)

	globals.Log.Debug("CreateCommentHandler called")

	if err = json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	vars := mux.Vars(r)

	if commentId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	comment.CommentId = int64(commentId)

	globals.Log.Debug("Calling CreateComment method")

	if comment, err = env.DB.UpdateComment(comment); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when updating the comment",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Comment updated")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(comment); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the comment id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	DeleteCommentHandler
/*	The handler called by the following endpoint : DELETE /comments/{id}
	This method is used to delete an existing comment.
*/
func (env *Env) DeleteCommentHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		commentId int
	)

	globals.Log.Debug("DeleteCommentHandler called")

	vars := mux.Vars(r)

	if commentId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Calling CreateComment method")

	if err = env.DB.DeleteComment(int64(commentId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when deleting the comment",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Comment deleted")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}
