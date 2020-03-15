package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/types"
)

// GetUserQuery returns User field information
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var users []types.User
			userOne := types.User{
				ID:        1,
				Firstname: "Matt",
				Lastname:  "Hwang",
			}
			users = append(users, userOne)
			return users, nil
		},
	}
}
