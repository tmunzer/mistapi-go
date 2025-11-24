// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSitesWANUsagesTestCountSiteWanUsage tests the behavior of the SitesWANUsages
func TestSitesWANUsagesTestCountSiteWanUsage(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mac := "001122334455"
	peerMac := "001122334455"
	portId := "ge-0/0/0"
	peerPortId := "ge-0/0/0"

	pathType := "primary"
	distinct := models.WanUsagesCountDistinctEnum("policy")

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesWanUsages.CountSiteWanUsage(ctx, siteId, &mac, &peerMac, &portId, &peerPortId, nil, nil, &pathType, &distinct, nil, nil, &duration, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWANUsagesTestCountSiteWanUsage1 tests the behavior of the SitesWANUsages
func TestSitesWANUsagesTestCountSiteWanUsage1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mac := "001122334455"
	peerMac := "001122334455"
	portId := "ge-0/0/0"
	peerPortId := "ge-0/0/0"

	pathType := "primary"
	distinct := models.WanUsagesCountDistinctEnum("policy")

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesWanUsages.CountSiteWanUsage(ctx, siteId, &mac, &peerMac, &portId, &peerPortId, nil, nil, &pathType, &distinct, nil, nil, &duration, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWANUsagesTestSearchSiteWanUsage tests the behavior of the SitesWANUsages
func TestSitesWANUsagesTestSearchSiteWanUsage(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mac := "001122334455"
	peerMac := "001122334455"
	portId := "ge-0/0/0"
	peerPortId := "ge-0/0/0"
	policy := "primary"

	pathType := "primary"
	limit := int(100)

	duration := "1d"
	sort := "timestamp"
	apiResponse, err := sitesWanUsages.SearchSiteWanUsage(ctx, siteId, &mac, &peerMac, &portId, &peerPortId, &policy, nil, &pathType, &limit, nil, nil, &duration, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesWANUsagesTestSearchSiteWanUsage1 tests the behavior of the SitesWANUsages
func TestSitesWANUsagesTestSearchSiteWanUsage1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mac := "001122334455"
	peerMac := "001122334455"
	portId := "ge-0/0/0"
	peerPortId := "ge-0/0/0"
	policy := "primary"

	pathType := "primary"
	limit := int(100)

	duration := "1d"
	sort := "timestamp"
	apiResponse, err := sitesWanUsages.SearchSiteWanUsage(ctx, siteId, &mac, &peerMac, &portId, &peerPortId, &policy, nil, &pathType, &limit, nil, nil, &duration, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
