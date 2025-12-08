// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestSitesDeviceProfilesTestListSiteDeviceProfilesDerived tests the behavior of the SitesDeviceProfiles
func TestSitesDeviceProfilesTestListSiteDeviceProfilesDerived(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resolve := bool(false)
	apiResponse, err := sitesDeviceProfiles.ListSiteDeviceProfilesDerived(ctx, siteId, &resolve)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"aeroscout":{"enabled":true,"host":"string"},"created_time":0,"disable_eth1":true,"disable_module":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ip_config":{"dns":["string"],"dns_suffix":["string"],"gateway":"192.168.0.1","gateway6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","ip":"192.168.0.1","ip6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","mtu":1500,"netmask":"192.168.0.1","netmask6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","type":"static","type6":"static","vlan_id":1},"mesh":{"enabled":true,"group":1,"role":"base"},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","poe_passthrough":true,"switch_config":{"enabled":true,"eth0":{"port_vlan_id":1,"vlan_ids":[1,10]},"eth1":{"port_vlan_id":1,"vlan_ids":[10]},"eth2":{"port_vlan_id":1,"vlan_ids":[10]},"eth3":{"port_vlan_id":1,"vlan_ids":[10]},"module":{"port_vlan_id":1,"vlan_ids":[10]},"wds":{"port_vlan_id":1,"vlan_ids":[10]}},"type":"ap","usb_config":{"cacert":"string","enabled":true,"host":"string","type":"imagotag","verify_cert":true}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDeviceProfilesTestListSiteDeviceProfilesDerived1 tests the behavior of the SitesDeviceProfiles
func TestSitesDeviceProfilesTestListSiteDeviceProfilesDerived1(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resolve := bool(false)
	apiResponse, err := sitesDeviceProfiles.ListSiteDeviceProfilesDerived(ctx, siteId, &resolve)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"aeroscout":{"enabled":true,"host":"string"},"created_time":0,"disable_eth1":true,"disable_module":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ip_config":{"dns":["string"],"dns_suffix":["string"],"gateway":"192.168.0.1","gateway6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","ip":"192.168.0.1","ip6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","mtu":1500,"netmask":"192.168.0.1","netmask6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","type":"static","type6":"static","vlan_id":1},"mesh":{"enabled":true,"group":1,"role":"base"},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","poe_passthrough":true,"switch_config":{"enabled":true,"eth0":{"port_vlan_id":1,"vlan_ids":[1,10]},"eth1":{"port_vlan_id":1,"vlan_ids":[10]},"eth2":{"port_vlan_id":1,"vlan_ids":[10]},"eth3":{"port_vlan_id":1,"vlan_ids":[10]},"module":{"port_vlan_id":1,"vlan_ids":[10]},"wds":{"port_vlan_id":1,"vlan_ids":[10]}},"type":"ap","usb_config":{"cacert":"string","enabled":true,"host":"string","type":"imagotag","verify_cert":true}}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
