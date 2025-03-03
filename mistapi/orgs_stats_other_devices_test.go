package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsStatsOtherDevicesTestGetOrgOtherDeviceStats tests the behavior of the OrgsStatsOtherDevices
func TestOrgsStatsOtherDevicesTestGetOrgOtherDeviceStats(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceMac := "0000000000ab"
    apiResponse, err := orgsStatsOtherDevices.GetOrgOtherDeviceStats(ctx, orgId, deviceMac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"cached_stats":true,"connected_devices":{"0200010edbca":{"mac":"020001abcdef","name":"DNT-NTR-GWE","port_id":"ge-0/0/1","type":"gateway"}},"last_seen":1740996902,"lldp_enabled":true,"mac":"00304498a1e8","uptime":622828,"vendor":"cradlepoint","vendor_specific":{"interfaces":{"ethernet-IPPT":{"bytes_in":506,"bytes_out":0,"carrier":"","imei":"","imsi":"","ip":"192.168.5.1","link":false,"mode":"lan","rsrp":0,"rsrq":0,"rssi":0,"service_mode":"Ethernet","sinr":0,"state":"","type":"ethernet","uptime":0},"ethernet-lan":{"bytes_in":13072566048,"bytes_out":5617915438,"carrier":"","imei":"","imsi":"","ip":"192.168.0.1","link":false,"mode":"lan","rsrp":0,"rsrq":0,"rssi":0,"service_mode":"Ethernet","sinr":0,"state":"","type":"ethernet","uptime":0},"mdm-8a1084c9":{"bytes_in":0,"bytes_out":0,"carrier":"Unknown Service","imei":"866401234567894","imsi":"","ip":"","link":false,"mode":"wan","rsrp":0,"rsrq":0,"rssi":0,"service_mode":"Not Available","sinr":0,"state":"NOSIM","type":"mdm","uptime":0},"mdm-8a1fc70c":{"bytes_in":5623096929,"bytes_out":12372750366,"carrier":"AT&T","imei":"866401234567893","imsi":"208001234567893","ip":"12.68.86.17","link":true,"mode":"wan","rsrp":-108,"rsrq":-14,"rssi":-74,"service_mode":"5G NSA","sinr":-1.2,"state":"READY","type":"mdm","uptime":2095779}},"ports":{"ethernet-IPPT":{"bytes_in":506,"bytes_out":0,"carrier":"","imei":"","imsi":"","ip":"192.168.5.1","link":false,"mode":"lan","rsrp":0,"rsrq":0,"rssi":0,"service_mode":"Ethernet","sinr":0,"state":"","type":"ethernet","uptime":0},"ethernet-lan":{"bytes_in":13072566048,"bytes_out":5617915438,"carrier":"","imei":"","imsi":"","ip":"192.168.0.1","link":false,"mode":"lan","rsrp":0,"rsrq":0,"rssi":0,"service_mode":"Ethernet","sinr":0,"state":"","type":"ethernet","uptime":0},"mdm-8a1084c9":{"bytes_in":0,"bytes_out":0,"carrier":"Unknown Service","imei":"866401234567892","imsi":"","ip":"","link":false,"mode":"wan","rsrp":0,"rsrq":0,"rssi":0,"service_mode":"Not Available","sinr":0,"state":"NOSIM","type":"mdm","uptime":0},"mdm-8a1fc70c":{"bytes_in":5623096929,"bytes_out":12372750366,"carrier":"Orange","imei":"866401234567891","imsi":"208001234567891","ip":"10.134.237.57","link":true,"mode":"wan","rsrp":-108,"rsrq":-14,"rssi":-74,"service_mode":"5G NSA","sinr":-1.2,"state":"READY","type":"mdm","uptime":2095779}}},"version":"7.24.80"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
