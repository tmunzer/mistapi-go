package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestMSPsSSORolesTestListMspSsoRoles tests the behavior of the MSPsSSORoles
func TestMSPsSSORolesTestListMspSsoRoles(t *testing.T) {
	ctx := context.Background()
	mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := msPsSsoRoles.ListMspSsoRoles(ctx, mspId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestMSPsSSORolesTestCreateMspSsoRole tests the behavior of the MSPsSSORoles
func TestMSPsSSORolesTestCreateMspSsoRole(t *testing.T) {
	ctx := context.Background()
	mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := msPsSsoRoles.CreateMspSsoRole(ctx, mspId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","privileges":[{"orggroup_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","role":"read","scope":"orggroup"}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsSSORolesTestDeleteMspSsoRole tests the behavior of the MSPsSSORoles
func TestMSPsSSORolesTestDeleteMspSsoRole(t *testing.T) {
	ctx := context.Background()
	mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	ssoroleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := msPsSsoRoles.DeleteMspSsoRole(ctx, mspId, ssoroleId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestMSPsSSORolesTestUpdateMspSsoRole tests the behavior of the MSPsSSORoles
func TestMSPsSSORolesTestUpdateMspSsoRole(t *testing.T) {
	ctx := context.Background()
	mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	ssoroleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := msPsSsoRoles.UpdateMspSsoRole(ctx, mspId, ssoroleId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","privileges":[{"orggroup_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","role":"read","scope":"orggroup"}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
