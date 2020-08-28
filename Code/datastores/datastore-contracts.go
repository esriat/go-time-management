package datastores

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//  GetContracts() (model.Contracts, error)
/*	This method is used to get all the existing contracts.
	Returns the list of the contracts or an error
*/
func (db *ConcreteDatastore) GetContracts() (model.Contracts, error) {
	var (
		rows *sqlx.Rows
		err  error
	)

	// Setting up the request and executing it
	request := "SELECT * FROM Contract;"
	if rows, err = db.Queryx(request); err != nil {
		return nil, err
	}

	// Formatting the data
	ContractsList := model.Contracts{}

	for rows.Next() {
		Contract := model.Contract{}
		if err = rows.StructScan(&Contract); err != nil {
			return nil, err
		}
		ContractsList = append(ContractsList, Contract)
	}

	defer rows.Close()
	return ContractsList, nil
}

//  GetContract(ContractId int64)
/*	This method is used to get a specific contract from the database, identified by its id.
	Returns the wanted contract or an error
*/
func (db *ConcreteDatastore) GetContract(ContractId int64) (model.Contract, error) {
	var (
		err      error
		contract model.Contract
	)

	// Setting up the request and executing it
	request := `SELECT * FROM Contract WHERE contract_id=?`
	if err = db.Get(&contract, request, ContractId); err != nil {
		return model.Contract{}, err
	}

	return contract, nil
}

//  GetContractsOfUser(UserId int64)
/*	This method fetches the contract of a given user.
	Returns the contract of the user or an error
*/
func (db *ConcreteDatastore) GetContractOfUser(UserId int64) (model.Contract, error) {
	var (
		err      error
		contract model.Contract
	)

	// Setting up the request and executing it
	request := `SELECT * 
	FROM Contract
	WHERE contract_id = (SELECT contract_id
						 FROM User
						 WHERE user_id = ?)`
	if err = db.Get(&contract, request, UserId); err != nil {
		return model.Contract{}, err
	}

	return contract, nil
}

//  CreateContract(ContractId int64)
/*	This method is used to create a bew contract.
	Returns the id of the new contract or an error
*/
func (db *ConcreteDatastore) CreateContract(Contract model.Contract) (int64, error) {
	var (
		tx         *sql.Tx
		err        error
		res        sql.Result
		contractId int64
	)

	// Preparing the request
	if tx, err = db.Begin(); err != nil {
		return -1, err
	}

	// Setting up the request and executing it
	request := `INSERT INTO Contract(contract_name) VALUES (?)`
	if res, err = tx.Exec(request, Contract.ContractName); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Getting the id of the new inserted contract
	if contractId, err = res.LastInsertId(); err != nil {
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

	return contractId, nil
}

//  DeleteContract(ContractId int64) error
/*	This method is used to delete a contract
	Can return an error
*/
func (db *ConcreteDatastore) DeleteContract(ContractId int64) error {
	request := `DELETE FROM Contract 
	WHERE contract_id=?`
	if _, err := db.Exec(request, ContractId); err != nil {
		return err
	}
	return nil
}

//  UpdateContract(Contract model.Contract) (model.Contract, error)
/*	This method is used to update an existing contract
	Returns the new contract or an error
*/
func (db *ConcreteDatastore) UpdateContract(Contract model.Contract) (model.Contract, error) {
	var (
		tx  *sql.Tx
		err error
	)

	// Preparing the request
	if tx, err = db.Begin(); err != nil {
		return model.Contract{}, err
	}

	// Setting up the request and executing it
	request := `UPDATE Contract 
	SET contract_name=?
	WHERE contract_id=?`
	if _, err = tx.Exec(request, Contract.ContractName, Contract.ContractId); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Contract{}, errr
		}
		return model.Contract{}, err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Contract{}, errr
		}
		return model.Contract{}, err
	}

	return Contract, nil
}
