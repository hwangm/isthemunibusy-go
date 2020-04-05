package queries

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
	"github.com/hwangm/isthemunibusy-go/types"
)

// GetUserQuery returns User field information
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var users []types.User
			err := dal.DB.Model(&users).Select()
			if err != nil {
				fmt.Printf("Error retrieving all Users: %v", err)
				return nil, err
			}
			return users, nil
		},
	}
}
