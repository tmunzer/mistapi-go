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

// TestSitesSkyatpTestCountSiteSkyatpEvents tests the behavior of the SitesSkyatp
func TestSitesSkyatpTestCountSiteSkyatpEvents(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SiteSkyAtpEventsCountDistinctEnum("type")

	ipAddress := "192.168.1.1"

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesSkyatp.CountSiteSkyatpEvents(ctx, siteId, &distinct, nil, nil, nil, nil, &ipAddress, nil, nil, &duration, &limit)
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

// TestSitesSkyatpTestCountSiteSkyatpEvents1 tests the behavior of the SitesSkyatp
func TestSitesSkyatpTestCountSiteSkyatpEvents1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SiteSkyAtpEventsCountDistinctEnum("type")

	ipAddress := "192.168.1.1"

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesSkyatp.CountSiteSkyatpEvents(ctx, siteId, &distinct, nil, nil, nil, nil, &ipAddress, nil, nil, &duration, &limit)
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

// TestSitesSkyatpTestSearchSiteSkyatpEvents tests the behavior of the SitesSkyatp
func TestSitesSkyatpTestSearchSiteSkyatpEvents(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	ipAddress := "192.168.1.1"
	limit := int(100)

	duration := "1d"
	sort := "timestamp"
	apiResponse, err := sitesSkyatp.SearchSiteSkyatpEvents(ctx, siteId, nil, nil, nil, nil, &ipAddress, &limit, nil, nil, &duration, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1513176951,"limit":10,"results":[{"device_mac":"658279bb1fa4","ip":"172.16.0.11","mac":"b019c66c8348","org_id":"3139f2c2-fac6-11e5-8156-0242ac110006","site_id":"70e0f468-fc13-11e5-85ad-0242ac110008","threat_level":7,"timestamp":1592524478,"type":"cc"}],"start":1512572151,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesSkyatpTestSearchSiteSkyatpEvents1 tests the behavior of the SitesSkyatp
func TestSitesSkyatpTestSearchSiteSkyatpEvents1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	ipAddress := "192.168.1.1"
	limit := int(100)

	duration := "1d"
	sort := "timestamp"
	apiResponse, err := sitesSkyatp.SearchSiteSkyatpEvents(ctx, siteId, nil, nil, nil, nil, &ipAddress, &limit, nil, nil, &duration, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1513176951,"limit":10,"results":[{"device_mac":"658279bb1fa4","ip":"172.16.0.11","mac":"b019c66c8348","org_id":"3139f2c2-fac6-11e5-8156-0242ac110006","site_id":"70e0f468-fc13-11e5-85ad-0242ac110008","threat_level":7,"timestamp":1592524478,"type":"cc"}],"start":1512572151,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
