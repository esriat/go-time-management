package datastores

import (
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

type IDatastore interface {
	CloseDatabase()

	// Users
	GetUsers() (model.Users, error)
	GetUser(UserId int64) (model.User, error)
	GetUserFromEmail(Email string) (model.User, error)
	GetUsersOfCompany(CompanyId int64) (model.Users, error)
	GetUsersOfProject(ProjectId int64) (model.Users, error)
	GetUsersOfSchedule(ScheduleId int64) (model.Users, error)
	CreateUser(User model.User) (int64, error)
	DeleteUser(UserId int64) error
	UpdateUser(User model.User) (model.User, error)

	//Companies
	GetCompanies() (model.Companies, error)
	GetCompany(CompanyId int64) (model.Company, error)
	CreateCompany(Company model.Company) (int64, error)
	DeleteCompany(CompanyId int64) error
	UpdateCompany(Company model.Company) (model.Company, error)

	//Projects
	GetProjects() (model.Projects, error)
	GetProject(ProjectId int64) (model.Project, error)
	GetProjectsOfCompany(CompanyId int64) (model.Projects, error)
	GetProjectsOfUser(UserId int64) (model.Projects, error)
	GetVacationProject() (model.Project, error)
	CreateProject(Project model.Project) (int64, error)
	DeleteProject(ProjectId int64) error
	UpdateProject(Project model.Project) (model.Project, error)

	//Comments
	GetComments() (model.Comments, error)
	GetComment(CommentId int64) (model.Comment, error)
	GetCommentsOfUser(UserId int64) (model.Comments, error)
	GetCommentsOfSchedule(ScheduleId int64) (model.Comments, error)
	GetCommentsOfProject(ProjectId int64) (model.Comments, error)
	CreateComment(Comment model.Comment) (int64, error)
	DeleteComment(CommentId int64) error
	UpdateComment(Comment model.Comment) (model.Comment, error)

	//Vacations
	GetVacationsOfUser(UserId int64) (model.Schedules, error)
	GetVacation(VacationId int64) (model.Schedule, error)
	CreateVacation(Schedule model.Schedule) (int64, error)
	DeleteVacation(VacationId int64) error
	UpdateVacation(Vacation model.Schedule) (model.Schedule, error)

	//Schedules
	GetSchedule(ScheduleId int64) (model.Schedule, error)
	GetSchedulesOfUser(UserId int64) (model.Schedules, error)
	GetSchedulesOfProject(ProjectId int64) (model.Schedules, error)
	CreateSchedule(Schedule model.Schedule) (int64, error)
	DeleteSchedule(ScheduleId int64) error
	UpdateSchedule(Schedule model.Schedule) (model.Schedule, error)

	//Roles
	GetRoles() (model.Roles, error)
	GetRole(RoleId int64) (model.Role, error)
	GetRoleOfUser(UserId int64) (model.Role, error)
	GetRoleByName(RoleName string) (model.Role, error)
	CreateRole(Role model.Role) (int64, error)
	DeleteRole(RoleId int64) error
	UpdateRole(Role model.Role) (model.Role, error)

	//Contracts
	GetContracts() (model.Contracts, error)
	GetContract(ContractId int64) (model.Contract, error)
	GetContractOfUser(UserId int64) (model.Contract, error)
	CreateContract(Contract model.Contract) (int64, error)
	DeleteContract(ContractId int64) error
	UpdateContract(Contract model.Contract) (model.Contract, error)

	//Function
	GetFunctions() (model.Functions, error)
	GetFunction(FunctionId int64) (model.Function, error)
	GetFunctionsOfUser(UserId int64) (model.Functions, error)
	CreateFunction(Function model.Function) (int64, error)
	DeleteFunction(FunctionId int64) error
	UpdateFunction(Function model.Function) (model.Function, error)

	//Intermediate tables
	CreateCompanyProject(CP model.CompanyProject) error
	CreateCompanyUser(CU model.CompanyUser) error
	CreateUserSchedule(US model.UserSchedule) error
	CreateUserFunction(UF model.UserFunction) error

	DeleteCompanyProject(CP model.CompanyProject) error
	DeleteCompanyUser(CU model.CompanyUser) error
	DeleteUserSchedule(US model.UserSchedule) error
	DeleteUserFunction(UF model.UserFunction) error
}
