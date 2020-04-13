package service

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/hwangm/isthemunibusy-go/types"
)

// Commonly used methods relating to feature tests should go here

// DeleteUserFeatureTestVariantsByVariantID deletes user test variants given a variant ID
func DeleteUserFeatureTestVariantsByVariantID(tx pg.Tx, variantID int) error {
	userFeatureTestVariant := types.UserToFeatureUserTestVariant{}
	_, err := tx.Model(&userFeatureTestVariant).Where("feature_testvariant_id = ?", variantID).Delete()
	if err != nil {
		fmt.Printf("Error deleting user feature test variants from variant id: %v", err)
		return err
	}
	return nil
}

// DeleteFeatureTestVariantsByTestID deletes test variants given a test ID
func DeleteFeatureTestVariantsByTestID(tx pg.Tx, testID int) error {
	featureTestVariant := types.FeatureTestVariant{}
	_, err := tx.Model(&featureTestVariant).Where("feature_test_id = ?", testID).Delete()
	if err != nil {
		fmt.Printf("Error deleting feature test variants from feature test: %v", err)
		return err
	}
	return nil
}
