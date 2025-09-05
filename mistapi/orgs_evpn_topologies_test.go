// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
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
	forSite := models.MxedgeForSiteEnum("any")
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsEvpnTopologies.ListOrgEvpnTopologies(ctx, orgId, &forSite, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":1736421230,"evpn_options":{"auto_loopback_subnet":"172.16.192.0/24","auto_loopback_subnet6":"fd33:ab00:2::/64","auto_router_id_subnet":"172.16.254.0/23","core_as_border":true,"overlay":{"as":65000},"per_vlan_vga_v4_mac":false,"routed_at":"core","underlay":{"as_base":65001,"subnet":"10.255.240.0/20","use_ipv6":false}},"for_site":false,"id":"764fb173-94f9-447c-8454-def62e5a999f","modified_time":1736421230,"name":"tert","org_id":"3a2627d7-bfbc-45af-b85d-8841581c6d63","pod_names":{"1":"Pod 1"},"site_id":"00000000-0000-0000-0000-000000000000","version":6}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsEVPNTopologiesTestListOrgEvpnTopologies1 tests the behavior of the OrgsEVPNTopologies
func TestOrgsEVPNTopologiesTestListOrgEvpnTopologies1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	forSite := models.MxedgeForSiteEnum("any")
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsEvpnTopologies.ListOrgEvpnTopologies(ctx, orgId, &forSite, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":1736421230,"evpn_options":{"auto_loopback_subnet":"172.16.192.0/24","auto_loopback_subnet6":"fd33:ab00:2::/64","auto_router_id_subnet":"172.16.254.0/23","core_as_border":true,"overlay":{"as":65000},"per_vlan_vga_v4_mac":false,"routed_at":"core","underlay":{"as_base":65001,"subnet":"10.255.240.0/20","use_ipv6":false}},"for_site":false,"id":"764fb173-94f9-447c-8454-def62e5a999f","modified_time":1736421230,"name":"tert","org_id":"3a2627d7-bfbc-45af-b85d-8841581c6d63","pod_names":{"1":"Pod 1"},"site_id":"00000000-0000-0000-0000-000000000000","version":6}]`
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":"9197ec96-4c8d-529f-c595-035895e688b2","name":"CC","overwrite":true,"pod_names":{"1":"default","2":"default"},"switches":[{"deviceprofile_id":"6a1deab1-96df-4fa2-8455-d5253f943d06","downlink_ips":["10.255.240.6","10.255.240.8"],"downlinks":["5c5b35000007","5c5b35000008"],"esilaglinks":["5c5b3500000f"],"evpn_id":1,"mac":"5c5b35000003","model":"QFX10002-36Q","role":"collapsed-core","site_id":"1916d52a-4a90-11e5-8b45-1258369c38a9","uplinks":["5c5b35000005","5c5b35000006"]}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsEVPNTopologiesTestCreateOrgEvpnTopology1 tests the behavior of the OrgsEVPNTopologies
func TestOrgsEVPNTopologiesTestCreateOrgEvpnTopology1(t *testing.T) {
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":"9197ec96-4c8d-529f-c595-035895e688b2","name":"CC","overwrite":true,"pod_names":{"1":"default","2":"default"},"switches":[{"deviceprofile_id":"6a1deab1-96df-4fa2-8455-d5253f943d06","downlink_ips":["10.255.240.6","10.255.240.8"],"downlinks":["5c5b35000007","5c5b35000008"],"esilaglinks":["5c5b3500000f"],"evpn_id":1,"mac":"5c5b35000003","model":"QFX10002-36Q","role":"collapsed-core","site_id":"1916d52a-4a90-11e5-8b45-1258369c38a9","uplinks":["5c5b35000005","5c5b35000006"]}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsEVPNTopologiesTestGetOrgEvpnTopology1 tests the behavior of the OrgsEVPNTopologies
func TestOrgsEVPNTopologiesTestGetOrgEvpnTopology1(t *testing.T) {
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":"9197ec96-4c8d-529f-c595-035895e688b2","name":"CC","overwrite":true,"pod_names":{"1":"default","2":"default"},"switches":[{"deviceprofile_id":"6a1deab1-96df-4fa2-8455-d5253f943d06","downlink_ips":["10.255.240.6","10.255.240.8"],"downlinks":["5c5b35000007","5c5b35000008"],"esilaglinks":["5c5b3500000f"],"evpn_id":1,"mac":"5c5b35000003","model":"QFX10002-36Q","role":"collapsed-core","site_id":"1916d52a-4a90-11e5-8b45-1258369c38a9","uplinks":["5c5b35000005","5c5b35000006"]}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsEVPNTopologiesTestUpdateOrgEvpnTopology1 tests the behavior of the OrgsEVPNTopologies
func TestOrgsEVPNTopologiesTestUpdateOrgEvpnTopology1(t *testing.T) {
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
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":"9197ec96-4c8d-529f-c595-035895e688b2","name":"CC","overwrite":true,"pod_names":{"1":"default","2":"default"},"switches":[{"deviceprofile_id":"6a1deab1-96df-4fa2-8455-d5253f943d06","downlink_ips":["10.255.240.6","10.255.240.8"],"downlinks":["5c5b35000007","5c5b35000008"],"esilaglinks":["5c5b3500000f"],"evpn_id":1,"mac":"5c5b35000003","model":"QFX10002-36Q","role":"collapsed-core","site_id":"1916d52a-4a90-11e5-8b45-1258369c38a9","uplinks":["5c5b35000005","5c5b35000006"]}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
