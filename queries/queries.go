package queries

import (
	"github.com/graphql-go/graphql"
)

// GetRootFields returns basic fields
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"users":                   GetUsersQuery(),
		"user":                    GetUserQuery(),
		"featureTests":            GetFeatureTestsQuery(),
		"featureTest":             GetFeatureTestQuery(),
		"featureTestVariants":     GetFeatureTestVariantsQuery(),
		"userFeatureTestVariants": GetUserFeatureTestVariantsQuery(),
	}
}
