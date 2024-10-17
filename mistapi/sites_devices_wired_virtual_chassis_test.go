package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSitesDevicesWiredVirtualChassisTestDeleteSiteVirtualChassis tests the behavior of the SitesDevicesWiredVirtualChassis
func TestSitesDevicesWiredVirtualChassisTestDeleteSiteVirtualChassis(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesDevicesWiredVirtualChassis.DeleteSiteVirtualChassis(ctx, siteId, deviceId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesDevicesWiredVirtualChassisTestGetSiteDeviceVirtualChassis tests the behavior of the SitesDevicesWiredVirtualChassis
func TestSitesDevicesWiredVirtualChassisTestGetSiteDeviceVirtualChassis(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesDevicesWiredVirtualChassis.GetSiteDeviceVirtualChassis(ctx, siteId, deviceId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":{"members":[{"mac":"string","member":0,"vc_ports":["string"],"vc_role":"master"}],"op":"add"}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesWiredVirtualChassisTestCreateSiteVirtualChassis tests the behavior of the SitesDevicesWiredVirtualChassis
func TestSitesDevicesWiredVirtualChassisTestCreateSiteVirtualChassis(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.VirtualChassisConfig
	errBody := json.Unmarshal([]byte(`{"members":[{"mac":"aff827549235","vc_ports":["xe-0/1/0"],"vc_role":"master"},{"mac":"8396cd006c8c","vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"backup"},{"mac":"8396cd00888c","vc_ports":["xe-0/1/0"],"vc_role":"linecard"}]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesDevicesWiredVirtualChassis.CreateSiteVirtualChassis(ctx, siteId, deviceId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":{"members":[{"mac":"string","member":0,"vc_ports":["string"],"vc_role":"master"}],"op":"add"}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesWiredVirtualChassisTestUpdateSiteVirtualChassisMember tests the behavior of the SitesDevicesWiredVirtualChassis
func TestSitesDevicesWiredVirtualChassisTestUpdateSiteVirtualChassisMember(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.VirtualChassisUpdate
	errBody := json.Unmarshal([]byte(`{"members":[{"mac":"aff827549235","member":0,"vc_ports":["xe-0/1/1"],"vc_role":"linecard"},{"mac":"8396cd00777c","vc_ports":["xe-0/1/0"],"vc_role":"linecard"}],"op":"add"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesDevicesWiredVirtualChassis.UpdateSiteVirtualChassisMember(ctx, siteId, deviceId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"id":{"members":[{"mac":"string","member":0,"vc_ports":["string"],"vc_role":"master"}],"op":"add"}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesWiredVirtualChassisTestSetSiteVcPort tests the behavior of the SitesDevicesWiredVirtualChassis
func TestSitesDevicesWiredVirtualChassisTestSetSiteVcPort(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.VirtualChassisPort
	errBody := json.Unmarshal([]byte(`{"members":[{"member":0,"vc_ports":["xe-0/1/1"]},{"member":2,"vc_ports":["xe-0/1/1"]}],"op":"delete"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	resp, err := sitesDevicesWiredVirtualChassis.SetSiteVcPort(ctx, siteId, deviceId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
