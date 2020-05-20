package queries

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
	"github.com/hwangm/isthemunibusy-go/types"
)

// GetUsersQuery returns a list of Users
func GetUsersQuery() *graphql.Field {
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

// GetUserQuery returns a single user by id
func GetUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			userID := params.Args["id"].(int)
			user := types.User{ID: userID}
			err := dal.DB.Select(&user)
			if err != nil {
				fmt.Printf("Error retrieving user by id: %v", err)
				return nil, err
			}
			return user, nil
		},
	}
}

// GetUserFeatureTestVariantsQuery returns all user feature test variants
func GetUserFeatureTestVariantsQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserFeatureTestVariantType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var userTestVariants []types.UserToFeatureUserTestVariant
			err := dal.DB.Model(&userTestVariants).Select()
			if err != nil {
				fmt.Printf("Error retrieving all user feature test variants: %v", err)
				return nil, err
			}

			return userTestVariants, nil
		},
	}
}
