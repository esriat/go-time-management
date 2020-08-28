package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/datastores"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

/*
	TESTED : GetFunctions() (model.Functions, error)
	TESTED : GetFunction(FunctionId int64) (model.Function, error)
	TESTED : GetFunctionsOfUser(UserId int64) (model.Functions, error)
	TESTED : CreateFunction(Function model.Function) (int64, error)
	TESTED : DeleteFunction(FunctionId int64) error
	TESTED : UpdateFunction(Function model.Function) (model.Function, error)
*/
func TestFunction(t *testing.T) {
	var (
		err           error
		testDatastore *datastores.ConcreteDatastore
	)

	// Initializing variables
	if testDatastore, err = datastores.NewDatabase("myTestDatabase.db"); err != nil {
		t.Error(err)
	}

	function1 := model.Function{
		FunctionId:   0,
		FunctionName: "Function 1",
	}

	function2 := model.Function{
		FunctionId:   0,
		FunctionName: "Function 2",
	}

	function3 := model.Function{
		FunctionId:   0,
		FunctionName: "Function 3",
	}

	//
	//	CreateFunction(Function)
	//
	if function1.FunctionId, err = testDatastore.CreateFunction(function1); err != nil {
		t.Error(err)
	}

	if function2.FunctionId, err = testDatastore.CreateFunction(function2); err != nil {
		t.Error(err)
	}

	if function3.FunctionId, err = testDatastore.CreateFunction(function3); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("CreateFunction test - PASSED")

	//
	//	GetFunctions()
	//

	// Formatting the existing data
	var allFunctions model.Functions

	functionList := model.Functions{}
	functionList = append(functionList, function1)
	functionList = append(functionList, function2)
	functionList = append(functionList, function3)

	// Getting the list of functions
	if allFunctions, err = testDatastore.GetFunctions(); err != nil {
		t.Error(err)
	}

	// Comparing the data
	if !cmp.Equal(allFunctions, functionList) {
		t.Error(err)
	}

	globals.Log.Debug("GetFunctions test - PASSED")

	//
	// Test GetFunction(FunctionId)
	//

	// Fetching the function
	var function model.Function
	if function, err = testDatastore.GetFunction(function1.FunctionId); err != nil {
		t.Error(err)
	}

	// Comparing it
	if !cmp.Equal(function1, function) {
		t.Error(err)
	}

	globals.Log.Debug("GetFunction test - PASSED")

	//
	//	GetFunctionOfUser
	//

	// Creating some data
	contract := model.Contract{
		ContractId:   0,
		ContractName: "Test contract",
	}

	if contract.ContractId, err = testDatastore.CreateContract(contract); err != nil {
		t.Error(err)
	}

	role := model.Role{
		RoleId:   0,
		RoleName: "Test role",
	}

	if role.RoleId, err = testDatastore.CreateRole(role); err != nil {
		t.Error(err)
	}

	user1 := model.User{
		UserId:               0,
		ContractId:           contract.ContractId,
		RoleId:               role.RoleId,
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
		ContractId:           contract.ContractId,
		RoleId:               role.RoleId,
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

	if err = testDatastore.CreateUserFunction(model.UserFunction{
		UserId:     user1.UserId,
		FunctionId: function1.FunctionId,
	}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserFunction(model.UserFunction{
		UserId:     user1.UserId,
		FunctionId: function2.FunctionId,
	}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserFunction(model.UserFunction{
		UserId:     user2.UserId,
		FunctionId: function2.FunctionId,
	}); err != nil {
		t.Error(err)
	}

	if err = testDatastore.CreateUserFunction(model.UserFunction{
		UserId:     user2.UserId,
		FunctionId: function3.FunctionId,
	}); err != nil {
		t.Error(err)
	}
	functionsOfUser1 := model.Functions{}
	functionsOfUser1 = append(functionsOfUser1, function1)
	functionsOfUser1 = append(functionsOfUser1, function2)

	functionsOfUser2 := model.Functions{}
	functionsOfUser2 = append(functionsOfUser2, function2)
	functionsOfUser2 = append(functionsOfUser2, function3)

	dbFU1 := model.Functions{}
	dbFU2 := model.Functions{}

	// Now getting functions of users
	if dbFU1, err = testDatastore.GetFunctionsOfUser(user1.UserId); err != nil {
		t.Error(err)
	}

	if dbFU2, err = testDatastore.GetFunctionsOfUser(user2.UserId); err != nil {
		t.Error(err)
	}

	// Comparing the data
	if !cmp.Equal(dbFU1, functionsOfUser1) {
		t.Error(err)
	}

	if !cmp.Equal(dbFU2, functionsOfUser2) {
		t.Error(err)
	}

	globals.Log.Debug("GetFunctionOfuUer test - PASSED")

	//
	//	UpdateFunction(Function)
	//

	// Updating the function
	var updatedFunction model.Function
	function1.FunctionName = "New function name"
	testDatastore.UpdateFunction(function1)

	if updatedFunction, err = testDatastore.GetFunction(function1.FunctionId); err != nil {
		t.Error(err)
	}

	// Comparing
	if !cmp.Equal(function1, updatedFunction) {
		t.Error(err)
	}

	globals.Log.Debug("UpdateFunction test - PASSED")

	//
	// Test DeleteFunction(FunctionId)
	//

	// Creating a function without dependency
	var ind int64
	if ind, err = testDatastore.CreateFunction(model.Function{
		FunctionId:   0,
		FunctionName: "Olala",
	}); err != nil {
		t.Error(err)
	}

	// So we can try our telete
	if err = testDatastore.DeleteFunction(ind); err != nil {
		t.Error(err)
	}

	if _, err = testDatastore.GetFunction(ind); err == nil {
		t.Error()
	}

	globals.Log.Debug("DeleteFunction test - PASSED")

	testDatastore.CloseDatabase()
}
