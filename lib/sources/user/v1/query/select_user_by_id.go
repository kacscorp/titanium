package query

import (
	"errors"

	goqu "github.com/doug-martin/goqu/v9"
)

// SelectUserByID ...
type SelectUserByID struct {
	Query string
	Args  []interface{}
}

// NewSelectUserByID ...
func NewSelectUserByID(param int64) (
	*SelectUserByID,
	error,
) {

	dataset := goqu.From(tableUser).Prepared(true).
		SelectDistinct(
			columnUserID,
			columnUserName,
			columnCreatedAt,
			columnIsActive,
			columnUserRoleID,
		).Where(
		columnUserID.Eq(param),
	)

	if dataset == nil {
		return nil, errors.New("dataset")
	}
	template, args, _ := dataset.Prepared(true).ToSQL()
	return &SelectUserByID{Query: template, Args: args}, nil
}
