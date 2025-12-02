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

// TestOrgsPskPortalsTestListOrgPskPortals tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestListOrgPskPortals(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsPskPortals.ListOrgPskPortals(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsPskPortalsTestListOrgPskPortals1 tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestListOrgPskPortals1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsPskPortals.ListOrgPskPortals(ctx, orgId, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsPskPortalsTestCreateOrgPskPortal tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestCreateOrgPskPortal(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.PskPortal
	errBody := json.Unmarshal([]byte(`{"auth":"sso","expire_time":262800,"expiry_notification_time":2,"hide_psks_created_by_other_admins":false,"max_usage":1,"name":"USR-test","notify_expiry":false,"passphrase_rules":{"alphabets_enabled":true,"length":8,"numerics_enabled":true,"symbols":"()[]{}_%@#&$","symbols_enabled":true},"ssid":"USR-test","sso":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIDqDCCApCgA...-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://dev-00000000.okta.com/app/dev-00000000/a6b88bec5ab7/sso/saml","issuer":"http://www.okta.com/a6b88bec5ab7","nameid_format":"email"},"type":"byod","vlan_id":42}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsPskPortals.CreateOrgPskPortal(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"auth":"sso","created_time":1727272316,"expire_time":262800,"expiry_notification_time":2,"hide_psks_created_by_other_admins":false,"id":"86ef73e5-e360-4e33-abd8-c2da7b11cbd6","max_usage":1,"modified_time":1727272316,"name":"USR-test","notify_expiry":false,"org_id":"8f6c24d5-b618-4c30-b6ff-b2febc1078ab","passphrase_rules":{"alphabets_enabled":true,"length":8,"numerics_enabled":true,"symbols":"()[]{}_%@#&$","symbols_enabled":true},"role":"","ssid":"USR-test","sso":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIDqDCCApCgA...-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://dev-00000000.okta.com/app/dev-00000000/a6b88bec5ab7/sso/saml","issuer":"http://www.okta.com/a6b88bec5ab7","nameid_format":"email"},"type":"byod","ui_url":"https://pskportal.gc1.mist.com/#!byod/a6b88bec-5ab7-41ae-a34c-44ff52a8b81d","vlan_id":123}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestCreateOrgPskPortal1 tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestCreateOrgPskPortal1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.PskPortal
	errBody := json.Unmarshal([]byte(`{"auth":"sso","expire_time":262800,"expiry_notification_time":2,"hide_psks_created_by_other_admins":false,"max_usage":1,"name":"USR-test","notify_expiry":false,"passphrase_rules":{"alphabets_enabled":true,"length":8,"numerics_enabled":true,"symbols":"()[]{}_%@#&$","symbols_enabled":true},"ssid":"USR-test","sso":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIDqDCCApCgA...-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://dev-00000000.okta.com/app/dev-00000000/a6b88bec5ab7/sso/saml","issuer":"http://www.okta.com/a6b88bec5ab7","nameid_format":"email"},"type":"byod","vlan_id":42}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsPskPortals.CreateOrgPskPortal(ctx, orgId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"auth":"sso","created_time":1727272316,"expire_time":262800,"expiry_notification_time":2,"hide_psks_created_by_other_admins":false,"id":"86ef73e5-e360-4e33-abd8-c2da7b11cbd6","max_usage":1,"modified_time":1727272316,"name":"USR-test","notify_expiry":false,"org_id":"8f6c24d5-b618-4c30-b6ff-b2febc1078ab","passphrase_rules":{"alphabets_enabled":true,"length":8,"numerics_enabled":true,"symbols":"()[]{}_%@#&$","symbols_enabled":true},"role":"","ssid":"USR-test","sso":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIDqDCCApCgA...-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://dev-00000000.okta.com/app/dev-00000000/a6b88bec5ab7/sso/saml","issuer":"http://www.okta.com/a6b88bec5ab7","nameid_format":"email"},"type":"byod","ui_url":"https://pskportal.gc1.mist.com/#!byod/a6b88bec-5ab7-41ae-a34c-44ff52a8b81d","vlan_id":123}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestListOrgPskPortalLogs tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestListOrgPskPortalLogs(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsPskPortals.ListOrgPskPortalLogs(ctx, orgId, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1428954000,"limit":100,"results":[{"id":"8a3dcaa7-80e3-4bb0-a75b-7bc6322cfd09","message":"Rotate PSK test@mist.com","name_id":"test@mist.com","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","psk_id":"608fe603-f9f0-4ce9-9473-04ef6c6ea749","psk_name":"test@mist.com","pskportal_id":"c1742c09-af35-4161-96ef-7dc65c6d5674","timestamp":1686346104.096}],"start":1428939600,"total":135}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestListOrgPskPortalLogs1 tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestListOrgPskPortalLogs1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsPskPortals.ListOrgPskPortalLogs(ctx, orgId, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1428954000,"limit":100,"results":[{"id":"8a3dcaa7-80e3-4bb0-a75b-7bc6322cfd09","message":"Rotate PSK test@mist.com","name_id":"test@mist.com","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","psk_id":"608fe603-f9f0-4ce9-9473-04ef6c6ea749","psk_name":"test@mist.com","pskportal_id":"c1742c09-af35-4161-96ef-7dc65c6d5674","timestamp":1686346104.096}],"start":1428939600,"total":135}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestCountOrgPskPortalLogs tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestCountOrgPskPortalLogs(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgPskPortalLogsCountDistinctEnum("pskportal_id")

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsPskPortals.CountOrgPskPortalLogs(ctx, orgId, &distinct, nil, nil, &duration, &limit)
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

// TestOrgsPskPortalsTestCountOrgPskPortalLogs1 tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestCountOrgPskPortalLogs1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	distinct := models.OrgPskPortalLogsCountDistinctEnum("pskportal_id")

	duration := "1d"
	limit := int(100)
	apiResponse, err := orgsPskPortals.CountOrgPskPortalLogs(ctx, orgId, &distinct, nil, nil, &duration, &limit)
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

// TestOrgsPskPortalsTestSearchOrgPskPortalLogs tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestSearchOrgPskPortalLogs(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsPskPortals.SearchOrgPskPortalLogs(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1428954000,"limit":100,"results":[{"id":"8a3dcaa7-80e3-4bb0-a75b-7bc6322cfd09","message":"Rotate PSK test@mist.com","name_id":"test@mist.com","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","psk_id":"608fe603-f9f0-4ce9-9473-04ef6c6ea749","psk_name":"test@mist.com","pskportal_id":"c1742c09-af35-4161-96ef-7dc65c6d5674","timestamp":1686346104.096}],"start":1428939600,"total":135}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestSearchOrgPskPortalLogs1 tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestSearchOrgPskPortalLogs1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)

	duration := "1d"
	sort := "timestamp"

	apiResponse, err := orgsPskPortals.SearchOrgPskPortalLogs(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1428954000,"limit":100,"results":[{"id":"8a3dcaa7-80e3-4bb0-a75b-7bc6322cfd09","message":"Rotate PSK test@mist.com","name_id":"test@mist.com","org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","psk_id":"608fe603-f9f0-4ce9-9473-04ef6c6ea749","psk_name":"test@mist.com","pskportal_id":"c1742c09-af35-4161-96ef-7dc65c6d5674","timestamp":1686346104.096}],"start":1428939600,"total":135}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestDeleteOrgPskPortal tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestDeleteOrgPskPortal(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	pskportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsPskPortals.DeleteOrgPskPortal(ctx, orgId, pskportalId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsPskPortalsTestGetOrgPskPortal tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestGetOrgPskPortal(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	pskportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsPskPortals.GetOrgPskPortal(ctx, orgId, pskportalId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"auth":"sso","created_time":1727272316,"expire_time":262800,"expiry_notification_time":2,"hide_psks_created_by_other_admins":false,"id":"86ef73e5-e360-4e33-abd8-c2da7b11cbd6","max_usage":1,"modified_time":1727272316,"name":"USR-test","notify_expiry":false,"org_id":"8f6c24d5-b618-4c30-b6ff-b2febc1078ab","passphrase_rules":{"alphabets_enabled":true,"length":8,"numerics_enabled":true,"symbols":"()[]{}_%@#&$","symbols_enabled":true},"role":"","ssid":"USR-test","sso":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIDqDCCApCgA...-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://dev-00000000.okta.com/app/dev-00000000/a6b88bec5ab7/sso/saml","issuer":"http://www.okta.com/a6b88bec5ab7","nameid_format":"email"},"type":"byod","ui_url":"https://pskportal.gc1.mist.com/#!byod/a6b88bec-5ab7-41ae-a34c-44ff52a8b81d","vlan_id":123}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestGetOrgPskPortal1 tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestGetOrgPskPortal1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	pskportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := orgsPskPortals.GetOrgPskPortal(ctx, orgId, pskportalId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"auth":"sso","created_time":1727272316,"expire_time":262800,"expiry_notification_time":2,"hide_psks_created_by_other_admins":false,"id":"86ef73e5-e360-4e33-abd8-c2da7b11cbd6","max_usage":1,"modified_time":1727272316,"name":"USR-test","notify_expiry":false,"org_id":"8f6c24d5-b618-4c30-b6ff-b2febc1078ab","passphrase_rules":{"alphabets_enabled":true,"length":8,"numerics_enabled":true,"symbols":"()[]{}_%@#&$","symbols_enabled":true},"role":"","ssid":"USR-test","sso":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIDqDCCApCgA...-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://dev-00000000.okta.com/app/dev-00000000/a6b88bec5ab7/sso/saml","issuer":"http://www.okta.com/a6b88bec5ab7","nameid_format":"email"},"type":"byod","ui_url":"https://pskportal.gc1.mist.com/#!byod/a6b88bec-5ab7-41ae-a34c-44ff52a8b81d","vlan_id":123}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestUpdateOrgPskPortal tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestUpdateOrgPskPortal(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	pskportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.PskPortal
	errBody := json.Unmarshal([]byte(`{"auth":"sso","expire_time":262800,"expiry_notification_time":2,"hide_psks_created_by_other_admins":false,"max_usage":1,"name":"USR-test","notify_expiry":false,"passphrase_rules":{"alphabets_enabled":true,"length":8,"numerics_enabled":true,"symbols":"()[]{}_%@#&$","symbols_enabled":true},"sponsors":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIDqDCCApCgA...-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://dev-00000000.okta.com/app/dev-00000000/a6b88bec5ab7/sso/saml","issuer":"http://www.okta.com/a6b88bec5ab7","nameid_format":"email"},"ssid":"USR-test","type":"byod","vlan_id":"nu42ll"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsPskPortals.UpdateOrgPskPortal(ctx, orgId, pskportalId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"auth":"sso","created_time":1727272316,"expire_time":262800,"expiry_notification_time":2,"hide_psks_created_by_other_admins":false,"id":"86ef73e5-e360-4e33-abd8-c2da7b11cbd6","max_usage":1,"modified_time":1727272316,"name":"USR-test","notify_expiry":false,"org_id":"8f6c24d5-b618-4c30-b6ff-b2febc1078ab","passphrase_rules":{"alphabets_enabled":true,"length":8,"numerics_enabled":true,"symbols":"()[]{}_%@#&$","symbols_enabled":true},"role":"","ssid":"USR-test","sso":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIDqDCCApCgA...-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://dev-00000000.okta.com/app/dev-00000000/a6b88bec5ab7/sso/saml","issuer":"http://www.okta.com/a6b88bec5ab7","nameid_format":"email"},"type":"byod","ui_url":"https://pskportal.gc1.mist.com/#!byod/a6b88bec-5ab7-41ae-a34c-44ff52a8b81d","vlan_id":123}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestUpdateOrgPskPortal1 tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestUpdateOrgPskPortal1(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	pskportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.PskPortal
	errBody := json.Unmarshal([]byte(`{"auth":"sso","expire_time":262800,"expiry_notification_time":2,"hide_psks_created_by_other_admins":false,"max_usage":1,"name":"USR-test","notify_expiry":false,"passphrase_rules":{"alphabets_enabled":true,"length":8,"numerics_enabled":true,"symbols":"()[]{}_%@#&$","symbols_enabled":true},"sponsors":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIDqDCCApCgA...-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://dev-00000000.okta.com/app/dev-00000000/a6b88bec5ab7/sso/saml","issuer":"http://www.okta.com/a6b88bec5ab7","nameid_format":"email"},"ssid":"USR-test","type":"byod","vlan_id":"nu42ll"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := orgsPskPortals.UpdateOrgPskPortal(ctx, orgId, pskportalId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/vnd.api+json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"auth":"sso","created_time":1727272316,"expire_time":262800,"expiry_notification_time":2,"hide_psks_created_by_other_admins":false,"id":"86ef73e5-e360-4e33-abd8-c2da7b11cbd6","max_usage":1,"modified_time":1727272316,"name":"USR-test","notify_expiry":false,"org_id":"8f6c24d5-b618-4c30-b6ff-b2febc1078ab","passphrase_rules":{"alphabets_enabled":true,"length":8,"numerics_enabled":true,"symbols":"()[]{}_%@#&$","symbols_enabled":true},"role":"","ssid":"USR-test","sso":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIDqDCCApCgA...-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://dev-00000000.okta.com/app/dev-00000000/a6b88bec5ab7/sso/saml","issuer":"http://www.okta.com/a6b88bec5ab7","nameid_format":"email"},"type":"byod","ui_url":"https://pskportal.gc1.mist.com/#!byod/a6b88bec-5ab7-41ae-a34c-44ff52a8b81d","vlan_id":123}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsPskPortalsTestDeleteOrgPskPortalImage tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestDeleteOrgPskPortalImage(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	pskportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := orgsPskPortals.DeleteOrgPskPortalImage(ctx, orgId, pskportalId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsPskPortalsTestUploadOrgPskPortalImage tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestUploadOrgPskPortalImage(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	pskportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	resp, err := orgsPskPortals.UploadOrgPskPortalImage(ctx, orgId, pskportalId, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsPskPortalsTestUpdateOrgPskPortalTemplate tests the behavior of the OrgsPskPortals
func TestOrgsPskPortalsTestUpdateOrgPskPortalTemplate(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	pskportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	resp, err := orgsPskPortals.UpdateOrgPskPortalTemplate(ctx, orgId, pskportalId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
