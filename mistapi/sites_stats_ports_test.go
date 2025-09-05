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

// TestSitesStatsPortsTestCountSiteSwOrGwPorts tests the behavior of the SitesStatsPorts
func TestSitesStatsPortsTestCountSiteSwOrGwPorts(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SitePortsCountDistinctEnum("mac")
	fullDuplex := bool(true)
	mac := "5c5b350e0410"
	neighborMac := "5c5b350e0410"
	neighborPortDesc := "ge-2/0/39"
	neighborSystemName := "Kumar-Acc-SW.mist.local"
	poeDisabled := bool(false)
	poeMode := "802.3at"
	poeOn := bool(true)
	portId := "ge-2/0/39"
	portMac := "5c5b350e0410"
	powerDraw := float64(15.4)
	txPkts := int(1000000)
	rxPkts := int(1000000)
	rxBytes := int(1000000)
	txBps := int(1000000)
	rxBps := int(1000000)
	txMcastPkts := int(1000000)
	txBcastPkts := int(1000000)
	rxMcastPkts := int(1000000)
	rxBcastPkts := int(1000000)
	speed := int(1000000000)

	up := bool(true)

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesStatsPorts.CountSiteSwOrGwPorts(ctx, siteId, &distinct, &fullDuplex, &mac, &neighborMac, &neighborPortDesc, &neighborSystemName, &poeDisabled, &poeMode, &poeOn, &portId, &portMac, &powerDraw, &txPkts, &rxPkts, &rxBytes, &txBps, &rxBps, &txMcastPkts, &txBcastPkts, &rxMcastPkts, &rxBcastPkts, &speed, nil, nil, nil, &up, nil, nil, &duration, &limit)
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

// TestSitesStatsPortsTestCountSiteSwOrGwPorts1 tests the behavior of the SitesStatsPorts
func TestSitesStatsPortsTestCountSiteSwOrGwPorts1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SitePortsCountDistinctEnum("mac")
	fullDuplex := bool(true)
	mac := "5c5b350e0410"
	neighborMac := "5c5b350e0410"
	neighborPortDesc := "ge-2/0/39"
	neighborSystemName := "Kumar-Acc-SW.mist.local"
	poeDisabled := bool(false)
	poeMode := "802.3at"
	poeOn := bool(true)
	portId := "ge-2/0/39"
	portMac := "5c5b350e0410"
	powerDraw := float64(15.4)
	txPkts := int(1000000)
	rxPkts := int(1000000)
	rxBytes := int(1000000)
	txBps := int(1000000)
	rxBps := int(1000000)
	txMcastPkts := int(1000000)
	txBcastPkts := int(1000000)
	rxMcastPkts := int(1000000)
	rxBcastPkts := int(1000000)
	speed := int(1000000000)

	up := bool(true)

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesStatsPorts.CountSiteSwOrGwPorts(ctx, siteId, &distinct, &fullDuplex, &mac, &neighborMac, &neighborPortDesc, &neighborSystemName, &poeDisabled, &poeMode, &poeOn, &portId, &portMac, &powerDraw, &txPkts, &rxPkts, &rxBytes, &txBps, &rxBps, &txMcastPkts, &txBcastPkts, &rxMcastPkts, &rxBcastPkts, &speed, nil, nil, nil, &up, nil, nil, &duration, &limit)
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

// TestSitesStatsPortsTestSearchSiteSwOrGwPorts tests the behavior of the SitesStatsPorts
func TestSitesStatsPortsTestSearchSiteSwOrGwPorts(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	fullDuplex := bool(true)
	disabled := bool(false)
	mac := "5c5b350e0410"

	neighborMac := "5c5b350e0410"
	neighborPortDesc := "ge-2/0/39"
	neighborSystemName := "Kumar-Acc-SW.mist.local"
	poeDisabled := bool(false)
	poeMode := "802.3at"
	poeOn := bool(true)
	portId := "ge-2/0/39"
	portMac := "5c5b350e0410"
	powerDraw := float64(15.4)
	txPkts := int(1000000)
	rxPkts := int(1000000)
	rxBytes := int(1000000)
	txBps := int(1000000)
	rxBps := int(1000000)
	txErrors := int(0)
	rxErrors := int(0)
	txMcastPkts := int(100)
	txBcastPkts := int(100)
	rxMcastPkts := int(100)
	rxBcastPkts := int(100)
	speed := int(1000)
	macLimit := int(1000)
	macCount := int(10)
	up := bool(true)
	active := bool(true)
	jitter := float64(0.456)
	loss := float64(0.01)
	latency := float64(0.123)

	xcvrPartNumber := "SFP-10G-SR"

	lteImsi := "310260000000001"
	lteIccid := "89014103211118510720"
	lteImei := "123456789012345"

	limit := int(100)

	duration := "1d"
	sort := "timestamp"
	apiResponse, err := sitesStatsPorts.SearchSiteSwOrGwPorts(ctx, siteId, &fullDuplex, &disabled, &mac, nil, &neighborMac, &neighborPortDesc, &neighborSystemName, &poeDisabled, &poeMode, &poeOn, &portId, &portMac, &powerDraw, &txPkts, &rxPkts, &rxBytes, &txBps, &rxBps, &txErrors, &rxErrors, &txMcastPkts, &txBcastPkts, &rxMcastPkts, &rxBcastPkts, &speed, &macLimit, &macCount, &up, &active, &jitter, &loss, &latency, nil, nil, &xcvrPartNumber, nil, &lteImsi, &lteIccid, &lteImei, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1513177200,"limit":10,"results":[{"active":true,"auth_state":"init","for_site":true,"full_duplex":true,"jitter":0,"latency":0,"loss":0,"lte_iccid":"string","lte_imei":"string","lte_imsi":"string","mac":"5c4527a96580","mac_count":0,"mac_limit":0,"neighbor_mac":"64d814353400","neighbor_port_desc":"GigabitEthernet1/0/21","neighbor_system_name":"CORP-D-SW-2","org_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","poe_disabled":true,"poe_mode":"802.3af","poe_on":true,"port_id":"ge-0/0/0","port_mac":"5c4527a96580","port_usage":"lan","power_draw":0,"rx_bcast_pkts":0,"rx_bps":0,"rx_bytes":4563443626,"rx_errors":0,"rx_mcast_pkts":0,"rx_pkts":0,"site_id":"c1698122-c14c-11e5-8e81-1258369c38a9","speed":1000,"stp_role":"designated","stp_state":"forwarding","tx_bcast_pkts":0,"tx_bps":0,"tx_bytes":11299516780,"tx_errors":0,"tx_mcast_pkts":0,"tx_pkts":492176,"type":"gateway","up":true,"xcvr_part_number":"string"}],"start":1511967600,"total":100}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsPortsTestSearchSiteSwOrGwPorts1 tests the behavior of the SitesStatsPorts
func TestSitesStatsPortsTestSearchSiteSwOrGwPorts1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	fullDuplex := bool(true)
	disabled := bool(false)
	mac := "5c5b350e0410"

	neighborMac := "5c5b350e0410"
	neighborPortDesc := "ge-2/0/39"
	neighborSystemName := "Kumar-Acc-SW.mist.local"
	poeDisabled := bool(false)
	poeMode := "802.3at"
	poeOn := bool(true)
	portId := "ge-2/0/39"
	portMac := "5c5b350e0410"
	powerDraw := float64(15.4)
	txPkts := int(1000000)
	rxPkts := int(1000000)
	rxBytes := int(1000000)
	txBps := int(1000000)
	rxBps := int(1000000)
	txErrors := int(0)
	rxErrors := int(0)
	txMcastPkts := int(100)
	txBcastPkts := int(100)
	rxMcastPkts := int(100)
	rxBcastPkts := int(100)
	speed := int(1000)
	macLimit := int(1000)
	macCount := int(10)
	up := bool(true)
	active := bool(true)
	jitter := float64(0.456)
	loss := float64(0.01)
	latency := float64(0.123)

	xcvrPartNumber := "SFP-10G-SR"

	lteImsi := "310260000000001"
	lteIccid := "89014103211118510720"
	lteImei := "123456789012345"

	limit := int(100)

	duration := "1d"
	sort := "timestamp"
	apiResponse, err := sitesStatsPorts.SearchSiteSwOrGwPorts(ctx, siteId, &fullDuplex, &disabled, &mac, nil, &neighborMac, &neighborPortDesc, &neighborSystemName, &poeDisabled, &poeMode, &poeOn, &portId, &portMac, &powerDraw, &txPkts, &rxPkts, &rxBytes, &txBps, &rxBps, &txErrors, &rxErrors, &txMcastPkts, &txBcastPkts, &rxMcastPkts, &rxBcastPkts, &speed, &macLimit, &macCount, &up, &active, &jitter, &loss, &latency, nil, nil, &xcvrPartNumber, nil, &lteImsi, &lteIccid, &lteImei, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1513177200,"limit":10,"results":[{"active":true,"auth_state":"init","for_site":true,"full_duplex":true,"jitter":0,"latency":0,"loss":0,"lte_iccid":"string","lte_imei":"string","lte_imsi":"string","mac":"5c4527a96580","mac_count":0,"mac_limit":0,"neighbor_mac":"64d814353400","neighbor_port_desc":"GigabitEthernet1/0/21","neighbor_system_name":"CORP-D-SW-2","org_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","poe_disabled":true,"poe_mode":"802.3af","poe_on":true,"port_id":"ge-0/0/0","port_mac":"5c4527a96580","port_usage":"lan","power_draw":0,"rx_bcast_pkts":0,"rx_bps":0,"rx_bytes":4563443626,"rx_errors":0,"rx_mcast_pkts":0,"rx_pkts":0,"site_id":"c1698122-c14c-11e5-8e81-1258369c38a9","speed":1000,"stp_role":"designated","stp_state":"forwarding","tx_bcast_pkts":0,"tx_bps":0,"tx_bytes":11299516780,"tx_errors":0,"tx_mcast_pkts":0,"tx_pkts":492176,"type":"gateway","up":true,"xcvr_part_number":"string"}],"start":1511967600,"total":100}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
