package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsJSITestListOrgJsiDevices tests the behavior of the OrgsJSI
func TestOrgsJSITestListOrgJsiDevices(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)

	apiResponse, err := orgsJsi.ListOrgJsiDevices(ctx, orgId, &limit, &page, nil, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"ext_ip":"73.92.124.103","last_seen":1654636867,"mac":"c15353123096","model":"EX2300-C-12P","serial":"DGCOO0015"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsJSITestAdoptOrgJsiDevice tests the behavior of the OrgsJSI
func TestOrgsJSITestAdoptOrgJsiDevice(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsJsi.AdoptOrgJsiDevice(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"cmd":"set system services ssh...\n...\nset system services outbound-ssh client mist ..."}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsJSITestCreateOrgJsiDeviceShellSession tests the behavior of the OrgsJSI
func TestOrgsJSITestCreateOrgJsiDeviceShellSession(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceMac := "0000000000ab"
	apiResponse, err := orgsJsi.CreateOrgJsiDeviceShellSession(ctx, orgId, deviceMac)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsJSITestListOrgJsiPastPurchases tests the behavior of the OrgsJSI
func TestOrgsJSITestListOrgJsiPastPurchases(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)

	apiResponse, err := orgsJsi.ListOrgJsiPastPurchases(ctx, orgId, &limit, &page, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"contract_end_time":1708214400,"contract_id":"68247710","contract_sku":"PAR-SDCE-EX4300T48","contract_start_time":1629331200,"customer_po":"4501942932","distributor":"dist1","eol_time":1671062400,"eos_time":1828828800,"installed_address":"Address","model":"EX4300-48T","order_id":"15909647","reseller":"reseller1","serial":"PE3721050223","sku":"EX4300-48T-AFI","type":"switch","warranty_end_time":2258236800,"warranty_start_time":1619395200,"warranty_type":"Enhanced Hardware Warranty"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
