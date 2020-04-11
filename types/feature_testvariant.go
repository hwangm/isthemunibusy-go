package types

import (
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/hwangm/isthemunibusy-go/dal"
)

// FeatureTestVariant type definition
type FeatureTestVariant struct {
	tableName     struct{} `pg:"feature_testvariants"`
	ID            int
	Name          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	FeatureTestID int
	FeatureTest   FeatureTest
}

// FeatureTestVariantType is the graphql type for feature test variants
var FeatureTestVariantType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeatureTestVariant",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.ID},
		"name":       &graphql.Field{Type: graphql.String},
		"created_at": &graphql.Field{Type: graphql.DateTime},
		"updated_at": &graphql.Field{Type: graphql.DateTime},
		"feature_test": &graphql.Field{
			Type: FeatureTest,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				featureTestID := params.Source.(FeatureTestVariant).FeatureTestID
				featureTest := new(FeatureTest)
				err := dal.DB.Model(featureTest).Where("id = ?", featureTestID).First()
				if err != nil {
					fmt.Printf("Error retrieving feature test from test variant: %v", err)
					return nil, err
				}
				return featureTest, nil
			},
		},
	},
})
