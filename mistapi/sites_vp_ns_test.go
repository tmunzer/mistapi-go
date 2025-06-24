package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestSitesVPNsTestListSiteVpnsDerived tests the behavior of the SitesVPNs
func TestSitesVPNsTestListSiteVpnsDerived(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resolve := bool(false)
    apiResponse, err := sitesVpNs.ListSiteVpnsDerived(ctx, siteId, &resolve)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"name":"string","paths":{"property1":{"bfd_profile":"broadband","ip":"string"},"property2":{"bfd_profile":"lte","ip":"string"}}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesVPNsTestListSiteVpnsDerived1 tests the behavior of the SitesVPNs
func TestSitesVPNsTestListSiteVpnsDerived1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resolve := bool(false)
    apiResponse, err := sitesVpNs.ListSiteVpnsDerived(ctx, siteId, &resolve)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"name":"string","paths":{"property1":{"bfd_profile":"broadband","ip":"string"},"property2":{"bfd_profile":"lte","ip":"string"}}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
