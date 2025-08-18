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

// TestOrgsSCEPTestDisableOrgMistScep tests the behavior of the OrgsSCEP
func TestOrgsSCEPTestDisableOrgMistScep(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsScep.DisableOrgMistScep(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cert_providers":["jamf","intune","byod"],"enabled":false,"intune_scep_url":"https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep","jamf_access_token":"1Z4QqEnCt05Jjt3TV5LgPJ4V_WL_RWnJ7dqVMLYHj81=","jamf_scep_url":"https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep","jamf_webhook_url":"https://scep.mistsys.com/api/v1/webhook/jamf/:org_id/scep"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSCEPTestDisableOrgMistScep1 tests the behavior of the OrgsSCEP
func TestOrgsSCEPTestDisableOrgMistScep1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsScep.DisableOrgMistScep(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cert_providers":["jamf","intune","byod"],"enabled":false,"intune_scep_url":"https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep","jamf_access_token":"1Z4QqEnCt05Jjt3TV5LgPJ4V_WL_RWnJ7dqVMLYHj81=","jamf_scep_url":"https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep","jamf_webhook_url":"https://scep.mistsys.com/api/v1/webhook/jamf/:org_id/scep"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSCEPTestGetOrgMistScep tests the behavior of the OrgsSCEP
func TestOrgsSCEPTestGetOrgMistScep(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsScep.GetOrgMistScep(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cert_providers":["jamf","intune","byod"],"enabled":false,"intune_scep_url":"https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep","jamf_access_token":"1Z4QqEnCt05Jjt3TV5LgPJ4V_WL_RWnJ7dqVMLYHj81=","jamf_scep_url":"https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep","jamf_webhook_url":"https://scep.mistsys.com/api/v1/webhook/jamf/:org_id/scep"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSCEPTestGetOrgMistScep1 tests the behavior of the OrgsSCEP
func TestOrgsSCEPTestGetOrgMistScep1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsScep.GetOrgMistScep(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cert_providers":["jamf","intune","byod"],"enabled":false,"intune_scep_url":"https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep","jamf_access_token":"1Z4QqEnCt05Jjt3TV5LgPJ4V_WL_RWnJ7dqVMLYHj81=","jamf_scep_url":"https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep","jamf_webhook_url":"https://scep.mistsys.com/api/v1/webhook/jamf/:org_id/scep"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSCEPTestUpdateOrgMistScep tests the behavior of the OrgsSCEP
func TestOrgsSCEPTestUpdateOrgMistScep(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.OrgSettingScep
    errBody := json.Unmarshal([]byte(`{"enabled":true}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsScep.UpdateOrgMistScep(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cert_providers":["jamf","intune","byod"],"enabled":false,"intune_scep_url":"https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep","jamf_access_token":"1Z4QqEnCt05Jjt3TV5LgPJ4V_WL_RWnJ7dqVMLYHj81=","jamf_scep_url":"https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep","jamf_webhook_url":"https://scep.mistsys.com/api/v1/webhook/jamf/:org_id/scep"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSCEPTestUpdateOrgMistScep1 tests the behavior of the OrgsSCEP
func TestOrgsSCEPTestUpdateOrgMistScep1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.OrgSettingScep
    errBody := json.Unmarshal([]byte(`{"enabled":true}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsScep.UpdateOrgMistScep(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cert_providers":["jamf","intune","byod"],"enabled":false,"intune_scep_url":"https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep","jamf_access_token":"1Z4QqEnCt05Jjt3TV5LgPJ4V_WL_RWnJ7dqVMLYHj81=","jamf_scep_url":"https://scep.mistsys.com/api/v1/incoming/intune/:org_id/scep","jamf_webhook_url":"https://scep.mistsys.com/api/v1/webhook/jamf/:org_id/scep"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSCEPTestListOrgIssuedClientCertificates tests the behavior of the OrgsSCEP
func TestOrgsSCEPTestListOrgIssuedClientCertificates(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    apiResponse, err := orgsScep.ListOrgIssuedClientCertificates(ctx, orgId, nil, nil, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"results":[{"common_name":"john@corp.com","created_time":1431382121,"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","expire_time":1718921115,"serial_number":"13 00 13 03 23 EE D5 84 01"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSCEPTestListOrgIssuedClientCertificates1 tests the behavior of the OrgsSCEP
func TestOrgsSCEPTestListOrgIssuedClientCertificates1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    apiResponse, err := orgsScep.ListOrgIssuedClientCertificates(ctx, orgId, nil, nil, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"results":[{"common_name":"john@corp.com","created_time":1431382121,"device_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","expire_time":1718921115,"serial_number":"13 00 13 03 23 EE D5 84 01"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSCEPTestRevokeOrgIssuedClientCertificates tests the behavior of the OrgsSCEP
func TestOrgsSCEPTestRevokeOrgIssuedClientCertificates(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.ClientCertSerialNumbers
    errBody := json.Unmarshal([]byte(`{"serial_numbers":["13 00 13 03 23 EE D5 84 01"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsScep.RevokeOrgIssuedClientCertificates(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
