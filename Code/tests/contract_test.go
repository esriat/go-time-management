package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/datastores"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

/*
	TESTED : GetContracts() (model.Contracts, error)
	TESTED : GetContract(ContractId int64) (model.Contract, error)
	TESTED : GetContractOfUser(UserId int64) (model.Contract, error)
	TESTED : CreateContract(Contract model.Contract) (int64, error)
	TESTED : DeleteContract(ContractId int64) error
	TESTED : UpdateContract(Contract model.Contract) (model.Contract, error)
*/
func TestContract(t *testing.T) {
	// Initializing variables
	var (
		err             error
		testDatastore   *datastores.ConcreteDatastore
		defaultContract model.Contract
	)

	if testDatastore, err = datastores.NewDatabase("myTestDatabase.db"); err != nil {
		t.Error(err)
	}

	if defaultContract, err = testDatastore.GetContract(1); err != nil {
		t.Error(err)
	}

	contract1 := model.Contract{
		ContractId:   0,
		ContractName: "Contract 1",
	}

	contract2 := model.Contract{
		ContractId:   0,
		ContractName: "Contract 2",
	}

	contract3 := model.Contract{
		ContractId:   0,
		ContractName: "Contract 3",
	}

	//
	//	CreateContract(Contract)
	//

	if contract1.ContractId, err = testDatastore.CreateContract(contract1); err != nil {
		t.Error(err)
	}

	if contract2.ContractId, err = testDatastore.CreateContract(contract2); err != nil {
		t.Error(err)
	}

	if contract3.ContractId, err = testDatastore.CreateContract(contract3); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("CreateContract test = PASSED")

	//
	//	GetContracts()
	//

	// Formatting data
	contractList := model.Contracts{}
	contractList = append(contractList, defaultContract)
	contractList = append(contractList, contract1)
	contractList = append(contractList, contract2)
	contractList = append(contractList, contract3)

	// Getting the contracts
	var allContracts model.Contracts
	if allContracts, err = testDatastore.GetContracts(); err != nil {
		t.Error(err)
	}

	if !cmp.Equal(allContracts, contractList) {
		t.Error(err)
	}

	globals.Log.Debug("GetContracts test = PASSED")

	//
	// Test GetContract(ContractId)
	//
	var contract model.Contract

	// Getting the contract
	if contract, err = testDatastore.GetContract(contract1.ContractId); err != nil {
		t.Error(err)
	}

	// Comparing
	if !cmp.Equal(contract1, contract) {
		t.Error(err)
	}

	globals.Log.Debug("GetContract test = PASSED")

	//
	//	GetContractOfUser
	//

	// Doing some data stuff
	roleId, _ := testDatastore.CreateRole(model.Role{
		RoleId:   0,
		RoleName: "Test Role",
	})

	user1 := model.User{
		UserId:               0,
		ContractId:           contract1.ContractId,
		RoleId:               roleId,
		Username:             "First User",
		Password:             "Password",
		LastName:             "User",
		FirstName:            "First",
		Mail:                 "FirstUser@user.com",
		TheoricalHoursWorked: 50,
		VacationHours:        50,
	}

	user2 := model.User{
		UserId:               0,
		ContractId:           contract2.ContractId,
		RoleId:               roleId,
		Username:             "Second User",
		Password:             "Password",
		LastName:             "User",
		FirstName:            "Second",
		Mail:                 "SecondUser@user.com",
		TheoricalHoursWorked: 50,
		VacationHours:        50,
	}

	if user1.UserId, err = testDatastore.CreateUser(user1); err != nil {
		t.Error(err)
	}

	if user2.UserId, err = testDatastore.CreateUser(user2); err != nil {
		t.Error(err)
	}

	user1Contract := model.Contract{}
	user2Contract := model.Contract{}

	// Getting the contracts of the 2 users

	if user1Contract, err = testDatastore.GetContractOfUser(user1.UserId); err != nil {
		t.Error(err)
	}

	if user2Contract, err = testDatastore.GetContractOfUser(user2.UserId); err != nil {
		t.Error(err)
	}

	// Comparing the data

	if !cmp.Equal(user1Contract, contract1) {
		t.Error(err)
	}

	if !cmp.Equal(user2Contract, contract2) {
		t.Error(err)
	}

	globals.Log.Debug("GetContractOfUser test = PASSED")

	//
	//	UpdateContract(Contract)
	//

	// Updating the contract
	var updatedContract model.Contract

	contract1.ContractName = "New contract name"
	testDatastore.UpdateContract(contract1)

	// Getting the updated contract
	if updatedContract, err = testDatastore.GetContract(contract1.ContractId); err != nil {
		t.Error(err)
	}

	// Comparing
	if !cmp.Equal(contract1, updatedContract) {
		t.Error(err)
	}

	globals.Log.Debug("UpdateContract test = PASSED")

	//
	//	DeleteContract(ContractId)
	//

	var ind int64
	if ind, err = testDatastore.CreateContract(model.Contract{
		ContractId:   0,
		ContractName: "Truc",
	}); err != nil {
		t.Error(err)
	}

	// Deleting the contract
	if err = testDatastore.DeleteContract(ind); err != nil {
		t.Error(err)
	}

	// The contract got deleted ?
	if _, err = testDatastore.GetContract(ind); err == nil {
		t.Error()
	}

	globals.Log.Debug("DeleteContract test = PASSED")

	testDatastore.CloseDatabase()
}
