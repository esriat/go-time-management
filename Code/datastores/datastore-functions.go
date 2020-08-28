package datastores

import (
	"database/sql"

	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//  GetFunctions() (model.Functions, error)
/*	This method fetches all the existing functions.
	Returns the list of the functions or an error
*/
func (db *ConcreteDatastore) GetFunctions() (model.Functions, error) {
	// Setting up the request and executing it
	request := "SELECT * FROM Function;"
	rows, err := db.Queryx(request)
	if err != nil {
		return nil, err
	}

	// Formatting the data
	functionsList := model.Functions{}

	for rows.Next() {
		function := model.Function{}
		err := rows.StructScan(&function)
		if err != nil {
			return nil, err
		}
		functionsList = append(functionsList, function)
	}

	defer rows.Close()
	return functionsList, nil
}

//  GetFunction(FunctionId int64) (model.Function, error)
/*	This method fetches the function with the id given in parameters.
	Returns the wanted function or an error
*/
func (db *ConcreteDatastore) GetFunction(FunctionId int64) (model.Function, error) {
	var (
		err      error
		function model.Function
	)

	// Setting up the request and executing it
	request := `SELECT * FROM Function WHERE function_id=?`
	if err = db.Get(&function, request, FunctionId); err != nil {
		return model.Function{}, err
	}

	return function, nil
}

//  GetFunctionsOfUser(UserId int64) (model.Functions, error)
/*	This method fetches the function of a given user.
	Returns the list of the wanted function or an error
*/
func (db *ConcreteDatastore) GetFunctionsOfUser(UserId int64) (model.Functions, error) {
	// Setting up the request and executing it
	request := `SELECT * 
	FROM Function F, (SELECT function_id
					  FROM UserFunction
					  WHERE user_id=?) UF
	WHERE F.function_id=UF.function_id`
	rows, err := db.Queryx(request, UserId)
	if err != nil {
		return nil, err
	}

	// Formatting the data
	functionsList := model.Functions{}

	for rows.Next() {
		function := model.Function{}
		err := rows.StructScan(&function)
		if err != nil {
			return nil, err
		}
		functionsList = append(functionsList, function)
	}

	defer rows.Close()
	return functionsList, nil
}

//  CreateFunction(Function model.Function) (int64, error)
/*	This method is used to create a new function
	Returns the id of the new function or an error
*/
func (db *ConcreteDatastore) CreateFunction(Function model.Function) (int64, error) {
	var (
		tx         *sql.Tx
		err        error
		res        sql.Result
		functionId int64
	)

	//Ppreparing the request
	if tx, err = db.Begin(); err != nil {
		return -1, err
	}

	// Setting up the request and executing it
	request := `INSERT INTO Function(function_name) VALUES (?)`
	if res, err = tx.Exec(request, Function.FunctionName); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Getting the id of the new element
	if functionId, err = res.LastInsertId(); err != nil {
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

	return functionId, nil
}

//  DeleteFunction(FunctionId int64) error
/*	This method is used to delete a new function.
	Can return an error
*/
func (db *ConcreteDatastore) DeleteFunction(FunctionId int64) error {

	// Setting up the request and executing it
	request := `DELETE FROM Function 
	WHERE function_id=?`
	if _, err := db.Exec(request, FunctionId); err != nil {
		return err
	}
	return nil
}

//  UpdateFunction(Function model.Function) (model.Function, error)
/*	This method is used to update an existing function.
	Returns the modified function or an error
*/
func (db *ConcreteDatastore) UpdateFunction(Function model.Function) (model.Function, error) {
	var (
		tx  *sql.Tx
		err error
	)

	// Preparing the request
	if tx, err = db.Begin(); err != nil {
		return model.Function{}, err
	}

	// Setting up the request and executing it
	request := `UPDATE Function 
	SET function_name=?
	WHERE function_id=?`
	if _, err = tx.Exec(request, Function.FunctionName, Function.FunctionId); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Function{}, errr
		}
		return model.Function{}, err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Function{}, errr
		}
		return model.Function{}, err
	}

	return Function, nil
}
