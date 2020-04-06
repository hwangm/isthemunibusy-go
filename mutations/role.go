package mutations

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
	"github.com/hwangm/isthemunibusy-go/types"
)

// GetCreateRoleMutation creates a new role with name
func GetCreateRoleMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.RoleType,
		Args: graphql.FieldConfigArgument{
			"rolename": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			name := params.Args["rolename"].(string)
			role := types.Role{Name: name}
			err := dal.DB.Insert(&role)
			if err != nil {
				fmt.Printf("Error creating new role: %v", err)
				return nil, err
			}

			return role, nil
		},
	}
}
