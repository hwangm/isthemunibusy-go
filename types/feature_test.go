package types

import (
	"time"

	"github.com/graphql-go/graphql"
)

// FeatureTest type definition
type FeatureTest struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// FeatureTestType is the GraphQL schema for feature tests
var FeatureTestType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FeatureTest",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.ID},
		"name":       &graphql.Field{Type: graphql.String},
		"created_at": &graphql.Field{Type: graphql.DateTime},
		"updated_at": &graphql.Field{Type: graphql.DateTime},
		"start_time": &graphql.Field{Type: graphql.DateTime},
		"end_time":   &graphql.Field{Type: graphql.DateTime},
	},
})
