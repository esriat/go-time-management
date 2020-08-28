package model

import (
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Project : Represents a company.
/*	ProjectName : The name of the project : Biorcell 3D/Lightspot...
 */
type Project struct {
	ProjectId   int64  `db:"project_id" json:"project_id"`
	ProjectName string `db:"project_name" json:"project_name"`
}

type Projects []Project
