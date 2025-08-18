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

// TestSitesWxTagsTestListSiteWxTags tests the behavior of the SitesWxTags
func TestSitesWxTagsTestListSiteWxTags(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesWxTags.ListSiteWxTags(ctx, siteId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxTagsTestListSiteWxTags1 tests the behavior of the SitesWxTags
func TestSitesWxTagsTestListSiteWxTags1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesWxTags.ListSiteWxTags(ctx, siteId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxTagsTestCreateSiteWxTag tests the behavior of the SitesWxTags
func TestSitesWxTagsTestCreateSiteWxTag(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.WxlanTag
	errBody := json.Unmarshal([]byte(`{"match":"app","name":"match app","type":"match","values":["gmail","dropbox"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesWxTags.CreateSiteWxTag(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxTagsTestCreateSiteWxTag1 tests the behavior of the SitesWxTags
func TestSitesWxTagsTestCreateSiteWxTag1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.WxlanTag
	errBody := json.Unmarshal([]byte(`{"match":"app","name":"match app","type":"match","values":["gmail","dropbox"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesWxTags.CreateSiteWxTag(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxTagsTestGetSiteApplicationList tests the behavior of the SitesWxTags
func TestSitesWxTagsTestGetSiteApplicationList(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesWxTags.GetSiteApplicationList(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"group":"Emails","key":"gmail","name":"Gmail - web/app"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxTagsTestGetSiteApplicationList1 tests the behavior of the SitesWxTags
func TestSitesWxTagsTestGetSiteApplicationList1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesWxTags.GetSiteApplicationList(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"group":"Emails","key":"gmail","name":"Gmail - web/app"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxTagsTestDeleteSiteWxTag tests the behavior of the SitesWxTags
func TestSitesWxTagsTestDeleteSiteWxTag(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesWxTags.DeleteSiteWxTag(ctx, siteId, wxtagId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesWxTagsTestGetSiteWxTag tests the behavior of the SitesWxTags
func TestSitesWxTagsTestGetSiteWxTag(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesWxTags.GetSiteWxTag(ctx, siteId, wxtagId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxTagsTestGetSiteWxTag1 tests the behavior of the SitesWxTags
func TestSitesWxTagsTestGetSiteWxTag1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesWxTags.GetSiteWxTag(ctx, siteId, wxtagId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxTagsTestUpdateSiteWxTag tests the behavior of the SitesWxTags
func TestSitesWxTagsTestUpdateSiteWxTag(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.WxlanTag
	errBody := json.Unmarshal([]byte(`{"match":"app","name":"match app","type":"match","values":["gmail","dropbox"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesWxTags.UpdateSiteWxTag(ctx, siteId, wxtagId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxTagsTestUpdateSiteWxTag1 tests the behavior of the SitesWxTags
func TestSitesWxTagsTestUpdateSiteWxTag1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtagId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.WxlanTag
	errBody := json.Unmarshal([]byte(`{"match":"app","name":"match app","type":"match","values":["gmail","dropbox"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesWxTags.UpdateSiteWxTag(ctx, siteId, wxtagId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","last_ips":["string"],"mac":"string","match":"wlan_id","modified_time":0,"name":"string","op":"in","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","resource_mac":"string","services":["string"],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","specs":[{"port_range":"string","protocol":"tcp","subnet":["string"]}],"subnet":"string","type":"match","values":["string"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
