package query

import (
	goqu "github.com/doug-martin/goqu/v9"
)

var (
	columnUserID     = goqu.I("user_id")
	columnUserName   = goqu.I("username")
	columnCreatedAt  = goqu.I("created_at")
	columnIsActive   = goqu.I("is_active")
	columnUserRoleID = goqu.I("user_role_id")

	tableUser = goqu.I("user")
)
