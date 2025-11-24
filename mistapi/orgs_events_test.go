// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsEventsTestSearchOrgEvents tests the behavior of the OrgsEvents
func TestOrgsEventsTestSearchOrgEvents(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsEvents.SearchOrgEvents(ctx, orgId, nil, &limit, nil, nil, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsEventsTestSearchOrgEvents1 tests the behavior of the OrgsEvents
func TestOrgsEventsTestSearchOrgEvents1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsEvents.SearchOrgEvents(ctx, orgId, nil, &limit, nil, nil, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsEventsTestCountOrgSystemEvents tests the behavior of the OrgsEvents
func TestOrgsEventsTestCountOrgSystemEvents(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := "site_id"
	limit := int(100)

	duration := "1d"
	apiResponse, err := orgsEvents.CountOrgSystemEvents(ctx, orgId, &distinct, &limit, nil, nil, &duration)
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

// TestOrgsEventsTestCountOrgSystemEvents1 tests the behavior of the OrgsEvents
func TestOrgsEventsTestCountOrgSystemEvents1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := "site_id"
	limit := int(100)

	duration := "1d"
	apiResponse, err := orgsEvents.CountOrgSystemEvents(ctx, orgId, &distinct, &limit, nil, nil, &duration)
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

// TestOrgsEventsTestSearchOrgSystemEvents tests the behavior of the OrgsEvents
func TestOrgsEventsTestSearchOrgSystemEvents(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsEvents.SearchOrgSystemEvents(ctx, orgId, &limit, nil, nil, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1741312800,"limit":1000,"results":[{"change_cat":"admin_action","metadata":"{\"\\wlan_id\\\": \\\"None\\\",\\\"admin_name\\\": \\\"API Automation api_automation@mistsys.com\\\",\\\"desc\\\": \\\"Delete WLAN \\\\\"Automation Test\\\\\"\\\"}","org_id":"268e599f-5379-473f-b18b-4384e9b3f244","scope":"org","site_id":"dc47920f-52d5-499c-af72-dcd646764f84","timestamp":1741309621,"type":"delete-wlan"},{"change_cat":"admin_action","metadata":"{\\\"wlan_id\\\": \\\"3855dc19-63fb-4490-b113-0176dc1cc8f4\\\",\\\"admin_name\\\": \\\"API Automation api_automation@mistsys.com\\\",\\\"desc\\\": \\\"Add WLAN \\\\\"Automation Test\\\\\"\\\"}","org_id":"268e599f-5379-473f-b18b-4384e9b3f244","scope":"org","site_id":"dc47920f-52d5-499c-af72-dcd646764f84","timestamp":1741309601,"type":"add-wlan"},{"change_cat":"admin_action","metadata":"{\\\"template_id\\\": \\\"7e49acf4-6841-4e56-ad7d-68d0801cbba8\\\",\\\"admin_name\\\": \\\"API Automation api_automation@mistsys.com\\\",\\\"desc\\\": \\\"Add Template \\\\\"Automation template\\\\\"\\\"}","org_id":"268e599f-5379-473f-b18b-4384e9b3f244","scope":"org","timestamp":1741309280,"type":"add-template"}],"start":1741309200,"total":3}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsEventsTestSearchOrgSystemEvents1 tests the behavior of the OrgsEvents
func TestOrgsEventsTestSearchOrgSystemEvents1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsEvents.SearchOrgSystemEvents(ctx, orgId, &limit, nil, nil, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1741312800,"limit":1000,"results":[{"change_cat":"admin_action","metadata":"{\"\\wlan_id\\\": \\\"None\\\",\\\"admin_name\\\": \\\"API Automation api_automation@mistsys.com\\\",\\\"desc\\\": \\\"Delete WLAN \\\\\"Automation Test\\\\\"\\\"}","org_id":"268e599f-5379-473f-b18b-4384e9b3f244","scope":"org","site_id":"dc47920f-52d5-499c-af72-dcd646764f84","timestamp":1741309621,"type":"delete-wlan"},{"change_cat":"admin_action","metadata":"{\\\"wlan_id\\\": \\\"3855dc19-63fb-4490-b113-0176dc1cc8f4\\\",\\\"admin_name\\\": \\\"API Automation api_automation@mistsys.com\\\",\\\"desc\\\": \\\"Add WLAN \\\\\"Automation Test\\\\\"\\\"}","org_id":"268e599f-5379-473f-b18b-4384e9b3f244","scope":"org","site_id":"dc47920f-52d5-499c-af72-dcd646764f84","timestamp":1741309601,"type":"add-wlan"},{"change_cat":"admin_action","metadata":"{\\\"template_id\\\": \\\"7e49acf4-6841-4e56-ad7d-68d0801cbba8\\\",\\\"admin_name\\\": \\\"API Automation api_automation@mistsys.com\\\",\\\"desc\\\": \\\"Add Template \\\\\"Automation template\\\\\"\\\"}","org_id":"268e599f-5379-473f-b18b-4384e9b3f244","scope":"org","timestamp":1741309280,"type":"add-template"}],"start":1741309200,"total":3}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
