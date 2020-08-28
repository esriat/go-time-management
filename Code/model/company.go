package model

import (
	"strconv"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Company : Represents a company.
/*	CompanyName : The name of the company : Biomarqueurs/Biopass...
 */
type Company struct {
	CompanyId   int64  `db:"company_id" json:"company_id"`
	CompanyName string `db:"company_name" json:"company_name"`
}

type Companies []Company

func (company *Company) String() string {
	return "Company id : " + strconv.FormatInt(company.CompanyId, 10) + "/ Company name : " + company.CompanyName
}
