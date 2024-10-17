package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestOrgsSettingZscalerTestDeleteOrgZscalerCredential tests the behavior of the OrgsSettingZscaler
func TestOrgsSettingZscalerTestDeleteOrgZscalerCredential(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsSettingZscaler.DeleteOrgZscalerCredential(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSettingZscalerTestGetOrgZscalerCredential tests the behavior of the OrgsSettingZscaler
func TestOrgsSettingZscalerTestGetOrgZscalerCredential(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsSettingZscaler.GetOrgZscalerCredential(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"cloud_name":"zscalerbeta.net","partner_key":"K35vrZcK3JvrZc","username":"john@nmo.com"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSettingZscalerTestSetupOrgZscalerCredential tests the behavior of the OrgsSettingZscaler
func TestOrgsSettingZscalerTestSetupOrgZscalerCredential(t *testing.T) {
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
	resp, err := orgsSettingZscaler.SetupOrgZscalerCredential(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
