package queries

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
	"github.com/hwangm/isthemunibusy-go/types"
)

// GetFeatureTestsQuery returns a list of Feature Tests
func GetFeatureTestsQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.FeatureTestType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var featureTests []types.FeatureTest
			err := dal.DB.Model(&featureTests).Select()
			if err != nil {
				fmt.Printf("Error retrieving all Feature Tests: %v", err)
				return nil, err
			}
			return featureTests, nil
		},
	}
}
