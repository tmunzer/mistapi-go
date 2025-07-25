// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestSitesLicensesTestGetSiteLicenseUsage tests the behavior of the SitesLicenses
func TestSitesLicensesTestGetSiteLicenseUsage(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesLicenses.GetSiteLicenseUsage(ctx, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"org_entitled":{"SUB-LOC":30,"SUB-MAN":60},"svna_enabled":true,"trial_enabled":true,"usages":{"SUB-LOC":30,"SUB-MAN":60},"vna_eligible":true,"vna_ui":true,"wvna_eligible":true}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesLicensesTestGetSiteLicenseUsage1 tests the behavior of the SitesLicenses
func TestSitesLicensesTestGetSiteLicenseUsage1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesLicenses.GetSiteLicenseUsage(ctx, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"org_entitled":{"SUB-LOC":30,"SUB-MAN":60},"svna_enabled":true,"trial_enabled":true,"usages":{"SUB-LOC":30,"SUB-MAN":60},"vna_eligible":true,"vna_ui":true,"wvna_eligible":true}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
