package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsVPNsTestListOrgVpns tests the behavior of the OrgsVPNs
func TestOrgsVPNsTestListOrgVpns(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsVpNs.ListOrgVpns(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"name":"string","paths":{"property1":{"bfd_profile":"broadband","ip":"string"},"property2":{"bfd_profile":"lte","ip":"string"}}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsVPNsTestCreateOrgVpn tests the behavior of the OrgsVPNs
func TestOrgsVPNsTestCreateOrgVpn(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Vpn
    errBody := json.Unmarshal([]byte(`{"name":"string","paths":{"property1":{"bfd_profile":"broadband","ip":"string"},"property2":{"bfd_profile":"lte","ip":"string"}}}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsVpNs.CreateOrgVpn(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"497f6eca-6276-5009-bfeb-53cbbbba6f1b","modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","paths":{"property1":{"bfd_profile":"broadband","ip":"string"},"property2":{"bfd_profile":"broadband","ip":"string"}}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsVPNsTestDeleteOrgVpn tests the behavior of the OrgsVPNs
func TestOrgsVPNsTestDeleteOrgVpn(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    vpnId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsVpNs.DeleteOrgVpn(ctx, orgId, vpnId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsVPNsTestGetOrgVpn tests the behavior of the OrgsVPNs
func TestOrgsVPNsTestGetOrgVpn(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    vpnId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsVpNs.GetOrgVpn(ctx, orgId, vpnId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"497f6eca-6276-5009-bfeb-53cbbbba6f1b","modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","paths":{"property1":{"bfd_profile":"broadband","ip":"string"},"property2":{"bfd_profile":"broadband","ip":"string"}}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsVPNsTestUpdateOrgVpn tests the behavior of the OrgsVPNs
func TestOrgsVPNsTestUpdateOrgVpn(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    vpnId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Vpn
    errBody := json.Unmarshal([]byte(`{"name":"string","paths":{"property1":{"bfd_profile":"broadband","ip":"string"},"property2":{"bfd_profile":"broadband","ip":"string"}}}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsVpNs.UpdateOrgVpn(ctx, orgId, vpnId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"id":"497f6eca-6276-5009-bfeb-53cbbbba6f1b","modified_time":0,"name":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","paths":{"property1":{"bfd_profile":"broadband","ip":"string"},"property2":{"bfd_profile":"broadband","ip":"string"}}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
