// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSitesRSSIZonesTestListSiteRssiZones tests the behavior of the SitesRSSIZones
func TestSitesRSSIZonesTestListSiteRssiZones(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesRssiZones.ListSiteRssiZones(ctx, siteId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesRSSIZonesTestCreateSiteRssiZone tests the behavior of the SitesRSSIZones
func TestSitesRSSIZonesTestCreateSiteRssiZone(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.RssiZone
	errBody := json.Unmarshal([]byte(`{"devices":[{"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","rssi":0}],"name":"string"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesRssiZones.CreateSiteRssiZone(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"devices":[{"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","rssi":0}],"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRSSIZonesTestDeleteSiteRssiZone tests the behavior of the SitesRSSIZones
func TestSitesRSSIZonesTestDeleteSiteRssiZone(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rssizoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesRssiZones.DeleteSiteRssiZone(ctx, siteId, rssizoneId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesRSSIZonesTestGetSiteRssiZone tests the behavior of the SitesRSSIZones
func TestSitesRSSIZonesTestGetSiteRssiZone(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rssizoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesRssiZones.GetSiteRssiZone(ctx, siteId, rssizoneId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":0,"devices":[{"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","rssi":0}],"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRSSIZonesTestGetSiteRssiZone1 tests the behavior of the SitesRSSIZones
func TestSitesRSSIZonesTestGetSiteRssiZone1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rssizoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesRssiZones.GetSiteRssiZone(ctx, siteId, rssizoneId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":0,"devices":[{"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","rssi":0}],"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRSSIZonesTestUpdateSiteRssiZone tests the behavior of the SitesRSSIZones
func TestSitesRSSIZonesTestUpdateSiteRssiZone(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rssizoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.RssiZone
	errBody := json.Unmarshal([]byte(`{"devices":[{"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","rssi":0}],"name":"string"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesRssiZones.UpdateSiteRssiZone(ctx, siteId, rssizoneId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"devices":[{"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","rssi":0}],"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRSSIZonesTestUpdateSiteRssiZone1 tests the behavior of the SitesRSSIZones
func TestSitesRSSIZonesTestUpdateSiteRssiZone1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rssizoneId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.RssiZone
	errBody := json.Unmarshal([]byte(`{"devices":[{"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","rssi":0}],"name":"string"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesRssiZones.UpdateSiteRssiZone(ctx, siteId, rssizoneId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"devices":[{"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","rssi":0}],"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
