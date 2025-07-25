// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsUserMACsTestCreateOrgUserMac tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestCreateOrgUserMac(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsUserMaCs.CreateOrgUserMac(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"111cafd2-ba1b-5169-bfcb-9cdf1d473ddb","labels":["flor1","bld4"],"mac":"921b638445cd","notes":"mac address refers to Canon printers","vlan":"30"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUserMACsTestCreateOrgUserMac1 tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestCreateOrgUserMac1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsUserMaCs.CreateOrgUserMac(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"111cafd2-ba1b-5169-bfcb-9cdf1d473ddb","labels":["flor1","bld4"],"mac":"921b638445cd","notes":"mac address refers to Canon printers","vlan":"30"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUserMACsTestUpdateOrgMultipleUserMacs tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestUpdateOrgMultipleUserMacs(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsUserMaCs.UpdateOrgMultipleUserMacs(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"errors":["2feacc8e-5893-418a-acaa-4d7c1afd01fe - invalid id"],"updated":["1041c16c-ca87-4d3f-bb94-b97c5819fc09","a016cc8e-5893-418a-acaa-4d7c1af6ac0f"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUserMACsTestUpdateOrgMultipleUserMacs1 tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestUpdateOrgMultipleUserMacs1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsUserMaCs.UpdateOrgMultipleUserMacs(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"errors":["2feacc8e-5893-418a-acaa-4d7c1afd01fe - invalid id"],"updated":["1041c16c-ca87-4d3f-bb94-b97c5819fc09","a016cc8e-5893-418a-acaa-4d7c1af6ac0f"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUserMACsTestDeleteOrgMultipleUserMacs tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestDeleteOrgMultipleUserMacs(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    resp, err := orgsUserMaCs.DeleteOrgMultipleUserMacs(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsUserMACsTestSearchOrgUserMacs tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestSearchOrgUserMacs(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsUserMaCs.SearchOrgUserMacs(ctx, orgId, nil, nil, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"id":"111cafd2-ba1b-5169-bfcb-9cdf1d473ddb","labels":["flor1","bld4"],"mac":"921b638445cd","notes":"mac address refers to Canon printers","vlan":"30"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUserMACsTestSearchOrgUserMacs1 tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestSearchOrgUserMacs1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsUserMaCs.SearchOrgUserMacs(ctx, orgId, nil, nil, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"id":"111cafd2-ba1b-5169-bfcb-9cdf1d473ddb","labels":["flor1","bld4"],"mac":"921b638445cd","notes":"mac address refers to Canon printers","vlan":"30"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUserMACsTestDeleteOrgUserMac tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestDeleteOrgUserMac(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    usermacId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsUserMaCs.DeleteOrgUserMac(ctx, orgId, usermacId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsUserMACsTestGetOrgUserMac tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestGetOrgUserMac(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    usermacId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsUserMaCs.GetOrgUserMac(ctx, orgId, usermacId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"111cafd2-ba1b-5169-bfcb-9cdf1d473ddb","labels":["flor1","bld4"],"mac":"921b638445cd","notes":"mac address refers to Canon printers","vlan":"30"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUserMACsTestGetOrgUserMac1 tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestGetOrgUserMac1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    usermacId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsUserMaCs.GetOrgUserMac(ctx, orgId, usermacId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"111cafd2-ba1b-5169-bfcb-9cdf1d473ddb","labels":["flor1","bld4"],"mac":"921b638445cd","notes":"mac address refers to Canon printers","vlan":"30"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUserMACsTestUpdateOrgUserMac tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestUpdateOrgUserMac(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    usermacId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsUserMaCs.UpdateOrgUserMac(ctx, orgId, usermacId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"111cafd2-ba1b-5169-bfcb-9cdf1d473ddb","labels":["flor1","bld4"],"mac":"921b638445cd","notes":"mac address refers to Canon printers","vlan":"30"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsUserMACsTestUpdateOrgUserMac1 tests the behavior of the OrgsUserMACs
func TestOrgsUserMACsTestUpdateOrgUserMac1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    usermacId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsUserMaCs.UpdateOrgUserMac(ctx, orgId, usermacId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"111cafd2-ba1b-5169-bfcb-9cdf1d473ddb","labels":["flor1","bld4"],"mac":"921b638445cd","notes":"mac address refers to Canon printers","vlan":"30"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
