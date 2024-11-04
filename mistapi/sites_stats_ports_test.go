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
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesStatsPorts.CountSiteSwOrGwPorts(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
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

// TestSitesStatsPortsTestSearchSiteSwOrGwPorts tests the behavior of the SitesStatsPorts
func TestSitesStatsPortsTestSearchSiteSwOrGwPorts(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesStatsPorts.SearchSiteSwOrGwPorts(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1513177200,"limit":10,"results":[{"active":true,"auth_state":"init","for_site":true,"full_duplex":true,"jitter":0,"latency":0,"loss":0,"lte_iccid":"string","lte_imei":"string","lte_imsi":"string","mac":"5c4527a96580","mac_count":0,"mac_limit":0,"neighbor_mac":"64d814353400","neighbor_port_desc":"GigabitEthernet1/0/21","neighbor_system_name":"CORP-D-SW-2","org_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","poe_disabled":true,"poe_mode":"802.3af","poe_on":true,"port_id":"ge-0/0/0","port_mac":"5c4527a96580","port_usage":"lan","power_draw":0,"rx_bcast_pkts":0,"rx_bps":0,"rx_bytes":4563443626,"rx_errors":0,"rx_mcast_pkts":0,"rx_pkts":0,"site_id":"c1698122-c14c-11e5-8e81-1258369c38a9","speed":1000,"stp_role":"designated","stp_state":"forwarding","tx_bcast_pkts":0,"tx_bps":0,"tx_bytes":11299516780,"tx_errors":0,"tx_mcast_pkts":0,"tx_pkts":492176,"type":"gateway","up":true,"xcvr_part_number":"string"}],"start":1511967600,"total":100}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsPortsTestCountSiteSwitchPorts tests the behavior of the SitesStatsPorts
func TestSitesStatsPortsTestCountSiteSwitchPorts(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteSwitchPortsCountDistinctEnum("mac")
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesStatsPorts.CountSiteSwitchPorts(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
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

// TestSitesStatsPortsTestSearchSiteSwitchPorts tests the behavior of the SitesStatsPorts
func TestSitesStatsPortsTestSearchSiteSwitchPorts(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesStatsPorts.SearchSiteSwitchPorts(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1513177200,"limit":10,"results":[{"active":true,"auth_state":"init","for_site":true,"full_duplex":true,"jitter":0,"latency":0,"loss":0,"lte_iccid":"string","lte_imei":"string","lte_imsi":"string","mac":"5c4527a96580","mac_count":0,"mac_limit":0,"neighbor_mac":"64d814353400","neighbor_port_desc":"GigabitEthernet1/0/21","neighbor_system_name":"CORP-D-SW-2","org_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","poe_disabled":true,"poe_mode":"802.3af","poe_on":true,"port_id":"ge-0/0/0","port_mac":"5c4527a96580","port_usage":"lan","power_draw":0,"rx_bcast_pkts":0,"rx_bps":0,"rx_bytes":4563443626,"rx_errors":0,"rx_mcast_pkts":0,"rx_pkts":0,"site_id":"c1698122-c14c-11e5-8e81-1258369c38a9","speed":1000,"stp_role":"designated","stp_state":"forwarding","tx_bcast_pkts":0,"tx_bps":0,"tx_bytes":11299516780,"tx_errors":0,"tx_mcast_pkts":0,"tx_pkts":492176,"type":"gateway","up":true,"xcvr_part_number":"string"}],"start":1511967600,"total":100}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
