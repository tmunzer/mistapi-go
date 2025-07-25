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

// TestAdminsLookupTestLookup tests the behavior of the AdminsLookup
func TestAdminsLookupTestLookup(t *testing.T) {
    ctx := context.Background()
    var body models.EmailString
    errBody := json.Unmarshal([]byte(`{"email":"test@mistsys.com"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := adminsLookup.Lookup(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"sso_url":"https://my.sso/idp_sso_url"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestAdminsLookupTestLookup1 tests the behavior of the AdminsLookup
func TestAdminsLookupTestLookup1(t *testing.T) {
    ctx := context.Background()
    var body models.EmailString
    errBody := json.Unmarshal([]byte(`{"email":"test@mistsys.com"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := adminsLookup.Lookup(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"sso_url":"https://my.sso/idp_sso_url"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
