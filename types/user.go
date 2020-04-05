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
	Roles     []Role `pg:"many2many:user_roles"`
}

// UserToRole type definition for many2many table joining User and Roles
type UserToRole struct {
	UserId int
	RoleId int
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
				user := &User{ID: userID}
				err := dal.DB.Model(user).Relation("Roles").First()
				if err != nil {
					fmt.Printf("Error retrieving user roles: %v", err)
				}
				return user.Roles, nil
			},
		},
	},
})
