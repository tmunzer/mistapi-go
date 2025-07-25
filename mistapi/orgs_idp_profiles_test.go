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

// TestOrgsIDPProfilesTestListOrgIdpProfiles tests the behavior of the OrgsIDPProfiles
func TestOrgsIDPProfilesTestListOrgIdpProfiles(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsIdpProfiles.ListOrgIdpProfiles(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"base_profile":"strict","created_time":0,"id":"874ca978-d736-4d4b-bc90-a49a29eec133","modified_time":0,"name":"relaxed","overwrites":[{"action":"alert","matching":{"attack_name":["HTTP:INVALID:HDR-FIELD"],"dst_subnet":["63.1.2.0/24"],"severity":["major"]}}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIDPProfilesTestListOrgIdpProfiles1 tests the behavior of the OrgsIDPProfiles
func TestOrgsIDPProfilesTestListOrgIdpProfiles1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsIdpProfiles.ListOrgIdpProfiles(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"base_profile":"strict","created_time":0,"id":"874ca978-d736-4d4b-bc90-a49a29eec133","modified_time":0,"name":"relaxed","overwrites":[{"action":"alert","matching":{"attack_name":["HTTP:INVALID:HDR-FIELD"],"dst_subnet":["63.1.2.0/24"],"severity":["major"]}}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIDPProfilesTestCreateOrgIdpProfile tests the behavior of the OrgsIDPProfiles
func TestOrgsIDPProfilesTestCreateOrgIdpProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.IdpProfile
    errBody := json.Unmarshal([]byte(`{"base_profile":"strict","name":"relaxed","overwrites":[{"action":"alert","matching":{"attack_name":["HTTP:INVALID:HDR-FIELD"],"dst_subnet":["63.1.2.0/24"],"severity":["major"]}}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsIdpProfiles.CreateOrgIdpProfile(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"base_profile":"strict","created_time":0,"id":"874ca978-d736-4d4b-bc90-a49a29eec133","modified_time":0,"name":"relaxed","overwrites":[{"action":"alert","matching":{"attack_name":["HTTP:INVALID:HDR-FIELD"],"dst_subnet":["63.1.2.0/24"],"severity":["major"]}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIDPProfilesTestCreateOrgIdpProfile1 tests the behavior of the OrgsIDPProfiles
func TestOrgsIDPProfilesTestCreateOrgIdpProfile1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.IdpProfile
    errBody := json.Unmarshal([]byte(`{"base_profile":"strict","name":"relaxed","overwrites":[{"action":"alert","matching":{"attack_name":["HTTP:INVALID:HDR-FIELD"],"dst_subnet":["63.1.2.0/24"],"severity":["major"]}}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsIdpProfiles.CreateOrgIdpProfile(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"base_profile":"strict","created_time":0,"id":"874ca978-d736-4d4b-bc90-a49a29eec133","modified_time":0,"name":"relaxed","overwrites":[{"action":"alert","matching":{"attack_name":["HTTP:INVALID:HDR-FIELD"],"dst_subnet":["63.1.2.0/24"],"severity":["major"]}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIDPProfilesTestDeleteOrgIdpProfile tests the behavior of the OrgsIDPProfiles
func TestOrgsIDPProfilesTestDeleteOrgIdpProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    idpprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsIdpProfiles.DeleteOrgIdpProfile(ctx, orgId, idpprofileId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsIDPProfilesTestGetOrgIdpProfile tests the behavior of the OrgsIDPProfiles
func TestOrgsIDPProfilesTestGetOrgIdpProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    idpprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsIdpProfiles.GetOrgIdpProfile(ctx, orgId, idpprofileId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"base_profile":"strict","created_time":0,"id":"874ca978-d736-4d4b-bc90-a49a29eec133","modified_time":0,"name":"relaxed","overwrites":[{"action":"alert","matching":{"attack_name":["HTTP:INVALID:HDR-FIELD"],"dst_subnet":["63.1.2.0/24"],"severity":["major"]}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIDPProfilesTestGetOrgIdpProfile1 tests the behavior of the OrgsIDPProfiles
func TestOrgsIDPProfilesTestGetOrgIdpProfile1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    idpprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsIdpProfiles.GetOrgIdpProfile(ctx, orgId, idpprofileId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"base_profile":"strict","created_time":0,"id":"874ca978-d736-4d4b-bc90-a49a29eec133","modified_time":0,"name":"relaxed","overwrites":[{"action":"alert","matching":{"attack_name":["HTTP:INVALID:HDR-FIELD"],"dst_subnet":["63.1.2.0/24"],"severity":["major"]}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIDPProfilesTestUpdateOrgIdpProfile tests the behavior of the OrgsIDPProfiles
func TestOrgsIDPProfilesTestUpdateOrgIdpProfile(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    idpprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.IdpProfile
    errBody := json.Unmarshal([]byte(`{"base_profile":"strict","name":"relaxed","overwrites":[{"action":"alert","matching":{"attack_name":["HTTP:INVALID:HDR-FIELD"],"dst_subnet":["63.1.2.0/24"],"severity":["major"]}}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsIdpProfiles.UpdateOrgIdpProfile(ctx, orgId, idpprofileId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"base_profile":"strict","created_time":0,"id":"874ca978-d736-4d4b-bc90-a49a29eec133","modified_time":0,"name":"relaxed","overwrites":[{"action":"alert","matching":{"attack_name":["HTTP:INVALID:HDR-FIELD"],"dst_subnet":["63.1.2.0/24"],"severity":["major"]}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsIDPProfilesTestUpdateOrgIdpProfile1 tests the behavior of the OrgsIDPProfiles
func TestOrgsIDPProfilesTestUpdateOrgIdpProfile1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    idpprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.IdpProfile
    errBody := json.Unmarshal([]byte(`{"base_profile":"strict","name":"relaxed","overwrites":[{"action":"alert","matching":{"attack_name":["HTTP:INVALID:HDR-FIELD"],"dst_subnet":["63.1.2.0/24"],"severity":["major"]}}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsIdpProfiles.UpdateOrgIdpProfile(ctx, orgId, idpprofileId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"base_profile":"strict","created_time":0,"id":"874ca978-d736-4d4b-bc90-a49a29eec133","modified_time":0,"name":"relaxed","overwrites":[{"action":"alert","matching":{"attack_name":["HTTP:INVALID:HDR-FIELD"],"dst_subnet":["63.1.2.0/24"],"severity":["major"]}}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
