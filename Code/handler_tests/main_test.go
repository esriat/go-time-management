package handler_tests

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/datastores"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/globals"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/handlers"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	env         *handlers.Env
	r           *mux.Router
	tokenCookie *http.Cookie

	fakeProjects            model.Projects
	fakeContracts           model.Contracts
	fakeCompanies           model.Companies
	fakeCompanyProjectLinks []model.CompanyProject
	fakeFunctions           model.Functions
	fakeUsers               model.Users
	fakeCompanyUserLinks    []model.CompanyUser
	fakeUserFunctionLinks   []model.UserFunction
	fakeSchedules           model.Schedules
	fakeUserScheduleLinks   []model.UserSchedule
	fakeComments            model.Comments
	allRoles                model.Roles
)

func TestMain(m *testing.M) {
	SetUp()
	globals.Log.Debug("Starting tests")
	os.Exit(m.Run())
}

func SetUp() {
	InitializeDatabase()
	PopulateDatabase()
}

func InitializeDatabase() {
	var (
		err       error
		datastore datastores.IDatastore

		request    *http.Request
		jsonObject []byte
	)

	globals.Init()

	// Initializing datastore and env
	if datastore, err = datastores.NewDatabase("myDatabase.db"); err != nil {
		os.Exit(-1)
	}

	env = &handlers.Env{
		DB: datastore,
	}

	r = mux.NewRouter()

	handlers.HandleRoutes(r, env)

	// Getting a token
	// Creating a new http request
	rr := httptest.NewRecorder()

	user := model.User{
		Mail:     "admin@mydb",
		Password: "Admin",
	}

	if jsonObject, err = json.Marshal(user); err != nil {
		panic(err)
	}

	request, _ = http.NewRequest("POST", "/get-token", strings.NewReader(string(jsonObject)))
	request.Header.Add("Content-Type", "text/plain")

	// Sending the request
	r.ServeHTTP(rr, request)

	// Retrieving the token and storing it
	tokenCookie = rr.Result().Cookies()[0]

	globals.Log.Debug("Initialized data for tests")
}

func PopulateDatabase() {
	var (
		err error
	)
	//
	// 	First, creating fake data
	//

	// Create fake projects
	fakeProjects = model.Projects{}
	vacationProject, _ := env.DB.GetVacationProject()
	fakeProjects = append(fakeProjects, vacationProject)
	fakeProjects = append(fakeProjects, model.Project{
		ProjectName: "Project 1",
	})
	fakeProjects = append(fakeProjects, model.Project{
		ProjectName: "Second project",
	})
	fakeProjects = append(fakeProjects, model.Project{
		ProjectName: "3rd project",
	})

	// Create fake contracts
	fakeContracts = model.Contracts{}
	adminContract, _ := env.DB.GetContract(1)
	fakeContracts = append(fakeContracts, adminContract)
	fakeContracts = append(fakeContracts, model.Contract{
		ContractName: "CDD",
	})
	fakeContracts = append(fakeContracts, model.Contract{
		ContractName: "CDI",
	})

	// Create fake companies
	fakeCompanies = model.Companies{}
	fakeCompanies = append(fakeCompanies, model.Company{
		CompanyName: "Biopass",
	})
	fakeCompanies = append(fakeCompanies, model.Company{
		CompanyName: "Biomarqueurs",
	})

	// Create fake Company-Project links
	fakeCompanyProjectLinks = []model.CompanyProject{}
	fakeCompanyProjectLinks = append(fakeCompanyProjectLinks, model.CompanyProject{
		CompanyId: 1,
		ProjectId: 1,
	})
	fakeCompanyProjectLinks = append(fakeCompanyProjectLinks, model.CompanyProject{
		CompanyId: 1,
		ProjectId: 2,
	})
	fakeCompanyProjectLinks = append(fakeCompanyProjectLinks, model.CompanyProject{
		CompanyId: 2,
		ProjectId: 2,
	})

	// Create fake functions
	fakeFunctions = model.Functions{}
	fakeFunctions = append(fakeFunctions, model.Function{
		FunctionName: "First function",
	})
	fakeFunctions = append(fakeFunctions, model.Function{
		FunctionName: "Function 2",
	})

	// Creating fake users
	fakeUsers = model.Users{}
	firstUser, _ := env.DB.GetUser(1)
	fakeUsers = append(fakeUsers, firstUser)
	var cryptedPassword []byte
	if cryptedPassword, err = bcrypt.GenerateFromPassword([]byte("Password"), bcrypt.DefaultCost); err != nil {
		panic(err)
	}
	fakeUsers = append(fakeUsers, model.User{
		ContractId: 1,
		RoleId:     1,
		Mail:       "SecondUser@mydb",
		Password:   string(cryptedPassword),
	})
	fakeUsers = append(fakeUsers, model.User{
		ContractId: 2,
		RoleId:     2,
		Mail:       "ThirdUser@mydb",
		Password:   string(cryptedPassword),
	})

	// Create fake Company-User links
	fakeCompanyUserLinks = []model.CompanyUser{}
	fakeCompanyUserLinks = append(fakeCompanyUserLinks, model.CompanyUser{
		UserId:    1,
		CompanyId: 1,
	})
	fakeCompanyUserLinks = append(fakeCompanyUserLinks, model.CompanyUser{
		UserId:    2,
		CompanyId: 2,
	})

	// Create fake User-Function links
	fakeUserFunctionLinks = []model.UserFunction{}
	fakeUserFunctionLinks = append(fakeUserFunctionLinks, model.UserFunction{
		UserId:     1,
		FunctionId: 1,
	})
	fakeUserFunctionLinks = append(fakeUserFunctionLinks, model.UserFunction{
		UserId:     2,
		FunctionId: 1,
	})
	fakeUserFunctionLinks = append(fakeUserFunctionLinks, model.UserFunction{
		UserId:     2,
		FunctionId: 2,
	})

	// Creating fake schedules
	fakeSchedules = model.Schedules{}
	fakeSchedules = append(fakeSchedules, model.Schedule{
		ProjectId: 1,
		StartDate: sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:   sql.NullTime{Valid: true, Time: time.Now()},
	})
	fakeSchedules = append(fakeSchedules, model.Schedule{
		ProjectId: 1,
		StartDate: sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:   sql.NullTime{Valid: true, Time: time.Now()},
	})
	fakeSchedules = append(fakeSchedules, model.Schedule{
		ProjectId: 2,
		StartDate: sql.NullTime{Valid: true, Time: time.Now()},
		EndDate:   sql.NullTime{Valid: true, Time: time.Now()},
	})

	// Create fake User-Schedule links
	fakeUserScheduleLinks = []model.UserSchedule{}
	fakeUserScheduleLinks = append(fakeUserScheduleLinks, model.UserSchedule{
		UserId:     1,
		ScheduleId: 1,
	})
	fakeUserScheduleLinks = append(fakeUserScheduleLinks, model.UserSchedule{
		UserId:     1,
		ScheduleId: 2,
	})
	fakeUserScheduleLinks = append(fakeUserScheduleLinks, model.UserSchedule{
		UserId:     2,
		ScheduleId: 2,
	})
	fakeUserScheduleLinks = append(fakeUserScheduleLinks, model.UserSchedule{
		UserId:     2,
		ScheduleId: 3,
	})

	// Create fake comments
	fakeComments = model.Comments{}
	fakeComments = append(fakeComments, model.Comment{
		ScheduleId:  1,
		Comment:     "Comment 1",
		IsImportant: false,
	})
	fakeComments = append(fakeComments, model.Comment{
		ScheduleId:  2,
		Comment:     "Comment 2",
		IsImportant: true,
	})
	fakeComments = append(fakeComments, model.Comment{
		ScheduleId:  3,
		Comment:     "Comment 3",
		IsImportant: false,
	})

	//
	//	Then, inserting the fake data
	//
	var (
		id int64
	)

	// Inserting the fake projects
	for index, project := range fakeProjects {
		// Don't create the vacation project
		if index != 0 {
			if id, err = env.DB.CreateProject(project); err != nil {
				panic(err)
			}
			fakeProjects[index].ProjectId = id
		}
	}

	// Inserting fake contracts
	for index, contract := range fakeContracts {
		// Don't duplicate the Admin function
		if index != 0 {
			if id, err = env.DB.CreateContract(contract); err != nil {
				panic(err)
			}
			fakeContracts[index].ContractId = id
		}
	}

	// Inserting fake companies
	for index, company := range fakeCompanies {
		if id, err = env.DB.CreateCompany(company); err != nil {
			panic(err)
		}
		fakeCompanies[index].CompanyId = id
	}

	// Inserting fake Project-Company links
	for _, CP := range fakeCompanyProjectLinks {
		if err = env.DB.CreateCompanyProject(CP); err != nil {
			panic(err)
		}
	}

	// Inserting fake functions
	for index, function := range fakeFunctions {
		if id, err = env.DB.CreateFunction(function); err != nil {
			panic(err)
		}
		fakeFunctions[index].FunctionId = id
	}

	// Inserting fake users
	for index, user := range fakeUsers {
		if index != 0 {
			if id, err = env.DB.CreateUser(user); err != nil {
				panic(err)
			}
			fakeUsers[index].UserId = id
		}
	}

	// Inserting fake User-Company links
	for _, CU := range fakeCompanyUserLinks {
		if err = env.DB.CreateCompanyUser(CU); err != nil {
			panic(err)
		}
	}

	// Inserting fake User-Function links
	for _, UF := range fakeUserFunctionLinks {
		if err = env.DB.CreateUserFunction(UF); err != nil {
			panic(err)
		}
	}

	// Inserting fake schedules
	for index, schedule := range fakeSchedules {
		if id, err = env.DB.CreateSchedule(schedule); err != nil {
			panic(err)
		}
		fakeSchedules[index].ScheduleId = id
	}

	// Inserting fake User-Schedule links
	for _, US := range fakeUserScheduleLinks {
		if err = env.DB.CreateUserSchedule(US); err != nil {
			panic(err)
		}
	}

	// Inserting fake comments
	for index, comment := range fakeComments {
		if id, err = env.DB.CreateComment(comment); err != nil {
			panic(err)
		}
		fakeComments[index].CommentId = id
	}

	//
	//	At last, fetching the 3 roles
	//
	if allRoles, err = env.DB.GetRoles(); err != nil {
		panic(err)
	}
}
