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
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"id":"684dfc5c-fe77-2290-eb1d-ef3d677fe168","name":"AlarmTemplate 1"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestListInstallerAlarmTemplates1 tests the behavior of the Installer
func TestInstallerTestListInstallerAlarmTemplates1(t *testing.T) {
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
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
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
    mType := models.DeviceTypeDefaultApEnum("ap")
    apiResponse, err := installer.ListInstallerDeviceProfiles(ctx, orgId, &mType)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9","name":"DeviceProfile 1"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestListInstallerDeviceProfiles1 tests the behavior of the Installer
func TestInstallerTestListInstallerDeviceProfiles1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mType := models.DeviceTypeDefaultApEnum("ap")
    apiResponse, err := installer.ListInstallerDeviceProfiles(ctx, orgId, &mType)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9","name":"DeviceProfile 1"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestListInstallerListOfRecentlyClaimedDevices tests the behavior of the Installer
func TestInstallerTestListInstallerListOfRecentlyClaimedDevices(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    model := "AP43"
    siteName := "SJ1"
    siteId, errUUID := uuid.Parse("72771e6a-6f5e-4de4-a5b9-1266c4197811")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := installer.ListInstallerListOfRecentlyClaimedDevices(ctx, orgId, &model, &siteName, &siteId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"connected":true,"deviceprofile_name":"SJ1","height":2.7,"mac":"5c5b35000018","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","model":"AP41","name":"hallway","orientation":90,"serial":"FXLH2015150025","site_name":"SJ1","x":150,"y":300}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestListInstallerListOfRecentlyClaimedDevices1 tests the behavior of the Installer
func TestInstallerTestListInstallerListOfRecentlyClaimedDevices1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    model := "AP43"
    siteName := "SJ1"
    siteId, errUUID := uuid.Parse("72771e6a-6f5e-4de4-a5b9-1266c4197811")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := installer.ListInstallerListOfRecentlyClaimedDevices(ctx, orgId, &model, &siteName, &siteId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
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
    body := []string{ "6JG8E-PTFV2-A9Z2N","DVH4V-SNMSZ-PDXBR" }
    apiResponse, err := installer.ClaimInstallerDevices(ctx, orgId, body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"added":["6JG8E-PTFV2-A9Z2N"],"duplicated":["DVH4V-SNMSZ-PDXBR"],"error":["PO1025335ohoh"],"inventory_added":[{"mac":"5c5b35000018","magic":"6JG8EPTFV2A9Z2N","model":"AP41","serial":"FXLH2015150025","type":"ap"}],"inventory_duplicated":[{"mac":"5c5b35000012","magic":"DVH4VSNMSZPDXBR","model":"AP41","serial":"FXLH2015150027","type":"ap"}],"reason":["belongs to another org ('e2f543f7-d6e1-409f-a565-e77a1f098d3b' (other) != '0de5d6fc-219a-414d-a840-67d6b919ad8f' (you))"]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestClaimInstallerDevices1 tests the behavior of the Installer
func TestInstallerTestClaimInstallerDevices1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    body := []string{ "6JG8E-PTFV2-A9Z2N","DVH4V-SNMSZ-PDXBR" }
    apiResponse, err := installer.ClaimInstallerDevices(ctx, orgId, body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
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

// TestInstallerTestGetInstallerDeviceVirtualChassis tests the behavior of the Installer
func TestInstallerTestGetInstallerDeviceVirtualChassis(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    fpc0Mac := "aff827549235"
    apiResponse, err := installer.GetInstallerDeviceVirtualChassis(ctx, orgId, fpc0Mac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"config_type":"nonprovisioned","id":"00000000-0000-0000-1000-52d9107af289","mac":"52d9107af289","members":[{"_idx":0,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":69,"interrupt":1,"load_avg":[0.8,1.09,1.07],"system":11,"user":19},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":0,"mac":"52d9107af289","memory_stat":{"usage":42},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":40.4,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW0000000000","temperatures":[{"celsius":33,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":1,"neighbor_port_id":"vcp-1/1/0","port_id":"vcp-0/1/0"}],"vc_mode":"HiGiG","vc_role":"master","vc_state":"present","version":"21.4R3-S4.18"},{"_idx":1,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":76,"interrupt":0,"load_avg":[0.96,0.87,0.76],"system":6,"user":17},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":1,"mac":"d0dd4991652d","memory_stat":{"usage":18},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":21.2,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW3619300922","temperatures":[{"celsius":32,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":0,"neighbor_port_id":"vcp-0/1/0","port_id":"vcp-1/1/0"}],"vc_mode":"HiGiG","vc_role":"backup","vc_state":"present","version":"21.4R3-S4.18"}],"model":"EX2300-48P","org_id":"1e9a61a9-bc42-42ca-bf67-1ad87443d9b8","serial":"JW3619300157","site_id":"ab0aca7a-d45c-469b-b3bb-4fe240642d0b","status":"connected","type":"switch","vc_mac":"52d9107af289"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestGetInstallerDeviceVirtualChassis1 tests the behavior of the Installer
func TestInstallerTestGetInstallerDeviceVirtualChassis1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    fpc0Mac := "aff827549235"
    apiResponse, err := installer.GetInstallerDeviceVirtualChassis(ctx, orgId, fpc0Mac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"config_type":"nonprovisioned","id":"00000000-0000-0000-1000-52d9107af289","mac":"52d9107af289","members":[{"_idx":0,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":69,"interrupt":1,"load_avg":[0.8,1.09,1.07],"system":11,"user":19},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":0,"mac":"52d9107af289","memory_stat":{"usage":42},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":40.4,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW0000000000","temperatures":[{"celsius":33,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":1,"neighbor_port_id":"vcp-1/1/0","port_id":"vcp-0/1/0"}],"vc_mode":"HiGiG","vc_role":"master","vc_state":"present","version":"21.4R3-S4.18"},{"_idx":1,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":76,"interrupt":0,"load_avg":[0.96,0.87,0.76],"system":6,"user":17},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":1,"mac":"d0dd4991652d","memory_stat":{"usage":18},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":21.2,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW3619300922","temperatures":[{"celsius":32,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":0,"neighbor_port_id":"vcp-0/1/0","port_id":"vcp-1/1/0"}],"vc_mode":"HiGiG","vc_role":"backup","vc_state":"present","version":"21.4R3-S4.18"}],"model":"EX2300-48P","org_id":"1e9a61a9-bc42-42ca-bf67-1ad87443d9b8","serial":"JW3619300157","site_id":"ab0aca7a-d45c-469b-b3bb-4fe240642d0b","status":"connected","type":"switch","vc_mac":"52d9107af289"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestCreateInstallerVirtualChassis tests the behavior of the Installer
func TestInstallerTestCreateInstallerVirtualChassis(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    fpc0Mac := "aff827549235"
    var body models.VirtualChassisConfig
    errBody := json.Unmarshal([]byte(`{"members":[{"mac":"aff827549235","vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"master"},{"mac":"8396cd006c8c","vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"backup"},{"mac":"8396cd00888c","vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"linecard"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := installer.CreateInstallerVirtualChassis(ctx, orgId, fpc0Mac, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"config_type":"nonprovisioned","id":"00000000-0000-0000-1000-52d9107af289","mac":"52d9107af289","members":[{"_idx":0,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":69,"interrupt":1,"load_avg":[0.8,1.09,1.07],"system":11,"user":19},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":0,"mac":"52d9107af289","memory_stat":{"usage":42},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":40.4,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW0000000000","temperatures":[{"celsius":33,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":1,"neighbor_port_id":"vcp-1/1/0","port_id":"vcp-0/1/0"}],"vc_mode":"HiGiG","vc_role":"master","vc_state":"present","version":"21.4R3-S4.18"},{"_idx":1,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":76,"interrupt":0,"load_avg":[0.96,0.87,0.76],"system":6,"user":17},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":1,"mac":"d0dd4991652d","memory_stat":{"usage":18},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":21.2,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW3619300922","temperatures":[{"celsius":32,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":0,"neighbor_port_id":"vcp-0/1/0","port_id":"vcp-1/1/0"}],"vc_mode":"HiGiG","vc_role":"backup","vc_state":"present","version":"21.4R3-S4.18"}],"model":"EX2300-48P","org_id":"1e9a61a9-bc42-42ca-bf67-1ad87443d9b8","serial":"JW3619300157","site_id":"ab0aca7a-d45c-469b-b3bb-4fe240642d0b","status":"connected","type":"switch","vc_mac":"52d9107af289"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestCreateInstallerVirtualChassis1 tests the behavior of the Installer
func TestInstallerTestCreateInstallerVirtualChassis1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    fpc0Mac := "aff827549235"
    var body models.VirtualChassisConfig
    errBody := json.Unmarshal([]byte(`{"members":[{"mac":"aff827549235","vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"master"},{"mac":"8396cd006c8c","vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"backup"},{"mac":"8396cd00888c","vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"linecard"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := installer.CreateInstallerVirtualChassis(ctx, orgId, fpc0Mac, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"config_type":"nonprovisioned","id":"00000000-0000-0000-1000-52d9107af289","mac":"52d9107af289","members":[{"_idx":0,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":69,"interrupt":1,"load_avg":[0.8,1.09,1.07],"system":11,"user":19},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":0,"mac":"52d9107af289","memory_stat":{"usage":42},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":40.4,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW0000000000","temperatures":[{"celsius":33,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":1,"neighbor_port_id":"vcp-1/1/0","port_id":"vcp-0/1/0"}],"vc_mode":"HiGiG","vc_role":"master","vc_state":"present","version":"21.4R3-S4.18"},{"_idx":1,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":76,"interrupt":0,"load_avg":[0.96,0.87,0.76],"system":6,"user":17},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":1,"mac":"d0dd4991652d","memory_stat":{"usage":18},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":21.2,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW3619300922","temperatures":[{"celsius":32,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":0,"neighbor_port_id":"vcp-0/1/0","port_id":"vcp-1/1/0"}],"vc_mode":"HiGiG","vc_role":"backup","vc_state":"present","version":"21.4R3-S4.18"}],"model":"EX2300-48P","org_id":"1e9a61a9-bc42-42ca-bf67-1ad87443d9b8","serial":"JW3619300157","site_id":"ab0aca7a-d45c-469b-b3bb-4fe240642d0b","status":"connected","type":"switch","vc_mac":"52d9107af289"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestUpdateInstallerVirtualChassisMember tests the behavior of the Installer
func TestInstallerTestUpdateInstallerVirtualChassisMember(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    fpc0Mac := "aff827549235"
    var body models.VirtualChassisUpdate
    errBody := json.Unmarshal([]byte(`{"members":[{"mac":"aff827549235","member_id":2,"vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"linecard"},{"mac":"8396cd00777c","member_id":3,"vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"linecard"}],"op":"add"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := installer.UpdateInstallerVirtualChassisMember(ctx, orgId, fpc0Mac, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"config_type":"nonprovisioned","id":"00000000-0000-0000-1000-52d9107af289","mac":"52d9107af289","members":[{"_idx":0,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":69,"interrupt":1,"load_avg":[0.8,1.09,1.07],"system":11,"user":19},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":0,"mac":"52d9107af289","memory_stat":{"usage":42},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":40.4,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW0000000000","temperatures":[{"celsius":33,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":1,"neighbor_port_id":"vcp-1/1/0","port_id":"vcp-0/1/0"}],"vc_mode":"HiGiG","vc_role":"master","vc_state":"present","version":"21.4R3-S4.18"},{"_idx":1,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":76,"interrupt":0,"load_avg":[0.96,0.87,0.76],"system":6,"user":17},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":1,"mac":"d0dd4991652d","memory_stat":{"usage":18},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":21.2,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW3619300922","temperatures":[{"celsius":32,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":0,"neighbor_port_id":"vcp-0/1/0","port_id":"vcp-1/1/0"}],"vc_mode":"HiGiG","vc_role":"backup","vc_state":"present","version":"21.4R3-S4.18"}],"model":"EX2300-48P","org_id":"1e9a61a9-bc42-42ca-bf67-1ad87443d9b8","serial":"JW3619300157","site_id":"ab0aca7a-d45c-469b-b3bb-4fe240642d0b","status":"connected","type":"switch","vc_mac":"52d9107af289"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestUpdateInstallerVirtualChassisMember1 tests the behavior of the Installer
func TestInstallerTestUpdateInstallerVirtualChassisMember1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    fpc0Mac := "aff827549235"
    var body models.VirtualChassisUpdate
    errBody := json.Unmarshal([]byte(`{"members":[{"mac":"aff827549235","member_id":2,"vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"linecard"},{"mac":"8396cd00777c","member_id":3,"vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"linecard"}],"op":"add"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := installer.UpdateInstallerVirtualChassisMember(ctx, orgId, fpc0Mac, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"config_type":"nonprovisioned","id":"00000000-0000-0000-1000-52d9107af289","mac":"52d9107af289","members":[{"_idx":0,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":69,"interrupt":1,"load_avg":[0.8,1.09,1.07],"system":11,"user":19},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":0,"mac":"52d9107af289","memory_stat":{"usage":42},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":40.4,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW0000000000","temperatures":[{"celsius":33,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":1,"neighbor_port_id":"vcp-1/1/0","port_id":"vcp-0/1/0"}],"vc_mode":"HiGiG","vc_role":"master","vc_state":"present","version":"21.4R3-S4.18"},{"_idx":1,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":76,"interrupt":0,"load_avg":[0.96,0.87,0.76],"system":6,"user":17},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":1,"mac":"d0dd4991652d","memory_stat":{"usage":18},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":21.2,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW3619300922","temperatures":[{"celsius":32,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":0,"neighbor_port_id":"vcp-0/1/0","port_id":"vcp-1/1/0"}],"vc_mode":"HiGiG","vc_role":"backup","vc_state":"present","version":"21.4R3-S4.18"}],"model":"EX2300-48P","org_id":"1e9a61a9-bc42-42ca-bf67-1ad87443d9b8","serial":"JW3619300157","site_id":"ab0aca7a-d45c-469b-b3bb-4fe240642d0b","status":"connected","type":"switch","vc_mac":"52d9107af289"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
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
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"id":"bb8a9017-1e36-5d6c-6f2b-551abe8a76a2","name":"RFTemplate 1"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestListInstallerRfTemplatesNames1 tests the behavior of the Installer
func TestInstallerTestListInstallerRfTemplatesNames1(t *testing.T) {
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
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
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
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"id":"581328b6-e382-f54e-c9dc-999983183a34","name":"SiteGroup 1"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestListInstallerSiteGroups1 tests the behavior of the Installer
func TestInstallerTestListInstallerSiteGroups1(t *testing.T) {
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
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
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
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"address":"1601 S. Deanza Blvd., Cupertino, CA, 95014","country_code":"US","id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","latlng":{"lat":37.295833,"lng":-122.032946},"name":"Mist Office","rftemplate_name":"rftemplate1","sitegroup_names":["sg1","sg2"],"timezone":"America/Los_Angeles"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestInstallerTestListInstallerSites1 tests the behavior of the Installer
func TestInstallerTestListInstallerSites1(t *testing.T) {
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
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"address":"1601 S. Deanza Blvd., Cupertino, CA, 95014","country_code":"US","id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","latlng":{"lat":37.295833,"lng":-122.032946},"name":"Mist Office","rftemplate_name":"rftemplate1","sitegroup_names":["sg1","sg2"],"timezone":"America/Los_Angeles"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
