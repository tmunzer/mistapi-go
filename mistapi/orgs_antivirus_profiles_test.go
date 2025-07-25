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

// TestOrgsAntivirusProfilesTestListOrgAntivirusProfiles tests the behavior of the OrgsAntivirusProfiles
func TestOrgsAntivirusProfilesTestListOrgAntivirusProfiles(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsAntivirusProfiles.ListOrgAntivirusProfiles(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAntivirusProfilesTestListOrgAntivirusProfiles1 tests the behavior of the OrgsAntivirusProfiles
func TestOrgsAntivirusProfilesTestListOrgAntivirusProfiles1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsAntivirusProfiles.ListOrgAntivirusProfiles(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAntivirusProfilesTestCreateOrgAntivirusProfile tests the behavior of the OrgsAntivirusProfiles
func TestOrgsAntivirusProfilesTestCreateOrgAntivirusProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Avprofile
    errBody := json.Unmarshal([]byte(`{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsAntivirusProfiles.CreateOrgAntivirusProfile(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAntivirusProfilesTestCreateOrgAntivirusProfile1 tests the behavior of the OrgsAntivirusProfiles
func TestOrgsAntivirusProfilesTestCreateOrgAntivirusProfile1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Avprofile
    errBody := json.Unmarshal([]byte(`{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsAntivirusProfiles.CreateOrgAntivirusProfile(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAntivirusProfilesTestDeleteOrgAntivirusProfile tests the behavior of the OrgsAntivirusProfiles
func TestOrgsAntivirusProfilesTestDeleteOrgAntivirusProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    avprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsAntivirusProfiles.DeleteOrgAntivirusProfile(ctx, orgId, avprofileId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAntivirusProfilesTestGetOrgAntivirusProfile tests the behavior of the OrgsAntivirusProfiles
func TestOrgsAntivirusProfilesTestGetOrgAntivirusProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    avprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsAntivirusProfiles.GetOrgAntivirusProfile(ctx, orgId, avprofileId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAntivirusProfilesTestGetOrgAntivirusProfile1 tests the behavior of the OrgsAntivirusProfiles
func TestOrgsAntivirusProfilesTestGetOrgAntivirusProfile1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    avprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsAntivirusProfiles.GetOrgAntivirusProfile(ctx, orgId, avprofileId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAntivirusProfilesTestUpdateOrgAntivirusProfile tests the behavior of the OrgsAntivirusProfiles
func TestOrgsAntivirusProfilesTestUpdateOrgAntivirusProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    avprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Avprofile
    errBody := json.Unmarshal([]byte(`{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsAntivirusProfiles.UpdateOrgAntivirusProfile(ctx, orgId, avprofileId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAntivirusProfilesTestUpdateOrgAntivirusProfile1 tests the behavior of the OrgsAntivirusProfiles
func TestOrgsAntivirusProfilesTestUpdateOrgAntivirusProfile1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    avprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Avprofile
    errBody := json.Unmarshal([]byte(`{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsAntivirusProfiles.UpdateOrgAntivirusProfile(ctx, orgId, avprofileId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"fallback_action":"permit","max_filesize":10000,"mime_whitelist":[],"name":"av-custom","protocols":["http"],"url_whitelist":[]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
