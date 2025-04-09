package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsSecurityPoliciesTestListOrgSecPolicies tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestListOrgSecPolicies(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsSecurityPolicies.ListOrgSecPolicies(ctx, orgId, &limit, &page)
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

// TestOrgsSecurityPoliciesTestCreateOrgSecPolicy tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestCreateOrgSecPolicy(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSecurityPolicies.CreateOrgSecPolicy(ctx, orgId, nil)
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

// TestOrgsSecurityPoliciesTestDeleteOrgSecPolicy tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestDeleteOrgSecPolicy(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    secpolicyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsSecurityPolicies.DeleteOrgSecPolicy(ctx, orgId, secpolicyId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSecurityPoliciesTestGetOrgSecPolicy tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestGetOrgSecPolicy(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    secpolicyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSecurityPolicies.GetOrgSecPolicy(ctx, orgId, secpolicyId)
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

// TestOrgsSecurityPoliciesTestUpdateOrgSecPolicy tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestUpdateOrgSecPolicy(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    secpolicyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSecurityPolicies.UpdateOrgSecPolicy(ctx, orgId, secpolicyId, nil)
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
