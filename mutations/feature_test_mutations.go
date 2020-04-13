package mutations

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
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
			name := params.Args["input"].(map[string]interface{})["name"].(string)
			currentTime := time.Now()
			endTime := currentTime.Add(time.Hour * 24 * 365) // 1 year
			featureTest := types.FeatureTest{
				Name:      name,
				StartTime: currentTime,
				EndTime:   endTime,
			}
			err := dal.DB.Insert(&featureTest)
			if err != nil {
				fmt.Printf("Error creating new feature test: %v", err)
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

				featureTestVariant := types.FeatureTestVariant{}
				_, err = tx.Model(&featureTestVariant).Where("feature_test_id = ?", featureTestID).Delete()
				if err != nil {
					fmt.Printf("Error deleting feature test variants from feature test: %v", err)
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
