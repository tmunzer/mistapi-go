package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsPremiumAnalyticsTestListOrgPmaDashboards tests the behavior of the OrgsPremiumAnalytics
func TestOrgsPremiumAnalyticsTestListOrgPmaDashboards(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsPremiumAnalytics.ListOrgPmaDashboards(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"name":"dashboard_1","url":"https://mist.looker.com/login/embed/%2Fembed%2Fdashboards%2F1?group_ids=%5B3%5D&last_name=%22%22&models=%5B%22generic%22%5D&....."}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
