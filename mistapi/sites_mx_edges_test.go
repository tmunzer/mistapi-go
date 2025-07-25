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

// TestSitesMxEdgesTestListSiteMxEdges tests the behavior of the SitesMxEdges
func TestSitesMxEdgesTestListSiteMxEdges(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesMxEdges.ListSiteMxEdges(ctx, siteId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"cpu_stat":{"cpus":{"cpu0":{"idle":79,"interrupt":0,"system":4,"usage":20,"user":16},"cpu1":{"idle":93,"interrupt":0,"system":4,"usage":6,"user":1}},"idle":87,"interrupt":0,"system":5,"usage":12,"user":7},"ext_ip":"116.187.144.16","id":"387804a7-3474-85ce-15a2-f9a9684c9c90","ip_stat":{"ip":"172.16.5.3","ips":{"ens192":"172.16.5.3/24,fe81::20c:29ff:fef8:d18e/64"}},"lag_stat":{"lag0":{"active_ports":["0","1"]}},"last_seen":1547437078,"magic":"ExNpT5gi-ADR8WTFd4EiQPY3cP5WdSoD","memory_stats":{"active":1061085184,"available":4124860416,"buffers":789495808,"cached":718016512,"free":2818838528,"inactive":458158080,"swap_cached":0,"swap_free":8161062912,"swap_total":8161062912,"total":7947616256,"usage":65},"model":"ME-S2019","mxagent_registered":false,"mxcluster_id":"572586b7-f97b-a22b-526c-8b97a3f609c4","name":"Guest","num_tunnels":31,"port_stat":{"eth0":{"full_duplex":true,"lldp_stats":{"mgmt_addr":"122.16.3.11","port_desc":"GigabitEthernet4/0/16","port_id":"\u0005Gi4/0/16","system_desc":"Cisco IOS Software","system_name":"ME-DC2-DIS-SW"},"rx_bytes":2056,"rx_errors":0,"rx_pkts":670,"speed":1000,"tx_bytes":2056,"tx_pkts":670,"up":true},"eth1":{"up":false},"module":{"up":false}},"status":"connected","tunterm_registered":false,"tunterm_stat":{"monitoring_failed":false},"uptime":884221,"version":"0.1.2","virtualization_type":"VirtualizationVMware"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMxEdgesTestListSiteMxEdges1 tests the behavior of the SitesMxEdges
func TestSitesMxEdgesTestListSiteMxEdges1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesMxEdges.ListSiteMxEdges(ctx, siteId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"cpu_stat":{"cpus":{"cpu0":{"idle":79,"interrupt":0,"system":4,"usage":20,"user":16},"cpu1":{"idle":93,"interrupt":0,"system":4,"usage":6,"user":1}},"idle":87,"interrupt":0,"system":5,"usage":12,"user":7},"ext_ip":"116.187.144.16","id":"387804a7-3474-85ce-15a2-f9a9684c9c90","ip_stat":{"ip":"172.16.5.3","ips":{"ens192":"172.16.5.3/24,fe81::20c:29ff:fef8:d18e/64"}},"lag_stat":{"lag0":{"active_ports":["0","1"]}},"last_seen":1547437078,"magic":"ExNpT5gi-ADR8WTFd4EiQPY3cP5WdSoD","memory_stats":{"active":1061085184,"available":4124860416,"buffers":789495808,"cached":718016512,"free":2818838528,"inactive":458158080,"swap_cached":0,"swap_free":8161062912,"swap_total":8161062912,"total":7947616256,"usage":65},"model":"ME-S2019","mxagent_registered":false,"mxcluster_id":"572586b7-f97b-a22b-526c-8b97a3f609c4","name":"Guest","num_tunnels":31,"port_stat":{"eth0":{"full_duplex":true,"lldp_stats":{"mgmt_addr":"122.16.3.11","port_desc":"GigabitEthernet4/0/16","port_id":"\u0005Gi4/0/16","system_desc":"Cisco IOS Software","system_name":"ME-DC2-DIS-SW"},"rx_bytes":2056,"rx_errors":0,"rx_pkts":670,"speed":1000,"tx_bytes":2056,"tx_pkts":670,"up":true},"eth1":{"up":false},"module":{"up":false}},"status":"connected","tunterm_registered":false,"tunterm_stat":{"monitoring_failed":false},"uptime":884221,"version":"0.1.2","virtualization_type":"VirtualizationVMware"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMxEdgesTestCountSiteMxEdgeEvents tests the behavior of the SitesMxEdges
func TestSitesMxEdgesTestCountSiteMxEdgeEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteMxedgeEventsCountDistinctEnum("mxedge_id")
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := sitesMxEdges.CountSiteMxEdgeEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, &duration, &limit)
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

// TestSitesMxEdgesTestCountSiteMxEdgeEvents1 tests the behavior of the SitesMxEdges
func TestSitesMxEdgesTestCountSiteMxEdgeEvents1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteMxedgeEventsCountDistinctEnum("mxedge_id")
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := sitesMxEdges.CountSiteMxEdgeEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, &duration, &limit)
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

// TestSitesMxEdgesTestSearchSiteMistEdgeEvents tests the behavior of the SitesMxEdges
func TestSitesMxEdgesTestSearchSiteMistEdgeEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := sitesMxEdges.SearchSiteMistEdgeEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1694708579,"limit":10,"page":3,"results":[{"mxcluster_id":"2815c917-58e7-472f-a190-bfd44fb58d05","mxedge_id":"00000000-0000-0000-1000-020000dc585c","org_id":"f2695c32-0e83-4936-b1b2-96fc88051213","service":"tunterm","timestamp":1694678225.927,"type":"ME_SERVICE_STOPPED"}],"start":1694622179}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMxEdgesTestSearchSiteMistEdgeEvents1 tests the behavior of the SitesMxEdges
func TestSitesMxEdgesTestSearchSiteMistEdgeEvents1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := sitesMxEdges.SearchSiteMistEdgeEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1694708579,"limit":10,"page":3,"results":[{"mxcluster_id":"2815c917-58e7-472f-a190-bfd44fb58d05","mxedge_id":"00000000-0000-0000-1000-020000dc585c","org_id":"f2695c32-0e83-4936-b1b2-96fc88051213","service":"tunterm","timestamp":1694678225.927,"type":"ME_SERVICE_STOPPED"}],"start":1694622179}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMxEdgesTestDeleteSiteMxEdge tests the behavior of the SitesMxEdges
func TestSitesMxEdgesTestDeleteSiteMxEdge(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesMxEdges.DeleteSiteMxEdge(ctx, siteId, mxedgeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesMxEdgesTestGetSiteMxEdge tests the behavior of the SitesMxEdges
func TestSitesMxEdgesTestGetSiteMxEdge(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesMxEdges.GetSiteMxEdge(ctx, siteId, mxedgeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesMxEdgesTestUpdateSiteMxEdge tests the behavior of the SitesMxEdges
func TestSitesMxEdgesTestUpdateSiteMxEdge(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := sitesMxEdges.UpdateSiteMxEdge(ctx, siteId, mxedgeId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"95ddd29a-6a3c-929e-a431-51a5b09f36a6","magic":"L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD","model":"ME-100","mxagent_registered":true,"mxcluster_id":"572586b7-f97b-a22b-526c-8b97a3f609c4","mxedge_mgmt":{"mist_password":"MIST_PASSWORD","root_password":"ROOT_PASSWORD"},"name":"Guest","ntp_servers":[],"oob_ip_config":{"dns":["8.8.8.8","4.4.4.4"],"gateway":"10.2.1.254","ip":"10.2.1.10","netmask":"255.255.255.0","type":"static"},"tunterm_dhcpd_config":{"2":{"enabled":true,"servers":["11.2.3.44"]},"enabled":false,"servers":["11.2.3.4"]},"tunterm_extra_routes":{"11.0.0.0/8":{"via":"10.3.3.1"}},"tunterm_ip_config":{"dns":["8.8.8.8"],"dns_suffix":[".mist.local"],"gateway":"10.2.1.254","ip":"10.2.1.1","netmask":"255.255.255.0"}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMxEdgesTestUpdateSiteMxEdge1 tests the behavior of the SitesMxEdges
func TestSitesMxEdgesTestUpdateSiteMxEdge1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := sitesMxEdges.UpdateSiteMxEdge(ctx, siteId, mxedgeId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"95ddd29a-6a3c-929e-a431-51a5b09f36a6","magic":"L-NpT5gi-ADR8WTFd4EiQPY3cP5WdSoD","model":"ME-100","mxagent_registered":true,"mxcluster_id":"572586b7-f97b-a22b-526c-8b97a3f609c4","mxedge_mgmt":{"mist_password":"MIST_PASSWORD","root_password":"ROOT_PASSWORD"},"name":"Guest","ntp_servers":[],"oob_ip_config":{"dns":["8.8.8.8","4.4.4.4"],"gateway":"10.2.1.254","ip":"10.2.1.10","netmask":"255.255.255.0","type":"static"},"tunterm_dhcpd_config":{"2":{"enabled":true,"servers":["11.2.3.44"]},"enabled":false,"servers":["11.2.3.4"]},"tunterm_extra_routes":{"11.0.0.0/8":{"via":"10.3.3.1"}},"tunterm_ip_config":{"dns":["8.8.8.8"],"dns_suffix":[".mist.local"],"gateway":"10.2.1.254","ip":"10.2.1.1","netmask":"255.255.255.0"}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMxEdgesTestUploadSiteMxEdgeSupportFiles tests the behavior of the SitesMxEdges
func TestSitesMxEdgesTestUploadSiteMxEdgeSupportFiles(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxedgeId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesMxEdges.UploadSiteMxEdgeSupportFiles(ctx, siteId, mxedgeId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
