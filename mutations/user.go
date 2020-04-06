package mutations

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
	"github.com/hwangm/isthemunibusy-go/types"
)

// GetCreateUserMutation creates a new user based on params
func GetCreateUserMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"firstname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"roleIds": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var user *types.User
			err := dal.DB.RunInTransaction(func(tx *pg.Tx) error {
				newUser := types.User{
					Firstname: params.Args["firstname"].(string),
					Lastname:  params.Args["lastname"].(string),
				}

				user = &newUser

				err := tx.Insert(user)
				if err != nil {
					fmt.Printf("Error creating new user: %v", err)
					return err
				}

				roleIds := params.Args["roleIds"].([]interface{})

				if roleIds != nil {
					ids := make([]int, len(roleIds))
					for i, v := range roleIds {
						ids[i] = v.(int)
					}
					for _, roleID := range ids {
						userToRole := types.UserToRole{UserID: user.ID, RoleID: roleID}
						err := tx.Insert(&userToRole)
						if err != nil {
							fmt.Printf("Error creating new user roles: %v", err)
							return err
						}
					}
				}

				return nil
			})

			if err != nil {
				fmt.Printf("Error creating new user: %v", err)
				return nil, err
			}

			err = dal.DB.Model(user).Where("id = ?", user.ID).Relation("Roles").First()
			if err != nil {
				fmt.Printf("Error retrieving new user roles after creating roles: %v", err)
				return nil, err
			}
			return *user, nil
		},
	}
}

// GetDeleteUserByIDMutation deletes a user by id
func GetDeleteUserByIDMutation() *graphql.Field {
	return &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			userID := params.Args["id"].(int)
			user := types.User{ID: userID}
			err := dal.DB.Delete(&user)
			if err != nil {
				fmt.Printf("Error deleting user: %v", err)
				return false, err
			}

			return true, nil

		},
	}
}
