package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsClientsWirelessTestCountOrgWirelessClients tests the behavior of the OrgsClientsWireless
func TestOrgsClientsWirelessTestCountOrgWirelessClients(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgClientsCountDistinctEnum("device")
    mac := "5c5b53010101"
    hostname := "my-hostname"
    device := "iPhone"
    os := "Windows 10"
    model := "iPhone 8"
    ap := "5c5b53010101"
    vlan := "10"
    ssid := "MySSID"
    ipAddress := "192.168.1.1"
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsWireless.CountOrgWirelessClients(ctx, orgId, &distinct, &mac, &hostname, &device, &os, &model, &ap, &vlan, &ssid, &ipAddress, nil, nil, &duration, &limit)
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

// TestOrgsClientsWirelessTestCountOrgWirelessClients1 tests the behavior of the OrgsClientsWireless
func TestOrgsClientsWirelessTestCountOrgWirelessClients1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgClientsCountDistinctEnum("device")
    mac := "5c5b53010101"
    hostname := "my-hostname"
    device := "iPhone"
    os := "Windows 10"
    model := "iPhone 8"
    ap := "5c5b53010101"
    vlan := "10"
    ssid := "MySSID"
    ipAddress := "192.168.1.1"
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsWireless.CountOrgWirelessClients(ctx, orgId, &distinct, &mac, &hostname, &device, &os, &model, &ap, &vlan, &ssid, &ipAddress, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsWirelessTestCountOrgWirelessClientEvents tests the behavior of the OrgsClientsWireless
func TestOrgsClientsWirelessTestCountOrgWirelessClientEvents(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteClientEventsCountDistinctEnum("type")
    
    
    
    
    
    
    
    siteId, errUUID := uuid.Parse("72771e6a-6f5e-4de4-a5b9-1266c4197811")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsWireless.CountOrgWirelessClientEvents(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, &siteId, nil, nil, &duration, &limit)
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

// TestOrgsClientsWirelessTestCountOrgWirelessClientEvents1 tests the behavior of the OrgsClientsWireless
func TestOrgsClientsWirelessTestCountOrgWirelessClientEvents1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteClientEventsCountDistinctEnum("type")
    
    
    
    
    
    
    
    siteId, errUUID := uuid.Parse("72771e6a-6f5e-4de4-a5b9-1266c4197811")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsWireless.CountOrgWirelessClientEvents(ctx, orgId, &distinct, nil, nil, nil, nil, nil, nil, nil, &siteId, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsWirelessTestSearchOrgWirelessClientEvents tests the behavior of the OrgsClientsWireless
func TestOrgsClientsWirelessTestSearchOrgWirelessClientEvents(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    reasonCode := int(7)
    ssid := "MySSID"
    ap := "5c5b53010101"
    
    
    wlanId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacruleId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    apiResponse, err := orgsClientsWireless.SearchOrgWirelessClientEvents(ctx, orgId, nil, &reasonCode, &ssid, &ap, nil, nil, &wlanId, &nacruleId, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"results":[{"ap":"string","band":"24","bssid":"string","channel":0,"proto":"a","ssid":"string","text":"string","timestamp":0,"type":"string","type_code":0,"wlan_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsWirelessTestSearchOrgWirelessClientEvents1 tests the behavior of the OrgsClientsWireless
func TestOrgsClientsWirelessTestSearchOrgWirelessClientEvents1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    reasonCode := int(7)
    ssid := "MySSID"
    ap := "5c5b53010101"
    
    
    wlanId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacruleId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    apiResponse, err := orgsClientsWireless.SearchOrgWirelessClientEvents(ctx, orgId, nil, &reasonCode, &ssid, &ap, nil, nil, &wlanId, &nacruleId, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":0,"limit":0,"results":[{"ap":"string","band":"24","bssid":"string","channel":0,"proto":"a","ssid":"string","text":"string","timestamp":0,"type":"string","type_code":0,"wlan_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsWirelessTestSearchOrgWirelessClients tests the behavior of the OrgsClientsWireless
func TestOrgsClientsWirelessTestSearchOrgWirelessClients(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    siteId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mac := "5c5b53010101"
    ipAddress := "192.168.1.1"
    hostname := "my-hostname"
    band := "5"
    device := "iPhone"
    os := "Windows 10"
    model := "iPhone 8"
    ap := "5c5b53010101"
    pskId := "000000ab-00ab-00ab-00ab-0000000000ab"
    pskName := "MyPPSK"
    username := "john.doe"
    vlan := "10"
    ssid := "MySSID"
    text := "5c5b530"
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := orgsClientsWireless.SearchOrgWirelessClients(ctx, orgId, &siteId, &mac, &ipAddress, &hostname, &band, &device, &os, &model, &ap, &pskId, &pskName, &username, &vlan, &ssid, &text, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":17141231418.812,"limit":118,"next":"next8","results":[{"ap":["a83a79a947ee","003e73170b4c"],"app_version":["0.100.3"],"band":"5","device":["Mac"],"ftc":false,"hardware":"Apple Wi-Fi adapter","hostname":["hostname-a","hostname-b"],"ip":["10.5.23.43","192.168.0.2"],"last_ap":"a83a79a947ee","last_device":"Mac","last_firmware":"wl0: Jan 20 2024 04:08:41 version 20.103.12.0.8.7.171 FWID 01-e09d2675","last_hostname":"hostname-a","last_ip":"10.5.23.43","last_model":"MBP 16\\\" M1 2021","last_os":"Sonoma","last_os_version":"14.4.1 (Build 23E224)","last_psk_id":"abf7dc5c-bb51-4bb7-93b6-5547400ffe11","last_psk_name":"iot","last_ssid":"IoT SSID","last_username":"user@corp.com","last_vlan":10,"last_wlan_id":"e5d67b07-aae8-494b-8584-cbc20c8110aa","mac":"bcd074000000","mfg":"Apple","model":"MBP 16\\\" M1 2021","org_id":"1abff1aa-4571-4c1f-a409-153a1e7a7a24","os":["Sonoma"],"os_version":["14.4.1 (Build 23E224)"],"protocol":"ax","psk_id":["abf7dc5c-bb51-4bb7-93b6-5547400ffe11"],"psk_name":["iot"],"sdk_version":["0.100.3"],"site_id":"25ff5219-9be7-4db9-907d-0c9b60445147","site_ids":["25ff5219-9be7-4db9-907d-0c9b60445147"],"ssid":["IoT SSID"],"timestamp":1714124722.113,"username":["user@corp.com"],"vlan":[10]}],"start":10,"total":44}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsWirelessTestSearchOrgWirelessClients1 tests the behavior of the OrgsClientsWireless
func TestOrgsClientsWirelessTestSearchOrgWirelessClients1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    siteId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mac := "5c5b53010101"
    ipAddress := "192.168.1.1"
    hostname := "my-hostname"
    band := "5"
    device := "iPhone"
    os := "Windows 10"
    model := "iPhone 8"
    ap := "5c5b53010101"
    pskId := "000000ab-00ab-00ab-00ab-0000000000ab"
    pskName := "MyPPSK"
    username := "john.doe"
    vlan := "10"
    ssid := "MySSID"
    text := "5c5b530"
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := orgsClientsWireless.SearchOrgWirelessClients(ctx, orgId, &siteId, &mac, &ipAddress, &hostname, &band, &device, &os, &model, &ap, &pskId, &pskName, &username, &vlan, &ssid, &text, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":17141231418.812,"limit":118,"next":"next8","results":[{"ap":["a83a79a947ee","003e73170b4c"],"app_version":["0.100.3"],"band":"5","device":["Mac"],"ftc":false,"hardware":"Apple Wi-Fi adapter","hostname":["hostname-a","hostname-b"],"ip":["10.5.23.43","192.168.0.2"],"last_ap":"a83a79a947ee","last_device":"Mac","last_firmware":"wl0: Jan 20 2024 04:08:41 version 20.103.12.0.8.7.171 FWID 01-e09d2675","last_hostname":"hostname-a","last_ip":"10.5.23.43","last_model":"MBP 16\\\" M1 2021","last_os":"Sonoma","last_os_version":"14.4.1 (Build 23E224)","last_psk_id":"abf7dc5c-bb51-4bb7-93b6-5547400ffe11","last_psk_name":"iot","last_ssid":"IoT SSID","last_username":"user@corp.com","last_vlan":10,"last_wlan_id":"e5d67b07-aae8-494b-8584-cbc20c8110aa","mac":"bcd074000000","mfg":"Apple","model":"MBP 16\\\" M1 2021","org_id":"1abff1aa-4571-4c1f-a409-153a1e7a7a24","os":["Sonoma"],"os_version":["14.4.1 (Build 23E224)"],"protocol":"ax","psk_id":["abf7dc5c-bb51-4bb7-93b6-5547400ffe11"],"psk_name":["iot"],"sdk_version":["0.100.3"],"site_id":"25ff5219-9be7-4db9-907d-0c9b60445147","site_ids":["25ff5219-9be7-4db9-907d-0c9b60445147"],"ssid":["IoT SSID"],"timestamp":1714124722.113,"username":["user@corp.com"],"vlan":[10]}],"start":10,"total":44}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsWirelessTestCountOrgWirelessClientsSessions tests the behavior of the OrgsClientsWireless
func TestOrgsClientsWirelessTestCountOrgWirelessClientsSessions(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgClientSessionsCountDistinctEnum("device")
    ap := "5c5b53010101"
    
    clientFamily := "iPhone"
    clientManufacture := "Apple"
    clientModel := "iPhone 8"
    clientOs := "Windows 10"
    ssid := "MySSID"
    wlanId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsWireless.CountOrgWirelessClientsSessions(ctx, orgId, &distinct, &ap, nil, &clientFamily, &clientManufacture, &clientModel, &clientOs, &ssid, &wlanId, nil, nil, &duration, &limit)
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

// TestOrgsClientsWirelessTestCountOrgWirelessClientsSessions1 tests the behavior of the OrgsClientsWireless
func TestOrgsClientsWirelessTestCountOrgWirelessClientsSessions1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.OrgClientSessionsCountDistinctEnum("device")
    ap := "5c5b53010101"
    
    clientFamily := "iPhone"
    clientManufacture := "Apple"
    clientModel := "iPhone 8"
    clientOs := "Windows 10"
    ssid := "MySSID"
    wlanId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsClientsWireless.CountOrgWirelessClientsSessions(ctx, orgId, &distinct, &ap, nil, &clientFamily, &clientManufacture, &clientModel, &clientOs, &ssid, &wlanId, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsWirelessTestSearchOrgWirelessClientSessions tests the behavior of the OrgsClientsWireless
func TestOrgsClientsWirelessTestSearchOrgWirelessClientSessions(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ap := "5c5b53010101"
    
    clientFamily := "iPhone"
    clientManufacture := "Apple"
    clientModel := "iPhone 8"
    clientUsername := "john.doe"
    clientOs := "Windows 10"
    ssid := "MySSID"
    wlanId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
    if errUUID != nil {
        t.Error(errUUID)
    }
    pskId := "000000ab-00ab-00ab-00ab-0000000000ab"
    pskName := "MyPPSK"
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := orgsClientsWireless.SearchOrgWirelessClientSessions(ctx, orgId, &ap, nil, &clientFamily, &clientManufacture, &clientModel, &clientUsername, &clientOs, &ssid, &wlanId, &pskId, &pskName, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1513177200,"limit":10,"results":[{"ap":"5c5b350e0262","band":"5","client_manufacture":"Apple","connect":1565208388,"disconnect":1565208448,"duration":60.09423865,"mac":"b019c66c8348","org_id":"3139f2c2-fac6-11e5-8156-0242ac110006","site_id":"70e0f468-fc13-11e5-85ad-0242ac110008","ssid":"Dummy WLAN 2","tags":["disassociate"],"timestamp":1565208448.662,"wlan_id":"99bb4c74-f954-4f36-b844-6b030faffabc"}],"start":1511967600,"total":100}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsClientsWirelessTestSearchOrgWirelessClientSessions1 tests the behavior of the OrgsClientsWireless
func TestOrgsClientsWirelessTestSearchOrgWirelessClientSessions1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ap := "5c5b53010101"
    
    clientFamily := "iPhone"
    clientManufacture := "Apple"
    clientModel := "iPhone 8"
    clientUsername := "john.doe"
    clientOs := "Windows 10"
    ssid := "MySSID"
    wlanId, errUUID := uuid.Parse("7dae216d-7c98-a51b-e068-dd7d477b7216")
    if errUUID != nil {
        t.Error(errUUID)
    }
    pskId := "000000ab-00ab-00ab-00ab-0000000000ab"
    pskName := "MyPPSK"
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := orgsClientsWireless.SearchOrgWirelessClientSessions(ctx, orgId, &ap, nil, &clientFamily, &clientManufacture, &clientModel, &clientUsername, &clientOs, &ssid, &wlanId, &pskId, &pskName, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1513177200,"limit":10,"results":[{"ap":"5c5b350e0262","band":"5","client_manufacture":"Apple","connect":1565208388,"disconnect":1565208448,"duration":60.09423865,"mac":"b019c66c8348","org_id":"3139f2c2-fac6-11e5-8156-0242ac110006","site_id":"70e0f468-fc13-11e5-85ad-0242ac110008","ssid":"Dummy WLAN 2","tags":["disassociate"],"timestamp":1565208448.662,"wlan_id":"99bb4c74-f954-4f36-b844-6b030faffabc"}],"start":1511967600,"total":100}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
