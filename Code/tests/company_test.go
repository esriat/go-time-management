package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/datastores"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

/*
	TESTED : GetCompanies() (model.Companies, error)
	TESTED : GetCompany(CompanyId int64) (model.Company, error)
	TESTED : CreateCompany(Company model.Company) (int64, error)
	TESTED : DeleteCompany(CompanyId int64) error
	TESTED : UpdateCompany(Company model.Company) (model.Company, error)
*/
func TestCompany(t *testing.T) {
	// Initializing variables
	var (
		err           error
		testDatastore *datastores.ConcreteDatastore
		allCompanies  model.Companies
	)

	if testDatastore, err = datastores.NewDatabase("myTestDatabase.db"); err != nil {
		t.Error(err)
	}

	company1 := model.Company{
		CompanyId:   0,
		CompanyName: "Company 1",
	}

	company2 := model.Company{
		CompanyId:   0,
		CompanyName: "Company 2",
	}

	company3 := model.Company{
		CompanyId:   0,
		CompanyName: "Company 3",
	}

	//
	//	CreateCompany(Company) (CompanyId, error)
	//

	// Creating the 3 companies
	if company1.CompanyId, err = testDatastore.CreateCompany(company1); err != nil {
		t.Error(err)
	}

	if company2.CompanyId, err = testDatastore.CreateCompany(company2); err != nil {
		t.Error(err)
	}

	if company3.CompanyId, err = testDatastore.CreateCompany(company3); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("CreateCompany test - PASSED")

	//
	//	GetCompanies() (Companies, error)
	//

	// List of all companies
	companyList := model.Companies{}
	companyList = append(companyList, company1)
	companyList = append(companyList, company2)
	companyList = append(companyList, company3)

	// Fetching all companies
	if allCompanies, err = testDatastore.GetCompanies(); err != nil {
		t.Error(err)
	}

	// Comparing them
	if !cmp.Equal(allCompanies, companyList) {
		t.Error(err)
	}

	globals.Log.Debug("GetCompanies test - PASSED")

	//
	//	GetCompany(CompanyId)
	//
	var company model.Company

	// Getting the company
	if company, err = testDatastore.GetCompany(company1.CompanyId); err != nil {
		t.Error(err)
	}

	// Comparing it
	if !cmp.Equal(company1, company) {
		t.Error(err)
	}

	globals.Log.Debug("GetCompany test - PASSED")

	//
	//	UpdateCompany(Company)
	//
	var updatedCompany model.Company

	// Updating the company
	company1.CompanyName = "New company name"
	if _, err = testDatastore.UpdateCompany(company1); err != nil {
		t.Error(err)
	}

	// Getting the updated company
	if updatedCompany, err = testDatastore.GetCompany(company1.CompanyId); err != nil {
		t.Error(err)
	}

	// Comparing
	if !cmp.Equal(company1, updatedCompany) {
		t.Error(err)
	}

	globals.Log.Debug("UpdateCompany test - PASSED")

	//
	//	DeleteCompany(CompanyId)
	//

	// Trying to delete
	if err = testDatastore.DeleteCompany(company1.CompanyId); err != nil {
		t.Error(err)
	}

	// Trying to get the deleted company
	if _, err = testDatastore.GetCompany(company1.CompanyId); err == nil {
		// If we don't have any error, the test failed
		t.Error()
	}

	testDatastore.CloseDatabase()
}
