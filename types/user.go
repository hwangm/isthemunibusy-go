package types

import (
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
)

// User type definition
type User struct {
	ID                  int                  `json:"id"`
	Firstname           string               `json:"firstname"`
	Lastname            string               `json:"lastname"`
	Roles               []Role               `pg:"many2many:user_roles,joinFK:role_id"`
	FeatureTestVariants []FeatureTestVariant `pg:"many2many:feature_usertestvariants,joinFK:feature_testvariant_id"`
	DeletedAt           time.Time            `pg:",soft_delete"`
}

// UserToRole type definition - join table for users and roles
type UserToRole struct {
	tableName struct{} `pg:"user_roles"`
	UserID    int
	RoleID    int
}

// UserToFeatureUserTestVariant - join table for users and feature test variants
type UserToFeatureUserTestVariant struct {
	tableName            struct{} `pg:"feature_usertestvariants"`
	UserID               int
	FeatureTestVariantID int       `pg:feature_testvariant_id`
	DeletedAt            time.Time `pg:",soft_delete"`
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
		"featureTestVariants": &graphql.Field{
			Type: graphql.NewList(FeatureTestVariantType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				userID := params.Source.(User).ID
				user := new(User)
				err := dal.DB.Model(user).Where("id = ?", userID).Relation("FeatureTestVariants").First()
				if err != nil {
					fmt.Printf("Error retrieving user feature test variants: %v", err)
					return nil, err
				}
				return user.FeatureTestVariants, nil
			},
		},
	},
})
