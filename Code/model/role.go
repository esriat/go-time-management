package model

import (
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Role : Defines a role a user can have, and its given permissions.
/*	RoleName : The name of the role : Admin/User...
	CanAddAndModifyUsers : Wether this role can add new Users and Modify their rights.
	CanSeeOtherSchedules : Wether this role can see other Users' schedules.
	CanAddProjects : Wether this role can add data in the Projects table.
	CanSeeReports : Wether this role can see the reports (a report includes all schedules and their comments and users)
*/
type Role struct {
	RoleId               int64  `db:"role_id" json:"role_id"`
	RoleName             string `db:"role_name" json:"role_name"`
	CanAddAndModifyUsers bool   `db:"can_add_and_modify_users" json:"can_add_and_modify_users"`
	CanSeeOtherSchedules bool   `db:"can_see_other_schedules" json:"can_see_other_schedules"`
	CanAddProjects       bool   `db:"can_add_projects" json:"can_add_projects"`
	CanSeeReports        bool   `db:"can_see_reports" json:"can_see_reports"`
}

type Roles []Role
