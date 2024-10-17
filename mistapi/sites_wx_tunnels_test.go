package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSitesWxTunnelsTestListSiteWxTunnels tests the behavior of the SitesWxTunnels
func TestSitesWxTunnelsTestListSiteWxTunnels(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := sitesWxTunnels.ListSiteWxTunnels(ctx, siteId, &limit, &page)
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

// TestSitesWxTunnelsTestCreateSiteWxTunnel tests the behavior of the SitesWxTunnels
func TestSitesWxTunnelsTestCreateSiteWxTunnel(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.WxlanTunnel
	errBody := json.Unmarshal([]byte(`{"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"mtu":0,"name":"string","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"udp_port":0,"use_udp":true}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesWxTunnels.CreateSiteWxTunnel(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"modified_time":0,"mtu":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","udp_port":0,"use_udp":true}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxTunnelsTestDeleteSiteWxTunnel tests the behavior of the SitesWxTunnels
func TestSitesWxTunnelsTestDeleteSiteWxTunnel(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesWxTunnels.DeleteSiteWxTunnel(ctx, siteId, wxtunnelId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesWxTunnelsTestGetSiteWxTunnel tests the behavior of the SitesWxTunnels
func TestSitesWxTunnelsTestGetSiteWxTunnel(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesWxTunnels.GetSiteWxTunnel(ctx, siteId, wxtunnelId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"modified_time":0,"mtu":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","udp_port":0,"use_udp":true}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWxTunnelsTestUpdateSiteWxTunnel tests the behavior of the SitesWxTunnels
func TestSitesWxTunnelsTestUpdateSiteWxTunnel(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	wxtunnelId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.WxlanTunnel
	errBody := json.Unmarshal([]byte(`{"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"mtu":0,"name":"string","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"udp_port":0,"use_udp":true}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesWxTunnels.UpdateSiteWxTunnel(ctx, siteId, wxtunnelId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"created_time":0,"dmvpn":{"enabled":true,"holding_time":0,"host_routes":["string"]},"for_mgmt":true,"hello_interval":1,"hello_retries":3,"hostname":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ipsec":{"enabled":true,"psk":"string123"},"is_static":true,"modified_time":0,"mtu":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","peers":["string"],"router_id":"string","secret":"string","sessions":[{"ap_as_session_id":"string","comment":"string","enable_cookie":true,"ethertype":"ethernet","local_session_id":1,"pseudo_802.1ad_enabled":true,"remote_id":"string","remote_session_id":1,"use_ap_as_session_ids":true}],"site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","udp_port":0,"use_udp":true}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
