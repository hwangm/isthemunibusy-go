package mutations

import (
	"github.com/graphql-go/graphql"
)

// GetRootFields returns available mutation fields
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"createUser": GetCreateUserMutation(),
	}
}
