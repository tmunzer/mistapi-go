// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestMSPsInventoryTestGetMspInventoryByMac tests the behavior of the MSPsInventory
func TestMSPsInventoryTestGetMspInventoryByMac(t *testing.T) {
	ctx := context.Background()
	mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceMac := "0000000000ab"
	apiResponse, err := msPsInventory.GetMspInventoryByMac(ctx, mspId, deviceMac)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"mac":"5c5b35000018","model":"AP200","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","serial":"FXLH2015150025","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","type":"ap"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsInventoryTestGetMspInventoryByMac1 tests the behavior of the MSPsInventory
func TestMSPsInventoryTestGetMspInventoryByMac1(t *testing.T) {
	ctx := context.Background()
	mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceMac := "0000000000ab"
	apiResponse, err := msPsInventory.GetMspInventoryByMac(ctx, mspId, deviceMac)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"mac":"5c5b35000018","model":"AP200","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","serial":"FXLH2015150025","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","type":"ap"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
