package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsSiteTemplatesTestListOrgSiteTemplates tests the behavior of the OrgsSiteTemplates
func TestOrgsSiteTemplatesTestListOrgSiteTemplates(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsSiteTemplates.ListOrgSiteTemplates(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"auto_upgrade":{"day_of_week":"mon","enabled":true,"time_of_day":"string","version":"string"},"name":"string","vars":{"SSID_STR":"string","VLAN_ID":"string"}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSiteTemplatesTestListOrgSiteTemplates1 tests the behavior of the OrgsSiteTemplates
func TestOrgsSiteTemplatesTestListOrgSiteTemplates1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsSiteTemplates.ListOrgSiteTemplates(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"auto_upgrade":{"day_of_week":"mon","enabled":true,"time_of_day":"string","version":"string"},"name":"string","vars":{"SSID_STR":"string","VLAN_ID":"string"}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSiteTemplatesTestCreateOrgSiteTemplate tests the behavior of the OrgsSiteTemplates
func TestOrgsSiteTemplatesTestCreateOrgSiteTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSiteTemplates.CreateOrgSiteTemplate(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"auto_upgrade":{"day_of_week":"mon","enabled":true,"time_of_day":"string","version":"string"},"name":"string","vars":{"SSID_STR":"string","VLAN_ID":"string"}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSiteTemplatesTestCreateOrgSiteTemplate1 tests the behavior of the OrgsSiteTemplates
func TestOrgsSiteTemplatesTestCreateOrgSiteTemplate1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSiteTemplates.CreateOrgSiteTemplate(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"auto_upgrade":{"day_of_week":"mon","enabled":true,"time_of_day":"string","version":"string"},"name":"string","vars":{"SSID_STR":"string","VLAN_ID":"string"}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSiteTemplatesTestDeleteOrgSiteTemplate tests the behavior of the OrgsSiteTemplates
func TestOrgsSiteTemplatesTestDeleteOrgSiteTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sitetemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsSiteTemplates.DeleteOrgSiteTemplate(ctx, orgId, sitetemplateId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSiteTemplatesTestGetOrgSiteTemplate tests the behavior of the OrgsSiteTemplates
func TestOrgsSiteTemplatesTestGetOrgSiteTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sitetemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSiteTemplates.GetOrgSiteTemplate(ctx, orgId, sitetemplateId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"auto_upgrade":{"day_of_week":"mon","enabled":true,"time_of_day":"string","version":"string"},"name":"string","vars":{"SSID_STR":"string","VLAN_ID":"string"}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSiteTemplatesTestGetOrgSiteTemplate1 tests the behavior of the OrgsSiteTemplates
func TestOrgsSiteTemplatesTestGetOrgSiteTemplate1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sitetemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSiteTemplates.GetOrgSiteTemplate(ctx, orgId, sitetemplateId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"auto_upgrade":{"day_of_week":"mon","enabled":true,"time_of_day":"string","version":"string"},"name":"string","vars":{"SSID_STR":"string","VLAN_ID":"string"}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSiteTemplatesTestUpdateOrgSiteTemplate tests the behavior of the OrgsSiteTemplates
func TestOrgsSiteTemplatesTestUpdateOrgSiteTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sitetemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSiteTemplates.UpdateOrgSiteTemplate(ctx, orgId, sitetemplateId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"auto_upgrade":{"day_of_week":"mon","enabled":true,"time_of_day":"string","version":"string"},"name":"string","vars":{"SSID_STR":"string","VLAN_ID":"string"}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSiteTemplatesTestUpdateOrgSiteTemplate1 tests the behavior of the OrgsSiteTemplates
func TestOrgsSiteTemplatesTestUpdateOrgSiteTemplate1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sitetemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSiteTemplates.UpdateOrgSiteTemplate(ctx, orgId, sitetemplateId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"auto_upgrade":{"day_of_week":"mon","enabled":true,"time_of_day":"string","version":"string"},"name":"string","vars":{"SSID_STR":"string","VLAN_ID":"string"}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
