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
	ID                   int
	UserID               int
	FeatureTestVariantID int       `pg:"feature_testvariant_id"`
	DeletedAt            time.Time `pg:",soft_delete"`
}

// UserFeatureTestVariantType is GraphQL schema for the user feature test variant type
var UserFeatureTestVariantType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserFeatureTestVariant",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"userID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return params.Source.(UserToFeatureUserTestVariant).UserID, nil
			},
		},
		"featureTestVariantID": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return params.Source.(UserToFeatureUserTestVariant).FeatureTestVariantID, nil
			},
		},
		"user": &graphql.Field{
			Type: UserType,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				userID := params.Source.(UserToFeatureUserTestVariant).UserID
				user := User{ID: userID}
				err := dal.DB.Select(&user)
				if err != nil {
					fmt.Printf("Error retrieving user from user feature test variant: %v", err)
					return nil, err
				}

				return user, nil
			},
		},
		"featureTestVariant": &graphql.Field{
			Type: FeatureTestVariantType,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				variantID := params.Source.(UserToFeatureUserTestVariant).FeatureTestVariantID
				variant := FeatureTestVariant{ID: variantID}
				err := dal.DB.Select(&variant)
				if err != nil {
					fmt.Printf("Error retrieving feature test variant from user feature test variant: %v", err)
					return nil, err
				}

				return variant, nil
			},
		},
	},
})

// UserFeatureTestVariantInputType is the Graphql schema for user feature test variants,
// used in the user feature test variant input object schema
var UserFeatureTestVariantInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserFeatureTestVariantInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"userID": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"featureTestVariantID": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})

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
