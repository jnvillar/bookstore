package permissions

import "bookstore/roles"

type Action = string

const (
	WriteBook Action = "writebook"
)

var actions = map[string]Action{

}

var permissions = map[roles.Role][]Action{
	roles.RoleAdmin: {
		WriteBook,
	},
}

func HasPermissions(role roles.Role, action Action) bool {
	allowedActions, found := permissions[role]
	if !found {
		return false
	}
	for _, allowedAction := range allowedActions {
		if allowedAction == action {
			return true
		}
	}
	return false
}
