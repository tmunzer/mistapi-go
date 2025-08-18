// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesSiteTemplatesTestListSiteSiteTemplatesDerived tests the behavior of the SitesSiteTemplates
func TestSitesSiteTemplatesTestListSiteSiteTemplatesDerived(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesSiteTemplates.ListSiteSiteTemplatesDerived(ctx, siteId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"auto_upgrade":{"day_of_week":"mon","enabled":true,"time_of_day":"string","version":"string"},"name":"string","vars":{"SSID_STR":"string","VLAN_ID":"string"}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesSiteTemplatesTestListSiteSiteTemplatesDerived1 tests the behavior of the SitesSiteTemplates
func TestSitesSiteTemplatesTestListSiteSiteTemplatesDerived1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesSiteTemplates.ListSiteSiteTemplatesDerived(ctx, siteId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"auto_upgrade":{"day_of_week":"mon","enabled":true,"time_of_day":"string","version":"string"},"name":"string","vars":{"SSID_STR":"string","VLAN_ID":"string"}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
