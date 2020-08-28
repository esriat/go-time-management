package handler_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

/*
	TESTED : GET /companies
	TESTED : GET /companies/{id}
	TESTED : POST /companies
	TESTED : PATCH /companies/{id}
	TESTED : DELETE /companies/{id}
*/
func TestCompanyHandler(t *testing.T) {
	var (
		err     error
		request *http.Request

		jsonObject []byte
	)

	// Creating a temporary struct to decode the json.
	var tmp struct {
		CompanyId int `json:"company_id"`
	}

	// Preparing for HTTP requests
	rr := httptest.NewRecorder()

	//
	//	GET /companies
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/companies", nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing the request
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbAllCompanies model.Companies
	if err = json.NewDecoder(rr.Body).Decode(&dbAllCompanies); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(fakeCompanies, dbAllCompanies) {
		t.Error("Companies are not the same")
	}

	globals.Log.Debug("GET /companies - PASSED")

	//
	//	POST /companies
	//

	company1 := model.Company{
		CompanyName: "New test company",
	}
	// Turning the object into JSON
	if jsonObject, err = json.Marshal(company1); err != nil {
		t.Error(err.Error())
	}

	if request, err = http.NewRequest(http.MethodPost, "/companies", bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)

	// Sending the request
	r.ServeHTTP(rr, request)

	// Checking the result
	if err = json.NewDecoder(rr.Body).Decode(&tmp); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("POST /companies - PASSED")

	company1.CompanyId = int64(tmp.CompanyId)

	//
	//	GET /companies/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodGet, "/companies/"+strconv.FormatInt(company1.CompanyId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbCompany1 model.Company
	if err = json.NewDecoder(rr.Body).Decode(&dbCompany1); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(company1, dbCompany1) {
		t.Error("Companies are not the same")
	}
	globals.Log.Debug("GET /companies/{id} - PASSED")

	//
	//	PATCH /companies/{id}
	//

	// Fetching a company
	company := fakeCompanies[1]
	// Modifying it
	company.CompanyName = "Modified company"
	// JSON-ing the object
	if jsonObject, err = json.Marshal(company); err != nil {
		t.Error(err.Error())
	}

	// Creating the request
	if request, err = http.NewRequest(http.MethodPatch, "/companies/"+strconv.FormatInt(company.CompanyId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new company so we can see if he got changed : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/companies/"+strconv.FormatInt(company.CompanyId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	var dbCompany model.Company
	if err = json.NewDecoder(rr.Body).Decode(&dbCompany); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(company, dbCompany) {
		t.Error("Companies are not the same")
	}
	globals.Log.Debug("PATCH /companies/{id} - PASSED")

	//
	//	DELETE /companies/{id}
	//

	// Creating the request
	if request, err = http.NewRequest(http.MethodDelete, "/companies/"+strconv.FormatInt(company.CompanyId, 10), bytes.NewBuffer(jsonObject)); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Now we need to get the new company so we can see if it got deleted : creating the GET request
	if request, err = http.NewRequest(http.MethodGet, "/companies/"+strconv.FormatInt(company.CompanyId, 10), nil); err != nil {
		t.Error(err.Error())
	}
	request.AddCookie(tokenCookie)
	// Executing it
	r.ServeHTTP(rr, request)

	// Reading the result
	if err = json.NewDecoder(rr.Body).Decode(&dbCompany); err == nil {
		t.Error(err)
	}

	globals.Log.Debug("DELETE /companies/{id} - PASSED")
}
