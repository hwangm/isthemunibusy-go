package mutations

import (
	"github.com/go-pg/pg/v9"
	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
	"github.com/hwangm/isthemunibusy-go/service"
	"github.com/hwangm/isthemunibusy-go/types"
)

// GetCreateUserFeatureTestVariantMutation creates a user feature test variant
func GetCreateUserFeatureTestVariantMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.UserFeatureTestVariantType,
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: types.UserFeatureTestVariantInputType,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			userTestVariant := params.Args["input"].(map[string]interface{})
			userID := userTestVariant["userID"].(int)
			featureTestVariantID := userTestVariant["featureTestVariantID"].(int)
			var featureTestVariantReturn types.UserToFeatureUserTestVariant

			err := dal.DB.RunInTransaction(func(tx *pg.Tx) error {
				userFTV, err := service.CreateUserFeatureTestVariant(tx, userID, featureTestVariantID)
				if err != nil {
					return err
				}

				featureTestVariantReturn = *userFTV

				return nil
			})

			if err != nil {
				return nil, err
			}

			return featureTestVariantReturn, nil
		},
	}
}

// GetUpdateUserFeatureTestVariantMutation updates an existing user feature test variant
func GetUpdateUserFeatureTestVariantMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.UserFeatureTestVariantType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"userFeatureTestVariant": &graphql.ArgumentConfig{
				Type: types.UserFeatureTestVariantInputType,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			userTestVariantID := params.Args["id"].(int)
			userTestVariant := params.Args["userFeatureTestVariant"].(map[string]interface{})
			userID := userTestVariant["userID"].(int)
			featureTestVariantID := userTestVariant["featureTestVariantID"].(int)
			var featureTestVariantReturn types.UserToFeatureUserTestVariant

			err := dal.DB.RunInTransaction(func(tx *pg.Tx) error {
				userFTV, err := service.UpdateUserFeatureTestVariant(tx, userTestVariantID, userID, featureTestVariantID)
				if err != nil {
					return err
				}

				featureTestVariantReturn = *userFTV

				return nil
			})

			if err != nil {
				return nil, err
			}

			return featureTestVariantReturn, nil
		},
	}
}

// GetDeleteUserFeatureTestVariantMutation deletes a user feature test variant
func GetDeleteUserFeatureTestVariantMutation() *graphql.Field {
	return &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			userTestVariantID := params.Args["id"].(int)
			err := dal.DB.RunInTransaction(func(tx *pg.Tx) error {
				err := service.DeleteUserFeatureTestVariant(tx, userTestVariantID)
				if err != nil {
					return err
				}

				return nil
			})

			if err != nil {
				return false, err
			}

			return true, err
		},
	}
}
