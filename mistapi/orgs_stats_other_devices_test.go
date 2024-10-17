package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsStatsOtherDevicesTestGetOrgOtherDeviceStats tests the behavior of the OrgsStatsOtherDevices
func TestOrgsStatsOtherDevicesTestGetOrgOtherDeviceStats(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceMac := "0000000000ab"
	apiResponse, err := orgsStatsOtherDevices.GetOrgOtherDeviceStats(ctx, orgId, deviceMac)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"config_status":"synced","last_config":1675392788,"last_seen":1675843629,"mac":"5c5b35000018","status":"online","uptime":20296,"vendor":"cradlepoint","vendor_specific":{"ports":{"mdm-4d0e073b":{"bytes_in":33004879,"bytes_out":41103393,"health_category":"","health_score":0,"id":"101027967","mode":"wan","model":"Internal 5GB (SIM1)","state":"READY","type":"5G","uptime":252371.34149021498}},"router_id":null,"target_version":"7.23.40"},"version":"7.22.70"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
