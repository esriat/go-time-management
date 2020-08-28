package model

import (
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Comment : A comment is linked to a schedule
/*	ScheduleId : The id the comment is linked to.
	Comment : The text of the comment. Contains whatever the user wants to say about a given Schedule.
	IsImportant : Wether the comment is important or not. Generally, a final comment will be more important
		than a random comment when nothing special happened.
*/
type Comment struct {
	CommentId   int64  `db:"comment_id" json:"comment_id"`
	ScheduleId  int64  `db:"schedule_id" json:"schedule_id"`
	Comment     string `db:"comment" json:"comment"`
	IsImportant bool   `db:"is_important" json:"is_important"`
}

type Comments []Comment
