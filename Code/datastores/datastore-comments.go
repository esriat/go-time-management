package datastores

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//  GetComments() (model.Comments, error)
/*	This method is used to get all the comment in the database.
 */
func (db *ConcreteDatastore) GetComments() (model.Comments, error) {
	var (
		rows *sqlx.Rows
		err  error
	)

	// Setting up and executing the request
	request := "SELECT * FROM Comment;"
	if rows, err = db.Queryx(request); err != nil {
		return nil, err
	}

	// Formatting the data we got from the request
	commentList := model.Comments{}
	for rows.Next() {
		comment := model.Comment{}
		if err = rows.StructScan(&comment); err != nil {
			return nil, err
		}
		commentList = append(commentList, comment)
	}

	defer rows.Close()
	return commentList, nil
}

//  GetComment(CommentId int64) (model.Comment, error)
/*	This method is used to get a specific comment from the database.
	It fetches the comment with the Id that's given in parameters.
*/
func (db *ConcreteDatastore) GetComment(CommentId int64) (model.Comment, error) {
	var (
		err     error
		comment model.Comment
	)

	// Setting up and executing the request
	request := `SELECT * FROM Comment WHERE Comment.comment_id=?`
	if err = db.Get(&comment, request, CommentId); err != nil {
		return model.Comment{}, err
	}

	return comment, nil
}

//  GetCommentsOfUser(UserId int64) (model.Comments, error)
/*	This method is used to get all the comments a given user posted.
	UserId represents the unique Id of the user we want to get the comments of.
*/
func (db *ConcreteDatastore) GetCommentsOfUser(UserId int64) (model.Comments, error) {
	var (
		err  error
		rows *sqlx.Rows
	)

	// Setting up and executing the request
	request := `SELECT c.comment_id, c.schedule_id, c.comment, c.is_important
	FROM Comment c, UserSchedule us
	WHERE c.schedule_id=us.schedule_id
	AND us.user_id=?
	`
	if rows, err = db.Queryx(request, UserId); err != nil {
		return nil, err
	}

	// Formatting the data we got from the request
	commentsList := model.Comments{}
	for rows.Next() {
		comment := model.Comment{}
		if err = rows.StructScan(&comment); err != nil {
			return nil, err
		}
		commentsList = append(commentsList, comment)
	}

	defer rows.Close()
	return commentsList, nil
}

//  GetCommentsOfSchedule(ScheduleId int64) (model.Comments, error)
/*	This method is used to get all the comments linked to a given schedule.
	ScheduleId represents the unique Id of the schedule we want to get the comments linked to.
*/
func (db *ConcreteDatastore) GetCommentsOfSchedule(ScheduleId int64) (model.Comments, error) {
	var (
		err  error
		rows *sqlx.Rows
	)

	// Setting up and executing the request
	request := "SELECT * FROM Comment WHERE Comment.schedule_id=?;"
	if rows, err = db.Queryx(request, ScheduleId); err != nil {
		return nil, err
	}

	// Formatting the data we got from the request
	commentsList := model.Comments{}
	for rows.Next() {
		comment := model.Comment{}
		if err = rows.StructScan(&comment); err != nil {
			return nil, err
		}
		commentsList = append(commentsList, comment)
	}

	defer rows.Close()
	return commentsList, nil
}

//  GetCommentsOfProject(ProjectId int64) (model.Comments, error)
/*	This method is used to get all the comments linked to a given schedule.
	ScheduleId represents the unique Id of the schedule we want to get the comments linked to.
*/
func (db *ConcreteDatastore) GetCommentsOfProject(ProjectId int64) (model.Comments, error) {
	var (
		err  error
		rows *sqlx.Rows
	)

	// Setting up and executing the request
	request := `SELECT Comment.comment_id, Comment.schedule_id, Comment.comment, Comment.is_important
	FROM Comment, Schedule
	WHERE Schedule.schedule_id=Comment.schedule_id
	AND Schedule.project_id=?
	`
	if rows, err = db.Queryx(request, ProjectId); err != nil {
		return nil, err
	}

	// Formatting the data
	commentsList := model.Comments{}
	for rows.Next() {
		comment := model.Comment{}
		if err = rows.StructScan(&comment); err != nil {
			return nil, err
		}
		commentsList = append(commentsList, comment)
	}

	defer rows.Close()
	return commentsList, nil
}

//  CreateComment(Comment model.Comment) (int64, error)
/*	This method is used to create a new comment.
	It returns the id of the created comment, or an error.
*/
func (db *ConcreteDatastore) CreateComment(Comment model.Comment) (int64, error) {
	var (
		tx        *sql.Tx
		err       error
		res       sql.Result
		commentId int64
	)

	// Preparing to request
	if tx, err = db.Begin(); err != nil {
		return -1, err
	}

	// Setting up the request and executing it
	request := `INSERT INTO Comment(schedule_id, comment, is_important) VALUES (?, ?, ?)`
	if res, err = tx.Exec(request, Comment.ScheduleId, Comment.Comment, Comment.IsImportant); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Getting the id of the last item inserted
	if commentId, err = res.LastInsertId(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	return commentId, nil
}

//  DeleteComment(CommentId int64) error
/*	This method is used to delete the comment of the given id.
	Returns an error if there is one
*/
func (db *ConcreteDatastore) DeleteComment(CommentId int64) error {
	// Setting up and executing the request
	request := `DELETE FROM Comment 
	WHERE comment_id=?`
	if _, err := db.Exec(request, CommentId); err != nil {
		return err
	}
	return nil
}

//  UpdateComment(Comment model.Comment) (model.Comments, error)
/*	This method is used to update a comment.
	The comment taken as a parameter is the new version of the comment.
	Returns the new comment, or an error
*/
func (db *ConcreteDatastore) UpdateComment(Comment model.Comment) (model.Comment, error) {
	var (
		tx  *sql.Tx
		err error
	)

	// Preparing the request
	if tx, err = db.Begin(); err != nil {
		return model.Comment{}, err
	}

	// Setting up the request and executing it
	request := `UPDATE Comment 
	SET schedule_id=?, comment=?, is_important=? 
	WHERE comment_id=?`

	if _, err = tx.Exec(request, Comment.ScheduleId, Comment.Comment, Comment.IsImportant, Comment.CommentId); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Comment{}, errr
		}
		return model.Comment{}, err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Comment{}, errr
		}
		return model.Comment{}, err
	}

	return Comment, nil
}
