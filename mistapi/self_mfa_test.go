// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSelfMFATestGenerateSecretFor2faVerification tests the behavior of the SelfMFA
func TestSelfMFATestGenerateSecretFor2faVerification(t *testing.T) {
	ctx := context.Background()
	by := models.MfaSecretTypeEnum("qrcode")
	apiResponse, err := selfMfa.GenerateSecretFor2faVerification(ctx, &by)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"two_factor_secret":"NRMTSTRWNBVECY3GJVYEY3DDJFRGSNCZGJUDO4RVN5FDM3DUMJSA"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSelfMFATestGenerateSecretFor2faVerification1 tests the behavior of the SelfMFA
func TestSelfMFATestGenerateSecretFor2faVerification1(t *testing.T) {
	ctx := context.Background()
	by := models.MfaSecretTypeEnum("qrcode")
	apiResponse, err := selfMfa.GenerateSecretFor2faVerification(ctx, &by)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/octet-stream"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSelfMFATestGenerateSecretFor2faVerification2 tests the behavior of the SelfMFA
func TestSelfMFATestGenerateSecretFor2faVerification2(t *testing.T) {
	ctx := context.Background()
	by := models.MfaSecretTypeEnum("qrcode")
	apiResponse, err := selfMfa.GenerateSecretFor2faVerification(ctx, &by)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"two_factor_secret":"NRMTSTRWNBVECY3GJVYEY3DDJFRGSNCZGJUDO4RVN5FDM3DUMJSA"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSelfMFATestVerifyTwoFactor tests the behavior of the SelfMFA
func TestSelfMFATestVerifyTwoFactor(t *testing.T) {
	ctx := context.Background()

	resp, err := selfMfa.VerifyTwoFactor(ctx, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
