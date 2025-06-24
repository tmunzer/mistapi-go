package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsMxTunnelsTestListOrgMxTunnels tests the behavior of the OrgsMxTunnels
func TestOrgsMxTunnelsTestListOrgMxTunnels(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsMxTunnels.ListOrgMxTunnels(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"hello_interval":60,"hello_retries":3,"ipsec":{"dns_servers":["172.16.0.8"],"enabled":true,"extra_routes":[{"dest":"172.16.0.0/12","next_hop":"172.16.0.1"}],"split_tunnel":true},"mxcluster_ids":["572586b7-f97b-a22b-526c-8b97a3f609c4"],"name":"HQ","protocol":"udp","vlan_ids":[3,4,5]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxTunnelsTestListOrgMxTunnels1 tests the behavior of the OrgsMxTunnels
func TestOrgsMxTunnelsTestListOrgMxTunnels1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsMxTunnels.ListOrgMxTunnels(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"hello_interval":60,"hello_retries":3,"ipsec":{"dns_servers":["172.16.0.8"],"enabled":true,"extra_routes":[{"dest":"172.16.0.0/12","next_hop":"172.16.0.1"}],"split_tunnel":true},"mxcluster_ids":["572586b7-f97b-a22b-526c-8b97a3f609c4"],"name":"HQ","protocol":"udp","vlan_ids":[3,4,5]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxTunnelsTestCreateOrgMxTunnel tests the behavior of the OrgsMxTunnels
func TestOrgsMxTunnelsTestCreateOrgMxTunnel(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Mxtunnel
    errBody := json.Unmarshal([]byte(`{"cluster_ids":["string"],"hello_interval":60,"hello_retries":7,"ipsec":{"dns_servers":["string"],"enabled":true,"extra_routes":[{"dest":"string","next_hop":"192.168.0.1"}],"split_tunnel":true,"use_mxedge":true},"vlan_ids":[0]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsMxTunnels.CreateOrgMxTunnel(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cluster_ids":["string"],"created_time":0,"for_site":true,"hello_interval":60,"hello_retries":7,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"dns_servers":["string"],"enabled":true,"extra_routes":[{"dest":"string","next_hop":"192.168.0.1"}],"split_tunnel":true,"use_mxedge":true},"modified_time":0,"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","vlan_ids":[0]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxTunnelsTestCreateOrgMxTunnel1 tests the behavior of the OrgsMxTunnels
func TestOrgsMxTunnelsTestCreateOrgMxTunnel1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Mxtunnel
    errBody := json.Unmarshal([]byte(`{"cluster_ids":["string"],"hello_interval":60,"hello_retries":7,"ipsec":{"dns_servers":["string"],"enabled":true,"extra_routes":[{"dest":"string","next_hop":"192.168.0.1"}],"split_tunnel":true,"use_mxedge":true},"vlan_ids":[0]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsMxTunnels.CreateOrgMxTunnel(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cluster_ids":["string"],"created_time":0,"for_site":true,"hello_interval":60,"hello_retries":7,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"dns_servers":["string"],"enabled":true,"extra_routes":[{"dest":"string","next_hop":"192.168.0.1"}],"split_tunnel":true,"use_mxedge":true},"modified_time":0,"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","vlan_ids":[0]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxTunnelsTestDeleteOrgMxTunnel tests the behavior of the OrgsMxTunnels
func TestOrgsMxTunnelsTestDeleteOrgMxTunnel(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsMxTunnels.DeleteOrgMxTunnel(ctx, orgId, mxtunnelId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsMxTunnelsTestGetOrgMxTunnel tests the behavior of the OrgsMxTunnels
func TestOrgsMxTunnelsTestGetOrgMxTunnel(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsMxTunnels.GetOrgMxTunnel(ctx, orgId, mxtunnelId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cluster_ids":["string"],"created_time":0,"for_site":true,"hello_interval":60,"hello_retries":7,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"dns_servers":["string"],"enabled":true,"extra_routes":[{"dest":"string","next_hop":"192.168.0.1"}],"split_tunnel":true,"use_mxedge":true},"modified_time":0,"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","vlan_ids":[0]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxTunnelsTestGetOrgMxTunnel1 tests the behavior of the OrgsMxTunnels
func TestOrgsMxTunnelsTestGetOrgMxTunnel1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsMxTunnels.GetOrgMxTunnel(ctx, orgId, mxtunnelId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cluster_ids":["string"],"created_time":0,"for_site":true,"hello_interval":60,"hello_retries":7,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"dns_servers":["string"],"enabled":true,"extra_routes":[{"dest":"string","next_hop":"192.168.0.1"}],"split_tunnel":true,"use_mxedge":true},"modified_time":0,"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","vlan_ids":[0]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxTunnelsTestUpdateOrgMxTunnel tests the behavior of the OrgsMxTunnels
func TestOrgsMxTunnelsTestUpdateOrgMxTunnel(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsMxTunnels.UpdateOrgMxTunnel(ctx, orgId, mxtunnelId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cluster_ids":["string"],"created_time":0,"for_site":true,"hello_interval":60,"hello_retries":7,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"dns_servers":["string"],"enabled":true,"extra_routes":[{"dest":"string","next_hop":"192.168.0.1"}],"split_tunnel":true,"use_mxedge":true},"modified_time":0,"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","vlan_ids":[0]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsMxTunnelsTestUpdateOrgMxTunnel1 tests the behavior of the OrgsMxTunnels
func TestOrgsMxTunnelsTestUpdateOrgMxTunnel1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsMxTunnels.UpdateOrgMxTunnel(ctx, orgId, mxtunnelId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cluster_ids":["string"],"created_time":0,"for_site":true,"hello_interval":60,"hello_retries":7,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"dns_servers":["string"],"enabled":true,"extra_routes":[{"dest":"string","next_hop":"192.168.0.1"}],"split_tunnel":true,"use_mxedge":true},"modified_time":0,"org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","vlan_ids":[0]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
