package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesDevicesTestListSiteDevices tests the behavior of the SitesDevices
func TestSitesDevicesTestListSiteDevices(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mType := models.DeviceTypeWithAllEnum("ap")
    
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesDevices.ListSiteDevices(ctx, siteId, &mType, nil, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesDevicesTestCountSiteDeviceConfigHistory tests the behavior of the SitesDevices
func TestSitesDevicesTestCountSiteDeviceConfigHistory(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesDevices.CountSiteDeviceConfigHistory(ctx, siteId, nil, nil, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesTestSearchSiteDeviceConfigHistory tests the behavior of the SitesDevices
func TestSitesDevicesTestSearchSiteDeviceConfigHistory(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mType := models.DeviceTypeEnum("ap")
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesDevices.SearchSiteDeviceConfigHistory(ctx, siteId, &mType, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1531862583,"limit":10,"results":[{"channel_24":11,"channel_5":100,"radio_macs":["5c5b352e000a","5c5b352e000b","5c5b352e000c"],"radios":[{"band":"24","channel":11},{"band":"5","channel":100}],"secpolicy_violated":false,"ssids":["test24","test5"],"ssids_24":["test24"],"ssids_5":["test5"],"timestamp":1531855856.643369,"version":"apfw-0.2.14754-cersei-75c8","wlans":[{"auth":"psk","bands":["24"],"id":"be22bba7-8e22-e1cf-5185-b880816fe2cf","ssid":"test24","vlan_ids":["1"]},{"auth":"psk","bands":["5"],"id":"f8c18724-4118-3487-811a-f98964988604","ssid":"test5","vlan_ids":["1"]}]}],"start":1531776183,"total":1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesTestCountSiteDevices tests the behavior of the SitesDevices
func TestSitesDevicesTestCountSiteDevices(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteDevicesCountDistinctEnum("model")
    
    
    
    
    
    
    
    
    
    
    
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesDevices.CountSiteDevices(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesTestCountSiteDeviceEvents tests the behavior of the SitesDevices
func TestSitesDevicesTestCountSiteDeviceEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteDeviceEventsCountDistinctEnum("model")
    
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesDevices.CountSiteDeviceEvents(ctx, siteId, &distinct, nil, nil, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesTestSearchSiteDeviceEvents tests the behavior of the SitesDevices
func TestSitesDevicesTestSearchSiteDeviceEvents(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    
    
    
    lastBy := "port_id"
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesDevices.SearchSiteDeviceEvents(ctx, siteId, nil, nil, nil, nil, nil, &lastBy, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1531862583,"limit":2,"next":"/api/v1/sites/8aaba0aa-09cc-44bd-9709-33b98040550c/devices/events/search?ap=5c5b350e0001&end=1531855849.000&limit=2&start=1531776183.0","results":[{"last_reboot_time":1531854327,"text":"Success","timestamp":1531855849.226722,"type":"AP_CONNECT_STATUS","type_code":2002},{"timestamp":1531854326,"type":"AP_CONFIGURED"}],"start":1531776183,"total":14}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesTestExportSiteDevices tests the behavior of the SitesDevices
func TestSitesDevicesTestExportSiteDevices(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesDevices.ExportSiteDevices(ctx, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesDevicesTestCountSiteDeviceLastConfig tests the behavior of the SitesDevices
func TestSitesDevicesTestCountSiteDeviceLastConfig(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteDeviceLastConfigCountDistinctEnum("mac")
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesDevices.CountSiteDeviceLastConfig(ctx, siteId, &distinct, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesTestSearchSiteDeviceLastConfigs tests the behavior of the SitesDevices
func TestSitesDevicesTestSearchSiteDeviceLastConfigs(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mType := models.DeviceTypeEnum("ap")
    
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesDevices.SearchSiteDeviceLastConfigs(ctx, siteId, &mType, nil, nil, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1531862583,"limit":10,"results":[{"channel_24":11,"channel_5":100,"radio_macs":["5c5b352e000a","5c5b352e000b","5c5b352e000c"],"radios":[{"band":"24","channel":11},{"band":"5","channel":100}],"secpolicy_violated":false,"ssids":["test24","test5"],"ssids_24":["test24"],"ssids_5":["test5"],"timestamp":1531855856.643369,"version":"apfw-0.2.14754-cersei-75c8","wlans":[{"auth":"psk","bands":["24"],"id":"be22bba7-8e22-e1cf-5185-b880816fe2cf","ssid":"test24","vlan_ids":["1"]},{"auth":"psk","bands":["5"],"id":"f8c18724-4118-3487-811a-f98964988604","ssid":"test5","vlan_ids":["1"]}]}],"start":1531776183,"total":1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesTestRestartSiteMultipleDevices tests the behavior of the SitesDevices
func TestSitesDevicesTestRestartSiteMultipleDevices(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.UtilsDevicesRestartMulti
    errBody := json.Unmarshal([]byte(`{"device_ids":["00000000-0000-0000-1000-5c5b35584a6f","00000000-0000-0000-1000-5c5b350ea3b3"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := sitesDevices.RestartSiteMultipleDevices(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesDevicesTestGetSiteDevice tests the behavior of the SitesDevices
func TestSitesDevicesTestGetSiteDevice(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesDevices.GetSiteDevice(ctx, siteId, deviceId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"aeroscout":{"enabled":false,"host":"aero.pvt.net","locate_connected":true},"ble_config":{"beacon_enabled":false,"beacon_rate":3,"beacon_rate_mode":"custom","beam_disabled":[1,3,6],"custom_ble_packet_enabled":false,"custom_ble_packet_frame":"0x........","custom_ble_packet_freq_msec":300,"eddystone_uid_adv_power":-65,"eddystone_uid_beams":"2-4,7","eddystone_uid_enabled":false,"eddystone_uid_freq_msec":200,"eddystone_uid_instance":"5c5b35000001","eddystone_uid_namespace":"2818e3868dec25629ede","eddystone_url_adv_power":-65,"eddystone_url_beams":"2-4,7","eddystone_url_enabled":true,"eddystone_url_freq_msec":1000,"eddystone_url_url":"https://www.abc.com","ibeacon_adv_power":-65,"ibeacon_beams":"2-4,7","ibeacon_enabled":false,"ibeacon_freq_msec":0,"ibeacon_major":13,"ibeacon_minor":138,"ibeacon_uuid":"f3f17139-704a-f03a-2786-0400279e37c3","power":6,"power_mode":"custom"},"centrak":{"enabled":false},"client_bridge":{"auth":{"psk":"foryoureyesonly","type":"psk"},"enabled":false,"ssid":"Uplink-SSID"},"created_time":0,"deviceprofile_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9","disable_eth1":false,"disable_eth2":false,"disable_eth3":false,"disable_module":false,"esl_config":{"cacert":"string","channel":3,"enabled":false,"host":"1.1.1.1","port":0,"type":"imagotag","verify_cert":true,"vlan_id":1},"for_site":true,"height":2.75,"id":"497f6eca-6276-4993-bfeb-53cbbbba6008","image1_url":"string","image2_url":"string","image3_url":"string","iot_config":{"A1":{"enabled":false,"name":"motion","output":true,"pullup":"internal","value":0},"A2":{"enabled":false,"name":"motion","output":true,"pullup":"internal","value":0},"A3":{"enabled":false,"name":"motion","output":true,"pullup":"internal","value":0},"A4":{"enabled":false,"name":"motion","output":true,"pullup":"internal","value":0},"DI1":{"enabled":false,"name":"string","pullup":"internal"},"DI2":{"enabled":false,"name":"string","pullup":"internal"},"DO":{"enabled":false,"name":"motion","output":true,"pullup":"internal","value":0}},"ip_config":{"dns":["8.8.8.8","4.4.4.4"],"dns_suffix":[".mist.local",".mist.com"],"gateway":"10.2.1.254","gateway6":"2607:f8b0:4005:808::1","ip":"10.2.1.1","ip6":"2607:f8b0:4005:808::2004","mtu":0,"netmask":"255.255.255.0","netmask6":"/32","type":"static","type6":"static","vlan_id":1},"led":{"brightness":255,"enabled":true},"locked":true,"map_id":"63eda950-c6da-11e4-a628-60f81dd250cc","mesh":{"enabled":false,"group":1,"role":"base"},"modified_time":0,"name":"conference room","notes":"slightly off center","ntp_servers":["string"],"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","orientation":45,"poe_passthrough":false,"pwr_config":{"base":2000,"prefer_usb_over_wifi":false},"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","type":"ap","uplink_port_config":{"dot1x":false,"keep_wlans_up_if_down":false},"usb_config":{"cacert":"string","channel":3,"enabled":true,"host":"1.1.1.1","port":0,"type":"imagotag","verify_cert":true,"vlan_id":1},"vars":{"RADIUS_IP1":"172.31.2.5","RADIUS_SECRET":"11s64632d"},"x":53.5,"y":173.1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesDevicesTestUpdateSiteDevice tests the behavior of the SitesDevices
func TestSitesDevicesTestUpdateSiteDevice(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    deviceId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := sitesDevices.UpdateSiteDevice(ctx, siteId, deviceId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"aeroscout":{"enabled":false,"host":"aero.pvt.net","locate_connected":true},"ble_config":{"beacon_enabled":false,"beacon_rate":3,"beacon_rate_mode":"custom","beam_disabled":[1,3,6],"custom_ble_packet_enabled":false,"custom_ble_packet_frame":"0x........","custom_ble_packet_freq_msec":300,"eddystone_uid_adv_power":-65,"eddystone_uid_beams":"2-4,7","eddystone_uid_enabled":false,"eddystone_uid_freq_msec":200,"eddystone_uid_instance":"5c5b35000001","eddystone_uid_namespace":"2818e3868dec25629ede","eddystone_url_adv_power":-65,"eddystone_url_beams":"2-4,7","eddystone_url_enabled":true,"eddystone_url_freq_msec":1000,"eddystone_url_url":"https://www.abc.com","ibeacon_adv_power":-65,"ibeacon_beams":"2-4,7","ibeacon_enabled":false,"ibeacon_freq_msec":0,"ibeacon_major":13,"ibeacon_minor":138,"ibeacon_uuid":"f3f17139-704a-f03a-2786-0400279e37c3","power":6,"power_mode":"custom"},"centrak":{"enabled":false},"client_bridge":{"auth":{"psk":"foryoureyesonly","type":"psk"},"enabled":false,"ssid":"Uplink-SSID"},"created_time":0,"deviceprofile_id":"6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9","disable_eth1":false,"disable_eth2":false,"disable_eth3":false,"disable_module":false,"esl_config":{"cacert":"string","channel":3,"enabled":false,"host":"1.1.1.1","port":0,"type":"imagotag","verify_cert":true,"vlan_id":1},"for_site":true,"height":2.75,"id":"497f6eca-6276-4993-bfeb-53cbbbba6008","image1_url":"string","image2_url":"string","image3_url":"string","iot_config":{"A1":{"enabled":false,"name":"motion","output":true,"pullup":"internal","value":0},"A2":{"enabled":false,"name":"motion","output":true,"pullup":"internal","value":0},"A3":{"enabled":false,"name":"motion","output":true,"pullup":"internal","value":0},"A4":{"enabled":false,"name":"motion","output":true,"pullup":"internal","value":0},"DI1":{"enabled":false,"name":"string","pullup":"internal"},"DI2":{"enabled":false,"name":"string","pullup":"internal"},"DO":{"enabled":false,"name":"motion","output":true,"pullup":"internal","value":0}},"ip_config":{"dns":["8.8.8.8","4.4.4.4"],"dns_suffix":[".mist.local",".mist.com"],"gateway":"10.2.1.254","gateway6":"2607:f8b0:4005:808::1","ip":"10.2.1.1","ip6":"2607:f8b0:4005:808::2004","mtu":0,"netmask":"255.255.255.0","netmask6":"/32","type":"static","type6":"static","vlan_id":1},"led":{"brightness":255,"enabled":true},"locked":true,"map_id":"63eda950-c6da-11e4-a628-60f81dd250cc","mesh":{"enabled":false,"group":1,"role":"base"},"modified_time":0,"name":"conference room","notes":"slightly off center","ntp_servers":["string"],"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","orientation":45,"poe_passthrough":false,"pwr_config":{"base":2000,"prefer_usb_over_wifi":false},"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","type":"ap","uplink_port_config":{"dot1x":false,"keep_wlans_up_if_down":false},"usb_config":{"cacert":"string","channel":3,"enabled":true,"host":"1.1.1.1","port":0,"type":"imagotag","verify_cert":true,"vlan_id":1},"vars":{"RADIUS_IP1":"172.31.2.5","RADIUS_SECRET":"11s64632d"},"x":53.5,"y":173.1}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
