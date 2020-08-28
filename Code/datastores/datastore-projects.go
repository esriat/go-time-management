package datastores

import (
	"database/sql"

	"gitlab.iut-clermont.uca.fr/esriat/gestion-tps-projet/Code/model"
)

//  GetProjects() (model.Projects, error)
/*  This method is used to get the list of all projects
    Returns the list of all projects or an error
*/
func (db *ConcreteDatastore) GetProjects() (model.Projects, error) {
	// Setting up and executing the request
	rows, err := db.Queryx("SELECT * FROM Project;")
	if err != nil {
		return nil, err
	}

	// Formatting the data
	projectsList := model.Projects{}

	for rows.Next() {
		project := model.Project{}
		err := rows.StructScan(&project)
		if err != nil {
			return nil, err
		}
		projectsList = append(projectsList, project)
	}

	defer rows.Close()
	return projectsList, nil
}

//  GetProject(ProjectId int64) (model.Project, error)
/*  This method is used to get a project
    Returns the wanted project or an error
*/
func (db *ConcreteDatastore) GetProject(ProjectId int64) (model.Project, error) {
	var (
		err     error
		project model.Project
	)

	// Setting up and executing the request
	request := `SELECT * FROM Project WHERE project_id=?`
	if err = db.Get(&project, request, ProjectId); err != nil {
		return model.Project{}, err
	}

	return project, nil
}

//  GetProjectsOfCompany(CompanyId int64) (model.Projects, error)
/*  This method is used to get the list of the projects of a company
    Returns the list of the projects of the wanted company or an error
*/
func (db *ConcreteDatastore) GetProjectsOfCompany(CompanyId int64) (model.Projects, error) {
	// Setting up and executing the request
	request := `SELECT *
	FROM Project, CompanyProject
	WHERE Project.project_id = CompanyProject.project_id
	AND CompanyProject.company_id=?`
	rows, err := db.Queryx(request, CompanyId)
	if err != nil {
		return nil, err
	}

	// Formatting the data
	projectsList := model.Projects{}

	for rows.Next() {
		project := model.Project{}
		err := rows.StructScan(&project)
		if err != nil {
			return nil, err
		}
		projectsList = append(projectsList, project)
	}

	defer rows.Close()
	return projectsList, nil
}

//  GetProjectsOfUser(UserId int64) (model.Projects, error)
/*  This method is used to get the list of the projects of a user
    Returns the list of the projects of the wanted user or an error
*/
func (db *ConcreteDatastore) GetProjectsOfUser(UserId int64) (model.Projects, error) {
	// Setting up and executing the request
	request := `SELECT DISTINCT Project.project_id, project.project_name
	FROM Project, Schedule, UserSchedule
	WHERE Project.project_id = Schedule.project_id
	AND Schedule.schedule_id=UserSchedule.schedule_id
	AND UserSchedule.user_id=?`
	rows, err := db.Queryx(request, UserId)
	if err != nil {
		return nil, err
	}

	// Formatting the data
	projectsList := model.Projects{}

	for rows.Next() {
		project := model.Project{}
		err := rows.StructScan(&project)
		if err != nil {
			return nil, err
		}
		projectsList = append(projectsList, project)
	}

	defer rows.Close()
	return projectsList, nil
}

//  CreateProject(Project model.Project) (int64, error)
/*	This method is used to create a new project
	Returns the id of the new project or an error
*/
func (db *ConcreteDatastore) CreateProject(Project model.Project) (int64, error) {
	var (
		tx        *sql.Tx
		err       error
		res       sql.Result
		projectId int64
	)

	// Preparing the request
	if tx, err = db.Begin(); err != nil {
		return -1, err
	}

	// Setting up and executing the request
	request := `INSERT INTO Project(project_name) VALUES (?)`
	if res, err = tx.Exec(request, Project.ProjectName); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Getting the id of the new element
	if projectId, err = res.LastInsertId(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return -1, errr
		}
		return -1, err
	}

	return projectId, nil
}

//  DeleteProject(ProjectId int64) error
/*	This method is used to delete a new project.
	Can return an error
*/
func (db *ConcreteDatastore) DeleteProject(ProjectId int64) error {
	request := `DELETE FROM Project 
	WHERE project_id=?`
	if _, err := db.Exec(request, ProjectId); err != nil {
		return err
	}
	return nil
}

//  UpdateProject(Project model.Project) (model.Project, error)
/*	This method is used to update an existing project.
	Returns the modified project or an error
*/
func (db *ConcreteDatastore) UpdateProject(Project model.Project) (model.Project, error) {
	var (
		tx  *sql.Tx
		err error
	)

	// Preparing
	if tx, err = db.Begin(); err != nil {
		return model.Project{}, err
	}

	// Executing the request
	request := `UPDATE Project 
	SET project_name=?
	WHERE project_id=?`
	if _, err = tx.Exec(request, Project.ProjectName, Project.ProjectId); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Project{}, errr
		}
		return model.Project{}, err
	}

	// Saving
	if err = tx.Commit(); err != nil {
		if errr := tx.Rollback(); errr != nil {
			return model.Project{}, errr
		}
		return model.Project{}, err
	}

	return Project, nil
}

//  GetVacationProject() (model.Project, error)
/*	This method is used to get the Vacations project.
 */
func (db *ConcreteDatastore) GetVacationProject() (model.Project, error) {
	var (
		err     error
		project model.Project
	)

	request := `SELECT * FROM Project WHERE project_name="Vacation"`
	if err = db.Get(&project, request); err != nil {
		return model.Project{}, err
	}

	return project, nil
}
