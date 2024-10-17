package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesStatsTestGetSiteStats tests the behavior of the SitesStats
func TestSitesStatsTestGetSiteStats(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesStats.GetSiteStats(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"address":"string","alarmtemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","country_code":"string","created_time":0,"id":"55c29ce5-7c0f-45b5-b99b-599f805fa3a1","lat":0,"latlng":{"lat":0,"lng":0},"lng":0,"modified_time":0,"msp_id":"dca3cad3-0c9b-439b-814f-8d5f23797972","name":"string","networktemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","num_ap":0,"num_ap_connected":0,"num_clients":0,"num_devices":0,"num_devices_connected":0,"num_gateway":0,"num_gateway_connected":0,"num_switch":0,"num_switch_connected":0,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","rftemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"],"timezone":"string","tzoffset":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
