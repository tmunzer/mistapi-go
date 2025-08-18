// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestMSPsLogoTestDeleteMspLogo tests the behavior of the MSPsLogo
func TestMSPsLogoTestDeleteMspLogo(t *testing.T) {
	ctx := context.Background()
	mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := msPsLogo.DeleteMspLogo(ctx, mspId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestMSPsLogoTestPostMspLogo tests the behavior of the MSPsLogo
func TestMSPsLogoTestPostMspLogo(t *testing.T) {
	ctx := context.Background()
	mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	resp, err := msPsLogo.PostMspLogo(ctx, mspId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
