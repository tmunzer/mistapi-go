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

// TestOrgsAlarmTemplatesTestListOrgAlarmTemplates tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestListOrgAlarmTemplates(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsAlarmTemplates.ListOrgAlarmTemplates(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsAlarmTemplatesTestListOrgAlarmTemplates1 tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestListOrgAlarmTemplates1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsAlarmTemplates.ListOrgAlarmTemplates(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsAlarmTemplatesTestCreateOrgAlarmTemplate tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestCreateOrgAlarmTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsAlarmTemplates.CreateOrgAlarmTemplate(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsAlarmTemplatesTestCreateOrgAlarmTemplate1 tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestCreateOrgAlarmTemplate1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsAlarmTemplates.CreateOrgAlarmTemplate(ctx, orgId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsAlarmTemplatesTestUnsuppressOrgSuppressedAlarms tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestUnsuppressOrgSuppressedAlarms(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsAlarmTemplates.UnsuppressOrgSuppressedAlarms(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAlarmTemplatesTestListOrgSuppressedAlarms tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestListOrgSuppressedAlarms(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	scope := models.SuppressedAlarmScopeEnum("site")
	apiResponse, err := orgsAlarmTemplates.ListOrgSuppressedAlarms(ctx, orgId, &scope)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"results":[{"duration":48,"expire_time":1678233080,"scheduled_time":1678232900,"scope":"site","site_id":"581328b6-e382-f54e-c9dc-9c998d183a34"},{"duration":48,"expire_time":1678233080,"scheduled_time":1678232900,"scope":"org","site_id":"581328b6-e382-f54e-c9dc-9c998d183a35"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAlarmTemplatesTestListOrgSuppressedAlarms1 tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestListOrgSuppressedAlarms1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	scope := models.SuppressedAlarmScopeEnum("site")
	apiResponse, err := orgsAlarmTemplates.ListOrgSuppressedAlarms(ctx, orgId, &scope)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"results":[{"duration":48,"expire_time":1678233080,"scheduled_time":1678232900,"scope":"site","site_id":"581328b6-e382-f54e-c9dc-9c998d183a34"},{"duration":48,"expire_time":1678233080,"scheduled_time":1678232900,"scope":"org","site_id":"581328b6-e382-f54e-c9dc-9c998d183a35"}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsAlarmTemplatesTestSuppressOrgAlarm tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestSuppressOrgAlarm(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.SuppressedAlarm
	errBody := json.Unmarshal([]byte(`{"duration":3600,"scheduled_time":1678232980,"scope":"org"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	resp, err := orgsAlarmTemplates.SuppressOrgAlarm(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAlarmTemplatesTestDeleteOrgAlarmTemplate tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestDeleteOrgAlarmTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	alarmtemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsAlarmTemplates.DeleteOrgAlarmTemplate(ctx, orgId, alarmtemplateId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsAlarmTemplatesTestGetOrgAlarmTemplate tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestGetOrgAlarmTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	alarmtemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsAlarmTemplates.GetOrgAlarmTemplate(ctx, orgId, alarmtemplateId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsAlarmTemplatesTestGetOrgAlarmTemplate1 tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestGetOrgAlarmTemplate1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	alarmtemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsAlarmTemplates.GetOrgAlarmTemplate(ctx, orgId, alarmtemplateId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsAlarmTemplatesTestUpdateOrgAlarmTemplate tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestUpdateOrgAlarmTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	alarmtemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.AlarmTemplate
	errBody := json.Unmarshal([]byte(`{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"name":"string","rules":{"adhoc_network":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"air_magnet_scan":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"ap_offline":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"bad_cable":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"beacon_flood":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"bssid_spoofing":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"device_down":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"device_restarted":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"dhcp_failure":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"disassociation_flood":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"dot1x_failure":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"eap_dictionary_attack":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"eap_failure_injection":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"eap_handshake_flood":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"eap_spoofed_success":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"eapol_logoff_attack":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"essid_jack":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"excessive_client":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"excessive_eapol_start":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"gateway_down":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"gw_bad_cable":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"gw_negotiation_mismatch":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"honeypot_ssid":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"krack_attack":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"missing_vlan":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"monkey_jack":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"negotiation_mismatch":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"non_compliant":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"out_of_sequence":{"enabled":true},"psk_failure":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"repeated_auth_failures":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"rogue_ap":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"rogue_client":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"secpolicy_violation":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"ssid_injection":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"switch_down":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"tkip_icv_attack":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"vendor_ie_missing":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"watched_station":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"zero_ssid_association":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true}}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsAlarmTemplates.UpdateOrgAlarmTemplate(ctx, orgId, alarmtemplateId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsAlarmTemplatesTestUpdateOrgAlarmTemplate1 tests the behavior of the OrgsAlarmTemplates
func TestOrgsAlarmTemplatesTestUpdateOrgAlarmTemplate1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	alarmtemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.AlarmTemplate
	errBody := json.Unmarshal([]byte(`{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"name":"string","rules":{"adhoc_network":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"air_magnet_scan":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"ap_offline":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"bad_cable":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"beacon_flood":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"bssid_spoofing":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"device_down":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"device_restarted":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"dhcp_failure":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"disassociation_flood":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"dot1x_failure":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"eap_dictionary_attack":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"eap_failure_injection":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"eap_handshake_flood":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"eap_spoofed_success":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"eapol_logoff_attack":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"essid_jack":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"excessive_client":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"excessive_eapol_start":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"gateway_down":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"gw_bad_cable":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"gw_negotiation_mismatch":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"honeypot_ssid":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"krack_attack":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"missing_vlan":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"monkey_jack":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"negotiation_mismatch":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"non_compliant":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"out_of_sequence":{"enabled":true},"psk_failure":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"repeated_auth_failures":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"rogue_ap":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"rogue_client":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"secpolicy_violation":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"ssid_injection":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"switch_down":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"tkip_icv_attack":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"vendor_ie_missing":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"watched_station":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true},"zero_ssid_association":{"delivery":{"additional_emails":["string"],"enabled":true,"to_org_admins":true,"to_site_admins":true},"enabled":true}}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsAlarmTemplates.UpdateOrgAlarmTemplate(ctx, orgId, alarmtemplateId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
