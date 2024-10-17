package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesStatsZonesTestListSiteZonesStats tests the behavior of the SitesStatsZones
func TestSitesStatsZonesTestListSiteZonesStats(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mapId := "00000000-0000-0000-0000-000000000000"
	apiResponse, err := sitesStatsZones.ListSiteZonesStats(ctx, siteId, &mapId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"assets_waits":{"avg":0,"max":0,"min":0,"p95":0},"clients_waits":{"avg":1200,"max":3610,"min":600,"p95":2800},"created_time":1616625211,"id":"123470c7-5d9d-424a-8475-8b344c621234","map_id":"123449d4-d12f-4feb-b40f-5be0e2ae1234","modified_time":1616625211,"name":"Zone A","num_assets":0,"num_clients":80,"num_sdkclients":10,"occupancy_limit":4,"org_id":"1234c1a0-6ef6-11e6-8bbf-02e208b21234","sdkclients_waits":{"avg":1200,"max":3610,"min":600,"p95":2800},"site_id":"123448e6-6ef6-11e6-8bbf-02e208b21234","vertices":[{"x":732,"y":1821},{"x":732.5,"y":1731},{"x":837.5,"y":1731.5},{"x":839,"y":1821}],"vertices_m":[{"x":24.1983341951072,"y":60.198314985369144},{"x":24.214863111907139,"y":57.223109961380558},{"x":27.685935639893827,"y":57.239638878180493},{"x":27.735522390293639,"y":60.198314985369144}]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
