package mutations

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
	"github.com/hwangm/isthemunibusy-go/service"
	"github.com/hwangm/isthemunibusy-go/types"
)

// GetCreateFeatureTestMutation creates a new feature test
func GetCreateFeatureTestMutation() *graphql.Field {
	return &graphql.Field{
		Type: types.FeatureTestType,
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: types.FeatureTestInputType,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			input := params.Args["input"].(map[string]interface{})
			name := input["name"].(string)
			variants := input["variants"].([]interface{})

			currentTime := time.Now()
			endTime := currentTime.Add(time.Hour * 24 * 365) // 1 year
			featureTest := types.FeatureTest{
				Name:      name,
				StartTime: currentTime,
				EndTime:   endTime,
			}
			err := dal.DB.RunInTransaction(func(tx *pg.Tx) error {
				err := tx.Insert(&featureTest)

				if err != nil {
					fmt.Printf("Error creating new feature test: %v", err)
					return err
				}

				if len(variants) > 0 {
					for _, variant := range variants {
						variantMap := variant.(map[string]interface{})
						testVariant := types.FeatureTestVariant{
							Name:          variantMap["name"].(string),
							IsControl:     variantMap["isControl"].(bool),
							Percentage:    variantMap["percentage"].(int),
							FeatureTestID: featureTest.ID,
						}
						err := tx.Insert(&testVariant)
						if err != nil {
							fmt.Printf("Error creating new feature test variants in feature test: %v", err)
							return err
						}
					}
				}

				return nil
			})

			if err != nil {
				return nil, err
			}

			return featureTest, nil
		},
	}
}

// GetDeleteFeatureTestMutation deletes a feature test by id
func GetDeleteFeatureTestMutation() *graphql.Field {
	return &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			featureTestID := params.Args["id"].(int)
			err := dal.DB.RunInTransaction(func(tx *pg.Tx) error {
				featureTest := types.FeatureTest{ID: featureTestID}
				err := tx.Delete(&featureTest)
				if err != nil {
					fmt.Printf("Error deleting feature test: %v", err)
					return err
				}

				featureTestVariants := []types.FeatureTestVariant{}
				err = tx.Model(&featureTestVariants).Where("feature_test_id = ?", featureTestID).Select()
				if err != nil {
					fmt.Printf("Error selecting feature test variants from feature test: %v", err)
					return err
				}

				if len(featureTestVariants) > 0 {
					for _, variant := range featureTestVariants {
						err := service.DeleteUserFeatureTestVariantsByVariantID(*tx, variant.ID)
						if err != nil {
							return err
						}
					}
				}

				err = service.DeleteFeatureTestVariantsByTestID(*tx, featureTestID)
				if err != nil {
					return err
				}

				return nil
			})

			if err != nil {
				fmt.Printf("Error deleting feature test in transaction: %v", err)
				return false, err
			}

			return true, nil
		},
	}
}
