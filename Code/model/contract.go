package model

import (
	"strconv"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Contract represents a type of contract.
/*	ContractName : Alternance/Stage/CDI/CDD...
 */
type Contract struct {
	ContractId   int64  `db:"contract_id" json:"contract_id"`
	ContractName string `db:"contract_name" json:"contract_name"`
}

type Contracts []Contract

func (contract *Contract) String() string {
	return "Contract id : " + strconv.FormatInt(contract.ContractId, 10) + "/ Contract name : " + contract.ContractName
}
