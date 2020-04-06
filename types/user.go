package types

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
)

// User type definition
type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Roles     []Role `pg:"many2many:user_roles, joinFK:role_id"`
}

// UserToRole type definition - join table for users and roles
type UserToRole struct {
	tableName struct{} `pg:"user_roles"`
	UserID    int
	RoleID    int
}

// UserType is GraphQL schema for the user type
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"firstname": &graphql.Field{Type: graphql.String},
		"lastname":  &graphql.Field{Type: graphql.String},
		"roles": &graphql.Field{
			Type: graphql.NewList(RoleType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				userID := params.Source.(User).ID
				user := new(User)
				err := dal.DB.Model(user).Where("id = ?", userID).Relation("Roles").First()
				if err != nil {
					fmt.Printf("Error retrieving user roles: %v", err)
					return nil, err
				}
				// fmt.Printf("User roles for user id %v are: %v", userID, user.Roles)
				return user.Roles, nil
			},
		},
	},
})
