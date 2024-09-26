package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsSecPoliciesTestListOrgSecPolicies tests the behavior of the OrgsSecPolicies
func TestOrgsSecPoliciesTestListOrgSecPolicies(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsSecPolicies.ListOrgSecPolicies(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"name":"corporate only","wlans":[{"auth":{"pairwise":["wpa1-tkip","wpa2-tkip"],"type":"psk"},"band":"both","ssid":"office"},{"auth":{"type":"open"},"band":"5","ssid":"office-guest"}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecPoliciesTestCreateOrgSecPolicies tests the behavior of the OrgsSecPolicies
func TestOrgsSecPoliciesTestCreateOrgSecPolicies(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSecPolicies.CreateOrgSecPolicies(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"name":"corporate only","wlans":[{"auth":{"pairwise":["wpa1-tkip","wpa2-tkip"],"type":"psk"},"band":"both","ssid":"office"},{"auth":{"type":"open"},"band":"5","ssid":"office-guest"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecPoliciesTestDeleteOrgSecPolicy tests the behavior of the OrgsSecPolicies
func TestOrgsSecPoliciesTestDeleteOrgSecPolicy(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    secpolicyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsSecPolicies.DeleteOrgSecPolicy(ctx, orgId, secpolicyId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSecPoliciesTestGetOrgSecPolicy tests the behavior of the OrgsSecPolicies
func TestOrgsSecPoliciesTestGetOrgSecPolicy(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    secpolicyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSecPolicies.GetOrgSecPolicy(ctx, orgId, secpolicyId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"name":"corporate only","wlans":[{"auth":{"pairwise":["wpa1-tkip","wpa2-tkip"],"type":"psk"},"band":"both","ssid":"office"},{"auth":{"type":"open"},"band":"5","ssid":"office-guest"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecPoliciesTestUpdateOrgSecPolicies tests the behavior of the OrgsSecPolicies
func TestOrgsSecPoliciesTestUpdateOrgSecPolicies(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    secpolicyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSecPolicies.UpdateOrgSecPolicies(ctx, orgId, secpolicyId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"name":"corporate only","wlans":[{"auth":{"pairwise":["wpa1-tkip","wpa2-tkip"],"type":"psk"},"band":"both","ssid":"office"},{"auth":{"type":"open"},"band":"5","ssid":"office-guest"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
