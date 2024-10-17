package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestInstallerTestListInstallerAlarmTemplates tests the behavior of the Installer
func TestInstallerTestListInstallerAlarmTemplates(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := installer.ListInstallerAlarmTemplates(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"id":"684dfc5c-fe77-2290-eb1d-ef3d677fe168","name":"AlarmTemplate 1"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestListInstallerDeviceProfiles tests the behavior of the Installer
func TestInstallerTestListInstallerDeviceProfiles(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	mType := models.DeviceTypeEnum("ap")
	apiResponse, err := installer.ListInstallerDeviceProfiles(ctx, orgId, &mType)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9","name":"DeviceProfile 1"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestListInstallerListOfRenctlyClaimedDevices tests the behavior of the Installer
func TestInstallerTestListInstallerListOfRenctlyClaimedDevices(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	limit := int(100)
	page := int(1)
	apiResponse, err := installer.ListInstallerListOfRenctlyClaimedDevices(ctx, orgId, nil, nil, nil, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"connected":true,"deviceprofile_name":"SJ1","height":2.7,"mac":"5c5b35000018","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","model":"AP41","name":"hallway","orientation":90,"serial":"FXLH2015150025","site_name":"SJ1","x":150,"y":300}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestClaimInstallerDevices tests the behavior of the Installer
func TestInstallerTestClaimInstallerDevices(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	body := []string{"6JG8E-PTFV2-A9Z2N", "DVH4V-SNMSZ-PDXBR"}
	apiResponse, err := installer.ClaimInstallerDevices(ctx, orgId, body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"added":["6JG8E-PTFV2-A9Z2N"],"duplicated":["DVH4V-SNMSZ-PDXBR"],"error":["PO1025335ohoh"],"inventory_added":[{"mac":"5c5b35000018","magic":"6JG8EPTFV2A9Z2N","model":"AP41","serial":"FXLH2015150025","type":"ap"}],"inventory_duplicated":[{"mac":"5c5b35000012","magic":"DVH4VSNMSZPDXBR","model":"AP41","serial":"FXLH2015150027","type":"ap"}],"reason":["belongs to another org ('e2f543f7-d6e1-409f-a565-e77a1f098d3b' (other) != '0de5d6fc-219a-414d-a840-67d6b919ad8f' (you))"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestUnassignInstallerRecentlyClaimedDevice tests the behavior of the Installer
func TestInstallerTestUnassignInstallerRecentlyClaimedDevice(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceMac := "0000000000ab"
	resp, err := installer.UnassignInstallerRecentlyClaimedDevice(ctx, orgId, deviceMac)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestInstallerTestProvisionInstallerDevices tests the behavior of the Installer
func TestInstallerTestProvisionInstallerDevices(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceMac := "0000000000ab"
	var body models.InstallerProvisionDevice
	errBody := json.Unmarshal([]byte(`{"deviceprofile_name":"SJ1","height":2.7,"name":"SJ1-AP1","orientation":90,"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","site_name":"SJ1","x":150,"y":300}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	resp, err := installer.ProvisionInstallerDevices(ctx, orgId, deviceMac, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestInstallerTestStartInstallerLocateDevice tests the behavior of the Installer
func TestInstallerTestStartInstallerLocateDevice(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceMac := "0000000000ab"
	resp, err := installer.StartInstallerLocateDevice(ctx, orgId, deviceMac)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestInstallerTestStopInstallerLocateDevice tests the behavior of the Installer
func TestInstallerTestStopInstallerLocateDevice(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	deviceMac := "0000000000ab"
	resp, err := installer.StopInstallerLocateDevice(ctx, orgId, deviceMac)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestInstallerTestListInstallerRfTemplatesNames tests the behavior of the Installer
func TestInstallerTestListInstallerRfTemplatesNames(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := installer.ListInstallerRfTemplatesNames(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"id":"bb8a9017-1e36-5d6c-6f2b-551abe8a76a2","name":"RFTemplate 1"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestListInstallerSiteGroups tests the behavior of the Installer
func TestInstallerTestListInstallerSiteGroups(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := installer.ListInstallerSiteGroups(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"id":"581328b6-e382-f54e-c9dc-999983183a34","name":"SiteGroup 1"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestListInstallerSites tests the behavior of the Installer
func TestInstallerTestListInstallerSites(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := installer.ListInstallerSites(ctx, orgId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `[{"address":"1601 S. Deanza Blvd., Cupertino, CA, 95014","country_code":"US","id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","latlng":{"lat":37.295833,"lng":-122.032946},"name":"Mist Office","rftemplate_name":"rftemplate1","sitegroup_names":["sg1","sg2"],"timezone":"America/Los_Angeles"}]`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
