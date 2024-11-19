package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsEVPNTopologiesTestListOrgEvpnTopologies tests the behavior of the OrgsEVPNTopologies
func TestOrgsEVPNTopologiesTestListOrgEvpnTopologies(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsEvpnTopologies.ListOrgEvpnTopologies(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"id":"9197ec96-4c8d-529f-c595-035895e688b2","name":"CC","overwrite":true,"pod_names":{"1":"default","2":"default"},"switches":[{"deviceprofile_id":"6a1deab1-96df-4fa2-8455-d5253f943d06","downlink_ips":["10.255.240.6","10.255.240.8"],"downlinks":["5c5b35000007","5c5b35000008"],"esilaglinks":["5c5b3500000f"],"evpn_id":1,"mac":"5c5b35000003","model":"QFX10002-36Q","role":"collapsed-core","site_id":"1916d52a-4a90-11e5-8b45-1258369c38a9","uplinks":["5c5b35000005","5c5b35000006"]}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsEVPNTopologiesTestCreateOrgEvpnTopology tests the behavior of the OrgsEVPNTopologies
func TestOrgsEVPNTopologiesTestCreateOrgEvpnTopology(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.EvpnTopology
    errBody := json.Unmarshal([]byte(`{"name":"CC","pod_names":{"1":"default","2":"default"},"switches":[{"mac":"5c5b35000003","role":"collapsed-core"},{"mac":"5c5b35000004","role":"collapsed-core"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsEvpnTopologies.CreateOrgEvpnTopology(ctx, orgId, &body)
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

// TestOrgsEVPNTopologiesTestDeleteOrgEvpnTopology tests the behavior of the OrgsEVPNTopologies
func TestOrgsEVPNTopologiesTestDeleteOrgEvpnTopology(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    evpnTopologyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsEvpnTopologies.DeleteOrgEvpnTopology(ctx, orgId, evpnTopologyId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsEVPNTopologiesTestGetOrgEvpnTopology tests the behavior of the OrgsEVPNTopologies
func TestOrgsEVPNTopologiesTestGetOrgEvpnTopology(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    evpnTopologyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsEvpnTopologies.GetOrgEvpnTopology(ctx, orgId, evpnTopologyId)
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

// TestOrgsEVPNTopologiesTestUpdateOrgEvpnTopology tests the behavior of the OrgsEVPNTopologies
func TestOrgsEVPNTopologiesTestUpdateOrgEvpnTopology(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
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
    apiResponse, err := orgsEvpnTopologies.UpdateOrgEvpnTopology(ctx, orgId, evpnTopologyId, &body)
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
