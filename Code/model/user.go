package model

import (
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// User represents a User of the application.
/*	ContractId : The id of the user's contract
	FunctionId : Same, for the function.
	RoleId : Same for the role.
	Username : The user's UCA username
	LastName : User's last name.
	Firstname : User's first name.
	Mail : User's UCA email address.
	TheoricalHoursWorked : The theorical number of hours the user has to work every week (probably 35).
	VacationHours : The remaining paid vacation hours the user has.
*/
type User struct {
	UserId               int64  `db:"user_id" json:"user_id"`
	ContractId           int64  `db:"contract_id" json:"contract_id"`
	RoleId               int64  `db:"role_id" json:"role_id"`
	Username             string `db:"username" json:"username"`
	Password             string `db:"password" json:"password"`
	LastName             string `db:"last_name" json:"last_name"`
	FirstName            string `db:"first_name" json:"first_name"`
	Mail                 string `db:"mail" json:"mail"`
	TheoricalHoursWorked int64  `db:"theorical_hours_worked" json:"theorical_hours_worked"`
	VacationHours        int64  `db:"vacation_hours" json:"vacation_hours"`
}

type Users []User
