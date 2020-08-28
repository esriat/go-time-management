package handlers

import (
	"github.com/gorilla/mux"
	_ "github.com/gorilla/schema"
	"github.com/justinas/alice"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func HandleRoutes(r *mux.Router, env *Env) {
	commonChain := alice.New(env.HeadersMiddleware)
	secureChain := alice.New(env.HeadersMiddleware, env.AuthenticateMiddleware, env.AuthorizeMiddleware)

	//
	// Routing login
	//
	r.Handle("/get-token", commonChain.Then(env.AppMiddleware(env.GetTokenHandler))).Methods("POST")

	//
	// Routing comments
	//
	r.Handle("/{item:comments}", secureChain.Then(env.AppMiddleware(env.GetCommentsHandler))).Methods("GET")
	r.Handle("/{item:comments}/{id}", secureChain.Then(env.AppMiddleware(env.GetCommentHandler))).Methods("GET")
	r.Handle("/{item:users}/{id}/{goal:comments}", secureChain.Then(env.AppMiddleware(env.GetCommentsOfUserHandler))).Methods("GET")
	r.Handle("/{item:schedules}/{id}/{goal:comments}", secureChain.Then(env.AppMiddleware(env.GetCommentsOfScheduleHandler))).Methods("GET")
	r.Handle("/{item:projects}/{id}/{goal:comments}", secureChain.Then(env.AppMiddleware(env.GetCommentsOfProjectHandler))).Methods("GET")
	r.Handle("/{item:comments}", secureChain.Then(env.AppMiddleware(env.CreateCommentHandler))).Methods("POST")
	r.Handle("/{item:comments}/{id}", secureChain.Then(env.AppMiddleware(env.UpdateCommentHandler))).Methods("PATCH")
	r.Handle("/{item:comments}/{id}", secureChain.Then(env.AppMiddleware(env.DeleteCommentHandler))).Methods("DELETE")

	//
	// Routing companies
	//
	r.Handle("/{item:companies}", secureChain.Then(env.AppMiddleware(env.GetCompaniesHandler))).Methods("GET")
	r.Handle("/{item:companies}/{id}", secureChain.Then(env.AppMiddleware(env.GetCompanyHandler))).Methods("GET")
	r.Handle("/{item:companies}", secureChain.Then(env.AppMiddleware(env.CreateCompanyHandler))).Methods("POST")
	r.Handle("/{item:companies}/{id}", secureChain.Then(env.AppMiddleware(env.UpdateCompanyHandler))).Methods("PATCH")
	r.Handle("/{item:companies}/{id}", secureChain.Then(env.AppMiddleware(env.DeleteCompanyHandler))).Methods("DELETE")

	//
	// Routing contracts
	//
	r.Handle("/{item:contracts}", secureChain.Then(env.AppMiddleware(env.GetContractsHandler))).Methods("GET")
	r.Handle("/{item:contracts}/{id}", secureChain.Then(env.AppMiddleware(env.GetContractHandler))).Methods("GET")
	r.Handle("/{item:users}/{id}/{goal:contract}", secureChain.Then(env.AppMiddleware(env.GetContractOfUserHandler))).Methods("GET")
	r.Handle("/{item:contracts}", secureChain.Then(env.AppMiddleware(env.CreateContractHandler))).Methods("POST")
	r.Handle("/{item:contracts}/{id}", secureChain.Then(env.AppMiddleware(env.UpdateContractHandler))).Methods("PATCH")
	r.Handle("/{item:contracts}/{id}", secureChain.Then(env.AppMiddleware(env.DeleteContractHandler))).Methods("DELETE")

	//
	// Routing functions
	//
	r.Handle("/{item:functions}", secureChain.Then(env.AppMiddleware(env.GetFunctionsHandler))).Methods("GET")
	r.Handle("/{item:functions}/{id}", secureChain.Then(env.AppMiddleware(env.GetFunctionHandler))).Methods("GET")
	r.Handle("/{item:users}/{id}/{goal:functions}", secureChain.Then(env.AppMiddleware(env.GetFunctionsOfUserHandler))).Methods("GET")
	r.Handle("/{item:functions}", secureChain.Then(env.AppMiddleware(env.CreateFunctionHandler))).Methods("POST")
	r.Handle("/{item:functions}/{id}", secureChain.Then(env.AppMiddleware(env.UpdateFunctionHandler))).Methods("PATCH")
	r.Handle("/{item:functions}/{id}", secureChain.Then(env.AppMiddleware(env.DeleteFunctionHandler))).Methods("DELETE")

	//
	// Routing projects
	//
	r.Handle("/{item:projects}", secureChain.Then(env.AppMiddleware(env.GetProjectsHandler))).Methods("GET")
	r.Handle("/{item:projects}/{id}", secureChain.Then(env.AppMiddleware(env.GetProjectHandler))).Methods("GET")
	r.Handle("/{item:companies}/{id}/{goal:projects}", secureChain.Then(env.AppMiddleware(env.GetProjectsOfCompanyHandler))).Methods("GET")
	r.Handle("/{item:users}/{id}/{goal:projects}", secureChain.Then(env.AppMiddleware(env.GetProjectsOfUserHandler))).Methods("GET")
	r.Handle("/{item:projects}", secureChain.Then(env.AppMiddleware(env.CreateProjectHandler))).Methods("POST")
	r.Handle("/{item:projects}/{id}", secureChain.Then(env.AppMiddleware(env.UpdateProjectHandler))).Methods("PATCH")
	r.Handle("/{item:projects}/{id}", secureChain.Then(env.AppMiddleware(env.DeleteProjectHandler))).Methods("DELETE")

	//
	// Routing roles
	//
	r.Handle("/{item:roles}", secureChain.Then(env.AppMiddleware(env.GetRolesHandler))).Methods("GET")
	r.Handle("/{item:roles}/{id}", secureChain.Then(env.AppMiddleware(env.GetRoleHandler))).Methods("GET")
	r.Handle("/{item:users}/{id}/{goal:role}", secureChain.Then(env.AppMiddleware(env.GetRoleOfUserHandler))).Methods("GET")
	r.Handle("/roles", secureChain.Then(env.AppMiddleware(env.CreateRoleHandler))).Methods("POST")
	r.Handle("/roles/{id}", secureChain.Then(env.AppMiddleware(env.UpdateRoleHandler))).Methods("PATCH")
	r.Handle("/roles/{id}", secureChain.Then(env.AppMiddleware(env.DeleteRoleHandler))).Methods("DELETE")

	//
	// Routing schedules
	//
	r.Handle("/{item:schedules}/{id}", secureChain.Then(env.AppMiddleware(env.GetScheduleHandler))).Methods("GET")
	r.Handle("/{item:users}/{id}/{goal:schedules}", secureChain.Then(env.AppMiddleware(env.GetSchedulesOfUserHandler))).Methods("GET")
	r.Handle("/{item:projects}/{id}/{goal:schedules}", secureChain.Then(env.AppMiddleware(env.GetSchedulesOfProjectHandler))).Methods("GET")
	r.Handle("/{item:schedules}", secureChain.Then(env.AppMiddleware(env.CreateScheduleHandler))).Methods("POST")
	r.Handle("/{item:schedules}/{id}", secureChain.Then(env.AppMiddleware(env.UpdateScheduleHandler))).Methods("PATCH")
	r.Handle("/{item:schedules}/{id}", secureChain.Then(env.AppMiddleware(env.DeleteScheduleHandler))).Methods("DELETE")

	//
	// Routing users
	//
	r.Handle("/{item:users}", secureChain.Then(env.AppMiddleware(env.GetUsersHandler))).Methods("GET")
	r.Handle("/{item:users}/{id}", secureChain.Then(env.AppMiddleware(env.GetUserHandler))).Methods("GET")
	r.Handle("/{item:companies}/{id}/{goal:users}", secureChain.Then(env.AppMiddleware(env.GetUsersOfCompanyHandler))).Methods("GET")
	r.Handle("/{item:schedules}/{id}/{goal:users}", secureChain.Then(env.AppMiddleware(env.GetUsersOfScheduleHandler))).Methods("GET")
	r.Handle("/{item:projects}/{id}/{goal:users}", secureChain.Then(env.AppMiddleware(env.GetUsersOfProjectHandler))).Methods("GET")
	r.Handle("/{item:users}", secureChain.Then(env.AppMiddleware(env.CreateUserHandler))).Methods("POST")
	r.Handle("/{item:users}/{id}", secureChain.Then(env.AppMiddleware(env.UpdateUserHandler))).Methods("PATCH")
	r.Handle("/{item:users}/{id}", secureChain.Then(env.AppMiddleware(env.DeleteUserHandler))).Methods("DELETE")

	//
	// Routing vacations
	//
	r.Handle("/{item:users}/{id}/{goal:vacations}", secureChain.Then(env.AppMiddleware(env.GetVacationsOfUserHandler))).Methods("GET")
	r.Handle("/{item:vacations}/{id}", secureChain.Then(env.AppMiddleware(env.GetVacationHandler))).Methods("GET")
	r.Handle("/{item:vacations}", secureChain.Then(env.AppMiddleware(env.CreateVacationHandler))).Methods("POST")
	r.Handle("/{item:vacations}/{id}", secureChain.Then(env.AppMiddleware(env.UpdateVacationHandler))).Methods("PATCH")
	r.Handle("/{item:vacations}/{id}", secureChain.Then(env.AppMiddleware(env.DeleteVacationHandler))).Methods("DELETE")

	//
	// Routing intermediate tables
	//
	r.Handle("/{item:companies}/{id}/{other_item:users}/{other_id}", secureChain.Then(env.AppMiddleware(env.CreateCompanyUserHandler))).Methods("POST")
	r.Handle("/{item:companies}/{id}/{other_item:users}/{other_id}", secureChain.Then(env.AppMiddleware(env.DeleteCompanyUserHandler))).Methods("DELETE")
	r.Handle("/{item:users}/{id}/{other_item:schedules}/{other_id}", secureChain.Then(env.AppMiddleware(env.CreateUserScheduleHandler))).Methods("POST")
	r.Handle("/{item:users}/{id}/{other_item:schedules}/{other_id}", secureChain.Then(env.AppMiddleware(env.DeleteUserScheduleHandler))).Methods("DELETE")
	r.Handle("/{item:companies}/{id}/{other_item:projects}/{other_id}", secureChain.Then(env.AppMiddleware(env.CreateCompanyProjectHandler))).Methods("POST")
	r.Handle("/{item:companies}/{id}/{other_item:projects}/{other_id}", secureChain.Then(env.AppMiddleware(env.DeleteCompanyProjectHandler))).Methods("DELETE")
	r.Handle("/{item:users}/{id}/{other_item:functions}/{other_id}", secureChain.Then(env.AppMiddleware(env.CreateUserFunctionHandler))).Methods("POST")
	r.Handle("/{item:users}/{id}/{other_item:functions}/{other_id}", secureChain.Then(env.AppMiddleware(env.DeleteUserFunctionHandler))).Methods("DELETE")
}
