package model

import (
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// CompanyProject : Represents a link between a Company and a Project
/* CompanyId : The Id of the Company
/* ProjectId : The Id of the Project
*/
type CompanyProject struct {
	CompanyId int64 `db:"company_id" json:"company_id"`
	ProjectId int64 `db:"project_id" json:"project_id"`
}

type CompaniesProjects []CompanyProject

// CompanyUser : Represents a link between a Company and a User
/* CompanyId : The Id of the Company
/* UserId : The Id of the User
*/
type CompanyUser struct {
	CompanyId int64 `db:"company_id" json:"company_id"`
	UserId    int64 `db:"user_id" json:"user_id"`
}

type CompaniesUsers []CompanyUser

// UserSchedule : Represents a link between a User and a Schedule
/* UserId : The Id of the User
/* ScheduleId : The Id of the Project
*/
type UserSchedule struct {
	UserId     int64 `db:"user_id" json:"user_id"`
	ScheduleId int64 `db:"schedule_id" json:"schedule_id"`
}

type UsersSchedules []UserSchedule

// UserFunction : Represents a link between a User and a Function
/* UserId : The Id of the User
/* FunctionId : The Id of the Project
*/
type UserFunction struct {
	UserId     int64 `db:"user_id" json:"user_id"`
	FunctionId int64 `db:"function_id" json:"function_id"`
}

type UsersFunctions []UserFunction
