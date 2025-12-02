// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package webhooks

import (
	"bytes"
	"net/http"
	"testing"
)

func newWebhooksHandlerRequest(bodyBytes []byte) *http.Request {
	req, _ := http.NewRequest("POST", "http://localhost", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func TestOnAlarms_WebhooksHandler_ReturnsAlarmsEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"aps":["string"],"bssids":["string"],"count":0,"event_id":"a7a26ff2-e851-45b6-9634-d595f45458b7","for_site":true,"id":"489f6eca-6276-4993-bfeb-c3cbbbba1f08","last_seen":0.0,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","ssids":["string"],"timestamp":0,"type":"string","update":true}],"topic":"alarms"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsAlarms(); !ok {
		t.Fatal("expected AlarmsEvent")
	}
}

func TestOnAssetRawRssi_WebhooksHandler_ReturnsAssetRawRssiEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"device_id":"53f10664-3ce8-4c27-b382-0ef66432349f","ibeacon_major":1234,"ibeacon_minor":1234,"ibeacon_uuid":"f3f17139-704a-f03a-2786-0400279e37c3","map_id":"53f10664-3ce8-4c27-b382-0ef66432349f","org_id":"a97c1b22-a4e9-411e-9bfd-d8695a0f9e61","site_id":"441a1214-6928-442a-8e92-e1d34b8ec6a6"}],"topic":"asset-raw-rssi"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsAssetRawRssi(); !ok {
		t.Fatal("expected AssetRawRssiEvent")
	}
}

func TestOnAudits_WebhooksHandler_ReturnsAuditsEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"admin_name":"admin_name8","device_id":"00000380-0000-0000-0000-000000000000","id":"0000122a-0000-0000-0000-000000000000","message":"message0","org_id":"00001302-0000-0000-0000-000000000000","site_id":"00000290-0000-0000-0000-000000000000","src_ip":"src_ip6","timestamp":157.68}],"topic":"audits"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsAudits(); !ok {
		t.Fatal("expected AuditsEvent")
	}
}

func TestOnClientInfo_WebhooksHandler_ReturnsClientInfoEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"hostname":"service.company.net","ip":"21.0.128.151","mac":"6ebaa47a3fd4","org_id":"0c160b7f-1027-4cd1-923b-744534c4b070","site_id":"6e77a2ea-d579-471c-9056-5ff5b4ed70ed","timestamp":1703003296.0}],"topic":"client-info"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsClientInfo(); !ok {
		t.Fatal("expected ClientInfoEvent")
	}
}

func TestOnClientJoin_WebhooksHandler_ReturnsClientJoinEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"ap":"string","ap_name":"string","band":"string","bssid":"string","connect":0,"connect_float":0.0,"mac":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","rssi":0.0,"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","site_name":"string","ssid":"string","timestamp":0,"version":0.0,"wlan_id":"5028e92b-fc59-4056-91d1-ea4b4ca1617a"}],"topic":"client-join"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsClientJoin(); !ok {
		t.Fatal("expected ClientJoinEvent")
	}
}

func TestOnClientLatency_WebhooksHandler_ReturnsClientLatencyEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"avg_auth":0.17170219,"avg_dhcp":0.017828934,"avg_dns":0.024532124,"max_auth":0.18170219,"max_dhcp":0.027828934,"max_dns":0.029532124,"min_auth":0.16050219,"min_dhcp":0.015828934,"min_dns":0.022532124,"org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1696401600.0}],"topic":"client-latency"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsClientLatency(); !ok {
		t.Fatal("expected ClientLatencyEvent")
	}
}

func TestOnClientSessions_WebhooksHandler_ReturnsClientSessionsEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"ap":"string","ap_name":"string","band":"string","bssid":"string","client_family":"string","client_manufacture":"string","client_model":"string","client_os":"string","connect":0,"connect_float":0.0,"disconnect":0,"disconnect_float":0.0,"duration":0,"mac":"string","next_ap":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","rssi":0.0,"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","site_name":"string","ssid":"string","termination_reason":0,"timestamp":0,"version":0.0,"wlan_id":"5028e92b-fc59-4056-91d1-ea4b4ca1617a"}],"topic":"client-sessions"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsClientSessions(); !ok {
		t.Fatal("expected ClientSessionsEvent")
	}
}

func TestOnDeviceEvents_WebhooksHandler_ReturnsDeviceEventsEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"ap":"string","ap_name":"string","audit_id":"78c04fa6-cfb4-46a0-9aa5-3681ba4f3897","device_name":"string","device_type":"ap","ev_type":"notice","mac":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","reason":"string","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","site_name":"string","text":"string","timestamp":0,"type":"string"}],"topic":"device_events"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsDeviceEvents(); !ok {
		t.Fatal("expected DeviceEventsEvent")
	}
}

func TestOnDeviceUpdowns_WebhooksHandler_ReturnsDeviceUpdownsEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"ap":"string","ap_name":"string","for_site":true,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","site_name":"string","timestamp":0,"type":"string"}],"topic":"device-updowns"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsDeviceUpdowns(); !ok {
		t.Fatal("expected DeviceUpdownsEvent")
	}
}

func TestOnDiscoveredRawRssi_WebhooksHandler_ReturnsDiscoveredRawRssiEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"ap_loc":[0.0],"beam":0,"device_id":"3bafab7b-4400-4bcf-8e6e-09f954699940","ibeacon_major":1,"ibeacon_minor":1,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","is_asset":true,"mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","mfg_company_id":"string","mfg_data":"string","org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","rssi":0.0,"service_packets":[{"service_data":"string","service_uuid":"7138cc00-c611-4dec-a05e-5c4b1cae13c0"}],"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","timestamp":0.0}],"topic":"string"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsDiscoveredRawRssi(); !ok {
		t.Fatal("expected DiscoveredRawRssiEvent")
	}
}

func TestOnGuestAuthorizations_WebhooksHandler_ReturnsGuestAuthorizationsEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"ap":"5c5b350e55c8","auth_method":"passphrase","authorized_expiring_time":1677076639,"authorized_time":1677076519,"carrier":"docomo","client":"ac2316eca70a","company":"MIST","email":"abcd@abcd.com","field1":"field1 value","field2":"field2 value","field3":"field3 value","field4":"field4 value","mobile":"+0123456789","name":"Dr Strange","org_id":"1688605f-916a-47a1-8c68-f19618300a08","site_id":"ec3b5624-73f1-4ed3-b3fd-5ba3ee40368a","sms_gateway":"Telstra","sponsor_email":"sponsor@gmail.com","ssid":"Portal Auth","wlan_id":"7681be9a-044a-4622-90cf-3accde5ad853"}],"topic":"guest-authorizations"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsGuestAuthorizations(); !ok {
		t.Fatal("expected GuestAuthorizationsEvent")
	}
}

func TestOnLocation_WebhooksHandler_ReturnsLocationEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"battery_voltage":0,"eddystone_uid_instance":"string","eddystone_uid_namespace":"string","eddystone_url_url":"string","ibeacon_major":1,"ibeacon_minor":1,"ibeacon_uuid":"1f89bc00-d0af-481b-82fe-a6629259a39f","id":"487f6eca-6276-4993-bfeb-e3cbbbba3f08","mac":"string","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","mfg_company_id":0,"mfg_data":"string","name":"string","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","timestamp":0,"type":"string","x":0.0,"y":0.0}],"topic":"location"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsLocation(); !ok {
		t.Fatal("expected LocationEvent")
	}
}

func TestOnLocationAsset_WebhooksHandler_ReturnsLocationAssetEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"battery_voltage":3370,"eddystone_uid_instance":"5c5b35000001","eddystone_uid_namespace":"2818e3868dec25629ede","eddystone_url_url":"https://www.abc.com","ibeacon_major":13,"ibeacon_minor":138,"ibeacon_uuid":"f3f17139-704a-f03a-2786-0400279e37c3","mac":"7fc2936fd243","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","mfg_company_id":935,"mfg_data":"648520a1020000","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1461220784.0,"type":"asset","x":13.5,"y":3.2}],"topic":"location_asset"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsLocationAsset(); !ok {
		t.Fatal("expected LocationAssetEvent")
	}
}

func TestOnLocationCentrak_WebhooksHandler_ReturnsLocationCentrakEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"mac":"5684dae9ac8b","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1461220784,"type":"wifi","wifi_beacon_extended_info":[{"frame_ctrl":776,"payload":"............","seq_ctrl":772}],"x":13.5,"y":3.2}],"topic":"location_centrak"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsLocationCentrak(); !ok {
		t.Fatal("expected LocationCentrakEvent")
	}
}

func TestOnLocationClient_WebhooksHandler_ReturnsLocationClientEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"mac":"5684dae9ac8b","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1461220784.0,"type":"wifi","wifi_beacon_extended_info":[{"frame_ctrl":776,"payload":"............","seq_ctrl":772}],"x":13.5,"y":3.2}],"topic":"location_client"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsLocationClient(); !ok {
		t.Fatal("expected LocationClientEvent")
	}
}

func TestOnLocationSdk_WebhooksHandler_ReturnsLocationSdkEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"id":"de87bf9d-183f-e383-cc68-6ba43947d403","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","name":"optional","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1461220784.0,"type":"sdk","x":13.5,"y":3.2}],"topic":"location_sdk"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsLocationSdk(); !ok {
		t.Fatal("expected LocationSdkEvent")
	}
}

func TestOnLocationUnclient_WebhooksHandler_ReturnsLocationUnclientEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"mac":"5684dae9ac8b","map_id":"845a23bf-bed9-e43c-4c86-6fa474be7ae5","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","timestamp":1461220784.0,"type":"wifi","wifi_beacon_extended_info":[{"frame_ctrl":776,"payload":"............","seq_ctrl":772}],"x":13.5,"y":3.2}],"topic":"location_unclient"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsLocationUnclient(); !ok {
		t.Fatal("expected LocationUnclientEvent")
	}
}

func TestOnMxedgeEvents_WebhooksHandler_ReturnsMxedgeEventsEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"audit_id":"8912e5cb-8ddd-41f7-be5f-476a7abbf658","component":null,"mxedge_id":"00000000-0000-0000-1000-020000230522","mxedge_name":"demo123","org_id":"203d3d02-dbc0-4c1b-9f41-76896a3330f4","timestamp":128.82,"type":"ME_CONFIG_CHANGED_BY_USER"},{"audit_id":"48efa5bf-d290-4e93-80ca-4dbf72f4187a","component":null,"mxedge_id":"00000000-0000-0000-1000-020000a5fca1","mxedge_name":"test123","org_id":"203d3d02-dbc0-4c1b-9f41-76896a3330f4","timestamp":128.82,"type":"ME_CONFIG_CHANGED_BY_USER"}],"topic":"mxedge-events"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsMxedgeEvents(); !ok {
		t.Fatal("expected MxedgeEventsEvent")
	}
}

func TestOnNacAccounting_WebhooksHandler_ReturnsNacAccountingEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"ap":"5c5b355005be","auth_type":"eap-tls","bssid":"5c5b35546bb4","client_ip":"172.16.87.4","client_type":"wireless","mac":"6e795836d5f9","nas_vendor":"juniper-mist","org_id":"625aba64-fe72-4b14-8985-cbf31ec3d78a","rx_pkts":9523,"site_id":"ec9d1e85-af24-43f9-8d65-d620580e8631","ssid":"Test-CMR SSID","timestamp":1547235620.89,"tx_pkts":5233,"type":"NAC_ACCOUNTING_STOP","username":"hi"}],"topic":"nac-accounting"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsNacAccounting(); !ok {
		t.Fatal("expected NacAccountingEvent")
	}
}

func TestOnNacEvents_WebhooksHandler_ReturnsNacEventsEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"ap":"5c5b35513227","auth_type":"eap-teap","bssid":"5c5b355fafcc","dryrun_nacrule_id":"32f27e7d-ff26-4a9b-b3d1-ff9bcb264012","dryrun_nacrule_matched":true,"idp_id":"912ef72e-2239-4996-b81e-469e87a27cd6","idp_role":["itsuperusers","vip"],"mac":"ac3eb179e535","nacrule_id":"32f27e7d-ff26-4a9b-b3d1-ff9bcb264c62","nacrule_matched":true,"nas_vendor":"juniper-mist","org_id":"27547ac2-d114-4e04-beb1-f3f1e6e81ec6","random_mac":"true","resp_attrs":["Tunnel-Type=VLAN","Tunnel-Medium-Type=IEEE-802","Tunnel-Private-Group-Id=750","User-Name=anonymous"],"site_id":"441a1214-6928-442a-8e92-e1d34b8ec6a6","ssid":"##mist_nac","timestamp":1691512031.358188,"type":"NAC_CLIENT_PERMIT","username":"user@deaflyz.net","vlan":"750"}],"topic":"string"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsNacEvents(); !ok {
		t.Fatal("expected NacEventsEvent")
	}
}

func TestOnOccupancyAlerts_WebhooksHandler_ReturnsOccupancyAlertsEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"alert_events":[{"current_occupancy":0,"map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","occupancy_limit":0,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","timestamp":0,"type":"COMPLIANCE-VIOLATION","zone_id":"4495020a-236f-46e0-9453-e3f9cc6476f4","zone_name":"string"}],"for_site":true,"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","site_name":"string"}],"topic":"occupancy-alerts"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsOccupancyAlerts(); !ok {
		t.Fatal("expected OccupancyAlertsEvent")
	}
}

func TestOnPing_WebhooksHandler_ReturnsPingEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"id":"487f6eca-6276-4993-bfeb-f3cbbbba4f08","name":"string","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","timestamp":0}],"topic":"ping"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsPing(); !ok {
		t.Fatal("expected PingEvent")
	}
}

func TestOnSdkclientScanData_WebhooksHandler_ReturnsSdkclientScanDataEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"connection_ap":"5c5b352f587e","connection_band":"2.4","connection_bssid":"5c5b352b51b4","connection_channel":11,"connection_rssi":-87.0,"last_seen":1592333828.0,"mac":"70ef0071535f","scan_data":[{"ap":"5c5b352f587e","band":"2.4","bssid":"5c5b352b51b4","channel":11,"rssi":-87.0,"ssid":"mist-wifi","timestamp":1592333828},{"ap":"5c5b352f587e","band":"5","bssid":"5c5b352b51b8","channel":36,"rssi":-75.0,"ssid":"mist-wifi","timestamp":1592333828}],"site_id":"93986f10-773b-42be-9438-8d3e6d127f1a"}],"topic":"sdkclient_scan_data"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsSdkclientScanData(); !ok {
		t.Fatal("expected SdkclientScanDataEvent")
	}
}

func TestOnSiteSle_WebhooksHandler_ReturnsSiteSleEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"org_id":"2818e386-8dec-2562-9ede-5b8a0fbbdc71","site_id":"4ac1dcf4-9d8b-7211-65c4-057819f0862b","sle":{"ap-availability":0.6,"successful-connect":0.7,"time-to-connect":0.9},"timestamp":1694620800.0}],"topic":"site_sle"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsSiteSle(); !ok {
		t.Fatal("expected SiteSleEvent")
	}
}

func TestOnWifiConnRaw_WebhooksHandler_ReturnsWifiConnRawEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"map_id":"53f10664-3ce8-4c27-b382-0ef66432349f","org_id":"a97c1b22-a4e9-411e-9bfd-d8695a0f9e61","site_id":"441a1214-6928-442a-8e92-e1d34b8ec6a6"}],"topic":"wifi-conn-raw"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsWifiConnRaw(); !ok {
		t.Fatal("expected WifiConnRawEvent")
	}
}

func TestOnWifiUnconnRaw_WebhooksHandler_ReturnsWifiUnconnRawEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"map_id":"53f10664-3ce8-4c27-b382-0ef66432349f","org_id":"a97c1b22-a4e9-411e-9bfd-d8695a0f9e61","site_id":"441a1214-6928-442a-8e92-e1d34b8ec6a6"}],"topic":"wifi-unconn-raw"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsWifiUnconnRaw(); !ok {
		t.Fatal("expected WifiUnconnRawEvent")
	}
}

func TestOnZone_WebhooksHandler_ReturnsZoneEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`{"events":[{"id":"487f6eca-6276-4993-bfeb-d3cbbbba2f08","map_id":"09d2b626-2e4e-45ef-a3c4-e6aeb6c83db1","name":"string","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","timestamp":0,"type":"string","trigger":"enter","zone_id":"00000f60-0000-0000-0000-000000000000"}],"topic":"zone"}`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if _, ok := result.AsZone(); !ok {
		t.Fatal("expected ZoneEvent")
	}
}

func TestOnUnknown_WebhooksHandler_ReturnsUnknownEvent(t *testing.T) {
	// Arrange
	bodyBytes := []byte(`null`)
	request := newWebhooksHandlerRequest(bodyBytes)
	handler := NewWebhooksHandler()

	// Act
	result := handler.ParseEvent(request)

	// Assert
	if ok := result.AsUnknown(); !ok {
		t.Fatal("expected UnknownEvent")
	}
}
