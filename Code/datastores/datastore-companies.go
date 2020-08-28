package datastores

import (
	"database/sql"

	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//  GetCompanies() (model.Companies, error)
/*	This method is used to get the list of all companies.
	Returns the list of companies, or an error
*/
func (db *ConcreteDatastore) GetCompanies() (model.Companies, error) {
	// Setting up the request and executing it
	request := "SELECT * FROM Company;"
	rows, err := db.Queryx(request)
	if err != nil {
		return nil, err
	}

	// Formatting the data
	companiesList := model.Companies{}

	for rows.Next() {
		company := model.Company{}
		err := rows.StructScan(&company)
		if err != nil {
			return nil, err
		}
		companiesList = append(companiesList, company)
	}

	defer rows.Close()
	return companiesList, nil
}

//  GetCompany(CompanyId int64) (model.Companies, error)
/*	This method is used to get a specific company.
	Takes the Id of the company as a parameter.
	Returns the wanted company, or an error
*/
func (db *ConcreteDatastore) GetCompany(CompanyId int64) (model.Company, error) {
	var (
		err     error
		company model.Company
	)

	// Setting up the request
	request := `SELECT * FROM Company WHERE company_id=?`
	if err = db.Get(&company, request, CompanyId); err != nil {
		return model.Company{}, err
	}

	return company, nil
}

//  CreateCompany(Company model.Company) (int64, error)
/*	This method is used to create a new company.
	Takes the new company as a parameter.
	Returns the id of the new company, or an error
*/
func (db *ConcreteDatastore) CreateCompany(Company model.Company) (int64, error) {
	var (
		tx        *sql.Tx
		err       error
		res       sql.Result
		companyId int64
	)

	// Preparing the request
	if tx, err = db.Begin(); err != nil {
		return -1, err
	}

	// Setting up the request and executing it
	request := `INSERT INTO Company(company_name) VALUES (?)`
	if res, err = tx.Exec(request, Company.CompanyName); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Retrieving the id of the new element
	if companyId, err = res.LastInsertId(); err != nil {
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

	return companyId, nil
}

//  DeleteCompany(CompanyId int64) (error)
/*	This method is used to delete a company.
	Can return an error
*/
func (db *ConcreteDatastore) DeleteCompany(CompanyId int64) error {
	// Setting up the request and executing it
	request := `DELETE FROM Company 
	WHERE company_id=?`
	if _, err := db.Exec(request, CompanyId); err != nil {
		return err
	}
	return nil
}

//  UpdateCompany(Company model.Company) (model.Company, error)
/*	This method is used to update an existing company.
	Takes the updated company as a parameter
	Returns the modified company, or an error
*/
func (db *ConcreteDatastore) UpdateCompany(Company model.Company) (model.Company, error) {
	var (
		tx  *sql.Tx
		err error
	)

	// Preparing the request
	if tx, err = db.Begin(); err != nil {
		return model.Company{}, err
	}

	// Setting up the request and executing it
	request := `UPDATE Company 
	SET company_name=?
	WHERE company_id=?`
	if _, err = tx.Exec(request, Company.CompanyName, Company.CompanyId); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Company{}, errr
		}
		return model.Company{}, err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Company{}, errr
		}
		return model.Company{}, err
	}

	return Company, nil
}
