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

// TestOrgsDevicesTestListOrgDevices tests the behavior of the OrgsDevices
func TestOrgsDevicesTestListOrgDevices(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsDevices.ListOrgDevices(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"results":[{"mac":"string","name":"string"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesTestListOrgDevices1 tests the behavior of the OrgsDevices
func TestOrgsDevicesTestListOrgDevices1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsDevices.ListOrgDevices(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"results":[{"mac":"string","name":"string"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesTestCountOrgDevices tests the behavior of the OrgsDevices
func TestOrgsDevicesTestCountOrgDevices(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgDevicesCountDistinctEnum("model")
	hostname := "my-hostname"
	siteId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
	if errUUID != nil {
		t.Error(errUUID)
	}
	model := "MR84"
	managed := "true"
	mac := "5c5b53010101"
	version := "10.0.0"
	ip := "192.168.1.1"

	mxedgeId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
	if errUUID != nil {
		t.Error(errUUID)
	}
	lldpSystemName := "my-lldp-system"
	lldpSystemDesc := "my-lldp-system-description"
	lldpPortId := "ge-0/0/1"
	lldpMgmtAddr := "10.4.2.3"
	mType := models.DeviceTypeDefaultApEnum("ap")

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsDevices.CountOrgDevices(ctx, orgId, &distinct, &hostname, &siteId, &model, &managed, &mac, &version, &ip, nil, &mxedgeId, &lldpSystemName, &lldpSystemDesc, &lldpPortId, &lldpMgmtAddr, &mType, nil, nil, &duration, &limit)
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

// TestOrgsDevicesTestCountOrgDevices1 tests the behavior of the OrgsDevices
func TestOrgsDevicesTestCountOrgDevices1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgDevicesCountDistinctEnum("model")
	hostname := "my-hostname"
	siteId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
	if errUUID != nil {
		t.Error(errUUID)
	}
	model := "MR84"
	managed := "true"
	mac := "5c5b53010101"
	version := "10.0.0"
	ip := "192.168.1.1"

	mxedgeId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
	if errUUID != nil {
		t.Error(errUUID)
	}
	lldpSystemName := "my-lldp-system"
	lldpSystemDesc := "my-lldp-system-description"
	lldpPortId := "ge-0/0/1"
	lldpMgmtAddr := "10.4.2.3"
	mType := models.DeviceTypeDefaultApEnum("ap")

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsDevices.CountOrgDevices(ctx, orgId, &distinct, &hostname, &siteId, &model, &managed, &mac, &version, &ip, nil, &mxedgeId, &lldpSystemName, &lldpSystemDesc, &lldpPortId, &lldpMgmtAddr, &mType, nil, nil, &duration, &limit)
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

// TestOrgsDevicesTestCountOrgDeviceEvents tests the behavior of the OrgsDevices
func TestOrgsDevicesTestCountOrgDeviceEvents(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgDevicesEventsCountDistinctEnum("model")
	siteId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
	if errUUID != nil {
		t.Error(errUUID)
	}
	ap := "5c5b53010101"
	apfw := "10.0.0"
	model := "AP43"
	text := "Device connected"
	timestamp := "1703003296"

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsDevices.CountOrgDeviceEvents(ctx, orgId, &distinct, &siteId, &ap, &apfw, &model, &text, &timestamp, nil, nil, nil, &duration, &limit)
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

// TestOrgsDevicesTestCountOrgDeviceEvents1 tests the behavior of the OrgsDevices
func TestOrgsDevicesTestCountOrgDeviceEvents1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgDevicesEventsCountDistinctEnum("model")
	siteId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
	if errUUID != nil {
		t.Error(errUUID)
	}
	ap := "5c5b53010101"
	apfw := "10.0.0"
	model := "AP43"
	text := "Device connected"
	timestamp := "1703003296"

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsDevices.CountOrgDeviceEvents(ctx, orgId, &distinct, &siteId, &ap, &apfw, &model, &text, &timestamp, nil, nil, nil, &duration, &limit)
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

// TestOrgsDevicesTestSearchOrgDeviceEvents tests the behavior of the OrgsDevices
func TestOrgsDevicesTestSearchOrgDeviceEvents(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mac := "5c5b53010101"
	model := "AP43"
	deviceType := models.DeviceTypeWithAllEnum("ap")
	text := "Device connected"
	timestamp := "1703003296"

	lastBy := "port_id"
	includes := "ext_tunnel"
	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsDevices.SearchOrgDeviceEvents(ctx, orgId, &mac, &model, &deviceType, &text, &timestamp, nil, &lastBy, &includes, &limit, nil, nil, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":0,"limit":0,"next":"string","results":[{"ap":"5c5b351e13b5","apfw":"5c5b351e13b5","model":"BT11-WW","org_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862a","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","text":"Succeeding DNS query from 172.29.101.134 to 172.29.101.7 for \"portal.mistsys.com\" on vlan 1, id 60224","timestamp":1547235620.89,"type":"CLIENT_DNS_OK"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesTestSearchOrgDeviceEvents1 tests the behavior of the OrgsDevices
func TestOrgsDevicesTestSearchOrgDeviceEvents1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mac := "5c5b53010101"
	model := "AP43"
	deviceType := models.DeviceTypeWithAllEnum("ap")
	text := "Device connected"
	timestamp := "1703003296"

	lastBy := "port_id"
	includes := "ext_tunnel"
	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsDevices.SearchOrgDeviceEvents(ctx, orgId, &mac, &model, &deviceType, &text, &timestamp, nil, &lastBy, &includes, &limit, nil, nil, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":0,"limit":0,"next":"string","results":[{"ap":"5c5b351e13b5","apfw":"5c5b351e13b5","model":"BT11-WW","org_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862a","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","text":"Succeeding DNS query from 172.29.101.134 to 172.29.101.7 for \"portal.mistsys.com\" on vlan 1, id 60224","timestamp":1547235620.89,"type":"CLIENT_DNS_OK"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesTestCountOrgDeviceLastConfigs tests the behavior of the OrgsDevices
func TestOrgsDevicesTestCountOrgDeviceLastConfigs(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mType := models.DeviceTypeDefaultApEnum("ap")

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsDevices.CountOrgDeviceLastConfigs(ctx, orgId, &mType, nil, nil, nil, &duration, &limit)
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

// TestOrgsDevicesTestCountOrgDeviceLastConfigs1 tests the behavior of the OrgsDevices
func TestOrgsDevicesTestCountOrgDeviceLastConfigs1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mType := models.DeviceTypeDefaultApEnum("ap")

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsDevices.CountOrgDeviceLastConfigs(ctx, orgId, &mType, nil, nil, nil, &duration, &limit)
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

// TestOrgsDevicesTestSearchOrgDeviceLastConfigs tests the behavior of the OrgsDevices
func TestOrgsDevicesTestSearchOrgDeviceLastConfigs(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceType := models.LastConfigDeviceTypeEnum("ap")
	mac := "5c5b53010101"
	name := "My AP"
	version := "10.0.0"
	certExpiryDuration := "2d"

	limit := int(100)
	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsDevices.SearchOrgDeviceLastConfigs(ctx, orgId, &deviceType, &mac, &name, &version, &certExpiryDuration, nil, nil, &limit, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1531862583,"limit":10,"results":[{"channel_24":11,"channel_5":100,"radio_macs":["5c5b352e000a","5c5b352e000b","5c5b352e000c"],"radios":[{"band":"24","channel":11},{"band":"5","channel":100}],"secpolicy_violated":false,"ssids":["test24","test5"],"ssids_24":["test24"],"ssids_5":["test5"],"timestamp":1531855856.643369,"version":"apfw-0.2.14754-cersei-75c8","wlans":[{"auth":"psk","bands":["24"],"id":"be22bba7-8e22-e1cf-5185-b880816fe2cf","ssid":"test24","vlan_ids":["1"]},{"auth":"psk","bands":["5"],"id":"f8c18724-4118-3487-811a-f98964988604","ssid":"test5","vlan_ids":["1"]}]}],"start":1531776183,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesTestSearchOrgDeviceLastConfigs1 tests the behavior of the OrgsDevices
func TestOrgsDevicesTestSearchOrgDeviceLastConfigs1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceType := models.LastConfigDeviceTypeEnum("ap")
	mac := "5c5b53010101"
	name := "My AP"
	version := "10.0.0"
	certExpiryDuration := "2d"

	limit := int(100)
	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsDevices.SearchOrgDeviceLastConfigs(ctx, orgId, &deviceType, &mac, &name, &version, &certExpiryDuration, nil, nil, &limit, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1531862583,"limit":10,"results":[{"channel_24":11,"channel_5":100,"radio_macs":["5c5b352e000a","5c5b352e000b","5c5b352e000c"],"radios":[{"band":"24","channel":11},{"band":"5","channel":100}],"secpolicy_violated":false,"ssids":["test24","test5"],"ssids_24":["test24"],"ssids_5":["test5"],"timestamp":1531855856.643369,"version":"apfw-0.2.14754-cersei-75c8","wlans":[{"auth":"psk","bands":["24"],"id":"be22bba7-8e22-e1cf-5185-b880816fe2cf","ssid":"test24","vlan_ids":["1"]},{"auth":"psk","bands":["5"],"id":"f8c18724-4118-3487-811a-f98964988604","ssid":"test5","vlan_ids":["1"]}]}],"start":1531776183,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesTestListOrgApsMacs tests the behavior of the OrgsDevices
func TestOrgsDevicesTestListOrgApsMacs(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsDevices.ListOrgApsMacs(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"mac":"5c5b35000001","radio_macs":["5c5b35000040","5c5b35000050","5c5b35000060"]},{"mac":"5c5b45000001","radio_macs":["5c5b45000040","5c5b45000050","5c5b45000060"]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesTestListOrgApsMacs1 tests the behavior of the OrgsDevices
func TestOrgsDevicesTestListOrgApsMacs1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsDevices.ListOrgApsMacs(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"mac":"5c5b35000001","radio_macs":["5c5b35000040","5c5b35000050","5c5b35000060"]},{"mac":"5c5b45000001","radio_macs":["5c5b45000040","5c5b45000050","5c5b45000060"]}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesTestSearchOrgDevices tests the behavior of the OrgsDevices
func TestOrgsDevicesTestSearchOrgDevices(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	band24Bandwidth := int(20)
	band24Channel := int(6)
	band24Power := int(8)
	band5Bandwidth := int(20)
	band5Channel := int(50)
	band5Power := int(8)
	band6Bandwidth := int(20)
	band6Channel := int(100)
	band6Power := int(8)
	cpu := "50"
	clustered := "true"
	eth0PortSpeed := int(1000)
	evpntopoId := "7dae216d-7c98-a51b-e068-dd7d477b7216"
	extIp := "83.42.53.1"
	hostname := "my-hostname"
	ip := "192.168.1.1"
	lastConfigStatus := "success"
	lastHostname := "my-last-hostname"
	lldpMgmtAddr := "10.4.2.3"
	lldpPortId := "ge-0/0/1"
	lldpPowerAllocated := int(15)
	lldpPowerDraw := int(12)
	lldpSystemDesc := "my-lldp-system-description"
	lldpSystemName := "my-lldp-system"
	mac := "5c5b53010101"
	model := "AP43"
	mxedgeId := "7dae216d-7c98-a51b-e068-dd7d477b7216"
	mxedgeIds := "7dae216d-7c98-a51b-e068-dd7d477b7216,7dae216d-7c98-a51b-e068-dd7d477b7217"

	node := "node0"
	node0Mac := "5c5b350e0001"
	node1Mac := "5c5b350e0002"
	powerConstrained := bool(true)
	siteId := "7dae216d-7c98-a51b-e068-dd7d477b7216"
	t128agentVersion := "1.2.3"
	version := "10.0.0"
	mType := models.DeviceTypeDefaultApEnum("ap")
	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsDevices.SearchOrgDevices(ctx, orgId, &band24Bandwidth, &band24Channel, &band24Power, &band5Bandwidth, &band5Channel, &band5Power, &band6Bandwidth, &band6Channel, &band6Power, &cpu, &clustered, &eth0PortSpeed, &evpntopoId, &extIp, &hostname, &ip, &lastConfigStatus, &lastHostname, &lldpMgmtAddr, &lldpPortId, &lldpPowerAllocated, &lldpPowerDraw, &lldpSystemDesc, &lldpSystemName, &mac, &model, &mxedgeId, &mxedgeIds, nil, &node, &node0Mac, &node1Mac, &powerConstrained, &siteId, &t128agentVersion, &version, &mType, &limit, nil, nil, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":0,"limit":0,"next":"string","results":[{"hostname":["AP41-STB-3E5299-WH-2001","AP41-STB-3E5299-WH-50","AP41-STB-3E5299","5c5b353e5299"],"ip":"10.2.16.205","lldp_mgmt_addr":"10.2.10.139","lldp_port_desc":"GigabitEthernet1/0/1","lldp_port_id":"Gi1/0/1","lldp_system_desc":"Cisco IOS Software, C2960S Software (C2960S-UNIVERSALK9-M), Version 15.2(1)E1, RELEASE SOFTWARE (fc2)\nTechnical Support: https://www.cisco.com/techsupport\nCopyright (c) 1986-2013 by Cisco Systems, Inc.\nCompiled Fri 22-Nov-13 07:10 by prod_rel_team","lldp_system_name":"ME-DC-1-ACC-SW","mac":"5c5b353e5299","model":"AP41","mxedge_id":"00000000-0000-0000-1000-43a81f238391","mxtunnel_status":"down","org_id":"6748cfa6-4e12-11e6-9188-0242ac110007","power_constrained":false,"power_opmode":"","site_id":"a8178443-ecb5-461c-b854-f16627619ab3","sku":"AP41-US","timestamp":1596588619.007,"uptime":85280,"version":"0.7.20216","wlans":[{"id":"28c36fc7-dc22-4960-9d81-34087511c2e5","ssid":"Live-Demo-NAC"},{"id":"51b82e2b-f9e8-470b-a32a-cecde5501b0f","ssid":"Live-Demo"}]}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesTestSearchOrgDevices1 tests the behavior of the OrgsDevices
func TestOrgsDevicesTestSearchOrgDevices1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	band24Bandwidth := int(20)
	band24Channel := int(6)
	band24Power := int(8)
	band5Bandwidth := int(20)
	band5Channel := int(50)
	band5Power := int(8)
	band6Bandwidth := int(20)
	band6Channel := int(100)
	band6Power := int(8)
	cpu := "50"
	clustered := "true"
	eth0PortSpeed := int(1000)
	evpntopoId := "7dae216d-7c98-a51b-e068-dd7d477b7216"
	extIp := "83.42.53.1"
	hostname := "my-hostname"
	ip := "192.168.1.1"
	lastConfigStatus := "success"
	lastHostname := "my-last-hostname"
	lldpMgmtAddr := "10.4.2.3"
	lldpPortId := "ge-0/0/1"
	lldpPowerAllocated := int(15)
	lldpPowerDraw := int(12)
	lldpSystemDesc := "my-lldp-system-description"
	lldpSystemName := "my-lldp-system"
	mac := "5c5b53010101"
	model := "AP43"
	mxedgeId := "7dae216d-7c98-a51b-e068-dd7d477b7216"
	mxedgeIds := "7dae216d-7c98-a51b-e068-dd7d477b7216,7dae216d-7c98-a51b-e068-dd7d477b7217"

	node := "node0"
	node0Mac := "5c5b350e0001"
	node1Mac := "5c5b350e0002"
	powerConstrained := bool(true)
	siteId := "7dae216d-7c98-a51b-e068-dd7d477b7216"
	t128agentVersion := "1.2.3"
	version := "10.0.0"
	mType := models.DeviceTypeDefaultApEnum("ap")
	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsDevices.SearchOrgDevices(ctx, orgId, &band24Bandwidth, &band24Channel, &band24Power, &band5Bandwidth, &band5Channel, &band5Power, &band6Bandwidth, &band6Channel, &band6Power, &cpu, &clustered, &eth0PortSpeed, &evpntopoId, &extIp, &hostname, &ip, &lastConfigStatus, &lastHostname, &lldpMgmtAddr, &lldpPortId, &lldpPowerAllocated, &lldpPowerDraw, &lldpSystemDesc, &lldpSystemName, &mac, &model, &mxedgeId, &mxedgeIds, nil, &node, &node0Mac, &node1Mac, &powerConstrained, &siteId, &t128agentVersion, &version, &mType, &limit, nil, nil, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":0,"limit":0,"next":"string","results":[{"hostname":["AP41-STB-3E5299-WH-2001","AP41-STB-3E5299-WH-50","AP41-STB-3E5299","5c5b353e5299"],"ip":"10.2.16.205","lldp_mgmt_addr":"10.2.10.139","lldp_port_desc":"GigabitEthernet1/0/1","lldp_port_id":"Gi1/0/1","lldp_system_desc":"Cisco IOS Software, C2960S Software (C2960S-UNIVERSALK9-M), Version 15.2(1)E1, RELEASE SOFTWARE (fc2)\nTechnical Support: https://www.cisco.com/techsupport\nCopyright (c) 1986-2013 by Cisco Systems, Inc.\nCompiled Fri 22-Nov-13 07:10 by prod_rel_team","lldp_system_name":"ME-DC-1-ACC-SW","mac":"5c5b353e5299","model":"AP41","mxedge_id":"00000000-0000-0000-1000-43a81f238391","mxtunnel_status":"down","org_id":"6748cfa6-4e12-11e6-9188-0242ac110007","power_constrained":false,"power_opmode":"","site_id":"a8178443-ecb5-461c-b854-f16627619ab3","sku":"AP41-US","timestamp":1596588619.007,"uptime":85280,"version":"0.7.20216","wlans":[{"id":"28c36fc7-dc22-4960-9d81-34087511c2e5","ssid":"Live-Demo-NAC"},{"id":"51b82e2b-f9e8-470b-a32a-cecde5501b0f","ssid":"Live-Demo"}]}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesTestListOrgDevicesSummary tests the behavior of the OrgsDevices
func TestOrgsDevicesTestListOrgDevicesSummary(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsDevices.ListOrgDevicesSummary(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"num_aps":630,"num_gateways":6,"num_mxedges":1,"num_switches":30,"num_unassigned_aps":5,"num_unassigned_gateways":0,"num_unassigned_switches":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesTestListOrgDevicesSummary1 tests the behavior of the OrgsDevices
func TestOrgsDevicesTestListOrgDevicesSummary1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsDevices.ListOrgDevicesSummary(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"num_aps":630,"num_gateways":6,"num_mxedges":1,"num_switches":30,"num_unassigned_aps":5,"num_unassigned_gateways":0,"num_unassigned_switches":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDevicesTestGetOrgJuniperDevicesCommand tests the behavior of the OrgsDevices
func TestOrgsDevicesTestGetOrgJuniperDevicesCommand(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsDevices.GetOrgJuniperDevicesCommand(ctx, orgId, nil)
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

// TestOrgsDevicesTestGetOrgJuniperDevicesCommand1 tests the behavior of the OrgsDevices
func TestOrgsDevicesTestGetOrgJuniperDevicesCommand1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsDevices.GetOrgJuniperDevicesCommand(ctx, orgId, nil)
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
