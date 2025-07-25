// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestSitesNetworksTestListSiteNetworksDerived tests the behavior of the SitesNetworks
func TestSitesNetworksTestListSiteNetworksDerived(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resolve := bool(false)
    apiResponse, err := sitesNetworks.ListSiteNetworksDerived(ctx, siteId, &resolve)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"disallow_mist_services":false,"gateway":"192.168.70.1","hosts":{"property1":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"},"property2":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"}},"id":"497f6eca-6276-4993-bfeb-53cbbbba6f13","internal_access":{"enabled":true},"internet_access":{"create_simple_service_policy":false,"destination_nat":{"property1":{"internal_ip":"192.168.70.30","name":"web server","port":"443"},"property2":{"internal_ip":"192.168.70.30","name":"web server","port":"443"}},"enabled":true,"restricted":false,"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"printer-1"},"property2":{"internal_ip":"192.168.70.3","name":"printer-1"}}},"isolation":true,"modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","subnet":"192.168.70.0/24","tenants":{"property1":{"addresses":["10.10.10.10"]},"property2":{"addresses":["10.10.10.45"]}},"vlan_id":10,"vpn_access":{"property1":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":"443","to":"192.168.70.5/30"},"property2":{"name":"web server","port":"443","to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"},"property2":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":"443","to":"192.168.70.5/30"},"property2":{"name":"web server","port":"443","to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"}}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesNetworksTestListSiteNetworksDerived1 tests the behavior of the SitesNetworks
func TestSitesNetworksTestListSiteNetworksDerived1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resolve := bool(false)
    apiResponse, err := sitesNetworks.ListSiteNetworksDerived(ctx, siteId, &resolve)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"disallow_mist_services":false,"gateway":"192.168.70.1","hosts":{"property1":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"},"property2":{"external_ips":"172.16.10.32-172.16.10.35","ips":"192.168.70.32-192.168.70.35"}},"id":"497f6eca-6276-4993-bfeb-53cbbbba6f13","internal_access":{"enabled":true},"internet_access":{"create_simple_service_policy":false,"destination_nat":{"property1":{"internal_ip":"192.168.70.30","name":"web server","port":"443"},"property2":{"internal_ip":"192.168.70.30","name":"web server","port":"443"}},"enabled":true,"restricted":false,"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"printer-1"},"property2":{"internal_ip":"192.168.70.3","name":"printer-1"}}},"isolation":true,"modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","subnet":"192.168.70.0/24","tenants":{"property1":{"addresses":["10.10.10.10"]},"property2":{"addresses":["10.10.10.45"]}},"vlan_id":10,"vpn_access":{"property1":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":"443","to":"192.168.70.5/30"},"property2":{"name":"web server","port":"443","to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"},"property2":{"allow_ping":true,"destination_nat":{"property1":{"name":"web server","port":"443","to":"192.168.70.5/30"},"property2":{"name":"web server","port":"443","to":"192.168.70.5/30"}},"nat_pool":"172.16.0.0/26","routed":true,"source_nat":{"external_ip":"172.16.0.8/30"},"static_nat":{"property1":{"internal_ip":"192.168.70.3","name":"pos_station-1"},"property2":{"internal_ip":"192.168.70.3","name":"pos_station-1"}},"summarized_subnet":"172.16.0.0/16"}}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
