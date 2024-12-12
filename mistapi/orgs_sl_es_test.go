package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsSLEsTestGetOrgSitesSle tests the behavior of the OrgsSLEs
func TestOrgsSLEsTestGetOrgSitesSle(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    duration := "1d"
    interval := "10m"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsSlEs.GetOrgSitesSle(ctx, orgId, nil, nil, nil, &duration, &interval, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1651323600,"interval":3600,"limit":1,"page":2,"results":[{"application_health":0.82500000479428659,"gateway-health":1,"num_clients":65,"num_gateways":1,"site_id":"f5fcbee5-1234-5678-9101-1619ede87879","wan-link-health":0.99884710892724837}],"start":1651269600,"total":4}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
