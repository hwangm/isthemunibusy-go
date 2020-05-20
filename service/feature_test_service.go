package service

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/hwangm/isthemunibusy-go/types"
)

// Commonly used methods relating to feature tests should go here

// DeleteUserFeatureTestVariantsByVariantID deletes user test variants given a variant ID
func DeleteUserFeatureTestVariantsByVariantID(tx *pg.Tx, variantID int) error {
	userFeatureTestVariant := types.UserToFeatureUserTestVariant{}
	_, err := tx.Model(&userFeatureTestVariant).Where("feature_testvariant_id = ?", variantID).Delete()
	if err != nil {
		fmt.Printf("Error deleting user feature test variants from variant id: %v", err)
		return err
	}
	return nil
}

// DeleteFeatureTestVariantsByTestID deletes test variants given a test ID
func DeleteFeatureTestVariantsByTestID(tx *pg.Tx, testID int) error {
	featureTestVariant := types.FeatureTestVariant{}
	_, err := tx.Model(&featureTestVariant).Where("feature_test_id = ?", testID).Delete()
	if err != nil {
		fmt.Printf("Error deleting feature test variants from feature test: %v", err)
		return err
	}
	return nil
}

// UpdateFeatureTestByID updates an existing feature test with given testID with new name and end time, if given
func UpdateFeatureTestByID(tx *pg.Tx, testID int, name *string, endTime *time.Time) (*types.FeatureTest, error) {
	featureTest := types.FeatureTest{}
	err := tx.Model(&featureTest).Where("id = ?", testID).Select()
	if err != nil {
		fmt.Printf("Error retrieving feature test by id: %v", err)
		return nil, err
	}

	if name != nil {
		featureTest.Name = *name
	}
	if endTime != nil {
		featureTest.EndTime = *endTime
	}

	_, err = tx.Model(&featureTest).WherePK().Update()
	if err != nil {
		fmt.Printf("Error updating feature test: %v", err)
		return nil, err
	}

	return &featureTest, nil
}

// CreateFeatureTestVariant creates a new feature test variant
func CreateFeatureTestVariant(tx *pg.Tx, featureTestID int, name string, isControl bool, percentage int) (*types.FeatureTestVariant, error) {
	testVariant := types.FeatureTestVariant{
		FeatureTestID: featureTestID,
		Name:          name,
		IsControl:     isControl,
		Percentage:    percentage,
	}

	err := tx.Insert(&testVariant)
	if err != nil {
		fmt.Printf("Error creating new feature test variant: %v", err)
		return nil, err
	}

	return &testVariant, nil
}

// UpdateFeatureTestVariant updates an existing feature test variant
func UpdateFeatureTestVariant(tx *pg.Tx, featureTestVariantID int, name string, isControl bool, percentage int) (*types.FeatureTestVariant, error) {
	testVariant := types.FeatureTestVariant{
		ID: featureTestVariantID,
	}

	err := tx.Select(&testVariant)
	if err != nil {
		fmt.Printf("Error selecting existing feature test variant during update: %v", err)
		return nil, err
	}

	testVariant.Name = name
	testVariant.IsControl = isControl
	testVariant.Percentage = percentage

	err = tx.Update(&testVariant)
	if err != nil {
		fmt.Printf("Error updating existing feature test variant during update: %v", err)
		return nil, err
	}

	return &testVariant, nil
}

// DeleteFeatureTestVariant deletes a feature test variant by ID
func DeleteFeatureTestVariant(tx *pg.Tx, variantID int) error {
	variant := types.FeatureTestVariant{
		ID: variantID,
	}

	err := tx.Delete(&variant)
	if err != nil {
		fmt.Printf("Error deleting feature test variant: %v", err)
		return err
	}

	return nil
}

// CreateUserFeatureTestVariant creates a new user feature test variant
func CreateUserFeatureTestVariant(tx *pg.Tx, userID int, variantID int) (*types.UserToFeatureUserTestVariant, error) {
	userFTV := types.UserToFeatureUserTestVariant{
		UserID:               userID,
		FeatureTestVariantID: variantID,
	}

	err := tx.Insert(&userFTV)
	if err != nil {
		fmt.Printf("Error creating user feature test variant: %v", err)
		return nil, err
	}

	return &userFTV, nil
}

// UpdateUserFeatureTestVariant updates an existing user feature test variant
func UpdateUserFeatureTestVariant(tx *pg.Tx, userTestVariantID int, userID int, variantID int) (*types.UserToFeatureUserTestVariant, error) {
	userFTV := types.UserToFeatureUserTestVariant{
		ID: userTestVariantID,
	}

	err := tx.Select(&userFTV)
	if err != nil {
		fmt.Printf("Error selecting existing user feature test variant during update: %v", err)
		return nil, err
	}

	userFTV.UserID = userID
	userFTV.FeatureTestVariantID = variantID

	err = tx.Update(&userFTV)
	if err != nil {
		fmt.Printf("Error updating user feature test variant: %v", err)
		return nil, err
	}

	return &userFTV, nil
}
