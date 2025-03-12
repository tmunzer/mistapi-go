package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesEVPNTopologiesTestListSiteEvpnTopologies tests the behavior of the SitesEVPNTopologies
func TestSitesEVPNTopologiesTestListSiteEvpnTopologies(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesEvpnTopologies.ListSiteEvpnTopologies(ctx, siteId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":1736421230,"evpn_options":{"auto_loopback_subnet":"172.16.192.0/24","auto_loopback_subnet6":"fd33:ab00:2::/64","auto_router_id_subnet":"172.16.254.0/23","core_as_border":true,"overlay":{"as":65000},"per_vlan_vga_v4_mac":false,"routed_at":"core","underlay":{"as_base":65001,"subnet":"10.255.240.0/20","use_ipv6":false}},"for_site":false,"id":"764fb173-94f9-447c-8454-def62e5a999f","modified_time":1736421230,"name":"tert","org_id":"3a2627d7-bfbc-45af-b85d-8841581c6d63","pod_names":{"1":"Pod 1"},"site_id":"00000000-0000-0000-0000-000000000000","version":6}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesEVPNTopologiesTestCreateSiteEvpnTopology tests the behavior of the SitesEVPNTopologies
func TestSitesEVPNTopologiesTestCreateSiteEvpnTopology(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.EvpnTopology
    errBody := json.Unmarshal([]byte(`{"name":"CC","overwrite":true,"pod_names":{"1":"default","2":"default"},"switches":[{"mac":"5c5b35000003","role":"collapsed-core"},{"mac":"5c5b35000004","role":"collapsed-core"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesEvpnTopologies.CreateSiteEvpnTopology(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"9197ec96-4c8d-529f-c595-035895e688b2","name":"CC","overwrite":true,"pod_names":{"1":"default","2":"default"},"switches":[{"deviceprofile_id":"6a1deab1-96df-4fa2-8455-d5253f943d06","downlink_ips":["10.255.240.6","10.255.240.8"],"downlinks":["5c5b35000007","5c5b35000008"],"esilaglinks":["5c5b3500000f"],"evpn_id":1,"mac":"5c5b35000003","model":"QFX10002-36Q","role":"collapsed-core","site_id":"1916d52a-4a90-11e5-8b45-1258369c38a9","uplinks":["5c5b35000005","5c5b35000006"]}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesEVPNTopologiesTestDeleteSiteEvpnTopology tests the behavior of the SitesEVPNTopologies
func TestSitesEVPNTopologiesTestDeleteSiteEvpnTopology(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    evpnTopologyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesEvpnTopologies.DeleteSiteEvpnTopology(ctx, siteId, evpnTopologyId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesEVPNTopologiesTestGetSiteEvpnTopology tests the behavior of the SitesEVPNTopologies
func TestSitesEVPNTopologiesTestGetSiteEvpnTopology(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    evpnTopologyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesEvpnTopologies.GetSiteEvpnTopology(ctx, siteId, evpnTopologyId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"9197ec96-4c8d-529f-c595-035895e688b2","name":"CC","overwrite":true,"pod_names":{"1":"default","2":"default"},"switches":[{"deviceprofile_id":"6a1deab1-96df-4fa2-8455-d5253f943d06","downlink_ips":["10.255.240.6","10.255.240.8"],"downlinks":["5c5b35000007","5c5b35000008"],"esilaglinks":["5c5b3500000f"],"evpn_id":1,"mac":"5c5b35000003","model":"QFX10002-36Q","role":"collapsed-core","site_id":"1916d52a-4a90-11e5-8b45-1258369c38a9","uplinks":["5c5b35000005","5c5b35000006"]}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesEVPNTopologiesTestUpdateSiteEvpnTopology tests the behavior of the SitesEVPNTopologies
func TestSitesEVPNTopologiesTestUpdateSiteEvpnTopology(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    evpnTopologyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.EvpnTopology
    errBody := json.Unmarshal([]byte(`{"overwrite":false,"switches":[{"mac":"5c5b35000003","role":"collapsed-core"},{"mac":"5c5b35000004","role":"none"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesEvpnTopologies.UpdateSiteEvpnTopology(ctx, siteId, evpnTopologyId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"id":"9197ec96-4c8d-529f-c595-035895e688b2","name":"CC","overwrite":true,"pod_names":{"1":"default","2":"default"},"switches":[{"deviceprofile_id":"6a1deab1-96df-4fa2-8455-d5253f943d06","downlink_ips":["10.255.240.6","10.255.240.8"],"downlinks":["5c5b35000007","5c5b35000008"],"esilaglinks":["5c5b3500000f"],"evpn_id":1,"mac":"5c5b35000003","model":"QFX10002-36Q","role":"collapsed-core","site_id":"1916d52a-4a90-11e5-8b45-1258369c38a9","uplinks":["5c5b35000005","5c5b35000006"]}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
