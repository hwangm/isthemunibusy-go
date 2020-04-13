package mutations

import (
	"fmt"
	"time"

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
