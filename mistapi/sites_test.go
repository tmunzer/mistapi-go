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

// TestSitesTestDeleteSite tests the behavior of the Sites
func TestSitesTestDeleteSite(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sites.DeleteSite(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesTestGetSiteInfo tests the behavior of the Sites
func TestSitesTestGetSiteInfo(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sites.GetSiteInfo(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"address":"1601 S. Deanza Blvd., Cupertino, CA, 95014","alarmtemplate_id":"684dfc5c-fe77-2290-eb1d-ef3d677fe168","apporttemplate_id":"string","aptemplate_id":"16bdf952-ade2-4491-80b0-85ce506c760b","country_code":"US","created_time":0,"gatewaytemplate_id":"6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f","id":"497f6eca-6276-5005-bfeb-53cbbbba6f17","latlng":{"lat":37.295833,"lng":-122.032946},"modified_time":0,"name":"Mist Office","networktemplate_id":"12ae9bd2-e0ab-107b-72e8-a7a005565ec2","notes":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","rftemplate_id":"bb8a9017-1e36-5d6c-6f2b-551abe8a76a2","secpolicy_id":"3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef","sitegroup_ids":["497f6eca-6276-5006-bfeb-53cbbbba6f18"],"timezone":"America/Los_Angeles"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesTestGetSiteInfo1 tests the behavior of the Sites
func TestSitesTestGetSiteInfo1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sites.GetSiteInfo(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"address":"1601 S. Deanza Blvd., Cupertino, CA, 95014","alarmtemplate_id":"684dfc5c-fe77-2290-eb1d-ef3d677fe168","apporttemplate_id":"string","aptemplate_id":"16bdf952-ade2-4491-80b0-85ce506c760b","country_code":"US","created_time":0,"gatewaytemplate_id":"6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f","id":"497f6eca-6276-5005-bfeb-53cbbbba6f17","latlng":{"lat":37.295833,"lng":-122.032946},"modified_time":0,"name":"Mist Office","networktemplate_id":"12ae9bd2-e0ab-107b-72e8-a7a005565ec2","notes":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","rftemplate_id":"bb8a9017-1e36-5d6c-6f2b-551abe8a76a2","secpolicy_id":"3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef","sitegroup_ids":["497f6eca-6276-5006-bfeb-53cbbbba6f18"],"timezone":"America/Los_Angeles"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesTestUpdateSiteInfo tests the behavior of the Sites
func TestSitesTestUpdateSiteInfo(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Site
	errBody := json.Unmarshal([]byte(`{"address":"string","alarmtemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","country_code":"string","latlng":{"lat":0,"lng":0},"name":"string","networktemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","notes":"string","rftemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","secpolicy_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"timezone":"string"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sites.UpdateSiteInfo(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"address":"1601 S. Deanza Blvd., Cupertino, CA, 95014","alarmtemplate_id":"684dfc5c-fe77-2290-eb1d-ef3d677fe168","apporttemplate_id":"string","aptemplate_id":"16bdf952-ade2-4491-80b0-85ce506c760b","country_code":"US","created_time":0,"gatewaytemplate_id":"6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f","id":"497f6eca-6276-5005-bfeb-53cbbbba6f17","latlng":{"lat":37.295833,"lng":-122.032946},"modified_time":0,"name":"Mist Office","networktemplate_id":"12ae9bd2-e0ab-107b-72e8-a7a005565ec2","notes":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","rftemplate_id":"bb8a9017-1e36-5d6c-6f2b-551abe8a76a2","secpolicy_id":"3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef","sitegroup_ids":["497f6eca-6276-5006-bfeb-53cbbbba6f18"],"timezone":"America/Los_Angeles"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesTestUpdateSiteInfo1 tests the behavior of the Sites
func TestSitesTestUpdateSiteInfo1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Site
	errBody := json.Unmarshal([]byte(`{"address":"string","alarmtemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","country_code":"string","latlng":{"lat":0,"lng":0},"name":"string","networktemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","notes":"string","rftemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","secpolicy_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"timezone":"string"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sites.UpdateSiteInfo(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"address":"1601 S. Deanza Blvd., Cupertino, CA, 95014","alarmtemplate_id":"684dfc5c-fe77-2290-eb1d-ef3d677fe168","apporttemplate_id":"string","aptemplate_id":"16bdf952-ade2-4491-80b0-85ce506c760b","country_code":"US","created_time":0,"gatewaytemplate_id":"6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f","id":"497f6eca-6276-5005-bfeb-53cbbbba6f17","latlng":{"lat":37.295833,"lng":-122.032946},"modified_time":0,"name":"Mist Office","networktemplate_id":"12ae9bd2-e0ab-107b-72e8-a7a005565ec2","notes":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","rftemplate_id":"bb8a9017-1e36-5d6c-6f2b-551abe8a76a2","secpolicy_id":"3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef","sitegroup_ids":["497f6eca-6276-5006-bfeb-53cbbbba6f18"],"timezone":"America/Los_Angeles"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
