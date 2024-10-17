package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSelfAPITokenTestListApiTokens tests the behavior of the SelfAPIToken
func TestSelfAPITokenTestListApiTokens(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := selfApiToken.ListApiTokens(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":1626875902,"id":"864f351a-1377-4ad9-83f8-72f3fe6199ba","key":"1qkb...QQCL","last_used":1690115110,"name":"org_token_xyz"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSelfAPITokenTestCreateApiToken tests the behavior of the SelfAPIToken
func TestSelfAPITokenTestCreateApiToken(t *testing.T) {
	ctx := context.Background()
	var body models.UserApitoken
	errBody := json.Unmarshal([]byte(`{"name":"org_token_xyz"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := selfApiToken.CreateApiToken(ctx, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":1626875902,"id":"864f351a-1377-4ad9-83f8-72f3fe6199ba","key":"1qkb...QQCL","last_used":1690115110,"name":"org_token_xyz"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSelfAPITokenTestDeleteApiToken tests the behavior of the SelfAPIToken
func TestSelfAPITokenTestDeleteApiToken(t *testing.T) {
	ctx := context.Background()
	apitokenId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := selfApiToken.DeleteApiToken(ctx, apitokenId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSelfAPITokenTestGetApiToken tests the behavior of the SelfAPIToken
func TestSelfAPITokenTestGetApiToken(t *testing.T) {
	ctx := context.Background()
	apitokenId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := selfApiToken.GetApiToken(ctx, apitokenId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1626875902,"id":"864f351a-1377-4ad9-83f8-72f3fe6199ba","key":"1qkb...QQCL","last_used":1690115110,"name":"org_token_xyz"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSelfAPITokenTestUpdateApiToken tests the behavior of the SelfAPIToken
func TestSelfAPITokenTestUpdateApiToken(t *testing.T) {
	ctx := context.Background()
	apitokenId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.UserApitoken
	errBody := json.Unmarshal([]byte(`{"name":"org_token_xyz"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := selfApiToken.UpdateApiToken(ctx, apitokenId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":1626875902,"id":"864f351a-1377-4ad9-83f8-72f3fe6199ba","key":"1qkb...QQCL","last_used":1690115110,"name":"org_token_xyz"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
