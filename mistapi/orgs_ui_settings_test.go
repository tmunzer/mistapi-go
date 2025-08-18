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

// TestOrgsUISettingsTestListOrgUiSettings tests the behavior of the OrgsUISettings
func TestOrgsUISettingsTestListOrgUiSettings(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsUiSettings.ListOrgUiSettings(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":1749083436,"description":"AP related stats","for_site":false,"id":"9a702097-0dd3-48af-909b-2be4ff94d139","isCustomDataboard":true,"modified_time":1749083436,"name":"AP Stats","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","purpose":"marvisdashboard","site_id":"00000000-0000-0000-0000-000000000000","tiles":[{"description":"User typed tile descr","id":"3eef7c83-3d33-417a-a729-4772d4a1013a","isAutoTitle":true,"name":"List top 10 APs by bandwidth","nl_query":"List top 10 APs by bandwidth","position":{"col":1,"colSpan":2,"row":1,"rowSpan":1}}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUISettingsTestListOrgUiSettings1 tests the behavior of the OrgsUISettings
func TestOrgsUISettingsTestListOrgUiSettings1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsUiSettings.ListOrgUiSettings(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":1749083436,"description":"AP related stats","for_site":false,"id":"9a702097-0dd3-48af-909b-2be4ff94d139","isCustomDataboard":true,"modified_time":1749083436,"name":"AP Stats","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","purpose":"marvisdashboard","site_id":"00000000-0000-0000-0000-000000000000","tiles":[{"description":"User typed tile descr","id":"3eef7c83-3d33-417a-a729-4772d4a1013a","isAutoTitle":true,"name":"List top 10 APs by bandwidth","nl_query":"List top 10 APs by bandwidth","position":{"col":1,"colSpan":2,"row":1,"rowSpan":1}}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUISettingsTestCreateOrgUiSettings tests the behavior of the OrgsUISettings
func TestOrgsUISettingsTestCreateOrgUiSettings(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.OrgUiSettings
    errBody := json.Unmarshal([]byte(`{"description":"AP related stats","isCustomDataboard":true,"name":"AP Stats","purpose":"marvisdashboard","tiles":[{"description":"User typed tile descr","isAutoTitle":true,"name":"List top 10 APs by bandwidth","nl_query":"List top 10 APs by bandwidth","position":{"col":1,"colSpan":2,"row":1,"rowSpan":1}}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsUiSettings.CreateOrgUiSettings(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":1749083436,"description":"AP related stats","for_site":false,"id":"9a702097-0dd3-48af-909b-2be4ff94d139","isCustomDataboard":true,"modified_time":1749083436,"name":"AP Stats","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","purpose":"marvisdashboard","site_id":"00000000-0000-0000-0000-000000000000","tiles":[{"description":"User typed tile descr","id":"3eef7c83-3d33-417a-a729-4772d4a1013a","isAutoTitle":true,"name":"List top 10 APs by bandwidth","nl_query":"List top 10 APs by bandwidth","position":{"col":1,"colSpan":2,"row":1,"rowSpan":1}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUISettingsTestCreateOrgUiSettings1 tests the behavior of the OrgsUISettings
func TestOrgsUISettingsTestCreateOrgUiSettings1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.OrgUiSettings
    errBody := json.Unmarshal([]byte(`{"description":"AP related stats","isCustomDataboard":true,"name":"AP Stats","purpose":"marvisdashboard","tiles":[{"description":"User typed tile descr","isAutoTitle":true,"name":"List top 10 APs by bandwidth","nl_query":"List top 10 APs by bandwidth","position":{"col":1,"colSpan":2,"row":1,"rowSpan":1}}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsUiSettings.CreateOrgUiSettings(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":1749083436,"description":"AP related stats","for_site":false,"id":"9a702097-0dd3-48af-909b-2be4ff94d139","isCustomDataboard":true,"modified_time":1749083436,"name":"AP Stats","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","purpose":"marvisdashboard","site_id":"00000000-0000-0000-0000-000000000000","tiles":[{"description":"User typed tile descr","id":"3eef7c83-3d33-417a-a729-4772d4a1013a","isAutoTitle":true,"name":"List top 10 APs by bandwidth","nl_query":"List top 10 APs by bandwidth","position":{"col":1,"colSpan":2,"row":1,"rowSpan":1}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUISettingsTestDeleteOrgUiSetting tests the behavior of the OrgsUISettings
func TestOrgsUISettingsTestDeleteOrgUiSetting(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    uisettingId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsUiSettings.DeleteOrgUiSetting(ctx, orgId, uisettingId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsUISettingsTestGetOrgUiSetting tests the behavior of the OrgsUISettings
func TestOrgsUISettingsTestGetOrgUiSetting(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    uisettingId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsUiSettings.GetOrgUiSetting(ctx, orgId, uisettingId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":1749083436,"description":"AP related stats","for_site":false,"id":"9a702097-0dd3-48af-909b-2be4ff94d139","isCustomDataboard":true,"modified_time":1749083436,"name":"AP Stats","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","purpose":"marvisdashboard","site_id":"00000000-0000-0000-0000-000000000000","tiles":[{"description":"User typed tile descr","id":"3eef7c83-3d33-417a-a729-4772d4a1013a","isAutoTitle":true,"name":"List top 10 APs by bandwidth","nl_query":"List top 10 APs by bandwidth","position":{"col":1,"colSpan":2,"row":1,"rowSpan":1}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUISettingsTestGetOrgUiSetting1 tests the behavior of the OrgsUISettings
func TestOrgsUISettingsTestGetOrgUiSetting1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    uisettingId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsUiSettings.GetOrgUiSetting(ctx, orgId, uisettingId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":1749083436,"description":"AP related stats","for_site":false,"id":"9a702097-0dd3-48af-909b-2be4ff94d139","isCustomDataboard":true,"modified_time":1749083436,"name":"AP Stats","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","purpose":"marvisdashboard","site_id":"00000000-0000-0000-0000-000000000000","tiles":[{"description":"User typed tile descr","id":"3eef7c83-3d33-417a-a729-4772d4a1013a","isAutoTitle":true,"name":"List top 10 APs by bandwidth","nl_query":"List top 10 APs by bandwidth","position":{"col":1,"colSpan":2,"row":1,"rowSpan":1}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUISettingsTestUpdateOrgUiSetting tests the behavior of the OrgsUISettings
func TestOrgsUISettingsTestUpdateOrgUiSetting(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    uisettingId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.OrgUiSettings
    errBody := json.Unmarshal([]byte(`{"description":"AP related stats","isCustomDataboard":true,"name":"AP Stats","purpose":"marvisdashboard","tiles":[{"description":"User typed tile descr","isAutoTitle":true,"name":"List top 10 APs by bandwidth","nl_query":"List top 10 APs by bandwidth","position":{"col":1,"colSpan":2,"row":1,"rowSpan":1}}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsUiSettings.UpdateOrgUiSetting(ctx, orgId, uisettingId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":1749083436,"description":"AP related stats","for_site":false,"id":"9a702097-0dd3-48af-909b-2be4ff94d139","isCustomDataboard":true,"modified_time":1749083436,"name":"AP Stats","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","purpose":"marvisdashboard","site_id":"00000000-0000-0000-0000-000000000000","tiles":[{"description":"User typed tile descr","id":"3eef7c83-3d33-417a-a729-4772d4a1013a","isAutoTitle":true,"name":"List top 10 APs by bandwidth","nl_query":"List top 10 APs by bandwidth","position":{"col":1,"colSpan":2,"row":1,"rowSpan":1}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUISettingsTestUpdateOrgUiSetting1 tests the behavior of the OrgsUISettings
func TestOrgsUISettingsTestUpdateOrgUiSetting1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    uisettingId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.OrgUiSettings
    errBody := json.Unmarshal([]byte(`{"description":"AP related stats","isCustomDataboard":true,"name":"AP Stats","purpose":"marvisdashboard","tiles":[{"description":"User typed tile descr","isAutoTitle":true,"name":"List top 10 APs by bandwidth","nl_query":"List top 10 APs by bandwidth","position":{"col":1,"colSpan":2,"row":1,"rowSpan":1}}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsUiSettings.UpdateOrgUiSetting(ctx, orgId, uisettingId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":1749083436,"description":"AP related stats","for_site":false,"id":"9a702097-0dd3-48af-909b-2be4ff94d139","isCustomDataboard":true,"modified_time":1749083436,"name":"AP Stats","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","purpose":"marvisdashboard","site_id":"00000000-0000-0000-0000-000000000000","tiles":[{"description":"User typed tile descr","id":"3eef7c83-3d33-417a-a729-4772d4a1013a","isAutoTitle":true,"name":"List top 10 APs by bandwidth","nl_query":"List top 10 APs by bandwidth","position":{"col":1,"colSpan":2,"row":1,"rowSpan":1}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
