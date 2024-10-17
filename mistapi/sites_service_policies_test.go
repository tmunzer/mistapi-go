package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesServicePoliciesTestListSiteServicePoliciesDerived tests the behavior of the SitesServicePolicies
func TestSitesServicePoliciesTestListSiteServicePoliciesDerived(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resolve := bool(false)
	apiResponse, err := sitesServicePolicies.ListSiteServicePoliciesDerived(ctx, siteId, &resolve)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"action":"allow","created_time":0,"id":"string","modified_time":0,"name":"string","org_id":"string","services":["string"],"tenants":["string"]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
