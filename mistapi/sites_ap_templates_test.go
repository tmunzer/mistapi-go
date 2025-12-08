// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesAPTemplatesTestListSiteApTemplatesDerived tests the behavior of the SitesAPTemplates
func TestSitesAPTemplatesTestListSiteApTemplatesDerived(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesApTemplates.ListSiteApTemplatesDerived(ctx, siteId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"ap_matching":{"enabled":true,"rules":[{"match_model":"string","name":"string","port_config":{"eth1,eth2":{"disabled":true,"dynamic_vlan":{"default_vlan_id":999,"enabled":true},"port_vlan_id":1,"vlan_id":9,"vlan_ids":"1, 10, 50"}}}]},"created_time":0,"for_site":true,"id":"497f6eca-6276-4993-bfeb-53cbbbba9f08","modified_time":0,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","wifi":{"cisco_enabled":true,"disable_11k":false,"disable_radios_when_power_constrained":true,"enable_arp_spoof":true,"enable_shared_radio_scanning":false,"enabled":true,"locate_connected":false,"locate_unconnected":false,"mesh_allow_dfs":false,"mesh_enable_crm":true,"mesh_enabled":true,"proxy_arp":false}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesAPTemplatesTestListSiteApTemplatesDerived1 tests the behavior of the SitesAPTemplates
func TestSitesAPTemplatesTestListSiteApTemplatesDerived1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesApTemplates.ListSiteApTemplatesDerived(ctx, siteId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"ap_matching":{"enabled":true,"rules":[{"match_model":"string","name":"string","port_config":{"eth1,eth2":{"disabled":true,"dynamic_vlan":{"default_vlan_id":999,"enabled":true},"port_vlan_id":1,"vlan_id":9,"vlan_ids":"1, 10, 50"}}}]},"created_time":0,"for_site":true,"id":"497f6eca-6276-4993-bfeb-53cbbbba9f08","modified_time":0,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","wifi":{"cisco_enabled":true,"disable_11k":false,"disable_radios_when_power_constrained":true,"enable_arp_spoof":true,"enable_shared_radio_scanning":false,"enabled":true,"locate_connected":false,"locate_unconnected":false,"mesh_allow_dfs":false,"mesh_enable_crm":true,"mesh_enabled":true,"proxy_arp":false}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
