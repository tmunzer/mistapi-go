// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsWxTunnelsTestListOrgWxTunnels tests the behavior of the OrgsWxTunnels
func TestOrgsWxTunnelsTestListOrgWxTunnels(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsWxTunnels.ListOrgWxTunnels(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":0,"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"modified_time":0,"mtu":1500,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","udp_port":0,"use_udp":true}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTunnelsTestListOrgWxTunnels1 tests the behavior of the OrgsWxTunnels
func TestOrgsWxTunnelsTestListOrgWxTunnels1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsWxTunnels.ListOrgWxTunnels(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"created_time":0,"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"modified_time":0,"mtu":1500,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","udp_port":0,"use_udp":true}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTunnelsTestCreateOrgWxTunnel tests the behavior of the OrgsWxTunnels
func TestOrgsWxTunnelsTestCreateOrgWxTunnel(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsWxTunnels.CreateOrgWxTunnel(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"modified_time":0,"mtu":1500,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","udp_port":0,"use_udp":true}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTunnelsTestCreateOrgWxTunnel1 tests the behavior of the OrgsWxTunnels
func TestOrgsWxTunnelsTestCreateOrgWxTunnel1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsWxTunnels.CreateOrgWxTunnel(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"modified_time":0,"mtu":1500,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","udp_port":0,"use_udp":true}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTunnelsTestDeleteOrgWxTunnel tests the behavior of the OrgsWxTunnels
func TestOrgsWxTunnelsTestDeleteOrgWxTunnel(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsWxTunnels.DeleteOrgWxTunnel(ctx, orgId, wxtunnelId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsWxTunnelsTestGetOrgWxTunnel tests the behavior of the OrgsWxTunnels
func TestOrgsWxTunnelsTestGetOrgWxTunnel(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsWxTunnels.GetOrgWxTunnel(ctx, orgId, wxtunnelId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"modified_time":0,"mtu":1500,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","udp_port":0,"use_udp":true}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTunnelsTestGetOrgWxTunnel1 tests the behavior of the OrgsWxTunnels
func TestOrgsWxTunnelsTestGetOrgWxTunnel1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsWxTunnels.GetOrgWxTunnel(ctx, orgId, wxtunnelId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"modified_time":0,"mtu":1500,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","udp_port":0,"use_udp":true}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTunnelsTestUpdateOrgWxTunnel tests the behavior of the OrgsWxTunnels
func TestOrgsWxTunnelsTestUpdateOrgWxTunnel(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsWxTunnels.UpdateOrgWxTunnel(ctx, orgId, wxtunnelId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"modified_time":0,"mtu":1500,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","udp_port":0,"use_udp":true}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWxTunnelsTestUpdateOrgWxTunnel1 tests the behavior of the OrgsWxTunnels
func TestOrgsWxTunnelsTestUpdateOrgWxTunnel1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsWxTunnels.UpdateOrgWxTunnel(ctx, orgId, wxtunnelId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"modified_time":0,"mtu":1500,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","udp_port":0,"use_udp":true}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
