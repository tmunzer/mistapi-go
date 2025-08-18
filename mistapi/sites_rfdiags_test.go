// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesRfdiagsTestGetSiteSiteRfdiagRecording tests the behavior of the SitesRfdiags
func TestSitesRfdiagsTestGetSiteSiteRfdiagRecording(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesRfdiags.GetSiteSiteRfdiagRecording(ctx, siteId, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[[{"asset_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","asset_name":"string","client_name":"string","duration":0,"end_time":0,"frame_count":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","next":"string","raw_events":"string","ready":true,"sdkclient_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sdkclient_name":"string","sdkclient_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","start_time":0,"type":"sdkclient","url":"string"}]]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRfdiagsTestGetSiteSiteRfdiagRecording1 tests the behavior of the SitesRfdiags
func TestSitesRfdiagsTestGetSiteSiteRfdiagRecording1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesRfdiags.GetSiteSiteRfdiagRecording(ctx, siteId, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[[{"asset_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","asset_name":"string","client_name":"string","duration":0,"end_time":0,"frame_count":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","next":"string","raw_events":"string","ready":true,"sdkclient_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sdkclient_name":"string","sdkclient_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","start_time":0,"type":"sdkclient","url":"string"}]]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRfdiagsTestStartSiteRecording tests the behavior of the SitesRfdiags
func TestSitesRfdiagsTestStartSiteRecording(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesRfdiags.StartSiteRecording(ctx, siteId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"asset_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","asset_name":"string","client_name":"string","duration":0,"end_time":0,"frame_count":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","next":"string","raw_events":"string","ready":true,"sdkclient_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sdkclient_name":"string","sdkclient_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","start_time":0,"type":"sdkclient","url":"string"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRfdiagsTestStartSiteRecording1 tests the behavior of the SitesRfdiags
func TestSitesRfdiagsTestStartSiteRecording1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesRfdiags.StartSiteRecording(ctx, siteId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"asset_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","asset_name":"string","client_name":"string","duration":0,"end_time":0,"frame_count":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","next":"string","raw_events":"string","ready":true,"sdkclient_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sdkclient_name":"string","sdkclient_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","start_time":0,"type":"sdkclient","url":"string"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRfdiagsTestDeleteSiteRfdiagRecording tests the behavior of the SitesRfdiags
func TestSitesRfdiagsTestDeleteSiteRfdiagRecording(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rfdiagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesRfdiags.DeleteSiteRfdiagRecording(ctx, siteId, rfdiagId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesRfdiagsTestGetSiteRfdiagRecording tests the behavior of the SitesRfdiags
func TestSitesRfdiagsTestGetSiteRfdiagRecording(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rfdiagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesRfdiags.GetSiteRfdiagRecording(ctx, siteId, rfdiagId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"asset_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","asset_name":"string","client_name":"string","duration":0,"end_time":0,"frame_count":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","next":"string","raw_events":"string","ready":true,"sdkclient_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sdkclient_name":"string","sdkclient_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","start_time":0,"type":"sdkclient","url":"string"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRfdiagsTestGetSiteRfdiagRecording1 tests the behavior of the SitesRfdiags
func TestSitesRfdiagsTestGetSiteRfdiagRecording1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rfdiagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesRfdiags.GetSiteRfdiagRecording(ctx, siteId, rfdiagId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"asset_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","asset_name":"string","client_name":"string","duration":0,"end_time":0,"frame_count":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","next":"string","raw_events":"string","ready":true,"sdkclient_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sdkclient_name":"string","sdkclient_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","start_time":0,"type":"sdkclient","url":"string"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRfdiagsTestUpdateSiteRfdiagRecording tests the behavior of the SitesRfdiags
func TestSitesRfdiagsTestUpdateSiteRfdiagRecording(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rfdiagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesRfdiags.UpdateSiteRfdiagRecording(ctx, siteId, rfdiagId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"asset_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","asset_name":"string","client_name":"string","duration":0,"end_time":0,"frame_count":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","next":"string","raw_events":"string","ready":true,"sdkclient_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sdkclient_name":"string","sdkclient_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","start_time":0,"type":"sdkclient","url":"string"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRfdiagsTestUpdateSiteRfdiagRecording1 tests the behavior of the SitesRfdiags
func TestSitesRfdiagsTestUpdateSiteRfdiagRecording1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rfdiagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesRfdiags.UpdateSiteRfdiagRecording(ctx, siteId, rfdiagId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"asset_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","asset_name":"string","client_name":"string","duration":0,"end_time":0,"frame_count":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","mac":"string","map_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","next":"string","raw_events":"string","ready":true,"sdkclient_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sdkclient_name":"string","sdkclient_uuid":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","start_time":0,"type":"sdkclient","url":"string"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesRfdiagsTestDownloadSiteRfdiagRecording tests the behavior of the SitesRfdiags
func TestSitesRfdiagsTestDownloadSiteRfdiagRecording(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rfdiagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesRfdiags.DownloadSiteRfdiagRecording(ctx, siteId, rfdiagId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesRfdiagsTestDownloadSiteRfdiagRecording1 tests the behavior of the SitesRfdiags
func TestSitesRfdiagsTestDownloadSiteRfdiagRecording1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rfdiagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesRfdiags.DownloadSiteRfdiagRecording(ctx, siteId, rfdiagId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesRfdiagsTestStopSiteRfdiagRecording tests the behavior of the SitesRfdiags
func TestSitesRfdiagsTestStopSiteRfdiagRecording(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	rfdiagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesRfdiags.StopSiteRfdiagRecording(ctx, siteId, rfdiagId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
