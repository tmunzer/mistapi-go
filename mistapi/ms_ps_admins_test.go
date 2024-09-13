package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestMSPsAdminsTestListMspAdmins tests the behavior of the MSPsAdmins
func TestMSPsAdminsTestListMspAdmins(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := msPsAdmins.ListMspAdmins(ctx, mspId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"admin_id":"456b7016-a916-a4b1-78dd-72b947c152b7","email":"jsmith@mycorp.org","first_name":"Joe","last_name":"Smith","privileges":[{"role":"admin","scope":"msp"},{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","role":"admin","scope":"org"},{"orggroup_ids":["507f1bab-13ba-73e2-f291-2bcb8d1362b0"],"role":"read","scope":"orggroup"}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsAdminsTestRevokeMspAdmin tests the behavior of the MSPsAdmins
func TestMSPsAdminsTestRevokeMspAdmin(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    adminId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := msPsAdmins.RevokeMspAdmin(ctx, mspId, adminId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestMSPsAdminsTestGetMspAdmin tests the behavior of the MSPsAdmins
func TestMSPsAdminsTestGetMspAdmin(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    adminId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := msPsAdmins.GetMspAdmin(ctx, mspId, adminId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"admin_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","email":"user@example.com","first_name":"string","last_name":"string","password_modified_time":1656353525,"privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"]}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsAdminsTestUpdateMspAdmin tests the behavior of the MSPsAdmins
func TestMSPsAdminsTestUpdateMspAdmin(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    adminId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Admin
    errBody := json.Unmarshal([]byte(`{"email":"jsnow@abc.com","first_name":"string","last_name":"string","privileges":[{"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","role":"admin","scope":"org","sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]},{"role":"admin","scope":"site","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := msPsAdmins.UpdateMspAdmin(ctx, mspId, adminId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"admin_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","email":"user@example.com","first_name":"string","last_name":"string","password_modified_time":1656353525,"privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"]}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsAdminsTestInviteMspAdmin tests the behavior of the MSPsAdmins
func TestMSPsAdminsTestInviteMspAdmin(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Admin
    errBody := json.Unmarshal([]byte(`{"email":"user@example.com","first_name":"string","last_name":"string","privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := msPsAdmins.InviteMspAdmin(ctx, mspId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"admin_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","email":"user@example.com","first_name":"string","last_name":"string","password_modified_time":1656353525,"privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"]}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestMSPsAdminsTestUninviteMspAdmin tests the behavior of the MSPsAdmins
func TestMSPsAdminsTestUninviteMspAdmin(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    inviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := msPsAdmins.UninviteMspAdmin(ctx, mspId, inviteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestMSPsAdminsTestUpdateMspAdminInvite tests the behavior of the MSPsAdmins
func TestMSPsAdminsTestUpdateMspAdminInvite(t *testing.T) {
    ctx := context.Background()
    mspId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    inviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Admin
    errBody := json.Unmarshal([]byte(`{"email":"user@example.com","first_name":"string","last_name":"string","privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := msPsAdmins.UpdateMspAdminInvite(ctx, mspId, inviteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"admin_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","email":"user@example.com","first_name":"string","last_name":"string","password_modified_time":1656353525,"privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"]}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
