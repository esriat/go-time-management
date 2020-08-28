package datastores

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
	"golang.org/x/crypto/bcrypt"
)

// This variable contains the table creation commands.
var schema = `
PRAGMA foreign_keys = ON;
PRAGMA encoding = "UTF-8"; 
PRAGMA temp_store = 2;
PRAGMA journal_mode = WAL;
PRAGMA temp_store = MEMORY;

DROP TABLE IF EXISTS UserFunction;
DROP TABLE IF EXISTS UserSchedule;
DROP TABLE IF EXISTS CompanyUser;
DROP TABLE IF EXISTS CompanyProject;
DROP TABLE IF EXISTS Comment;
DROP TABLE IF EXISTS Schedule;
DROP TABLE IF EXISTS User;
DROP TABLE IF EXISTS Role;
DROP TABLE IF EXISTS Project;
DROP TABLE IF EXISTS Company;
DROP TABLE IF EXISTS Function;
DROP TABLE IF EXISTS Contract;

CREATE TABLE IF NOT EXISTS Contract (
    contract_id integer PRIMARY KEY AUTOINCREMENT,
	contract_name text NOT NULL
);

CREATE TABLE IF NOT EXISTS Function (
    function_id integer PRIMARY KEY AUTOINCREMENT,
    function_name text NOT NULL
);

CREATE TABLE IF NOT EXISTS Company (
    company_id integer PRIMARY KEY AUTOINCREMENT,
    company_name text NOT NULL
);

CREATE TABLE IF NOT EXISTS Project (
    project_id integer PRIMARY KEY AUTOINCREMENT,
    project_name text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Role (
    role_id integer PRIMARY KEY AUTOINCREMENT,
    role_name text NOT NULL UNIQUE,  
    can_add_and_modify_users bool NOT NULL,
    can_see_other_schedules bool NOT NULL,
    can_add_projects bool NOT NULL,
    can_see_reports bool NOT NULL
);

CREATE TABLE IF NOT EXISTS User (
    user_id integer PRIMARY KEY AUTOINCREMENT,
    contract_id integer NOT NULL,
    role_id integer NOT NULL,
    username text NOT NULL,
    password text NOT NULL,
    last_name text NOT NULL,
    first_name text NOT NULL,
    mail text NOT NULL UNIQUE,
    theorical_hours_worked integer NOT NULL,
    vacation_hours integer NOT NULL,
    CONSTRAINT FK_User_Contract FOREIGN KEY (contract_id) REFERENCES Contract(contract_id),
    CONSTRAINT FK_User_Role FOREIGN KEY (role_id) REFERENCES Role(role_id)
);

CREATE TABLE IF NOT EXISTS Schedule (
    schedule_id integer PRIMARY KEY AUTOINCREMENT,
    project_id integer NOT NULL,
    start_date datetime NOT NULL,
    end_date datetime NOT NULL,
    CONSTRAINT FK_Schedule_Function FOREIGN KEY (project_id) REFERENCES Project(project_id)
);

CREATE TABLE IF NOT EXISTS Comment (
	comment_id integer PRIMARY KEY AUTOINCREMENT,
	schedule_id integer NOT NULL,
    comment text NOT NULL,
    is_important bool NOT NULL,
    CONSTRAINT FK_Comment_Shedule FOREIGN KEY (schedule_id) REFERENCES Schedule(schedule_id)
);

CREATE TABLE IF NOT EXISTS CompanyProject (
    company_id integer,
    project_id integer,
    CONSTRAINT FK_CP_Company FOREIGN KEY (company_id) REFERENCES Company(company_id),
    CONSTRAINT FK_CP_Project FOREIGN KEY (project_id) REFERENCES Project(project_id),
    CONSTRAINT PK_CompanyProject PRIMARY KEY (company_id, project_id)
);

CREATE TABLE IF NOT EXISTS CompanyUser (
    company_id integer,
    user_id integer,
    CONSTRAINT FK_CU_Company FOREIGN KEY (company_id) REFERENCES Company(company_id),
    CONSTRAINT FK_CU_User FOREIGN KEY (user_id) REFERENCES User(user_id),
    CONSTRAINT PK_CompanyUser PRIMARY KEY (company_id, user_id)
);

CREATE TABLE IF NOT EXISTS UserSchedule (
    user_id integer,
    schedule_id integer,
    CONSTRAINT FK_CS_User FOREIGN KEY (user_id) REFERENCES User(user_id),
    CONSTRAINT FK_CS_Schedule FOREIGN KEY (schedule_id) REFERENCES Schedule(schedule_id),
    CONSTRAINT PK_UserSchedule PRIMARY KEY (user_id, schedule_id)
);

CREATE TABLE IF NOT EXISTS UserFunction (
    user_id integer,
    function_id integer,
    CONSTRAINT FK_UF_User FOREIGN KEY (user_id) REFERENCES User(user_id),
    CONSTRAINT FK_UF_Function FOREIGN KEY (function_id) REFERENCES Function(function_id),
    CONSTRAINT PK_UserFunction PRIMARY KEY (user_id, function_id)
);
`

type ConcreteDatastore struct {
	*sqlx.DB
}

// This variable contains a link to the database
// var myDatabase *ConcreteDatastore

func NewDatabase(databaseName string) (*ConcreteDatastore, error) {
	var (
		db              *sqlx.DB
		err             error
		datastore       *ConcreteDatastore
		adminRoleId     int64
		adminContractId int64
		cryptedPassword []byte
	)

	if db, err = sqlx.Open("sqlite3", databaseName); err != nil {
		return nil, err
	}

	db.MustExec(schema)
	db = db.Unsafe()

	datastore = &ConcreteDatastore{db}

	// Creating the vacation project, as it is necessary
	if _, err = datastore.CreateProject(model.Project{
		ProjectName: "Vacation",
	}); err != nil {
		return nil, err
	}

	// Then, creating the 3 basic roles
	if adminRoleId, err = datastore.CreateRole(model.Role{
		RoleName:             "Superadmin",
		CanAddAndModifyUsers: true,
		CanSeeOtherSchedules: true,
		CanAddProjects:       true,
		CanSeeReports:        true,
	}); err != nil {
		return nil, err
	}

	if _, err = datastore.CreateRole(model.Role{
		RoleName:             "Admin",
		CanAddAndModifyUsers: false,
		CanSeeOtherSchedules: false,
		CanAddProjects:       true,
		CanSeeReports:        true,
	}); err != nil {
		return nil, err
	}

	if _, err = datastore.CreateRole(model.Role{
		RoleName:             "User",
		CanAddAndModifyUsers: false,
		CanSeeOtherSchedules: false,
		CanAddProjects:       false,
		CanSeeReports:        false,
	}); err != nil {
		return nil, err
	}

	// Creating a "default" user with all permissions.. Otherwise, we can't do anything
	// For that, we need a contract. This contract will only be used for this user.
	if adminContractId, err = datastore.CreateContract(model.Contract{
		ContractName: "Admin",
	}); err != nil {
		return nil, err
	}

	if cryptedPassword, err = bcrypt.GenerateFromPassword([]byte("Admin"), bcrypt.DefaultCost); err != nil {
		return nil, err
	}

	if _, err = datastore.CreateUser(model.User{
		ContractId: adminContractId,
		RoleId:     adminRoleId,
		Username:   "Admin",
		Mail:       "admin@mydb",
		Password:   string(cryptedPassword),
	}); err != nil {
		return nil, err
	}

	return datastore, nil
}

func (db *ConcreteDatastore) CloseDatabase() {
	defer db.Close()
}
