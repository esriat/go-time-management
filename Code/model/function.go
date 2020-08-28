package model

import (
	"strconv"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Function represent the Function of a user
/*	FunctionName : Chimiste/Biologiste/Chef de projet...
 */
type Function struct {
	FunctionId   int64  `db:"function_id" json:"function_id"`
	FunctionName string `db:"function_name" json:"function_name"`
}

type Functions []Function

func (function *Function) String() string {
	return "Function id : " + strconv.FormatInt(function.FunctionId, 10) + "/ Function name : " + function.FunctionName
}
