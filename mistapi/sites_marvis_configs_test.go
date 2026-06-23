// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesMarvisConfigsTestCountSiteMarvisConfigActions tests the behavior of the SitesMarvisConfigs
func TestSitesMarvisConfigsTestCountSiteMarvisConfigActions(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := "mac"

	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesMarvisConfigs.CountSiteMarvisConfigActions(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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

// TestSitesMarvisConfigsTestCountSiteMarvisConfigActions1 tests the behavior of the SitesMarvisConfigs
func TestSitesMarvisConfigsTestCountSiteMarvisConfigActions1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := "mac"

	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesMarvisConfigs.CountSiteMarvisConfigActions(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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

// TestSitesMarvisConfigsTestSearchSiteMarvisConfigActions tests the behavior of the SitesMarvisConfigs
func TestSitesMarvisConfigsTestSearchSiteMarvisConfigActions(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesMarvisConfigs.SearchSiteMarvisConfigActions(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1775122221,"limit":10,"results":[{"admin_id":"6d617276-0000-0000-3157-000000000000","id":"05b46288-37d4-4860-9de4-1edc6e8d5363","mac":"f8c1165aba00","op":"disable_port","org_id":"174260d5-cb22-4ea8-badb-c77a89acb0a9","port_id":"ge-0/0/2","reason":"rogue_dhcp_server_detected","site_id":"437ac5f0-fc76-4a2b-87ab-d8d1e5c00405","src":"marvis","timestamp":1775028130.405962,"type":"wired","vlan_ids":[]},{"admin_id":"6d617276-0000-0000-3157-000000000000","id":"b1a81ed4-a7a2-4945-b01d-f54afe6d5cc4","mac":"f8c1165aba00","op":"add_vlans_to_port","org_id":"174260d5-cb22-4ea8-badb-c77a89acb0a9","port_id":"ge-0/0/12","reason":"missing_vlans","site_id":"437ac5f0-fc76-4a2b-87ab-d8d1e5c00405","src":"marvis","timestamp":1774866716.15723,"type":"wired","vlan_ids":[100,200]}],"start":1775118621,"total":2}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMarvisConfigsTestSearchSiteMarvisConfigActions1 tests the behavior of the SitesMarvisConfigs
func TestSitesMarvisConfigsTestSearchSiteMarvisConfigActions1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesMarvisConfigs.SearchSiteMarvisConfigActions(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1775122221,"limit":10,"results":[{"admin_id":"6d617276-0000-0000-3157-000000000000","id":"05b46288-37d4-4860-9de4-1edc6e8d5363","mac":"f8c1165aba00","op":"disable_port","org_id":"174260d5-cb22-4ea8-badb-c77a89acb0a9","port_id":"ge-0/0/2","reason":"rogue_dhcp_server_detected","site_id":"437ac5f0-fc76-4a2b-87ab-d8d1e5c00405","src":"marvis","timestamp":1775028130.405962,"type":"wired","vlan_ids":[]},{"admin_id":"6d617276-0000-0000-3157-000000000000","id":"b1a81ed4-a7a2-4945-b01d-f54afe6d5cc4","mac":"f8c1165aba00","op":"add_vlans_to_port","org_id":"174260d5-cb22-4ea8-badb-c77a89acb0a9","port_id":"ge-0/0/12","reason":"missing_vlans","site_id":"437ac5f0-fc76-4a2b-87ab-d8d1e5c00405","src":"marvis","timestamp":1774866716.15723,"type":"wired","vlan_ids":[100,200]}],"start":1775118621,"total":2}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
