// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsIntegrationCradlepointTestDeleteOrgCradlepointConnection tests the behavior of the OrgsIntegrationCradlepoint
func TestOrgsIntegrationCradlepointTestDeleteOrgCradlepointConnection(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsIntegrationCradlepoint.DeleteOrgCradlepointConnection(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsIntegrationCradlepointTestTestOrgCradlepointConnection tests the behavior of the OrgsIntegrationCradlepoint
func TestOrgsIntegrationCradlepointTestTestOrgCradlepointConnection(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsIntegrationCradlepoint.TestOrgCradlepointConnection(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"error":"Cradlepoint API keys are no longer valid, please verify and update the keys under organization settings.","last_status":"inactive"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIntegrationCradlepointTestTestOrgCradlepointConnection1 tests the behavior of the OrgsIntegrationCradlepoint
func TestOrgsIntegrationCradlepointTestTestOrgCradlepointConnection1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsIntegrationCradlepoint.TestOrgCradlepointConnection(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"error":"Cradlepoint API keys are no longer valid, please verify and update the keys under organization settings.","last_status":"inactive"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIntegrationCradlepointTestSetupOrgCradlepointConnectionToMist tests the behavior of the OrgsIntegrationCradlepoint
func TestOrgsIntegrationCradlepointTestSetupOrgCradlepointConnectionToMist(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    resp, err := orgsIntegrationCradlepoint.SetupOrgCradlepointConnectionToMist(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsIntegrationCradlepointTestUpdateOrgCradlepointConnectionToMist tests the behavior of the OrgsIntegrationCradlepoint
func TestOrgsIntegrationCradlepointTestUpdateOrgCradlepointConnectionToMist(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.AccountCradlepointConfig
    errBody := json.Unmarshal([]byte(`{"cp_api_id":"84446d61-2206-4ea5-855a-0043f980be54","cp_api_key":"79c329da9893e34099c7d8ad5cb9c941","ecm_api_id":"73446d61-2206-4ea5-855a-0043f980be62","ecm_api_key":"68b329da9893e34099c7d8ad5cb9c9405"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsIntegrationCradlepoint.UpdateOrgCradlepointConnectionToMist(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsIntegrationCradlepointTestSyncOrgCradlepointRouters tests the behavior of the OrgsIntegrationCradlepoint
func TestOrgsIntegrationCradlepointTestSyncOrgCradlepointRouters(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsIntegrationCradlepoint.SyncOrgCradlepointRouters(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
