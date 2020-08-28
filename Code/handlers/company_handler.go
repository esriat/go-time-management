package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//	GetCompaniesHandler
/*	The handler called by the following endpoint : GET /companies
	This method is used to get the list of all existing comanies.
*/
func (env *Env) GetCompaniesHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		companies model.Companies
	)

	globals.Log.Debug("Calling GetCompaniesHandler")

	if companies, err = env.DB.GetCompanies(); err != nil {
		return &AppError{
			Error:   err,
			Code:    http.StatusInternalServerError,
			Message: "Internal error retrieving all the comanies",
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(companies)
	return nil
}

//	GetCompanyHandler
/*	The handler called by the following endpoint : GET /companies/{id}
	This method is used to get a company by its id.
*/
func (env *Env) GetCompanyHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err     error
		company model.Company
		id      int
	)

	globals.Log.Debug("Calling GetCompanyHandler")

	vars := mux.Vars(r)

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	if company, err = env.DB.GetCompany(int64(id)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when fetching company",
			Code:    http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(company)
	return nil
}

//	CreateCompanyHandler
/*	The handler called by the following endpoint : POST /companies
	This method is used to create a new company.
*/
func (env *Env) CreateCompanyHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		company   model.Company
		companyId int64
	)

	globals.Log.Debug("CreateCompanyHandler called")

	if err = json.NewDecoder(r.Body).Decode(&company); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	globals.Log.Debug("Decoded company : " + company.String())
	globals.Log.Debug("Calling CreateCompany method")

	if companyId, err = env.DB.CreateCompany(company); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error creating the company",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Company created")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(struct {
		CompanyId int64 `json:"company_id"`
	}{
		CompanyId: companyId,
	}); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the company id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	UpdateCompanyHandler
/*	The handler called by the following endpoint : PATCH /companies/{id}
	This method is used to update an existing company.
*/
func (env *Env) UpdateCompanyHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		company   model.Company
		companyId int
	)

	globals.Log.Debug("CreateCompanyHandler called")

	if err = json.NewDecoder(r.Body).Decode(&company); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when decoding the form",
			Code:    http.StatusBadRequest,
		}
	}

	vars := mux.Vars(r)

	if companyId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	company.CompanyId = int64(companyId)

	globals.Log.Debug("Decoded company : " + company.String())
	globals.Log.Debug("Calling CreateCompany method")

	if company, err = env.DB.UpdateCompany(company); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when updating the company",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Company updated")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(company); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when encoding the company id",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

//	DeleteCompanyHandler
/*	The handler called by the following endpoint : DELETE /companies/{id}
	This method is used to delete a company.
*/
func (env *Env) DeleteCompanyHandler(w http.ResponseWriter, r *http.Request) *AppError {
	var (
		err       error
		companyId int
	)

	globals.Log.Debug("DeleteCompanyHandler called")

	vars := mux.Vars(r)

	if companyId, err = strconv.Atoi(vars["id"]); err != nil {
		return &AppError{
			Error:   err,
			Message: "Id atoi conversion error",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Calling CreateCompany method")

	if err = env.DB.DeleteCompany(int64(companyId)); err != nil {
		return &AppError{
			Error:   err,
			Message: "Error when deleting the company",
			Code:    http.StatusInternalServerError,
		}
	}

	globals.Log.Debug("Company deleted")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return nil
}
