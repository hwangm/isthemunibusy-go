package queries

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
	"github.com/hwangm/isthemunibusy-go/types"
)

// File is named as such because files ending in _test are reserved

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

// GetFeatureTestQuery returns a single feature test by id
func GetFeatureTestQuery() *graphql.Field {
	return &graphql.Field{
		Type: types.FeatureTestType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			featureTestID := params.Args["id"].(int)
			featureTest := types.FeatureTest{ID: featureTestID}
			err := dal.DB.Select(&featureTest)
			if err != nil {
				fmt.Printf("Error retrieving feature test by id: %v", err)
				return nil, err
			}
			return featureTest, nil
		},
	}
}
