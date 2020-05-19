package mutations

import (
	"github.com/go-pg/pg/v9"
	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
	"github.com/hwangm/isthemunibusy-go/service"
	"github.com/hwangm/isthemunibusy-go/types"
)

// GetCreateFeatureTestVariantMutation returns the graphql field config for creating a feature test variant
func GetCreateFeatureTestVariantMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.FeatureTestVariantType,
		Args: graphql.FieldConfigArgument{
			"featureTestID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"featureTestVariant": &graphql.ArgumentConfig{
				Type: types.FeatureTestVariantInputType,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			featureTestID := params.Args["featureTestID"].(int)
			testVariant := params.Args["featureTestVariant"].(map[string]interface{})
			var testVariantReturn types.FeatureTestVariant

			err := dal.DB.RunInTransaction(func(tx *pg.Tx) error {
				featureTestVariant, err := service.CreateFeatureTestVariant(
					tx,
					featureTestID,
					testVariant["name"].(string),
					testVariant["isControl"].(bool),
					testVariant["percentage"].(int),
				)

				if err != nil {
					return err
				}

				testVariantReturn = *featureTestVariant

				return nil
			})

			if err != nil {
				return nil, err
			}

			return testVariantReturn, nil
		},
	}
}

// GetUpdateFeatureTestVariantMutation returns the graphql field config for the updating a feature test variant
func GetUpdateFeatureTestVariantMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.FeatureTestVariantType,
		Args: graphql.FieldConfigArgument{
			"featureTestVariantID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"featureTestVariant": &graphql.ArgumentConfig{
				Type: types.FeatureTestVariantInputType,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			variantID := params.Args["featureTestVariantID"].(int)
			variant := params.Args["featureTestVariant"].(map[string]interface{})
			var variantReturn types.FeatureTestVariant

			err := dal.DB.RunInTransaction(func(tx *pg.Tx) error {
				testVariant, err := service.UpdateFeatureTestVariant(
					tx,
					variantID,
					variant["name"].(string),
					variant["isControl"].(bool),
					variant["percentage"].(int),
				)

				if err != nil {
					return err
				}

				variantReturn = *testVariant

				return nil
			})

			if err != nil {
				return nil, err
			}

			return variantReturn, nil
		},
	}
}

// GetDeleteFeatureTestVariantMutation returns the graphql field config for deleting a feature test variant
func GetDeleteFeatureTestVariantMutation() *graphql.Field {
	return &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"featureTestVariantID": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			variantID := params.Args["featureTestVariantID"].(int)

			err := dal.DB.RunInTransaction(func(tx *pg.Tx) error {
				err := service.DeleteFeatureTestVariant(tx, variantID)
				if err != nil {
					return err
				}

				err = service.DeleteUserFeatureTestVariantsByVariantID(tx, variantID)
				if err != nil {
					return err
				}

				return nil
			})

			if err != nil {
				return false, err
			}

			return true, nil
		},
	}
}
