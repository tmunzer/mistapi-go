package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsSSORolesTestListOrgSsoRoles tests the behavior of the OrgsSSORoles
func TestOrgsSSORolesTestListOrgSsoRoles(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsSsoRoles.ListOrgSsoRoles(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsSSORolesTestListOrgSsoRoles1 tests the behavior of the OrgsSSORoles
func TestOrgsSSORolesTestListOrgSsoRoles1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsSsoRoles.ListOrgSsoRoles(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsSSORolesTestCreateOrgSsoRole tests the behavior of the OrgsSSORoles
func TestOrgsSSORolesTestCreateOrgSsoRole(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SsoRoleOrg
    errBody := json.Unmarshal([]byte(`{"name":"string","privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsSsoRoles.CreateOrgSsoRole(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","privileges":[{"role":"admin","scope":"sitegroup","sitegroup_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSSORolesTestCreateOrgSsoRole1 tests the behavior of the OrgsSSORoles
func TestOrgsSSORolesTestCreateOrgSsoRole1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SsoRoleOrg
    errBody := json.Unmarshal([]byte(`{"name":"string","privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsSsoRoles.CreateOrgSsoRole(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","privileges":[{"role":"admin","scope":"sitegroup","sitegroup_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSSORolesTestDeleteOrgSsoRole tests the behavior of the OrgsSSORoles
func TestOrgsSSORolesTestDeleteOrgSsoRole(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ssoroleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsSsoRoles.DeleteOrgSsoRole(ctx, orgId, ssoroleId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSSORolesTestGetOrgSsoRole tests the behavior of the OrgsSSORoles
func TestOrgsSSORolesTestGetOrgSsoRole(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ssoroleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSsoRoles.GetOrgSsoRole(ctx, orgId, ssoroleId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","privileges":[{"role":"admin","scope":"sitegroup","sitegroup_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSSORolesTestGetOrgSsoRole1 tests the behavior of the OrgsSSORoles
func TestOrgsSSORolesTestGetOrgSsoRole1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ssoroleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSsoRoles.GetOrgSsoRole(ctx, orgId, ssoroleId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","privileges":[{"role":"admin","scope":"sitegroup","sitegroup_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSSORolesTestUpdateOrgSsoRole tests the behavior of the OrgsSSORoles
func TestOrgsSSORolesTestUpdateOrgSsoRole(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ssoroleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SsoRoleOrg
    errBody := json.Unmarshal([]byte(`{"name":"string","privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsSsoRoles.UpdateOrgSsoRole(ctx, orgId, ssoroleId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","privileges":[{"role":"admin","scope":"sitegroup","sitegroup_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSSORolesTestUpdateOrgSsoRole1 tests the behavior of the OrgsSSORoles
func TestOrgsSSORolesTestUpdateOrgSsoRole1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ssoroleId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.SsoRoleOrg
    errBody := json.Unmarshal([]byte(`{"name":"string","privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"]}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsSsoRoles.UpdateOrgSsoRole(ctx, orgId, ssoroleId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","privileges":[{"role":"admin","scope":"sitegroup","sitegroup_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
