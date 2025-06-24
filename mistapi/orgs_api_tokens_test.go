package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsAPITokensTestListOrgApiTokens tests the behavior of the OrgsAPITokens
func TestOrgsAPITokensTestListOrgApiTokens(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsApiTokens.ListOrgApiTokens(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_by":"user@mycorp.com","created_time":1626875902,"id":"497f6eca-6276-4993-bfeb-53f0bbba6f08","key":"1qkb...QQCL","last_used":1690115110,"name":"org_token_xyz","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","privileges":[{"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","role":"admin","scope":"org"}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAPITokensTestListOrgApiTokens1 tests the behavior of the OrgsAPITokens
func TestOrgsAPITokensTestListOrgApiTokens1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsApiTokens.ListOrgApiTokens(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_by":"user@mycorp.com","created_time":1626875902,"id":"497f6eca-6276-4993-bfeb-53f0bbba6f08","key":"1qkb...QQCL","last_used":1690115110,"name":"org_token_xyz","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","privileges":[{"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","role":"admin","scope":"org"}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAPITokensTestCreateOrgApiToken tests the behavior of the OrgsAPITokens
func TestOrgsAPITokensTestCreateOrgApiToken(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsApiTokens.CreateOrgApiToken(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_by":"user@mycorp.com","created_time":1626875902,"id":"497f6eca-6276-4993-bfeb-53efbbba6f08","key":"1qkb...QQCL","last_used":1690115110,"name":"org_token_xyz","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","privileges":[{"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","role":"admin","scope":"org"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAPITokensTestCreateOrgApiToken1 tests the behavior of the OrgsAPITokens
func TestOrgsAPITokensTestCreateOrgApiToken1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsApiTokens.CreateOrgApiToken(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_by":"user@mycorp.com","created_time":1626875902,"id":"497f6eca-6276-4993-bfeb-53efbbba6f08","key":"1qkb...QQCL","last_used":1690115110,"name":"org_token_xyz","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","privileges":[{"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","role":"admin","scope":"org"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAPITokensTestDeleteOrgApiToken tests the behavior of the OrgsAPITokens
func TestOrgsAPITokensTestDeleteOrgApiToken(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apitokenId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsApiTokens.DeleteOrgApiToken(ctx, orgId, apitokenId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAPITokensTestGetOrgApiToken tests the behavior of the OrgsAPITokens
func TestOrgsAPITokensTestGetOrgApiToken(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apitokenId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsApiTokens.GetOrgApiToken(ctx, orgId, apitokenId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_by":"user@mycorp.com","created_time":1626875902,"id":"497f6eca-6276-4993-bfeb-53efbbba6f08","key":"1qkb...QQCL","last_used":1690115110,"name":"org_token_xyz","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","privileges":[{"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","role":"admin","scope":"org"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAPITokensTestGetOrgApiToken1 tests the behavior of the OrgsAPITokens
func TestOrgsAPITokensTestGetOrgApiToken1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apitokenId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsApiTokens.GetOrgApiToken(ctx, orgId, apitokenId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_by":"user@mycorp.com","created_time":1626875902,"id":"497f6eca-6276-4993-bfeb-53efbbba6f08","key":"1qkb...QQCL","last_used":1690115110,"name":"org_token_xyz","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","privileges":[{"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","role":"admin","scope":"org"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAPITokensTestUpdateOrgApiToken tests the behavior of the OrgsAPITokens
func TestOrgsAPITokensTestUpdateOrgApiToken(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apitokenId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.OrgApitoken
    errBody := json.Unmarshal([]byte(`{"name":"org_token_xyz","privileges":[{"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","role":"admin","scope":"org"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsApiTokens.UpdateOrgApiToken(ctx, orgId, apitokenId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
