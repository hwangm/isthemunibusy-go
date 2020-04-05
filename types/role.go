package types

import "github.com/graphql-go/graphql"

// Role type definition
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// RoleType is the GraphQL schema for the role type
var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.Int},
		"name": &graphql.Field{Type: graphql.String},
	},
})
