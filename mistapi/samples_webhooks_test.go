package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSamplesWebhooksTestAlarms tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestAlarms(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookAlarms
    errBody := json.Unmarshal([]byte(`{"events":[{"aps":["string"],"bssids":["string"],"count":0,"event_id":"a7a26ff2-e851-45b6-9634-d595f45458b7","for_site":true,"id":"489f6eca-6276-4993-bfeb-c3cbbbba1f08","last_seen":0,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","ssids":["string"],"timestamp":0,"type":"string","update":true}],"topic":"alarms"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.Alarms(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestAssetRaw tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestAssetRaw(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookAssetRaw
    errBody := json.Unmarshal([]byte(`{"events":[{"asset_id":"b4695157-0d1d-4da0-8f9e-5c53149389e4","beam":0,"device_id":"3bafab7b-4400-4bcf-8e6e-09f954699940","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","mfg_company_id":0,"mfg_data":"string","rssi":0,"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","timestamp":0}],"topic":"asset-raw"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.AssetRaw(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestAudits tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestAudits(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookAudits
    errBody := json.Unmarshal([]byte(`{"events":[{"admin_name":"admin_name8","device_id":"00000380-0000-0000-0000-000000000000","id":"0000122a-0000-0000-0000-000000000000","message":"message0","org_id":"00001302-0000-0000-0000-000000000000","site_id":"00000290-0000-0000-0000-000000000000","src_ip":"src_ip6","timestamp":157.68}],"topic":"audits"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.Audits(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestClientInfo tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestClientInfo(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookClientInfo
    errBody := json.Unmarshal([]byte(`{"events":[{"hostname":"service.company.net","ip":"21.0.128.151","mac":"6ebaa47a3fd4","org_id":"0c160b7f-1027-4cd1-923b-744534c4b070","site_id":"6e77a2ea-d579-471c-9056-5ff5b4ed70ed","timestamp":1703003296}],"topic":"client-info"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.ClientInfo(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestClientJoin tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestClientJoin(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookClientJoin
    errBody := json.Unmarshal([]byte(`{"events":[{"ap":"string","ap_name":"string","band":"string","bssid":"string","connect":0,"connect_float":0,"mac":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","rssi":0,"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","site_name":"string","ssid":"string","timestamp":0,"version":0,"wlan_id":"5028e92b-fc59-4056-91d1-ea4b4ca1617a"}],"topic":"client-join"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.ClientJoin(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestClientLatency tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestClientLatency(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookClientLatency
    errBody := json.Unmarshal([]byte(`{"events":[{"avg_auth":0.17170219,"avg_dhcp":0.017828934,"avg_dns":0.024532124,"max_auth":0.18170219,"max_dhcp":0.027828934,"max_dns":0.029532124,"min_auth":0.16050219,"min_dhcp":0.015828934,"min_dns":0.022532124,"org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1696401600}],"topic":"client-latency"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.ClientLatency(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestClientSessions tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestClientSessions(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookClientSessions
    errBody := json.Unmarshal([]byte(`{"events":[{"ap":"string","ap_name":"string","band":"string","bssid":"string","client_family":"string","client_manufacture":"string","client_model":"string","client_os":"string","connect":0,"connect_float":0,"disconnect":0,"disconnect_float":0,"duration":0,"mac":"string","next_ap":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","rssi":0,"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","site_name":"string","ssid":"string","termination_reason":0,"timestamp":0,"version":0,"wlan_id":"5028e92b-fc59-4056-91d1-ea4b4ca1617a"}],"topic":"client-sessions"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.ClientSessions(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestDeviceEvents tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestDeviceEvents(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookDeviceEvents
    errBody := json.Unmarshal([]byte(`{"events":[{"ap":"string","ap_name":"string","audit_id":"78c04fa6-cfb4-46a0-9aa5-3681ba4f3897","device_name":"string","device_type":"ap","ev_type":"notice","mac":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","reason":"string","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","site_name":"string","text":"string","timestamp":0,"type":"string"}],"topic":"device_events"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.DeviceEvents(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestDeviceUpDown tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestDeviceUpDown(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookDeviceUpdowns
    errBody := json.Unmarshal([]byte(`{"events":[{"ap":"string","ap_name":"string","for_site":true,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","site_name":"string","timestamp":0,"type":"string"}],"topic":"device-updowns"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.DeviceUpDown(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestDiscoveredRawRssi tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestDiscoveredRawRssi(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookDiscoveredRawRssi
    errBody := json.Unmarshal([]byte(`{"events":[{"ap_loc":[0],"beam":0,"device_id":"3bafab7b-4400-4bcf-8e6e-09f954699940","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","is_asset":true,"mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","mfg_company_id":"string","mfg_data":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","rssi":0,"service_packets":[{"service_data":"string","service_uuid":"7138cc00-c611-4dec-a05e-5c4b1cae13c0"}],"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","timestamp":0}],"topic":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.DiscoveredRawRssi(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestLocation tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestLocation(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookLocation
    errBody := json.Unmarshal([]byte(`{"events":[{"battery_voltage":0,"eddystone_uid_instance":"string","eddystone_uid_namespace":"string","eddystone_url_url":"string","ibeacon_major":0,"ibeacon_minor":0,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","id":"487f6eca-6276-4993-bfeb-e3cbbbba3f08","mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","mfg_company_id":0,"mfg_data":"string","name":"string","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","timestamp":0,"type":"string","x":0,"y":0}],"topic":"location"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.Location(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestLocationAsset tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestLocationAsset(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookLocationAsset
    errBody := json.Unmarshal([]byte(`{"events":[{"battery_voltage":3370,"eddystone_uid_instance":"5c5b35000001","eddystone_uid_namespace":"2818e3868dec25629ede","eddystone_url_url":"https://www.abc.com","ibeacon_major":13,"ibeacon_minor":138,"ibeacon_uuid":"f3f17139-704a-f03a-2786-0400279e37c3","mac":"7fc2936fd243","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","mfg_company_id":935,"mfg_data":"648520a1020000","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1461220784,"type":"asset","x":13.5,"y":3.2}],"topic":"location_asset"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.LocationAsset(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestLocationCentrak tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestLocationCentrak(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookLocationCentrak
    errBody := json.Unmarshal([]byte(`{"events":[{"mac":"5684dae9ac8b","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1461220784,"type":"wifi","wifi_beacon_extended_info":[{"frame_ctrl":776,"payload":"............","seq_ctrl":772}],"x":13.5,"y":3.2}],"topic":"location_centrak"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.LocationCentrak(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestLocationClient tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestLocationClient(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookLocationClient
    errBody := json.Unmarshal([]byte(`{"events":[{"mac":"5684dae9ac8b","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1461220784,"type":"wifi","wifi_beacon_extended_info":[{"frame_ctrl":776,"payload":"............","seq_ctrl":772}],"x":13.5,"y":3.2}],"topic":"location_client"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.LocationClient(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestLocationSdk tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestLocationSdk(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookLocationSdk
    errBody := json.Unmarshal([]byte(`{"events":[{"id":"de87bf9d-183f-e383-cc68-6ba43947d403","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","name":"optional","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1461220784,"type":"sdk","x":13.5,"y":3.2}],"topic":"location_sdk"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.LocationSdk(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestLocationUnclient tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestLocationUnclient(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookLocationUnclient
    errBody := json.Unmarshal([]byte(`{"events":[{"mac":"5684dae9ac8b","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1461220784,"type":"wifi","wifi_beacon_extended_info":[{"frame_ctrl":776,"payload":"............","seq_ctrl":772}],"x":13.5,"y":3.2}],"topic":"location_unclient"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.LocationUnclient(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestNacAccounting tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestNacAccounting(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookNacAccounting
    errBody := json.Unmarshal([]byte(`{"events":[{"ap":"5c5b355005be","auth_type":"eap-tls","bssid":"5c5b35546bb4","client_ip":"172.16.87.4","client_type":"wireless","mac":"6e795836d5f9","nas_vendor":"juniper-mist","org_id":"625aba64-fe72-4b14-8985-cbf31ec3d78a","rx_pkts":9523,"site_id":"ec9d1e85-af24-43f9-8d65-d620580e8631","ssid":"Test-CMR SSID","timestamp":1547235620.89,"tx_pkts":5233,"type":"NAC_ACCOUNTING_STOP","username":"hi"}],"topic":"nac-accounting"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.NacAccounting(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestNacEvents tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestNacEvents(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookNacEvents
    errBody := json.Unmarshal([]byte(`{"events":[{"ap":"5c5b35513227","auth_type":"802.1X","bssid":"5c5b355fafcc","dryrun_nacrule_id":"32f27e7d-ff26-4a9b-b3d1-ff9bcb264012","dryrun_nacrule_matched":true,"idp_id":"912ef72e-2239-4996-b81e-469e87a27cd6","idp_role":["itsuperusers","vip"],"mac":"ac3eb179e535","nacrule_id":"32f27e7d-ff26-4a9b-b3d1-ff9bcb264c62","nacrule_matched":true,"nas_vendor":"juniper-mist","org_id":"27547ac2-d114-4e04-beb1-f3f1e6e81ec6","random_mac":true,"resp_attrs":["Tunnel-Type=VLAN","Tunnel-Medium-Type=IEEE-802","Tunnel-Private-Group-Id=750","User-Name=anonymous"],"site_id":"441a1214-6928-442a-8e92-e1d34b8ec6a6","ssid":"##mist_nac","timestamp":1691512031.358188,"type":"NAC_CLIENT_PERMIT","username":"user@deaflyz.net","vlan":"750"}],"topic":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.NacEvents(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestOccupancyAlerts tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestOccupancyAlerts(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookOccupancyAlerts
    errBody := json.Unmarshal([]byte(`{"events":[{"alert_events":[{"current_occupancy":0,"map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","occupancy_limit":0,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","timestamp":0,"type":"COMPLIANCE-VIOLATION","zone_id":"4495020a-236f-46e0-9453-e3f9cc6476f4","zone_name":"string"}],"for_site":true,"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","site_name":"string"}],"topic":"occupancy-alerts"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.OccupancyAlerts(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestPing tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestPing(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookPing
    errBody := json.Unmarshal([]byte(`{"events":[{"id":"487f6eca-6276-4993-bfeb-f3cbbbba4f08","name":"string","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","timestamp":0}],"topic":"ping"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.Ping(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestSdkclientScanData tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestSdkclientScanData(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookSdkclientScanData
    errBody := json.Unmarshal([]byte(`{"events":[{"connection_ap":"5c5b352f587e","connection_band":"2.4","connection_bssid":"5c5b352b51b4","connection_channel":11,"connection_rssi":-87,"last_seen":1592333828,"mac":"70ef0071535f","scan_data":[{"ap":"5c5b352f587e","band":"2.4","bssid":"5c5b352b51b4","channel":11,"rssi":-87,"ssid":"mist-wifi","timestamp":1592333828},{"ap":"5c5b352f587e","band":"5","bssid":"5c5b352b51b8","channel":36,"rssi":-75,"ssid":"mist-wifi","timestamp":1592333828}],"site_id":"93986f10-773b-42be-9438-8d3e6d127f1a"}],"topic":"sdkclient_scan_data"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.SdkclientScanData(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestSiteSle tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestSiteSle(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookSiteSle
    errBody := json.Unmarshal([]byte(`{"events":[{"org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","sle":{"ap-availability":0.6,"successful-connect":0.7,"time-to-connect":0.9},"timestamp":1694620800}],"topic":"site_sle"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.SiteSle(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSamplesWebhooksTestZone tests the behavior of the SamplesWebhooks
func TestSamplesWebhooksTestZone(t *testing.T) {
    ctx := context.Background()
    var body models.WebhookZone
    errBody := json.Unmarshal([]byte(`{"events":[{"asset_id":"b4695157-0d1d-4da0-8f9e-5c53149389e4","id":"485f6eca-6276-4993-bfeb-54cbbbba5f08","mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","name":"string","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","timestamp":0,"trigger":"enter","type":"string","zone_id":"4495020a-236f-46e0-9453-e3f9cc6476f4"}],"topic":"zone"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := samplesWebhooks.Zone(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
