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
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
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
    
    
    
    
    
    
    ipAddress := "192.168.1.1"
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsDevices.CountOrgDevices(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, &ipAddress, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
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
    
    
    
    
    
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := orgsDevices.CountOrgDeviceEvents(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
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
    
    
    deviceType := models.DeviceTypeEnum("ap")
    
    
    
    lastBy := "port_id"
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := orgsDevices.SearchOrgDeviceEvents(ctx, orgId, nil, nil, &deviceType, nil, nil, nil, &lastBy, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
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
    mType := models.DeviceTypeEnum("ap")
    
    
    
    limit := int(100)
    apiResponse, err := orgsDevices.CountOrgDeviceLastConfigs(ctx, orgId, &mType, nil, nil, nil, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
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
    mType := models.DeviceTypeEnum("ap")
    
    
    
    
    
    limit := int(100)
    duration := "1d"
    apiResponse, err := orgsDevices.SearchOrgDeviceLastConfigs(ctx, orgId, &mType, nil, nil, nil, nil, nil, &limit, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
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
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"mac":"5c5b35000001","radio_macs":["5c5b35000040","5c5b35000050","5c5b35000060"]},{"mac":"5c5b45000001","radio_macs":["5c5b45000040","5c5b45000050","5c5b45000060"]}]`
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
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cmd":"set system services ssh...\n...\nset system services outbound-ssh client mist ..."}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
