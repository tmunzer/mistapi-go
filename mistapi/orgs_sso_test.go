package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsSSOTestListOrgSsos tests the behavior of the OrgsSSO
func TestOrgsSSOTestListOrgSsos(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsSso.ListOrgSsos(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"custom_logout_url":"string","default_role":"string","domain":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","idp_cert":"string","idp_sign_algo":"string","idp_sso_url":"string","ignore_unmatched_roles":true,"issuer":"string","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","nameid_format":"email","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","type":"string"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSSOTestCreateOrgSso tests the behavior of the OrgsSSO
func TestOrgsSSOTestCreateOrgSso(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Sso
    errBody := json.Unmarshal([]byte(`{"custom_logout_url":"string","idp_cert":"string","idp_sign_algo":"string","idp_sso_url":"string","ignore_unmatched_roles":true,"issuer":"string","name":"string","nameid_format":"email"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsSso.CreateOrgSso(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"custom_logout_url":"string","default_role":"string","domain":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","idp_cert":"string","idp_sign_algo":"string","idp_sso_url":"string","ignore_unmatched_roles":true,"issuer":"string","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","nameid_format":"email","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","type":"string"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSSOTestDeleteOrgSso tests the behavior of the OrgsSSO
func TestOrgsSSOTestDeleteOrgSso(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ssoId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsSso.DeleteOrgSso(ctx, orgId, ssoId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSSOTestGetOrgSso tests the behavior of the OrgsSSO
func TestOrgsSSOTestGetOrgSso(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ssoId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSso.GetOrgSso(ctx, orgId, ssoId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"custom_logout_url":"string","default_role":"string","domain":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","idp_cert":"string","idp_sign_algo":"string","idp_sso_url":"string","ignore_unmatched_roles":true,"issuer":"string","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","nameid_format":"email","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","type":"string"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSSOTestUpdateOrgSso tests the behavior of the OrgsSSO
func TestOrgsSSOTestUpdateOrgSso(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ssoId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSso.UpdateOrgSso(ctx, orgId, ssoId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"custom_logout_url":"string","default_role":"string","domain":"string","id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","idp_cert":"string","idp_sign_algo":"string","idp_sso_url":"string","ignore_unmatched_roles":true,"issuer":"string","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","nameid_format":"email","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","type":"string"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSSOTestListOrgSsoLatestFailures tests the behavior of the OrgsSSO
func TestOrgsSSOTestListOrgSsoLatestFailures(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ssoId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsSso.ListOrgSsoLatestFailures(ctx, orgId, ssoId, nil, nil, &duration, &limit, &page)
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

// TestOrgsSSOTestGetOrgSsoSamlMetadata tests the behavior of the OrgsSSO
func TestOrgsSSOTestGetOrgSsoSamlMetadata(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ssoId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSso.GetOrgSsoSamlMetadata(ctx, orgId, ssoId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"acs_url":"https://api.mist.com/api/v1/saml/llDfa13f/login","entity_id":"https://api.mist.com/api/v1/saml/llDfa13f/login","logout_url":"https://api.mist.com/api/v1/saml/llDfa13f/logout","metadata_xml":"<?xml version=\"1.0\" encoding=\"UTF-8\"?><md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://api.mist.com/api/v1/saml/llDfa13f/login\" validUntil=\"2027-10-12T21:59:01Z\" xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"><md:SPSSODescriptor AuthnRequestsSigned=\"false\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\"><md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat><md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://api.mist.com/api/v1/saml/llDfa13f/login\" index=\"0\" isDefault=\"true\"/></md:SPSSODescriptor></md:EntityDescriptor>"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSSOTestDownloadOrgSsoSamlMetadata tests the behavior of the OrgsSSO
func TestOrgsSSOTestDownloadOrgSsoSamlMetadata(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    ssoId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSso.DownloadOrgSsoSamlMetadata(ctx, orgId, ssoId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
