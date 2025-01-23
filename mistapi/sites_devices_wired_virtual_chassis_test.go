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
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"config_type":"nonprovisioned","id":"00000000-0000-0000-1000-52d9107af289","mac":"52d9107af289","members":[{"_idx":0,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":69,"interrupt":1,"load_avg":[0.8,1.09,1.07],"system":11,"user":19},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":0,"mac":"52d9107af289","memory_stat":{"usage":42},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":40.4,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW0000000000","temperatures":[{"celsius":33,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":1,"neighbor_port_id":"vcp-1/1/0","port_id":"vcp-0/1/0"}],"vc_mode":"HiGiG","vc_role":"master","vc_state":"present","version":"21.4R3-S4.18"},{"_idx":1,"boot_partition":"junos","cpld_version":"6","cpu_stat":{"idle":76,"interrupt":0,"load_avg":[0.96,0.87,0.76],"system":6,"user":17},"fans":[{"airflow":"out","name":"Fan Tray 0 Fan 0","rpm":0,"status":"ok"},{"airflow":"out","name":"Fan Tray 1 Fan 0","rpm":0,"status":"ok"}],"fpc_idx":1,"mac":"d0dd4991652d","memory_stat":{"usage":18},"model":"EX2300-48P","pics":[{"index":0,"model_number":"EX2300-48P","port_groups":[{"count":48,"type":"GE"}]},{"index":1,"model_number":"EX2300-48P","port_groups":[{"count":4,"type":"SFP/SFP+"}]}],"poe":{"max_power":750,"power_draw":21.2,"status":"AT_MODE"},"poe_version":"2.1.1.19.3 (type1)","psus":[{"name":"Power Supply 0","status":"ok"}],"recovery_version":"21.4R3-S4.18","serial":"JW3619300922","temperatures":[{"celsius":32,"name":"CPU Sensor","status":"ok"},{"celsius":29,"name":"PSU Sensor","status":"ok"}],"type":"fpc","uboot_version":"U-Boot 2016.01-rc1 (Nov 11 2016 - 14:26:00 -0800)  1.3.2","uptime":27636720,"vc_links":[{"neighbor_module_idx":0,"neighbor_port_id":"vcp-0/1/0","port_id":"vcp-1/1/0"}],"vc_mode":"HiGiG","vc_role":"backup","vc_state":"present","version":"21.4R3-S4.18"}],"model":"EX2300-48P","org_id":"1e9a61a9-bc42-42ca-bf67-1ad87443d9b8","serial":"JW3619300157","site_id":"ab0aca7a-d45c-469b-b3bb-4fe240642d0b","status":"connected","type":"switch","vc_mac":"52d9107af289"}`
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
    resp, err := sitesDevicesWiredVirtualChassis.CreateSiteVirtualChassis(ctx, siteId, deviceId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
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
    errBody := json.Unmarshal([]byte(`{"members":[{"mac":"aff827549235","member_id":2,"vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"linecard"},{"mac":"8396cd00777c","member_id":3,"vc_ports":["xe-0/1/0","xe-0/1/1"],"vc_role":"linecard"}],"op":"add"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := sitesDevicesWiredVirtualChassis.UpdateSiteVirtualChassisMember(ctx, siteId, deviceId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
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
