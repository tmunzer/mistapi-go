package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsNACPortalsTestListOrgNacPortals tests the behavior of the OrgsNACPortals
func TestOrgsNACPortalsTestListOrgNacPortals(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsNacPortals.ListOrgNacPortals(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"access_type":"wireless","cert_expire_time":365,"enable_telemetry":true,"expiry_notification_time":2,"name":"get-wifi","notify_expiry":true,"ssid":"Corp","sso":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\n-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://yourorg.onelogin.com/trust/saml2/http-post/sso/138130","issuer":"https://app.onelogin.com/saml/metadata/138130","nameid_format":"email","sso_role_matching":[{"assigned":"user","match":"Student"}],"use_sso_role_for_cert":true}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACPortalsTestCreateOrgNacPortal tests the behavior of the OrgsNACPortals
func TestOrgsNACPortalsTestCreateOrgNacPortal(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsNacPortals.CreateOrgNacPortal(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"access_type":"wireless","cert_expire_time":365,"enable_telemetry":true,"expiry_notification_time":2,"name":"get-wifi","notify_expiry":true,"ssid":"Corp","sso":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\n-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://yourorg.onelogin.com/trust/saml2/http-post/sso/138130","issuer":"https://app.onelogin.com/saml/metadata/138130","nameid_format":"email","sso_role_matching":[{"assigned":"user","match":"Student"}],"use_sso_role_for_cert":true}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACPortalsTestDeleteOrgNacPortal tests the behavior of the OrgsNACPortals
func TestOrgsNACPortalsTestDeleteOrgNacPortal(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsNacPortals.DeleteOrgNacPortal(ctx, orgId, nacportalId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsNACPortalsTestGetOrgNacPortal tests the behavior of the OrgsNACPortals
func TestOrgsNACPortalsTestGetOrgNacPortal(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsNacPortals.GetOrgNacPortal(ctx, orgId, nacportalId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"access_type":"wireless","cert_expire_time":365,"enable_telemetry":true,"expiry_notification_time":2,"name":"get-wifi","notify_expiry":true,"ssid":"Corp","sso":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\n-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://yourorg.onelogin.com/trust/saml2/http-post/sso/138130","issuer":"https://app.onelogin.com/saml/metadata/138130","nameid_format":"email","sso_role_matching":[{"assigned":"user","match":"Student"}],"use_sso_role_for_cert":true}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACPortalsTestUpdateOrgNacPortal tests the behavior of the OrgsNACPortals
func TestOrgsNACPortalsTestUpdateOrgNacPortal(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsNacPortals.UpdateOrgNacPortal(ctx, orgId, nacportalId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"access_type":"wireless","cert_expire_time":365,"enable_telemetry":true,"expiry_notification_time":2,"name":"get-wifi","notify_expiry":true,"ssid":"Corp","sso":{"idp_cert":"-----BEGIN CERTIFICATE-----\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\n-----END CERTIFICATE-----","idp_sign_algo":"sha256","idp_sso_url":"https://yourorg.onelogin.com/trust/saml2/http-post/sso/138130","issuer":"https://app.onelogin.com/saml/metadata/138130","nameid_format":"email","sso_role_matching":[{"assigned":"user","match":"Student"}],"use_sso_role_for_cert":true}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACPortalsTestListOrgNacPortalSsoLatestFailures tests the behavior of the OrgsNACPortals
func TestOrgsNACPortalsTestListOrgNacPortalSsoLatestFailures(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsNacPortals.ListOrgNacPortalSsoLatestFailures(ctx, orgId, nacportalId, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"results":[{"detail":"string","saml_assertion_xml":"string","timestamp":0}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACPortalsTestDeleteOrgNacPortalImage tests the behavior of the OrgsNACPortals
func TestOrgsNACPortalsTestDeleteOrgNacPortalImage(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsNacPortals.DeleteOrgNacPortalImage(ctx, orgId, nacportalId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsNACPortalsTestUploadOrgNacPortalImage tests the behavior of the OrgsNACPortals
func TestOrgsNACPortalsTestUploadOrgNacPortalImage(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    resp, err := orgsNacPortals.UploadOrgNacPortalImage(ctx, orgId, nacportalId, nil, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsNACPortalsTestUpdateOrgNacPortalTempalte tests the behavior of the OrgsNACPortals
func TestOrgsNACPortalsTestUpdateOrgNacPortalTempalte(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    resp, err := orgsNacPortals.UpdateOrgNacPortalTempalte(ctx, orgId, nacportalId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsNACPortalsTestGetOrgNacPortalSsoSamlMetadata tests the behavior of the OrgsNACPortals
func TestOrgsNACPortalsTestGetOrgNacPortalSsoSamlMetadata(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsNacPortals.GetOrgNacPortalSsoSamlMetadata(ctx, orgId, nacportalId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"acs_url":"https://api.mist.com/api/v1/saml/llDfa13f/login","entity_id":"https://api.mist.com/api/v1/saml/llDfa13f/login","logout_url":"https://api.mist.com/api/v1/saml/llDfa13f/logout","metadata":"<?xml version=\"1.0\" encoding=\"UTF-8\"?><md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://api.mist.com/api/v1/saml/llDfa13f/login\" validUntil=\"2027-10-12T21:59:01Z\" xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"><md:SPSSODescriptor AuthnRequestsSigned=\"false\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://api.mist.com/api/v1/saml/llDfa13f/login\" index=\"0\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACPortalsTestDownloadOrgNacPortalSsoSamlMetadata tests the behavior of the OrgsNACPortals
func TestOrgsNACPortalsTestDownloadOrgNacPortalSsoSamlMetadata(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    nacportalId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsNacPortals.DownloadOrgNacPortalSsoSamlMetadata(ctx, orgId, nacportalId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
