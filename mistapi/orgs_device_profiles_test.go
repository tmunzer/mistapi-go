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

// TestOrgsDeviceProfilesTestListOrgDeviceProfiles tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestListOrgDeviceProfiles(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mType := models.DeviceTypeDefaultApEnum("ap")
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsDeviceProfiles.ListOrgDeviceProfiles(ctx, orgId, &mType, &limit, &page)
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

// TestOrgsDeviceProfilesTestListOrgDeviceProfiles1 tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestListOrgDeviceProfiles1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mType := models.DeviceTypeDefaultApEnum("ap")
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsDeviceProfiles.ListOrgDeviceProfiles(ctx, orgId, &mType, &limit, &page)
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

// TestOrgsDeviceProfilesTestCreateOrgDeviceProfile tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestCreateOrgDeviceProfile(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Deviceprofile
	errBody := json.Unmarshal([]byte(`{"aeroscout":{"enabled":false,"host":"aero.pvt.net","locate_connected":true},"led":{"brightness":255,"enabled":true},"name":"string","ntp_servers":["10.10.10.10"],"type":"ap","usb_config":{"cacert":"string","channel":3,"enabled":true,"host":"1.1.1.1","port":0,"type":"imagotag","verify_cert":true,"vlan_id":1}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsDeviceProfiles.CreateOrgDeviceProfile(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"aeroscout":{"enabled":true,"host":"string"},"created_time":0,"disable_eth1":true,"disable_module":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ip_config":{"dns":["string"],"dns_suffix":["string"],"gateway":"192.168.0.1","gateway6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","ip":"192.168.0.1","ip6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","mtu":1500,"netmask":"192.168.0.1","netmask6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","type":"static","type6":"static","vlan_id":1},"mesh":{"enabled":true,"group":1,"role":"base"},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","poe_passthrough":true,"switch_config":{"enabled":true,"eth0":{"port_vlan_id":1,"vlan_ids":[1,10]},"eth1":{"port_vlan_id":1,"vlan_ids":[10]},"eth2":{"port_vlan_id":1,"vlan_ids":[10]},"eth3":{"port_vlan_id":1,"vlan_ids":[10]},"module":{"port_vlan_id":1,"vlan_ids":[10]},"wds":{"port_vlan_id":1,"vlan_ids":[10]}},"type":"ap","usb_config":{"cacert":"string","enabled":true,"host":"string","type":"imagotag","verify_cert":true}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestCreateOrgDeviceProfile1 tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestCreateOrgDeviceProfile1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Deviceprofile
	errBody := json.Unmarshal([]byte(`{"aeroscout":{"enabled":false,"host":"aero.pvt.net","locate_connected":true},"led":{"brightness":255,"enabled":true},"name":"string","ntp_servers":["10.10.10.10"],"type":"ap","usb_config":{"cacert":"string","channel":3,"enabled":true,"host":"1.1.1.1","port":0,"type":"imagotag","verify_cert":true,"vlan_id":1}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsDeviceProfiles.CreateOrgDeviceProfile(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"aeroscout":{"enabled":true,"host":"string"},"created_time":0,"disable_eth1":true,"disable_module":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ip_config":{"dns":["string"],"dns_suffix":["string"],"gateway":"192.168.0.1","gateway6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","ip":"192.168.0.1","ip6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","mtu":1500,"netmask":"192.168.0.1","netmask6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","type":"static","type6":"static","vlan_id":1},"mesh":{"enabled":true,"group":1,"role":"base"},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","poe_passthrough":true,"switch_config":{"enabled":true,"eth0":{"port_vlan_id":1,"vlan_ids":[1,10]},"eth1":{"port_vlan_id":1,"vlan_ids":[10]},"eth2":{"port_vlan_id":1,"vlan_ids":[10]},"eth3":{"port_vlan_id":1,"vlan_ids":[10]},"module":{"port_vlan_id":1,"vlan_ids":[10]},"wds":{"port_vlan_id":1,"vlan_ids":[10]}},"type":"ap","usb_config":{"cacert":"string","enabled":true,"host":"string","type":"imagotag","verify_cert":true}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestDeleteOrgDeviceProfile tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestDeleteOrgDeviceProfile(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsDeviceProfiles.DeleteOrgDeviceProfile(ctx, orgId, deviceprofileId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsDeviceProfilesTestGetOrgDeviceProfile tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestGetOrgDeviceProfile(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsDeviceProfiles.GetOrgDeviceProfile(ctx, orgId, deviceprofileId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"aeroscout":{"enabled":true,"host":"string"},"created_time":0,"disable_eth1":true,"disable_module":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ip_config":{"dns":["string"],"dns_suffix":["string"],"gateway":"192.168.0.1","gateway6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","ip":"192.168.0.1","ip6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","mtu":1500,"netmask":"192.168.0.1","netmask6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","type":"static","type6":"static","vlan_id":1},"mesh":{"enabled":true,"group":1,"role":"base"},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","poe_passthrough":true,"switch_config":{"enabled":true,"eth0":{"port_vlan_id":1,"vlan_ids":[1,10]},"eth1":{"port_vlan_id":1,"vlan_ids":[10]},"eth2":{"port_vlan_id":1,"vlan_ids":[10]},"eth3":{"port_vlan_id":1,"vlan_ids":[10]},"module":{"port_vlan_id":1,"vlan_ids":[10]},"wds":{"port_vlan_id":1,"vlan_ids":[10]}},"type":"ap","usb_config":{"cacert":"string","enabled":true,"host":"string","type":"imagotag","verify_cert":true}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestGetOrgDeviceProfile1 tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestGetOrgDeviceProfile1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsDeviceProfiles.GetOrgDeviceProfile(ctx, orgId, deviceprofileId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"aeroscout":{"enabled":true,"host":"string"},"created_time":0,"disable_eth1":true,"disable_module":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ip_config":{"dns":["string"],"dns_suffix":["string"],"gateway":"192.168.0.1","gateway6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","ip":"192.168.0.1","ip6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","mtu":1500,"netmask":"192.168.0.1","netmask6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","type":"static","type6":"static","vlan_id":1},"mesh":{"enabled":true,"group":1,"role":"base"},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","poe_passthrough":true,"switch_config":{"enabled":true,"eth0":{"port_vlan_id":1,"vlan_ids":[1,10]},"eth1":{"port_vlan_id":1,"vlan_ids":[10]},"eth2":{"port_vlan_id":1,"vlan_ids":[10]},"eth3":{"port_vlan_id":1,"vlan_ids":[10]},"module":{"port_vlan_id":1,"vlan_ids":[10]},"wds":{"port_vlan_id":1,"vlan_ids":[10]}},"type":"ap","usb_config":{"cacert":"string","enabled":true,"host":"string","type":"imagotag","verify_cert":true}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestUpdateOrgDeviceProfile tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestUpdateOrgDeviceProfile(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Deviceprofile
	errBody := json.Unmarshal([]byte(`{"aeroscout":{"enabled":true,"host":"aero.pvt.net"},"disable_eth1":true,"disable_module":true,"mesh":{"enabled":true,"group":1,"role":"base"},"name":"string","poe_passthrough":true,"radio_config":{"ant_gain_24":0,"ant_gain_5":0,"band_24":{"allow_rrm_disable":true,"antenna_mode":"default","bandwidth":20,"channel":6,"disabled":true,"power":8,"preamble":"auto"},"band_24_usage":"24","band_5":{"allow_rrm_disable":true,"antenna_mode":"default","bandwidth":20,"channel":50,"disabled":true,"power_max":8,"power_min":15,"preamble":"auto"},"scanning_enabled":true},"type":"ap"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsDeviceProfiles.UpdateOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"aeroscout":{"enabled":true,"host":"string"},"created_time":0,"disable_eth1":true,"disable_module":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ip_config":{"dns":["string"],"dns_suffix":["string"],"gateway":"192.168.0.1","gateway6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","ip":"192.168.0.1","ip6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","mtu":1500,"netmask":"192.168.0.1","netmask6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","type":"static","type6":"static","vlan_id":1},"mesh":{"enabled":true,"group":1,"role":"base"},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","poe_passthrough":true,"switch_config":{"enabled":true,"eth0":{"port_vlan_id":1,"vlan_ids":[1,10]},"eth1":{"port_vlan_id":1,"vlan_ids":[10]},"eth2":{"port_vlan_id":1,"vlan_ids":[10]},"eth3":{"port_vlan_id":1,"vlan_ids":[10]},"module":{"port_vlan_id":1,"vlan_ids":[10]},"wds":{"port_vlan_id":1,"vlan_ids":[10]}},"type":"ap","usb_config":{"cacert":"string","enabled":true,"host":"string","type":"imagotag","verify_cert":true}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestUpdateOrgDeviceProfile1 tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestUpdateOrgDeviceProfile1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.Deviceprofile
	errBody := json.Unmarshal([]byte(`{"aeroscout":{"enabled":true,"host":"aero.pvt.net"},"disable_eth1":true,"disable_module":true,"mesh":{"enabled":true,"group":1,"role":"base"},"name":"string","poe_passthrough":true,"radio_config":{"ant_gain_24":0,"ant_gain_5":0,"band_24":{"allow_rrm_disable":true,"antenna_mode":"default","bandwidth":20,"channel":6,"disabled":true,"power":8,"preamble":"auto"},"band_24_usage":"24","band_5":{"allow_rrm_disable":true,"antenna_mode":"default","bandwidth":20,"channel":50,"disabled":true,"power_max":8,"power_min":15,"preamble":"auto"},"scanning_enabled":true},"type":"ap"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsDeviceProfiles.UpdateOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"aeroscout":{"enabled":true,"host":"string"},"created_time":0,"disable_eth1":true,"disable_module":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","ip_config":{"dns":["string"],"dns_suffix":["string"],"gateway":"192.168.0.1","gateway6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","ip":"192.168.0.1","ip6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","mtu":1500,"netmask":"192.168.0.1","netmask6":"2001:0db8:85a3:0000:0000:8a2e:0370:7334","type":"static","type6":"static","vlan_id":1},"mesh":{"enabled":true,"group":1,"role":"base"},"modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","poe_passthrough":true,"switch_config":{"enabled":true,"eth0":{"port_vlan_id":1,"vlan_ids":[1,10]},"eth1":{"port_vlan_id":1,"vlan_ids":[10]},"eth2":{"port_vlan_id":1,"vlan_ids":[10]},"eth3":{"port_vlan_id":1,"vlan_ids":[10]},"module":{"port_vlan_id":1,"vlan_ids":[10]},"wds":{"port_vlan_id":1,"vlan_ids":[10]}},"type":"ap","usb_config":{"cacert":"string","enabled":true,"host":"string","type":"imagotag","verify_cert":true}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestAssignOrgDeviceProfile tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestAssignOrgDeviceProfile(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.MacAddresses
	errBody := json.Unmarshal([]byte(`{"macs":["5c5b350e0001","5c5b350e0003"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsDeviceProfiles.AssignOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"success":["5c5b350e0001"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestAssignOrgDeviceProfile1 tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestAssignOrgDeviceProfile1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.MacAddresses
	errBody := json.Unmarshal([]byte(`{"macs":["5c5b350e0001","5c5b350e0003"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsDeviceProfiles.AssignOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"success":["5c5b350e0001"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestUnassignOrgDeviceProfile tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestUnassignOrgDeviceProfile(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.MacAddresses
	errBody := json.Unmarshal([]byte(`{"macs":["5c5b350e0001","5c5b350e0003"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsDeviceProfiles.UnassignOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"success":["5c5b350e0001"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsDeviceProfilesTestUnassignOrgDeviceProfile1 tests the behavior of the OrgsDeviceProfiles
func TestOrgsDeviceProfilesTestUnassignOrgDeviceProfile1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceprofileId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.MacAddresses
	errBody := json.Unmarshal([]byte(`{"macs":["5c5b350e0001","5c5b350e0003"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsDeviceProfiles.UnassignOrgDeviceProfile(ctx, orgId, deviceprofileId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"success":["5c5b350e0001"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
