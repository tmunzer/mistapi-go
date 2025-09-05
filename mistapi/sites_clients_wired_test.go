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

// TestSitesClientsWiredTestCountSiteWiredClients tests the behavior of the SitesClientsWired
func TestSitesClientsWiredTestCountSiteWiredClients(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SiteWiredClientsCountDistinctEnum("mac")
	mac := "0123456789ab"
	deviceMac := "0123456789ab"
	portId := "ge-1/1/1"
	vlan := "10"

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesClientsWired.CountSiteWiredClients(ctx, siteId, &distinct, &mac, &deviceMac, &portId, &vlan, nil, nil, &duration, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesClientsWiredTestCountSiteWiredClients1 tests the behavior of the SitesClientsWired
func TestSitesClientsWiredTestCountSiteWiredClients1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.SiteWiredClientsCountDistinctEnum("mac")
	mac := "0123456789ab"
	deviceMac := "0123456789ab"
	portId := "ge-1/1/1"
	vlan := "10"

	duration := "1d"
	limit := int(100)
	apiResponse, err := sitesClientsWired.CountSiteWiredClients(ctx, siteId, &distinct, &mac, &deviceMac, &portId, &vlan, nil, nil, &duration, &limit)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesClientsWiredTestSearchSiteWiredClients tests the behavior of the SitesClientsWired
func TestSitesClientsWiredTestSearchSiteWiredClients(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceMac := "0123456789ab"
	mac := "0123456789ab"
	ip := "10.3.5.12"
	portId := "ge-1/1/1"

	vlan := "10"
	manufacture := "Juniper-Mist"
	text := "client-hostname"
	nacruleId := "8bfc2490-d726-3587-038d-cb2e71bd2330"
	dhcpHostname := "client-hostname"
	dhcpFqdn := "client.example.com"
	dhcpClientIdentifier := "01:23:45:67:89:ab"
	dhcpVendorClassIdentifier := "Juniper-Mist-AP,Juniper-Mist-Client"
	dhcpRequestParams := "hostname,domain-name,domain-name-servers"
	limit := int(100)

	duration := "1d"
	sort := "timestamp"
	apiResponse, err := sitesClientsWired.SearchSiteWiredClients(ctx, siteId, &deviceMac, &mac, &ip, &portId, nil, &vlan, &manufacture, &text, &nacruleId, &dhcpHostname, &dhcpFqdn, &dhcpClientIdentifier, &dhcpVendorClassIdentifier, &dhcpRequestParams, &limit, nil, nil, &duration, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1648529800.8221116,"limit":1000,"results":[{"auth_method":"mac_auth","auth_state":"authenticated","device_mac":["001122334455"],"dhcp_client_identifier":"MAC address 00155df6d500","dhcp_client_options":[{"code":"DHO_DHCP_MESSAGE_TYPE(53)","data":"DHCPREQUEST"},{"code":"DHO_DHCP_CLIENT_IDENTIFIER(61)","data":"MAC address 00155df6d500"},{"code":"DHO_DHCP_REQUESTED_ADDRESS(50)","data":"172.17.10.255"},{"code":"DHO_DHCP_SERVER_IDENTIFIER(54)","data":"172.17.8.1"},{"code":"DHO_DHCP_MAX_MESSAGE_SIZE(57)","data":"1280"},{"code":"DHO_DHCP_PARAMETER_REQUEST_LIST(55)","data":" 1 3 6 12 15 28 43 180"},{"code":"DHO_VENDOR_CLASS_IDENTIFIER(60)","data":"MSFT 5.0"},{"code":"DHO_HOST_NAME(12)","data":"ITS-VMMT0-D1N02"}],"dhcp_fqdn":"ITS-VMMT0-D1N02.mgthub.local","dhcp_hostname":"ITS-VMMT0-D1N02","dhcp_request_params":"1 3 6 15 31 33 43 44 46 47 119 121 249 252","dhcp_vendor_class_identifier":"MSFT 5.0","mac":"112233445566","org_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","port_id":["et-0/0/1"],"site_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","timestamp":1571174567.807,"vlan":[0,1001]}],"start":1648443400.8221116,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesClientsWiredTestSearchSiteWiredClients1 tests the behavior of the SitesClientsWired
func TestSitesClientsWiredTestSearchSiteWiredClients1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceMac := "0123456789ab"
	mac := "0123456789ab"
	ip := "10.3.5.12"
	portId := "ge-1/1/1"

	vlan := "10"
	manufacture := "Juniper-Mist"
	text := "client-hostname"
	nacruleId := "8bfc2490-d726-3587-038d-cb2e71bd2330"
	dhcpHostname := "client-hostname"
	dhcpFqdn := "client.example.com"
	dhcpClientIdentifier := "01:23:45:67:89:ab"
	dhcpVendorClassIdentifier := "Juniper-Mist-AP,Juniper-Mist-Client"
	dhcpRequestParams := "hostname,domain-name,domain-name-servers"
	limit := int(100)

	duration := "1d"
	sort := "timestamp"
	apiResponse, err := sitesClientsWired.SearchSiteWiredClients(ctx, siteId, &deviceMac, &mac, &ip, &portId, nil, &vlan, &manufacture, &text, &nacruleId, &dhcpHostname, &dhcpFqdn, &dhcpClientIdentifier, &dhcpVendorClassIdentifier, &dhcpRequestParams, &limit, nil, nil, &duration, &sort)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1648529800.8221116,"limit":1000,"results":[{"auth_method":"mac_auth","auth_state":"authenticated","device_mac":["001122334455"],"dhcp_client_identifier":"MAC address 00155df6d500","dhcp_client_options":[{"code":"DHO_DHCP_MESSAGE_TYPE(53)","data":"DHCPREQUEST"},{"code":"DHO_DHCP_CLIENT_IDENTIFIER(61)","data":"MAC address 00155df6d500"},{"code":"DHO_DHCP_REQUESTED_ADDRESS(50)","data":"172.17.10.255"},{"code":"DHO_DHCP_SERVER_IDENTIFIER(54)","data":"172.17.8.1"},{"code":"DHO_DHCP_MAX_MESSAGE_SIZE(57)","data":"1280"},{"code":"DHO_DHCP_PARAMETER_REQUEST_LIST(55)","data":" 1 3 6 12 15 28 43 180"},{"code":"DHO_VENDOR_CLASS_IDENTIFIER(60)","data":"MSFT 5.0"},{"code":"DHO_HOST_NAME(12)","data":"ITS-VMMT0-D1N02"}],"dhcp_fqdn":"ITS-VMMT0-D1N02.mgthub.local","dhcp_hostname":"ITS-VMMT0-D1N02","dhcp_request_params":"1 3 6 15 31 33 43 44 46 47 119 121 249 252","dhcp_vendor_class_identifier":"MSFT 5.0","mac":"112233445566","org_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","port_id":["et-0/0/1"],"site_id":"c168ddee-c14c-11e5-8e81-1258369c38a9","timestamp":1571174567.807,"vlan":[0,1001]}],"start":1648443400.8221116,"total":1}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
