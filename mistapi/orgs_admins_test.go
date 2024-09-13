package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsAdminsTestListOrgAdmins tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestListOrgAdmins(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsAdmins.ListOrgAdmins(ctx, orgId)
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

// TestOrgsAdminsTestRevokeOrgAdmin tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestRevokeOrgAdmin(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    adminId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsAdmins.RevokeOrgAdmin(ctx, orgId, adminId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAdminsTestUpdateOrgAdmin tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestUpdateOrgAdmin(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    adminId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Admin
    errBody := json.Unmarshal([]byte(`{"email":"jsnow@abc.com","expire_time":0,"first_name":"John","hours":24,"last_name":"Sno","phone":"string","phone2":"string","privileges":[{"msp_id":"c0cf23fc-d82f-4219-988c-82fb61d8c875","name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","org_name":"string","orggroup_ids":["497f6eca-6276-4993-bfeb-53d5bbba6f08"],"role":"admin","scope":"org","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","sitegroup_ids":["497f6eca-6276-4993-bfeb-53d6bbba6f08"],"views":"switch_admin"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsAdmins.UpdateOrgAdmin(ctx, orgId, adminId, &body)
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

// TestOrgsAdminsTestInviteOrgAdmin tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestInviteOrgAdmin(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Admin
    errBody := json.Unmarshal([]byte(`{"email":"user@example.com","first_name":"string","last_name":"string","privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsAdmins.InviteOrgAdmin(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAdminsTestUninviteOrgAdmin tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestUninviteOrgAdmin(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    inviteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsAdmins.UninviteOrgAdmin(ctx, orgId, inviteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAdminsTestUpdateOrgAdminInvite tests the behavior of the OrgsAdmins
func TestOrgsAdminsTestUpdateOrgAdminInvite(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
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
    resp, err := orgsAdmins.UpdateOrgAdminInvite(ctx, orgId, inviteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
