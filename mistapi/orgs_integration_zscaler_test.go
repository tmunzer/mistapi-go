package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsIntegrationZscalerTestDeleteOrgZscalerIntegration tests the behavior of the OrgsIntegrationZscaler
func TestOrgsIntegrationZscalerTestDeleteOrgZscalerIntegration(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsIntegrationZscaler.DeleteOrgZscalerIntegration(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsIntegrationZscalerTestGetOrgZscalerIntegration tests the behavior of the OrgsIntegrationZscaler
func TestOrgsIntegrationZscalerTestGetOrgZscalerIntegration(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsIntegrationZscaler.GetOrgZscalerIntegration(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cloud_name":"zscalerbeta.net","partner_key":"K35vrZcK3JvrZc","username":"john@nmo.com"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIntegrationZscalerTestGetOrgZscalerIntegration1 tests the behavior of the OrgsIntegrationZscaler
func TestOrgsIntegrationZscalerTestGetOrgZscalerIntegration1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsIntegrationZscaler.GetOrgZscalerIntegration(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cloud_name":"zscalerbeta.net","partner_key":"K35vrZcK3JvrZc","username":"john@nmo.com"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIntegrationZscalerTestSetupOrgZscalerIntegration tests the behavior of the OrgsIntegrationZscaler
func TestOrgsIntegrationZscalerTestSetupOrgZscalerIntegration(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.AccountZscalerConfig
    errBody := json.Unmarshal([]byte(`{"cloud_name":"zscalerbeta.net","partner_key":"K35vrZcK3JvrZc","password":"password","username":"john@nmo.com"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsIntegrationZscaler.SetupOrgZscalerIntegration(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
