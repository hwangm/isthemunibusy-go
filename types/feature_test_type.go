package types

import (
	"time"

	"github.com/graphql-go/graphql"
)

// File is named as such because files ending in _test are reserved

// FeatureTest type definition
type FeatureTest struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	StartTime time.Time
	EndTime   time.Time
	DeletedAt time.Time `pg:",soft_delete"`
}

// FeatureTestType is the GraphQL schema for feature tests
var FeatureTestType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeatureTest",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.ID},
		"name":      &graphql.Field{Type: graphql.String},
		"createdAt": &graphql.Field{Type: graphql.DateTime},
		"updatedAt": &graphql.Field{Type: graphql.DateTime},
		"startTime": &graphql.Field{Type: graphql.DateTime},
		"endTime":   &graphql.Field{Type: graphql.DateTime},
	},
})

// FeatureTestInputType is the GraphQL schema for creating a feature test
var FeatureTestInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "FeatureTestInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"startTime": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.DateTime),
		},
		"endTime": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.DateTime),
		},
		"variants": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.NewList(FeatureTestVariantInputType)),
		},
	},
})
