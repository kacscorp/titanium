package query

import (
	goqu "github.com/doug-martin/goqu/v9"
)

var (
	columnAge                  = goqu.I("age")
	columnFirstName            = goqu.I("first_name")
	columnIdentificationNumber = goqu.I("identification_number")
	columnID                   = goqu.I("id")
	columnLastName             = goqu.I("last_name")

	tableEmployees = goqu.I("employees")
)
