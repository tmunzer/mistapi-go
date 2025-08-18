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

// TestSitesClientsWirelessTestCountSiteWirelessClients tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestCountSiteWirelessClients(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SiteClientsCountDistinctEnum("device")

	ipAddress := "192.168.1.1"

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesClientsWireless.CountSiteWirelessClients(ctx, siteId, &distinct, nil, nil, &ipAddress, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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

// TestSitesClientsWirelessTestCountSiteWirelessClients1 tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestCountSiteWirelessClients1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SiteClientsCountDistinctEnum("device")

	ipAddress := "192.168.1.1"

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesClientsWireless.CountSiteWirelessClients(ctx, siteId, &distinct, nil, nil, &ipAddress, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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

// TestSitesClientsWirelessTestCountSiteWirelessClientEvents tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestCountSiteWirelessClientEvents(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SiteClientEventsCountDistinctEnum("type")

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesClientsWireless.CountSiteWirelessClientEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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

// TestSitesClientsWirelessTestCountSiteWirelessClientEvents1 tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestCountSiteWirelessClientEvents1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SiteClientEventsCountDistinctEnum("type")

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesClientsWireless.CountSiteWirelessClientEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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

// TestSitesClientsWirelessTestSearchSiteWirelessClientEvents tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestSearchSiteWirelessClientEvents(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesClientsWireless.SearchSiteWirelessClientEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":0,"limit":0,"results":[{"ap":"string","band":"24","bssid":"string","channel":0,"proto":"a","ssid":"string","text":"string","timestamp":0,"type":"string","type_code":0,"wlan_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesClientsWirelessTestSearchSiteWirelessClientEvents1 tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestSearchSiteWirelessClientEvents1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesClientsWireless.SearchSiteWirelessClientEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":0,"limit":0,"results":[{"ap":"string","band":"24","bssid":"string","channel":0,"proto":"a","ssid":"string","text":"string","timestamp":0,"type":"string","type_code":0,"wlan_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesClientsWirelessTestSearchSiteWirelessClients tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestSearchSiteWirelessClients(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	ipAddress := "192.168.1.1"

	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesClientsWireless.SearchSiteWirelessClients(ctx, siteId, nil, &ipAddress, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":17141231418.812,"limit":118,"next":"next8","results":[{"ap":["a83a79a947ee","003e73170b4c"],"app_version":["0.100.3"],"band":"5","device":["Mac"],"ftc":false,"hardware":"Apple Wi-Fi adapter","hostname":["hostname-a","hostname-b"],"ip":["10.5.23.43","192.168.0.2"],"last_ap":"a83a79a947ee","last_device":"Mac","last_firmware":"wl0: Jan 20 2024 04:08:41 version 20.103.12.0.8.7.171 FWID 01-e09d2675","last_hostname":"hostname-a","last_ip":"10.5.23.43","last_model":"MBP 16\\\" M1 2021","last_os":"Sonoma","last_os_version":"14.4.1 (Build 23E224)","last_psk_id":"abf7dc5c-bb51-4bb7-93b6-5547400ffe11","last_psk_name":"iot","last_ssid":"IoT SSID","last_username":"user@corp.com","last_vlan":10,"last_wlan_id":"e5d67b07-aae8-494b-8584-cbc20c8110aa","mac":"bcd074000000","mfg":"Apple","model":"MBP 16\\\" M1 2021","org_id":"1abff1aa-4571-4c1f-a409-153a1e7a7a24","os":["Sonoma"],"os_version":["14.4.1 (Build 23E224)"],"protocol":"ax","psk_id":["abf7dc5c-bb51-4bb7-93b6-5547400ffe11"],"psk_name":["iot"],"sdk_version":["0.100.3"],"site_id":"25ff5219-9be7-4db9-907d-0c9b60445147","site_ids":["25ff5219-9be7-4db9-907d-0c9b60445147"],"ssid":["IoT SSID"],"timestamp":1714124722.113,"username":["user@corp.com"],"vlan":[10]}],"start":10,"total":44}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesClientsWirelessTestSearchSiteWirelessClients1 tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestSearchSiteWirelessClients1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	ipAddress := "192.168.1.1"

	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesClientsWireless.SearchSiteWirelessClients(ctx, siteId, nil, &ipAddress, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":17141231418.812,"limit":118,"next":"next8","results":[{"ap":["a83a79a947ee","003e73170b4c"],"app_version":["0.100.3"],"band":"5","device":["Mac"],"ftc":false,"hardware":"Apple Wi-Fi adapter","hostname":["hostname-a","hostname-b"],"ip":["10.5.23.43","192.168.0.2"],"last_ap":"a83a79a947ee","last_device":"Mac","last_firmware":"wl0: Jan 20 2024 04:08:41 version 20.103.12.0.8.7.171 FWID 01-e09d2675","last_hostname":"hostname-a","last_ip":"10.5.23.43","last_model":"MBP 16\\\" M1 2021","last_os":"Sonoma","last_os_version":"14.4.1 (Build 23E224)","last_psk_id":"abf7dc5c-bb51-4bb7-93b6-5547400ffe11","last_psk_name":"iot","last_ssid":"IoT SSID","last_username":"user@corp.com","last_vlan":10,"last_wlan_id":"e5d67b07-aae8-494b-8584-cbc20c8110aa","mac":"bcd074000000","mfg":"Apple","model":"MBP 16\\\" M1 2021","org_id":"1abff1aa-4571-4c1f-a409-153a1e7a7a24","os":["Sonoma"],"os_version":["14.4.1 (Build 23E224)"],"protocol":"ax","psk_id":["abf7dc5c-bb51-4bb7-93b6-5547400ffe11"],"psk_name":["iot"],"sdk_version":["0.100.3"],"site_id":"25ff5219-9be7-4db9-907d-0c9b60445147","site_ids":["25ff5219-9be7-4db9-907d-0c9b60445147"],"ssid":["IoT SSID"],"timestamp":1714124722.113,"username":["user@corp.com"],"vlan":[10]}],"start":10,"total":44}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesClientsWirelessTestCountSiteWirelessClientSessions tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestCountSiteWirelessClientSessions(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SiteClientSessionsCountDistinctEnum("mac")

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesClientsWireless.CountSiteWirelessClientSessions(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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

// TestSitesClientsWirelessTestCountSiteWirelessClientSessions1 tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestCountSiteWirelessClientSessions1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SiteClientSessionsCountDistinctEnum("mac")

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesClientsWireless.CountSiteWirelessClientSessions(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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

// TestSitesClientsWirelessTestSearchSiteWirelessClientSessions tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestSearchSiteWirelessClientSessions(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	pskId := "000000ab-00ab-00ab-00ab-0000000000ab"

	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesClientsWireless.SearchSiteWirelessClientSessions(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, &pskId, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1513177200,"limit":10,"results":[{"ap":"5c5b350e0262","band":"5","client_manufacture":"Apple","connect":1565208388,"disconnect":1565208448,"duration":60.09423865,"mac":"b019c66c8348","org_id":"3139f2c2-fac6-11e5-8156-0242ac110006","site_id":"70e0f468-fc13-11e5-85ad-0242ac110008","ssid":"Dummy WLAN 2","tags":["disassociate"],"timestamp":1565208448.662,"wlan_id":"99bb4c74-f954-4f36-b844-6b030faffabc"}],"start":1511967600,"total":100}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesClientsWirelessTestSearchSiteWirelessClientSessions1 tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestSearchSiteWirelessClientSessions1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	pskId := "000000ab-00ab-00ab-00ab-0000000000ab"

	limit := int(100)

	duration := "1d"
	apiResponse, err := sitesClientsWireless.SearchSiteWirelessClientSessions(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, &pskId, nil, &limit, nil, nil, &duration)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1513177200,"limit":10,"results":[{"ap":"5c5b350e0262","band":"5","client_manufacture":"Apple","connect":1565208388,"disconnect":1565208448,"duration":60.09423865,"mac":"b019c66c8348","org_id":"3139f2c2-fac6-11e5-8156-0242ac110006","site_id":"70e0f468-fc13-11e5-85ad-0242ac110008","ssid":"Dummy WLAN 2","tags":["disassociate"],"timestamp":1565208448.662,"wlan_id":"99bb4c74-f954-4f36-b844-6b030faffabc"}],"start":1511967600,"total":100}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesClientsWirelessTestGetSiteEventsForClient tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestGetSiteEventsForClient(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	clientMac := "0000000000ab"

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesClientsWireless.GetSiteEventsForClient(ctx, siteId, clientMac, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1513176951,"limit":10,"results":[{"ap":"5c5b350eb31b","band":"5","bssid":"5c5b350918f1","channel":149,"proto":"ac","ssid":"Guest","text":"Status code 0 \"Successful\"","timestamp":1513358874.667,"type":"CLIENT_DNS_OK","type_code":15,"wlan_id":"be22bba7-8e22-e1cf-5185-b880816fe2cf"}],"start":1512572151,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesClientsWirelessTestGetSiteEventsForClient1 tests the behavior of the SitesClientsWireless
func TestSitesClientsWirelessTestGetSiteEventsForClient1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	clientMac := "0000000000ab"

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesClientsWireless.GetSiteEventsForClient(ctx, siteId, clientMac, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1513176951,"limit":10,"results":[{"ap":"5c5b350eb31b","band":"5","bssid":"5c5b350918f1","channel":149,"proto":"ac","ssid":"Guest","text":"Status code 0 \"Successful\"","timestamp":1513358874.667,"type":"CLIENT_DNS_OK","type_code":15,"wlan_id":"be22bba7-8e22-e1cf-5185-b880816fe2cf"}],"start":1512572151,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
