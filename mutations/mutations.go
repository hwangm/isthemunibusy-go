package mutations

import (
	"github.com/graphql-go/graphql"
)

// GetRootFields returns available mutation fields
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"createUser":               GetCreateUserMutation(),
		"deleteUser":               GetDeleteUserByIDMutation(),
		"createRole":               GetCreateRoleMutation(),
		"createFeatureTest":        GetCreateFeatureTestMutation(),
		"deleteFeatureTest":        GetDeleteFeatureTestMutation(),
		"updateFeatureTest":        GetUpdateFeatureTestMutation(),
		"createFeatureTestVariant": GetCreateFeatureTestVariantMutation(),
		"deleteFeatureTestVariant": GetDeleteFeatureTestVariantMutation(),
		"updateFeatureTestVariant": GetUpdateFeatureTestVariantMutation(),
		// "createUserTestVariant": GetCreateUserTestVariantMutation(),
		// "updateUserTestVariant": GetUpdateUserTestVariantMutation(),
		// "deleteUserTestVariant": GetDeleteUserTestVariantMutation(),
	}
}
