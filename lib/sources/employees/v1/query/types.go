package query

import (
	"database/sql"
)

// Employees ...
type Employees struct {
	ID                   int
	Age                  sql.NullInt64
	IdentificationNumber sql.NullString
	FirstName            sql.NullString
	LastName             sql.NullString
}
