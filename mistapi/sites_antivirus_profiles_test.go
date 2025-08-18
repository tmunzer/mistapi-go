// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesAntivirusProfilesTestListSiteAntivirusProfilesDerived tests the behavior of the SitesAntivirusProfiles
func TestSitesAntivirusProfilesTestListSiteAntivirusProfilesDerived(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesAntivirusProfiles.ListSiteAntivirusProfilesDerived(ctx, siteId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAntivirusProfilesTestListSiteAntivirusProfilesDerived1 tests the behavior of the SitesAntivirusProfiles
func TestSitesAntivirusProfilesTestListSiteAntivirusProfilesDerived1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesAntivirusProfiles.ListSiteAntivirusProfilesDerived(ctx, siteId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
