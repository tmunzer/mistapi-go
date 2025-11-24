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

// TestOrgsClientsWiredTestCountOrgWiredClients tests the behavior of the OrgsClientsWired
func TestOrgsClientsWiredTestCountOrgWiredClients(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgWiredClientsCountDistinctEnum("mac")

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsClientsWired.CountOrgWiredClients(ctx, orgId, &distinct, nil, nil, &duration, &limit)
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

// TestOrgsClientsWiredTestCountOrgWiredClients1 tests the behavior of the OrgsClientsWired
func TestOrgsClientsWiredTestCountOrgWiredClients1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgWiredClientsCountDistinctEnum("mac")

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsClientsWired.CountOrgWiredClients(ctx, orgId, &distinct, nil, nil, &duration, &limit)
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

// TestOrgsClientsWiredTestSearchOrgWiredClients tests the behavior of the OrgsClientsWired
func TestOrgsClientsWiredTestSearchOrgWiredClients(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	ip := "192.168.1.1"

	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsClientsWired.SearchOrgWiredClients(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, &ip, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
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

// TestOrgsClientsWiredTestSearchOrgWiredClients1 tests the behavior of the OrgsClientsWired
func TestOrgsClientsWiredTestSearchOrgWiredClients1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	ip := "192.168.1.1"

	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsClientsWired.SearchOrgWiredClients(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, &ip, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
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
