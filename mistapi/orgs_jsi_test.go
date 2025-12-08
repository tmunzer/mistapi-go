// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
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
	model := "AP43"
	serial := "FXLH2015150025"
	mac := "5c5b350e0001"
	apiResponse, err := orgsJsi.ListOrgJsiDevices(ctx, orgId, &limit, &page, &model, &serial, &mac)
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

// TestOrgsJSITestListOrgJsiDevices1 tests the behavior of the OrgsJSI
func TestOrgsJSITestListOrgJsiDevices1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	model := "AP43"
	serial := "FXLH2015150025"
	mac := "5c5b350e0001"
	apiResponse, err := orgsJsi.ListOrgJsiDevices(ctx, orgId, &limit, &page, &model, &serial, &mac)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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

// TestOrgsJSITestAdoptOrgJsiDevice1 tests the behavior of the OrgsJSI
func TestOrgsJSITestAdoptOrgJsiDevice1(t *testing.T) {
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
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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

// TestOrgsJSITestCreateOrgJsiDeviceShellSession1 tests the behavior of the OrgsJSI
func TestOrgsJSITestCreateOrgJsiDeviceShellSession1(t *testing.T) {
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
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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
	model := "AP43"
	serial := "FXLH2015150025"
	apiResponse, err := orgsJsi.ListOrgJsiPastPurchases(ctx, orgId, &limit, &page, &model, &serial)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"device_name":"name1","eol_time":1561507200,"eos_time":1672012800,"master":true,"model":"EX2300-24MP","org_id":"6e843b41-f953-4af9-80e5-e1a70f65754a","serial":"XN3123300095","sku":"EX2300","status":"connected","suggested_version":"Latest 21.4R3-Sx","type":"switch","version":"23.4R2-S4.11","version_eos_time":1672012800,"version_time":1561507200,"warranty":"Enhanced Hardware Warranty","warranty_time":1672012800,"warranty_type":"Enhanced Hardware Warranty"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsJSITestListOrgJsiPastPurchases1 tests the behavior of the OrgsJSI
func TestOrgsJSITestListOrgJsiPastPurchases1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	model := "AP43"
	serial := "FXLH2015150025"
	apiResponse, err := orgsJsi.ListOrgJsiPastPurchases(ctx, orgId, &limit, &page, &model, &serial)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"device_name":"name1","eol_time":1561507200,"eos_time":1672012800,"master":true,"model":"EX2300-24MP","org_id":"6e843b41-f953-4af9-80e5-e1a70f65754a","serial":"XN3123300095","sku":"EX2300","status":"connected","suggested_version":"Latest 21.4R3-Sx","type":"switch","version":"23.4R2-S4.11","version_eos_time":1672012800,"version_time":1561507200,"warranty":"Enhanced Hardware Warranty","warranty_time":1672012800,"warranty_type":"Enhanced Hardware Warranty"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsJSITestCountOrgJsiAssetsAndContracts tests the behavior of the OrgsJSI
func TestOrgsJSITestCountOrgJsiAssetsAndContracts(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)
	apiResponse, err := orgsJsi.CountOrgJsiAssetsAndContracts(ctx, orgId, nil, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsJSITestCountOrgJsiAssetsAndContracts1 tests the behavior of the OrgsJSI
func TestOrgsJSITestCountOrgJsiAssetsAndContracts1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)
	apiResponse, err := orgsJsi.CountOrgJsiAssetsAndContracts(ctx, orgId, nil, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsJSITestSearchOrgJsiAssetsAndContracts tests the behavior of the OrgsJSI
func TestOrgsJSITestSearchOrgJsiAssetsAndContracts(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	claimed := bool(true)
	model := "AP43"
	serial := "FXLH2015150025"
	sku := "EX2300"
	status := models.DeviceStatusEnum("all")

	eolDuration := "30d"
	eosDuration := "30d"
	hasSupport := bool(true)

	limit := int(100)
	sort := "timestamp"

	apiResponse, err := orgsJsi.SearchOrgJsiAssetsAndContracts(ctx, orgId, &claimed, &model, &serial, &sku, &status, nil, &eolDuration, &eosDuration, &hasSupport, nil, &limit, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1748023308,"limit":1000,"results":[{"claimed":true,"device_name":"name1","eol_time":1561507200,"eos_time":1672012800,"has_support":true,"master":true,"model":"EX2300-24MP","org_id":"6e843b41-f953-4af9-80e5-e1a70f65754a","serial":"XN3123300095","sku":"EX2300","status":"connected","suggested_version":"Latest 21.4R3-Sx","type":"switch","version":"23.4R2-S4.11","version_eos_time":1672012800,"version_time":1561507200,"warranty":"Enhanced Hardware Warranty","warranty_time":1672012800,"warranty_type":"Enhanced Hardware Warranty"}],"start":1748019708,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsJSITestSearchOrgJsiAssetsAndContracts1 tests the behavior of the OrgsJSI
func TestOrgsJSITestSearchOrgJsiAssetsAndContracts1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	claimed := bool(true)
	model := "AP43"
	serial := "FXLH2015150025"
	sku := "EX2300"
	status := models.DeviceStatusEnum("all")

	eolDuration := "30d"
	eosDuration := "30d"
	hasSupport := bool(true)

	limit := int(100)
	sort := "timestamp"

	apiResponse, err := orgsJsi.SearchOrgJsiAssetsAndContracts(ctx, orgId, &claimed, &model, &serial, &sku, &status, nil, &eolDuration, &eosDuration, &hasSupport, nil, &limit, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1748023308,"limit":1000,"results":[{"claimed":true,"device_name":"name1","eol_time":1561507200,"eos_time":1672012800,"has_support":true,"master":true,"model":"EX2300-24MP","org_id":"6e843b41-f953-4af9-80e5-e1a70f65754a","serial":"XN3123300095","sku":"EX2300","status":"connected","suggested_version":"Latest 21.4R3-Sx","type":"switch","version":"23.4R2-S4.11","version_eos_time":1672012800,"version_time":1561507200,"warranty":"Enhanced Hardware Warranty","warranty_time":1672012800,"warranty_type":"Enhanced Hardware Warranty"}],"start":1748019708,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
