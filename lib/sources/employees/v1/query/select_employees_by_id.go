package query

import (
	"errors"

	goqu "github.com/doug-martin/goqu/v9"
)

// SelectEmployeesByID ...
type SelectEmployeesByID struct {
	Query string
	Args  []interface{}
}

// NewSelectEmployeesByID ...
func NewSelectEmployeesByID(param int64) (
	*SelectEmployeesByID,
	error,
) {

	dataset := goqu.From(tableEmployees).Prepared(true).
		SelectDistinct(
			columnID,
			columnAge,
			columnIdentificationNumber,
			columnFirstName,
			columnLastName,
		).Where(
		columnID.Eq(param),
	)

	if dataset == nil {
		return nil, errors.New("dataset")
	}
	template, args, _ := dataset.Prepared(true).ToSQL()
	return &SelectEmployeesByID{Query: template, Args: args}, nil
}
