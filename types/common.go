package types

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
)

func init() {
	addVariantsToFeatureTestType()
}

// addVariantsToFeatureTestType adds the variants field to the feature test type
// this is necessary to avoid the cyclic compile references if both feature test
// and feature test variant types are initialized with the bidirectional associations
func addVariantsToFeatureTestType() {
	FeatureTestType.AddFieldConfig(
		"variants",
		&graphql.Field{
			Type: graphql.NewList(FeatureTestVariantType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				featureTestID := params.Source.(FeatureTest).ID
				var variants []FeatureTestVariant

				err := dal.DB.Model(&variants).Where("feature_test_id = ?", featureTestID).Select()
				if err != nil {
					fmt.Printf("Error getting variants for a feature test with id %d: %v", featureTestID, err)
					return nil, err
				}

				return variants, nil
			},
		},
	)
}
