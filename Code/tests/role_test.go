package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/datastores"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

/*
	TESTED : GetRoles() (model.Roles, error)
	TESTED : GetRole(RoleId int64) (model.Role, error)
	TESTED : GetRolesOfUser(UserId int64) (model.Roles, error)
	TESTED : CreateRole(Role model.Role) (int64, error)
	TESTED : DeleteRole(RoleId int64) error
	TESTED : UpdateRole(Role model.Role) (model.Role, error)
*/
func TestRole(t *testing.T) {
	// Initializing variables
	var (
		err           error
		testDatastore *datastores.ConcreteDatastore
	)

	if testDatastore, err = datastores.NewDatabase("myTestDatabase.db"); err != nil {
		t.Error(err)
	}

	role1 := model.Role{
		RoleId:               0,
		RoleName:             "Role 1",
		CanAddAndModifyUsers: true,
		CanSeeOtherSchedules: true,
		CanAddProjects:       true,
		CanSeeReports:        true,
	}

	role2 := model.Role{
		RoleId:               0,
		RoleName:             "Role 2",
		CanAddAndModifyUsers: false,
		CanSeeOtherSchedules: false,
		CanAddProjects:       true,
		CanSeeReports:        true,
	}

	role3 := model.Role{
		RoleId:               0,
		RoleName:             "Role 3",
		CanAddAndModifyUsers: false,
		CanSeeOtherSchedules: false,
		CanAddProjects:       false,
		CanSeeReports:        false,
	}

	//
	//	CreateRoles()
	//

	if role1.RoleId, err = testDatastore.CreateRole(role1); err != nil {
		t.Error(err)
	}

	if role2.RoleId, err = testDatastore.CreateRole(role2); err != nil {
		t.Error(err)
	}

	if role3.RoleId, err = testDatastore.CreateRole(role3); err != nil {
		t.Error(err)
	}

	globals.Log.Debug("CreateRole test - PASSED")

	//
	// Test GetRoles()
	//
	var (
		allRoles     model.Roles
		defaultRole1 model.Role
		defaultRole2 model.Role
		defaultRole3 model.Role
	)

	// Fetching the default roles
	if defaultRole1, err = testDatastore.GetRole(1); err != nil {
		t.Error(err)
	}

	if defaultRole2, err = testDatastore.GetRole(2); err != nil {
		t.Error(err)
	}

	if defaultRole3, err = testDatastore.GetRole(3); err != nil {
		t.Error(err)
	}

	// Formatting the data
	roleList := model.Roles{}
	roleList = append(roleList, defaultRole1)
	roleList = append(roleList, defaultRole2)
	roleList = append(roleList, defaultRole3)
	roleList = append(roleList, role1)
	roleList = append(roleList, role2)
	roleList = append(roleList, role3)

	// Fetching all roles
	if allRoles, err = testDatastore.GetRoles(); err != nil {
		t.Error(err)
	}

	// Verigying the result
	if !cmp.Equal(allRoles, roleList) {
		t.Error(err)
	}

	globals.Log.Debug("GetRoles test - PASSED")

	//
	// Test GetRole(RoleId)
	//
	var role model.Role

	// Fetching a role
	if role, err = testDatastore.GetRole(role1.RoleId); err != nil {
		t.Error(err)
	}

	// Verifying the data
	if !cmp.Equal(role1, role) {
		t.Error(err)
	}

	globals.Log.Debug("GetRole test - PASSED")

	//
	// Test GetRoleOfUser
	//

	// Creating some data
	contract := model.Contract{
		ContractId:   0,
		ContractName: "Contract",
	}

	if contract.ContractId, err = testDatastore.CreateContract(contract); err != nil {
		t.Error(err)
	}

	user := model.User{
		UserId:               0,
		ContractId:           contract.ContractId,
		RoleId:               role1.RoleId,
		Username:             "First User",
		Password:             "This is a password",
		LastName:             "User",
		FirstName:            "First",
		Mail:                 "FirstUser@user.com",
		TheoricalHoursWorked: 50,
		VacationHours:        50,
	}

	if user.UserId, err = testDatastore.CreateUser(user); err != nil {
		t.Error(err)
	}

	// Fetching the role of the user
	var dbRoleOfUser model.Role
	if dbRoleOfUser, err = testDatastore.GetRoleOfUser(user.UserId); err != nil {
		t.Error(err)
	}

	// Verifying the data
	if !cmp.Equal(dbRoleOfUser, role1) {
		t.Error(err)
	}

	globals.Log.Debug("GetRoleOfUser test - PASSED")

	//
	// Test UpdateRole(Role)
	//
	var updatedRole model.Role

	// Modifying the role
	role1.RoleName = "New role name"
	role1.CanAddAndModifyUsers = false
	role1.CanSeeOtherSchedules = false
	role1.CanAddProjects = false
	role1.CanSeeReports = false

	// Saving the changes
	if _, err = testDatastore.UpdateRole(role1); err != nil {
		t.Error(err)
	}

	// Getting the role so we can check the changes
	if updatedRole, err = testDatastore.GetRole(role1.RoleId); err != nil {
		t.Error(err)
	}

	// Checking changes
	if !cmp.Equal(role1, updatedRole) {
		t.Error(err)
	}

	globals.Log.Debug("UpdateRole test - PASSED")

	//
	// Test DeleteRole(RoleId)
	//

	// Creating a role so we can delete it
	var ind int64
	if ind, err = testDatastore.CreateRole(model.Role{
		RoleId:   0,
		RoleName: "This is a role",
	}); err != nil {
		t.Error(err)
	}

	// Deleting it
	if err = testDatastore.DeleteRole(ind); err != nil {
		t.Error(err)
	}

	// trying to get the role we just deleted
	if _, err = testDatastore.GetRole(ind); err == nil {
		t.Error()
	}

	globals.Log.Debug("DeleteRole test - PASSED")

	testDatastore.CloseDatabase()
}
