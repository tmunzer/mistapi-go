package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestMSPsOrgGroupsTestListMspOrgGroups tests the behavior of the MSPsOrgGroups
func TestMSPsOrgGroupsTestListMspOrgGroups(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := msPsOrgGroups.ListMspOrgGroups(ctx, mspId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","org_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsOrgGroupsTestCreateMspOrgGroup tests the behavior of the MSPsOrgGroups
func TestMSPsOrgGroupsTestCreateMspOrgGroup(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Orggroup
    errBody := json.Unmarshal([]byte(`{"name":"string","org_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := msPsOrgGroups.CreateMspOrgGroup(ctx, mspId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","org_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsOrgGroupsTestDeleteMspOrgGroup tests the behavior of the MSPsOrgGroups
func TestMSPsOrgGroupsTestDeleteMspOrgGroup(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    orggroupId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := msPsOrgGroups.DeleteMspOrgGroup(ctx, mspId, orggroupId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestMSPsOrgGroupsTestGetMspOrgGroup tests the behavior of the MSPsOrgGroups
func TestMSPsOrgGroupsTestGetMspOrgGroup(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    orggroupId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := msPsOrgGroups.GetMspOrgGroup(ctx, mspId, orggroupId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","org_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsOrgGroupsTestUpdateMspOrgGroup tests the behavior of the MSPsOrgGroups
func TestMSPsOrgGroupsTestUpdateMspOrgGroup(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    orggroupId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := msPsOrgGroups.UpdateMspOrgGroup(ctx, mspId, orggroupId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","org_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
