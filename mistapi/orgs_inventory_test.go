// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestOrgsInventoryTestGetOrgInventory tests the behavior of the OrgsInventory
func TestOrgsInventoryTestGetOrgInventory(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	serial := "FXLH2015150025"
	model := "AP43"

	mac := "5c5b350e0001"
	siteId, errUUID := uuid.Parse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")
	if errUUID != nil {
		t.Error(errUUID)
	}
	vcMac := "5c5b350e0001"
	vc := bool(false)
	unassigned := bool(true)
	modifiedAfter := int(1703003296)
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsInventory.GetOrgInventory(ctx, orgId, &serial, &model, nil, &mac, &siteId, &vcMac, &vc, &unassigned, &modifiedAfter, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"connected":true,"created_time":1542328276,"deviceprofile_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9","id":"00000000-0000-0000-0000-5c5b35000018","mac":"5c5b35000018","model":"AP41","modified_time":1542829778,"name":"hallway","serial":"FXLH2015150025","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","type":"ap"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsInventoryTestGetOrgInventory1 tests the behavior of the OrgsInventory
func TestOrgsInventoryTestGetOrgInventory1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	serial := "FXLH2015150025"
	model := "AP43"

	mac := "5c5b350e0001"
	siteId, errUUID := uuid.Parse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")
	if errUUID != nil {
		t.Error(errUUID)
	}
	vcMac := "5c5b350e0001"
	vc := bool(false)
	unassigned := bool(true)
	modifiedAfter := int(1703003296)
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsInventory.GetOrgInventory(ctx, orgId, &serial, &model, nil, &mac, &siteId, &vcMac, &vc, &unassigned, &modifiedAfter, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"connected":true,"created_time":1542328276,"deviceprofile_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9","id":"00000000-0000-0000-0000-5c5b35000018","mac":"5c5b35000018","model":"AP41","modified_time":1542829778,"name":"hallway","serial":"FXLH2015150025","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","type":"ap"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsInventoryTestAddOrgInventory tests the behavior of the OrgsInventory
func TestOrgsInventoryTestAddOrgInventory(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	body := []string{"6JG8E-PTFV2-A9Z2N", "DVH4V-SNMSZ-PDXBR"}
	apiResponse, err := orgsInventory.AddOrgInventory(ctx, orgId, body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"added":["6JG8E-PTFV2-A9Z2N"],"duplicated":["DVH4V-SNMSZ-PDXBR"],"error":["PO1025335ohoh"],"inventory_added":[{"mac":"5c5b35000018","magic":"6JG8EPTFV2A9Z2N","model":"AP41","serial":"FXLH2015150025","type":"ap"}],"inventory_duplicated":[{"mac":"5c5b35000012","magic":"DVH4VSNMSZPDXBR","model":"AP41","serial":"FXLH2015150027","type":"ap"}],"reason":["belongs to another org ('e2f543f7-d6e1-409f-a565-e77a1f098d3b' (other) != '0de5d6fc-219a-414d-a840-67d6b919ad8f' (you))"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsInventoryTestAddOrgInventory1 tests the behavior of the OrgsInventory
func TestOrgsInventoryTestAddOrgInventory1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	body := []string{"6JG8E-PTFV2-A9Z2N", "DVH4V-SNMSZ-PDXBR"}
	apiResponse, err := orgsInventory.AddOrgInventory(ctx, orgId, body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"added":["6JG8E-PTFV2-A9Z2N"],"duplicated":["DVH4V-SNMSZ-PDXBR"],"error":["PO1025335ohoh"],"inventory_added":[{"mac":"5c5b35000018","magic":"6JG8EPTFV2A9Z2N","model":"AP41","serial":"FXLH2015150025","type":"ap"}],"inventory_duplicated":[{"mac":"5c5b35000012","magic":"DVH4VSNMSZPDXBR","model":"AP41","serial":"FXLH2015150027","type":"ap"}],"reason":["belongs to another org ('e2f543f7-d6e1-409f-a565-e77a1f098d3b' (other) != '0de5d6fc-219a-414d-a840-67d6b919ad8f' (you))"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsInventoryTestUpdateOrgInventoryAssignment tests the behavior of the OrgsInventory
func TestOrgsInventoryTestUpdateOrgInventoryAssignment(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.InventoryUpdate
	errBody := json.Unmarshal([]byte(`{"disable_auto_config":false,"macs":["5c5b350e0001"],"managed":false,"no_reassign":false,"op":"assign","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsInventory.UpdateOrgInventoryAssignment(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"error":[],"op":"assign","reason":[],"success":["5c5b350e0001"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsInventoryTestUpdateOrgInventoryAssignment1 tests the behavior of the OrgsInventory
func TestOrgsInventoryTestUpdateOrgInventoryAssignment1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.InventoryUpdate
	errBody := json.Unmarshal([]byte(`{"disable_auto_config":false,"macs":["5c5b350e0001"],"managed":false,"no_reassign":false,"op":"assign","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsInventory.UpdateOrgInventoryAssignment(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"error":[],"op":"assign","reason":[],"success":["5c5b350e0001"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsInventoryTestCountOrgInventory tests the behavior of the OrgsInventory
func TestOrgsInventoryTestCountOrgInventory(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mType := models.DeviceTypeDefaultApEnum("ap")
	distinct := models.InventoryCountDistinctEnum("model")
	limit := int(100)
	apiResponse, err := orgsInventory.CountOrgInventory(ctx, orgId, &mType, &distinct, &limit)
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

// TestOrgsInventoryTestCountOrgInventory1 tests the behavior of the OrgsInventory
func TestOrgsInventoryTestCountOrgInventory1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mType := models.DeviceTypeDefaultApEnum("ap")
	distinct := models.InventoryCountDistinctEnum("model")
	limit := int(100)
	apiResponse, err := orgsInventory.CountOrgInventory(ctx, orgId, &mType, &distinct, &limit)
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

// TestOrgsInventoryTestCreateOrgGatewayHaCluster tests the behavior of the OrgsInventory
func TestOrgsInventoryTestCreateOrgGatewayHaCluster(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.HaClusterConfig
	errBody := json.Unmarshal([]byte(`{"disable_auto_config":true,"managed":true,"nodes":[{"mac":"aff827549235"},{"mac":"8396cd006c8c"}],"site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	resp, err := orgsInventory.CreateOrgGatewayHaCluster(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsInventoryTestDeleteOrgGatewayHaCluster tests the behavior of the OrgsInventory
func TestOrgsInventoryTestDeleteOrgGatewayHaCluster(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	resp, err := orgsInventory.DeleteOrgGatewayHaCluster(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsInventoryTestReevaluateOrgAutoAssignment tests the behavior of the OrgsInventory
func TestOrgsInventoryTestReevaluateOrgAutoAssignment(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsInventory.ReevaluateOrgAutoAssignment(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsInventoryTestReplaceOrgDevices tests the behavior of the OrgsInventory
func TestOrgsInventoryTestReplaceOrgDevices(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.ReplaceDevice
	errBody := json.Unmarshal([]byte(`{"discard":[],"inventory_mac":"5c5b35000301","mac":"5c5b35000101","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsInventory.ReplaceOrgDevices(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"error":[],"op":"assign","reason":[],"success":["5c5b350e0001"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsInventoryTestReplaceOrgDevices1 tests the behavior of the OrgsInventory
func TestOrgsInventoryTestReplaceOrgDevices1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.ReplaceDevice
	errBody := json.Unmarshal([]byte(`{"discard":[],"inventory_mac":"5c5b35000301","mac":"5c5b35000101","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsInventory.ReplaceOrgDevices(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"error":[],"op":"assign","reason":[],"success":["5c5b350e0001"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsInventoryTestSearchOrgInventory tests the behavior of the OrgsInventory
func TestOrgsInventoryTestSearchOrgInventory(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mType := models.DeviceTypeDefaultApEnum("ap")
	mac := "5c5b350e0001"
	vcMac := "5c5b350e0001"
	masterMac := "5c5b350e0001"
	siteId, errUUID := uuid.Parse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")
	if errUUID != nil {
		t.Error(errUUID)
	}
	serial := "FXLH2015150025"
	master := "true"
	sku := "AP43-WW"
	version := "1.0.0"
	status := "connected"

	limit := int(100)
	sort := "timestamp"

	apiResponse, err := orgsInventory.SearchOrgInventory(ctx, orgId, &mType, &mac, &vcMac, &masterMac, &siteId, &serial, &master, &sku, &version, &status, nil, &limit, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"limit":1000,"results":[{"mac":"f01c2df166e0","master":true,"members":[{"mac":"f01c2df166e0","model":"EX4300-48P","serial":"PD3714460200"}],"model":"EX4300-48P","name":"mist-wa-ex4300-VC","org_id":"9b853544-51e4-45fb-81ac-a442e4a111d0","serial":"PD3714460200","site_id":"01dc141d-b6af-4baa-b00f-0e31ef954c4f","sku":"EX4300-48P","status":"disconnected","type":"switch","vc_mac":"f01c2df166e0","version":"21.4R3.5"}],"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsInventoryTestSearchOrgInventory1 tests the behavior of the OrgsInventory
func TestOrgsInventoryTestSearchOrgInventory1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mType := models.DeviceTypeDefaultApEnum("ap")
	mac := "5c5b350e0001"
	vcMac := "5c5b350e0001"
	masterMac := "5c5b350e0001"
	siteId, errUUID := uuid.Parse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")
	if errUUID != nil {
		t.Error(errUUID)
	}
	serial := "FXLH2015150025"
	master := "true"
	sku := "AP43-WW"
	version := "1.0.0"
	status := "connected"

	limit := int(100)
	sort := "timestamp"

	apiResponse, err := orgsInventory.SearchOrgInventory(ctx, orgId, &mType, &mac, &vcMac, &masterMac, &siteId, &serial, &master, &sku, &version, &status, nil, &limit, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"limit":1000,"results":[{"mac":"f01c2df166e0","master":true,"members":[{"mac":"f01c2df166e0","model":"EX4300-48P","serial":"PD3714460200"}],"model":"EX4300-48P","name":"mist-wa-ex4300-VC","org_id":"9b853544-51e4-45fb-81ac-a442e4a111d0","serial":"PD3714460200","site_id":"01dc141d-b6af-4baa-b00f-0e31ef954c4f","sku":"EX4300-48P","status":"disconnected","type":"switch","vc_mac":"f01c2df166e0","version":"21.4R3.5"}],"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
