// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"testing"
)

// TestConstantsModelsTestGetGatewayDefaultConfig tests the behavior of the ConstantsModels
func TestConstantsModelsTestGetGatewayDefaultConfig(t *testing.T) {
	ctx := context.Background()
	model := "srx550"
	ha := "false"
	apiResponse, err := constantsModels.GetGatewayDefaultConfig(ctx, model, &ha)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"dhcpd_config":{"lan":{"ip_end":"192.168.1.254","ip_start":"192.168.1.2"}},"ip_configs":{"lan":{"ip":"192.168.1.1","type":"static"}},"networks":{"lan":{"name":"lan","subnet":"192.168.1.0/24","vlan_id":1}},"path_preferences":{"wan":{"paths":[{"name":"wan","type":"wan"}]}},"port_config":{"cl-1/0/0":{"ip_config":{"type":"dhcp"},"name":"lte","usage":"wan","wan_type":"lte"},"ge-0/0/0,ge-0/0/7":{"ip_config":{"type":"dhcp"},"name":"wan","usage":"wan"},"ge-0/0/1-6":{"port_network":"lan","usage":"lan"}},"service_policies":[{"action":"allow","name":"Internet","path_preference":"wan","services":["any"],"tenants":["lan"]}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsModelsTestGetGatewayDefaultConfig1 tests the behavior of the ConstantsModels
func TestConstantsModelsTestGetGatewayDefaultConfig1(t *testing.T) {
	ctx := context.Background()
	model := "srx550"
	ha := "false"
	apiResponse, err := constantsModels.GetGatewayDefaultConfig(ctx, model, &ha)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"dhcpd_config":{"lan":{"ip_end":"192.168.1.254","ip_start":"192.168.1.2"}},"ip_configs":{"lan":{"ip":"192.168.1.1","type":"static"}},"networks":{"lan":{"name":"lan","subnet":"192.168.1.0/24","vlan_id":1}},"path_preferences":{"wan":{"paths":[{"name":"wan","type":"wan"}]}},"port_config":{"cl-1/0/0":{"ip_config":{"type":"dhcp"},"name":"lte","usage":"wan","wan_type":"lte"},"ge-0/0/0,ge-0/0/7":{"ip_config":{"type":"dhcp"},"name":"wan","usage":"wan"},"ge-0/0/1-6":{"port_network":"lan","usage":"lan"}},"service_policies":[{"action":"allow","name":"Internet","path_preference":"wan","services":["any"],"tenants":["lan"]}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsModelsTestListMxEdgeModels tests the behavior of the ConstantsModels
func TestConstantsModelsTestListMxEdgeModels(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsModels.ListMxEdgeModels(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"display":"X10","model":"ME-X10","ports":{"0":{"display":"xe0","speed":10000},"1":{"display":"xe1","speed":10000},"2":{"display":"xe2","speed":10000},"3":{"display":"xe3","speed":10000}}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsModelsTestListMxEdgeModels1 tests the behavior of the ConstantsModels
func TestConstantsModelsTestListMxEdgeModels1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsModels.ListMxEdgeModels(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"display":"X10","model":"ME-X10","ports":{"0":{"display":"xe0","speed":10000},"1":{"display":"xe1","speed":10000},"2":{"display":"xe2","speed":10000},"3":{"display":"xe3","speed":10000}}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsModelsTestListSupportedOtherDeviceModels tests the behavior of the ConstantsModels
func TestConstantsModelsTestListSupportedOtherDeviceModels(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsModels.ListSupportedOtherDeviceModels(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"_vendor_model_id":"65","display":"W1850","model":"W1850","type":"router","vendor":"cradlepoint"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestConstantsModelsTestListSupportedOtherDeviceModels1 tests the behavior of the ConstantsModels
func TestConstantsModelsTestListSupportedOtherDeviceModels1(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := constantsModels.ListSupportedOtherDeviceModels(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"_vendor_model_id":"65","display":"W1850","model":"W1850","type":"router","vendor":"cradlepoint"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
