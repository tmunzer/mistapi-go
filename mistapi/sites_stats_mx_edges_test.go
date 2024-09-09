package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestSitesStatsMxEdgesTestListSiteMxEdgesStats tests the behavior of the SitesStatsMxEdges
func TestSitesStatsMxEdgesTestListSiteMxEdgesStats(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesStatsMxEdges.ListSiteMxEdgesStats(ctx, siteId, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"cpu_stat":{"cpus":{"cpu0":{"idle":89,"interrupt":0,"system":8,"usage":10,"user":1},"cpu1":{"idle":81,"interrupt":0,"system":4,"usage":18,"user":13},"cpu2":{"idle":81,"interrupt":0,"system":4,"usage":18,"user":13},"cpu3":{"idle":2,"interrupt":0,"system":50,"usage":97,"user":46}},"idle":62,"interrupt":0,"system":17,"usage":37,"user":19},"created_time":1632684398,"for_site":false,"id":"00000000-0000-0000-1000-020000a80cb4","ip_stat":{"ip":"192.168.1.244","ips":{"ens18":"192.168.1.244/24,fe80::104c:ffff:fee0:caf8/64"},"macs":{"ens18":"e4434b217044"}},"lag_stat":{"lacp0":{"active_ports":["port0","port1"]}},"last_seen":1633721215,"mac":"020000a80cb4","memory_stat":{"active":394936320,"available":4699291648,"buffers":107646976,"cached":478060544,"free":4330659840,"inactive":211980288,"swap_cached":0,"swap_free":1022357504,"swap_total":1022357504,"total":8365957120,"usage":48},"model":"ME-VM","modified_time":1633643629,"mxagent_registered":true,"mxcluster_id":"678bc339-7635-4556-bbc0-e77ad493ef8b","name":"me-vm-1","num_tunnels":0,"oob_ip_config":{"dns":["8.8.8.8","1.1.1.1"],"gateway":"10.0.0.1","ip":"10.0.0.10","netmask":"255.255.255.0","type":"static"},"org_id":"11b08247-b1ee-4152-9b25-312b323ce480","port_stat":{"port0":{"full_duplex":true,"mac":"9e294e49091d","rx_bytes":646898375700,"rx_errors":0,"rx_pkts":8784449574,"speed":10000,"state":"forwarding","tx_bytes":647200748038,"tx_errors":0,"tx_pkts":8788647466,"up":true},"port1":{"full_duplex":true,"mac":"a270fe53437e","rx_bytes":647200437652,"rx_errors":0,"rx_pkts":8788644886,"speed":10000,"state":"forwarding","tx_bytes":646898681650,"tx_errors":0,"tx_pkts":8784452092,"up":true}},"sensor_stat":{},"serial":"string","service_stat":{"mxagent":{"ext_ip":"99.0.86.164","last_seen":1633721215,"package_state":"Installed","package_version":"3.1.1037-1","running_state":"Running","uptime":21240},"tunterm":{"ext_ip":"99.0.86.164","last_seen":1633721203,"package_state":"Installed","package_version":"0.1.2449+deb10","running_state":"Running","uptime":76261}},"services":["tunterm"],"site_id":"00000000-0000-0000-0000-000000000000","status":"connected","tunterm_ip_config":{"gateway":"192.168.11.1","ip":"192.168.11.91","netmask":"255.255.255.0"},"tunterm_port_config":{"downstream_ports":["0","1"],"separate_upstream_downstream":false,"upstream_ports":["0","1"]},"tunterm_registered":true,"tunterm_stat":{"monitoring_failed":false},"uptime":76281,"virtualization_type":"KVM"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesStatsMxEdgesTestGetSiteMxEdgeStats tests the behavior of the SitesStatsMxEdges
func TestSitesStatsMxEdgesTestGetSiteMxEdgeStats(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    apiResponse, err := sitesStatsMxEdges.GetSiteMxEdgeStats(ctx, siteId, mxedgeId, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cpu_stat":{"cpus":{"cpu0":{"idle":89,"interrupt":0,"system":8,"usage":10,"user":1},"cpu1":{"idle":81,"interrupt":0,"system":4,"usage":18,"user":13},"cpu2":{"idle":81,"interrupt":0,"system":4,"usage":18,"user":13},"cpu3":{"idle":2,"interrupt":0,"system":50,"usage":97,"user":46}},"idle":62,"interrupt":0,"system":17,"usage":37,"user":19},"created_time":1632684398,"for_site":false,"id":"00000000-0000-0000-1000-020000a80cb4","ip_stat":{"ip":"192.168.1.244","ips":{"ens18":"192.168.1.244/24,fe80::104c:ffff:fee0:caf8/64"},"macs":{"ens18":"e4434b217044"}},"lag_stat":{"lacp0":{"active_ports":["port0","port1"]}},"last_seen":1633721215,"mac":"020000a80cb4","memory_stat":{"active":394936320,"available":4699291648,"buffers":107646976,"cached":478060544,"free":4330659840,"inactive":211980288,"swap_cached":0,"swap_free":1022357504,"swap_total":1022357504,"total":8365957120,"usage":48},"model":"ME-VM","modified_time":1633643629,"mxagent_registered":true,"mxcluster_id":"678bc339-7635-4556-bbc0-e77ad493ef8b","name":"me-vm-1","num_tunnels":0,"oob_ip_config":{"dns":["8.8.8.8","1.1.1.1"],"gateway":"10.0.0.1","ip":"10.0.0.10","netmask":"255.255.255.0","type":"static"},"org_id":"11b08247-b1ee-4152-9b25-312b323ce480","port_stat":{"port0":{"full_duplex":true,"mac":"9e294e49091d","rx_bytes":646898375700,"rx_errors":0,"rx_pkts":8784449574,"speed":10000,"state":"forwarding","tx_bytes":647200748038,"tx_errors":0,"tx_pkts":8788647466,"up":true},"port1":{"full_duplex":true,"mac":"a270fe53437e","rx_bytes":647200437652,"rx_errors":0,"rx_pkts":8788644886,"speed":10000,"state":"forwarding","tx_bytes":646898681650,"tx_errors":0,"tx_pkts":8784452092,"up":true}},"sensor_stat":{},"serial":"string","service_stat":{"mxagent":{"ext_ip":"99.0.86.164","last_seen":1633721215,"package_state":"Installed","package_version":"3.1.1037-1","running_state":"Running","uptime":21240},"tunterm":{"ext_ip":"99.0.86.164","last_seen":1633721203,"package_state":"Installed","package_version":"0.1.2449+deb10","running_state":"Running","uptime":76261}},"services":["tunterm"],"site_id":"00000000-0000-0000-0000-000000000000","status":"connected","tunterm_ip_config":{"gateway":"192.168.11.1","ip":"192.168.11.91","netmask":"255.255.255.0"},"tunterm_port_config":{"downstream_ports":["0","1"],"separate_upstream_downstream":false,"upstream_ports":["0","1"]},"tunterm_registered":true,"tunterm_stat":{"monitoring_failed":false},"uptime":76281,"virtualization_type":"KVM"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}