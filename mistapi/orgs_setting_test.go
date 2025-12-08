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

// TestOrgsSettingTestGetOrgSettings tests the behavior of the OrgsSetting
func TestOrgsSettingTestGetOrgSettings(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsSetting.GetOrgSettings(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"auto_device_naming":{"enable":true,"rules":[{"match_device":"ap","prefix":"MIST-","src":"lldp_port_desc"}]},"auto_deviceprofile_assignment":{"enable":true,"rules":[{"expression":"string","model":"string","prefix":"string","src":"name","subnet":"string","suffix":"string","value":"string"}]},"auto_site_assignment":{"enable":true,"rules":[{"expression":"string","model":"string","prefix":"string","src":"name","subnet":"string","suffix":"string","value":"string"}]},"cacerts":["string"],"cloudshark":{"apitoken":"string","url":"string"},"created_time":0,"device_cert":{"cert":"string","key":"string"},"device_updown_threshold":0,"disable_pcap":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","installer":{"allow_all_sites":true,"extra_site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"grace_period":0},"mgmt":{"mxtunnel_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"use_mxtunnel":true,"use_wxtunnel":true},"modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","password_policy":{"enabled":true,"expiry_in_days":180,"min_length":8,"requires_special_char":true,"requires_two_factor_auth":true},"pcap":{"bucket":"string","max_pkt_len":0},"pcap_bucket_verified":true,"remote_syslog":{"enabled":true,"send_to_all_servers":true,"servers":[{"facility":"conflict-log","host":"string","port":0,"protocol":"udp","severity":"any","tag":"string"}]},"security":{"disable_local_ssh":true,"fips_zeroize_password":"string","limit_ssh_access":true},"tags":["string"],"ui_idle_timeout":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSettingTestGetOrgSettings1 tests the behavior of the OrgsSetting
func TestOrgsSettingTestGetOrgSettings1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsSetting.GetOrgSettings(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"auto_device_naming":{"enable":true,"rules":[{"match_device":"ap","prefix":"MIST-","src":"lldp_port_desc"}]},"auto_deviceprofile_assignment":{"enable":true,"rules":[{"expression":"string","model":"string","prefix":"string","src":"name","subnet":"string","suffix":"string","value":"string"}]},"auto_site_assignment":{"enable":true,"rules":[{"expression":"string","model":"string","prefix":"string","src":"name","subnet":"string","suffix":"string","value":"string"}]},"cacerts":["string"],"cloudshark":{"apitoken":"string","url":"string"},"created_time":0,"device_cert":{"cert":"string","key":"string"},"device_updown_threshold":0,"disable_pcap":true,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","installer":{"allow_all_sites":true,"extra_site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"grace_period":0},"mgmt":{"mxtunnel_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"use_mxtunnel":true,"use_wxtunnel":true},"modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","password_policy":{"enabled":true,"expiry_in_days":180,"min_length":8,"requires_special_char":true,"requires_two_factor_auth":true},"pcap":{"bucket":"string","max_pkt_len":0},"pcap_bucket_verified":true,"remote_syslog":{"enabled":true,"send_to_all_servers":true,"servers":[{"facility":"conflict-log","host":"string","port":0,"protocol":"udp","severity":"any","tag":"string"}]},"security":{"disable_local_ssh":true,"fips_zeroize_password":"string","limit_ssh_access":true},"tags":["string"],"ui_idle_timeout":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSettingTestUpdateOrgSettings tests the behavior of the OrgsSetting
func TestOrgsSettingTestUpdateOrgSettings(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsSetting.UpdateOrgSettings(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"auto_device_naming":{"enable":true,"rules":[{"match_device":"ap","prefix":"MIST-","src":"lldp_port_desc"}]},"auto_deviceprofile_assignment":{"enable":true,"rules":[{"expression":"string","model":"string","prefix":"string","src":"name","subnet":"string","suffix":"string","value":"string"}]},"auto_site_assignment":{"enable":true,"rules":[{"expression":"string","model":"string","prefix":"string","src":"name","subnet":"string","suffix":"string","value":"string"}]},"cacerts":["string"],"cloudshark":{"apitoken":"string","url":"string"},"device_cert":{"cert":"string","key":"string"},"disable_pcap":true,"installer":{"allow_all_sites":true,"extra_site_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"grace_period":0},"mgmt":{"mxtunnel_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"use_mxtunnel":true,"use_wxtunnel":true},"modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","password_policy":{"enabled":true,"expiry_in_days":365,"min_length":8,"requires_special_char":true,"requires_two_factor_auth":true},"pcap":{"bucket":"string","max_pkt_len":0},"pcap_bucket_verified":true,"remote_syslog":{"enabled":true,"send_to_all_servers":true,"servers":[{"facility":"change-log","host":"string","port":0,"protocol":"udp","severity":"critical","tag":"string"}]},"security":{"disable_local_ssh":true,"fips_zeroize_password":"string","limit_ssh_access":true},"tags":["string"],"ui_idle_timeout":0}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSettingTestDeleteOrgWirelessClientsBlocklist tests the behavior of the OrgsSetting
func TestOrgsSettingTestDeleteOrgWirelessClientsBlocklist(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsSetting.DeleteOrgWirelessClientsBlocklist(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSettingTestCreateOrgWirelessClientsBlocklist tests the behavior of the OrgsSetting
func TestOrgsSettingTestCreateOrgWirelessClientsBlocklist(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.MacAddresses
	errBody := json.Unmarshal([]byte(`{"macs":["18-65-90-de-f4-c6","84-89-ad-5d-69-0d"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsSetting.CreateOrgWirelessClientsBlocklist(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"macs":["18-65-90-de-f4-c6","84-89-ad-5d-69-0d"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSettingTestCreateOrgWirelessClientsBlocklist1 tests the behavior of the OrgsSetting
func TestOrgsSettingTestCreateOrgWirelessClientsBlocklist1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.MacAddresses
	errBody := json.Unmarshal([]byte(`{"macs":["18-65-90-de-f4-c6","84-89-ad-5d-69-0d"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsSetting.CreateOrgWirelessClientsBlocklist(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"macs":["18-65-90-de-f4-c6","84-89-ad-5d-69-0d"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSettingTestSetOrgCustomBucket tests the behavior of the OrgsSetting
func TestOrgsSettingTestSetOrgCustomBucket(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.PcapBucket
	errBody := json.Unmarshal([]byte(`{"bucket":"company-private-pcap"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsSetting.SetOrgCustomBucket(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"bucket":"company-private-pcap","detail":"failed to write bucket - 403 AccessDenied"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSettingTestSetOrgCustomBucket1 tests the behavior of the OrgsSetting
func TestOrgsSettingTestSetOrgCustomBucket1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.PcapBucket
	errBody := json.Unmarshal([]byte(`{"bucket":"company-private-pcap"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsSetting.SetOrgCustomBucket(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"bucket":"company-private-pcap","detail":"failed to write bucket - 403 AccessDenied"}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSettingTestVerifyOrgCustomBucket tests the behavior of the OrgsSetting
func TestOrgsSettingTestVerifyOrgCustomBucket(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.PcapBucketVerify
	errBody := json.Unmarshal([]byte(`{"bucket":"company-private-pcap","verify_token":"eyJhbGciOiJIUzI1J9.eyJzdWIiOiIxMjM0joiMjgxOG5MDIyfQ.2rzcRvMA3Eg09NnjCAC-1EWMRtxAnFDM"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	resp, err := orgsSetting.VerifyOrgCustomBucket(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
