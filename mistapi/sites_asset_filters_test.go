// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesAssetFiltersTestListSiteAssetFilters tests the behavior of the SitesAssetFilters
func TestSitesAssetFiltersTestListSiteAssetFilters(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesAssetFilters.ListSiteAssetFilters(ctx, siteId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesAssetFiltersTestListSiteAssetFilters1 tests the behavior of the SitesAssetFilters
func TestSitesAssetFiltersTestListSiteAssetFilters1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesAssetFilters.ListSiteAssetFilters(ctx, siteId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesAssetFiltersTestCreateSiteAssetFilter tests the behavior of the SitesAssetFilters
func TestSitesAssetFiltersTestCreateSiteAssetFilter(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesAssetFilters.CreateSiteAssetFilter(ctx, siteId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","for_site":true,"ibeacon_major":1,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mfg_company_id":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAssetFiltersTestCreateSiteAssetFilter1 tests the behavior of the SitesAssetFilters
func TestSitesAssetFiltersTestCreateSiteAssetFilter1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesAssetFilters.CreateSiteAssetFilter(ctx, siteId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","for_site":true,"ibeacon_major":1,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mfg_company_id":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAssetFiltersTestDeleteSiteAssetFilter tests the behavior of the SitesAssetFilters
func TestSitesAssetFiltersTestDeleteSiteAssetFilter(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	assetfilterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesAssetFilters.DeleteSiteAssetFilter(ctx, siteId, assetfilterId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesAssetFiltersTestGetSiteAssetFilter tests the behavior of the SitesAssetFilters
func TestSitesAssetFiltersTestGetSiteAssetFilter(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	assetfilterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesAssetFilters.GetSiteAssetFilter(ctx, siteId, assetfilterId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","for_site":true,"ibeacon_major":1,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mfg_company_id":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAssetFiltersTestGetSiteAssetFilter1 tests the behavior of the SitesAssetFilters
func TestSitesAssetFiltersTestGetSiteAssetFilter1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	assetfilterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesAssetFilters.GetSiteAssetFilter(ctx, siteId, assetfilterId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","for_site":true,"ibeacon_major":1,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mfg_company_id":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAssetFiltersTestUpdateSiteAssetFilter tests the behavior of the SitesAssetFilters
func TestSitesAssetFiltersTestUpdateSiteAssetFilter(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	assetfilterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesAssetFilters.UpdateSiteAssetFilter(ctx, siteId, assetfilterId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","for_site":true,"ibeacon_major":1,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mfg_company_id":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAssetFiltersTestUpdateSiteAssetFilter1 tests the behavior of the SitesAssetFilters
func TestSitesAssetFiltersTestUpdateSiteAssetFilter1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	assetfilterId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesAssetFilters.UpdateSiteAssetFilter(ctx, siteId, assetfilterId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"disabled":true,"eddystone_uid_namespace":"string","eddystone_url":"string","for_site":true,"ibeacon_major":1,"ibeacon_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mfg_company_id":0,"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
