// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestAdminsLoginOAuth2TestUnlinkOauth2Provider tests the behavior of the AdminsLoginOAuth2
func TestAdminsLoginOAuth2TestUnlinkOauth2Provider(t *testing.T) {
    ctx := context.Background()
    provider := "google"
    resp, err := adminsLoginOAuth2.UnlinkOauth2Provider(ctx, provider)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestAdminsLoginOAuth2TestGetOauth2AuthorizationUrlForLogin tests the behavior of the AdminsLoginOAuth2
func TestAdminsLoginOAuth2TestGetOauth2AuthorizationUrlForLogin(t *testing.T) {
    ctx := context.Background()
    provider := "google"
    forward := "https://manage.mist.com/oauth/callback.html"
    apiResponse, err := adminsLoginOAuth2.GetOauth2AuthorizationUrlForLogin(ctx, provider, &forward)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"authorization_url":"https://accounts.google.com/o/oauth2/v2/auth?.....","client_id":"173131512-mpbnju32.apps.googleusercontent.com"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestAdminsLoginOAuth2TestGetOauth2AuthorizationUrlForLogin1 tests the behavior of the AdminsLoginOAuth2
func TestAdminsLoginOAuth2TestGetOauth2AuthorizationUrlForLogin1(t *testing.T) {
    ctx := context.Background()
    provider := "google"
    forward := "https://manage.mist.com/oauth/callback.html"
    apiResponse, err := adminsLoginOAuth2.GetOauth2AuthorizationUrlForLogin(ctx, provider, &forward)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"authorization_url":"https://accounts.google.com/o/oauth2/v2/auth?.....","client_id":"173131512-mpbnju32.apps.googleusercontent.com"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestAdminsLoginOAuth2TestLoginOauth2 tests the behavior of the AdminsLoginOAuth2
func TestAdminsLoginOAuth2TestLoginOauth2(t *testing.T) {
    ctx := context.Background()
    provider := "google"
    var body models.CodeString
    errBody := json.Unmarshal([]byte(`{"code":"4/S9tegDeLkrYg0L9pWNXV4cgMVbbr3SR9t693A2kSHzw"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := adminsLoginOAuth2.LoginOauth2(ctx, provider, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
