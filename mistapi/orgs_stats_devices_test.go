// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsStatsDevicesTestListOrgDevicesStats tests the behavior of the OrgsStatsDevices
func TestOrgsStatsDevicesTestListOrgDevicesStats(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mType := models.DeviceTypeWithAllEnum("ap")
    status := models.DeviceStatusEnum("all")
    
    
    
    
    fields := "field1,field2"
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsStatsDevices.ListOrgDevicesStats(ctx, orgId, &mType, &status, nil, nil, nil, nil, &fields, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"auto_upgrade_stat":{"lastcheck":1720594762},"ble_stat":{"beacon_enabled":true,"beacon_rate":4,"eddystone_uid_enabled":false,"eddystone_uid_freq_msec":1000,"eddystone_uid_instance":"5c5b35d0077b","eddystone_uid_namespace":"9777c1a06ef611e68bbf","eddystone_url_enabled":false,"eddystone_url_freq_msec":1000,"eddystone_url_url":"","ibeacon_enabled":false,"ibeacon_freq_msec":1000,"ibeacon_major":894,"ibeacon_minor":9328,"ibeacon_uuid":"af010e2b-f829-4975-b49e-2e896ed1d627","major":894,"minors":[9328,9329,9330,9331,9332,9333,9334,9335,-1],"power":8,"rx_bytes":158500843,"rx_pkts":3549163,"tx_bytes":509640,"tx_pkts":85411,"tx_resets":0,"uuid":"af010e2b-f829-4975-b49e-2e896ed1d627"},"config_reverted":false,"cpu_system":21921854,"cpu_user":7496631,"cpu_util":5,"created_time":1718228350,"env_stat":{"accel_x":-0.092,"accel_y":0.004,"accel_z":-1.02,"ambient_temp":43,"attitude":0,"cpu_temp":53,"humidity":9,"magne_x":0,"magne_y":0,"magne_z":0,"pressure":968,"vcore_voltage":0},"expiring_certs":{"375895068012951219417723061830777179121567314483":1535534392,"483412046338348876283259475264040058183366677871":1534534392},"ext_ip":"66.129.234.28","hw_rev":"C02","id":"00000000-0000-0000-1000-5c5b35d0077b","inactive_wired_vlans":[],"ip":"192.168.95.3","ip_stat":{"dhcp_server":"192.168.95.1","dns":["8.8.8.8"],"gateway":"192.168.95.1","ip":"192.168.95.3","ip6":"fe80:0:0:0:5e5b:35ff:fed0:77b","ips":{"vlan1":"192.168.95.3/24,fe80:0:0:0:5e5b:35ff:fed0:77b/64"},"netmask":"255.255.255.0","netmask6":"/64"},"last_seen":1720595866,"last_trouble":{"code":"07","timestamp":1720039666},"lldp_stat":{"chassis_id":"d0:07:ca:f5:21:00","lldp_med_supported":false,"mgmt_addr":"100.123.105.1","mgmt_addrs":["10.1.5.2","2002:A01:203:0:0:0:0:0"],"port_desc":"ge-0/0/4","port_id":"ge-0/0/4","power_allocated":0,"power_draw":0,"power_request_count":0,"power_requested":0,"system_desc":"Juniper Networks, Inc. ex4300-48t internet router, kernel JUNOS 20.4R3-S7.2, Build date: 2023-04-21 19:47:18 UTC Copyright (c) 1996-2023 Juniper Networks, Inc.","system_name":"Phoenix-Switch"},"lldp_stats":{"eth0":{"chassis_id":"63:68:61:73:73:69","lldp_med_supported":false,"mgmt_addr":"10.1.5.2","mgmt_addrs":["10.1.5.2","2002:A01:203:0:0:0:0:0"],"port_desc":"2/26","port_id":"port1","power_allocated":15500,"power_draw":15000,"power_request_count":3,"power_requested":25500,"system_desc":"HP J9729A 2920-48G-POE+ Switch","system_name":"TC2-OWL-Stack-01"},"eth1":{"chassis_id":"chassis2","lldp_med_supported":true,"mgmt_addr":"10.1.5.3","mgmt_addrs":["10.1.5.3"],"port_desc":"1/12","port_id":"port2","power_allocated":15400,"power_draw":14000,"power_request_count":2,"power_requested":15400,"system_desc":"HP J9728A 2920-24G-POE+ Switch","system_name":"TC2-OWL-Stack-02"}},"mac":"5c5b35d0077b","mem_total_kb":505468,"mem_used_kb":202096,"model":"AP43","modified_time":1718530662,"mount":"faceup","name":"Phoenix","notes":"","num_clients":1,"org_id":"af010e2b-f829-4975-b49e-2e896ed1d627","port_stat":{"eth0":{"full_duplex":true,"rx_bytes":1284143195,"rx_errors":0,"rx_peak_bps":17585,"rx_pkts":5199816,"speed":1000,"tx_bytes":1283744961,"tx_peak_bps":26484,"tx_pkts":3990463,"up":true},"eth1":{"full_duplex":false,"rx_bytes":0,"rx_errors":0,"rx_peak_bps":0,"rx_pkts":0,"speed":0,"tx_bytes":0,"tx_peak_bps":0,"tx_pkts":0,"up":false}},"power_budget":8400,"power_constrained":false,"power_src":"PoE 802.3at","power_srcs":["PoE 802.3at","PoE 802.3af"],"radio_config":{},"radio_stat":{"band_24":{"bandwidth":20,"channel":11,"mac":"5c5b35dea810","noise_floor":-80,"num_clients":0,"power":17,"rx_bytes":12948211,"rx_pkts":65292,"tx_bytes":19071943,"tx_pkts":76926,"usage":"24","util_all":24,"util_non_wifi":2,"util_rx_in_bss":0,"util_rx_other_bss":17,"util_tx":4,"util_undecodable_wifi":0,"util_unknown_wifi":1},"band_5":{"bandwidth":40,"channel":36,"mac":"5c5b35dea7f0","noise_floor":-90,"num_clients":1,"power":17,"rx_bytes":578362619,"rx_pkts":2687577,"tx_bytes":1199571353,"tx_pkts":2479302,"usage":"5","util_all":13,"util_non_wifi":0,"util_rx_in_bss":0,"util_rx_other_bss":10,"util_tx":1,"util_undecodable_wifi":0,"util_unknown_wifi":1}},"rx_bps":9276,"rx_bytes":591310830,"rx_pkts":2752869,"serial":"A0703200709E6","site_id":"46fc665e-9706-4296-8fe2-78f42f2e67e4","status":"connected","switch_redundancy":{"num_redundant_aps":1},"tx_bps":8067,"tx_bytes":1218643296,"tx_pkts":2556228,"type":"ap","uptime":1593120,"version":"0.14.29313"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsStatsDevicesTestListOrgDevicesStats1 tests the behavior of the OrgsStatsDevices
func TestOrgsStatsDevicesTestListOrgDevicesStats1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mType := models.DeviceTypeWithAllEnum("ap")
    status := models.DeviceStatusEnum("all")
    
    
    
    
    fields := "field1,field2"
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsStatsDevices.ListOrgDevicesStats(ctx, orgId, &mType, &status, nil, nil, nil, nil, &fields, nil, nil, &duration, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"auto_upgrade_stat":{"lastcheck":1720594762},"ble_stat":{"beacon_enabled":true,"beacon_rate":4,"eddystone_uid_enabled":false,"eddystone_uid_freq_msec":1000,"eddystone_uid_instance":"5c5b35d0077b","eddystone_uid_namespace":"9777c1a06ef611e68bbf","eddystone_url_enabled":false,"eddystone_url_freq_msec":1000,"eddystone_url_url":"","ibeacon_enabled":false,"ibeacon_freq_msec":1000,"ibeacon_major":894,"ibeacon_minor":9328,"ibeacon_uuid":"af010e2b-f829-4975-b49e-2e896ed1d627","major":894,"minors":[9328,9329,9330,9331,9332,9333,9334,9335,-1],"power":8,"rx_bytes":158500843,"rx_pkts":3549163,"tx_bytes":509640,"tx_pkts":85411,"tx_resets":0,"uuid":"af010e2b-f829-4975-b49e-2e896ed1d627"},"config_reverted":false,"cpu_system":21921854,"cpu_user":7496631,"cpu_util":5,"created_time":1718228350,"env_stat":{"accel_x":-0.092,"accel_y":0.004,"accel_z":-1.02,"ambient_temp":43,"attitude":0,"cpu_temp":53,"humidity":9,"magne_x":0,"magne_y":0,"magne_z":0,"pressure":968,"vcore_voltage":0},"expiring_certs":{"375895068012951219417723061830777179121567314483":1535534392,"483412046338348876283259475264040058183366677871":1534534392},"ext_ip":"66.129.234.28","hw_rev":"C02","id":"00000000-0000-0000-1000-5c5b35d0077b","inactive_wired_vlans":[],"ip":"192.168.95.3","ip_stat":{"dhcp_server":"192.168.95.1","dns":["8.8.8.8"],"gateway":"192.168.95.1","ip":"192.168.95.3","ip6":"fe80:0:0:0:5e5b:35ff:fed0:77b","ips":{"vlan1":"192.168.95.3/24,fe80:0:0:0:5e5b:35ff:fed0:77b/64"},"netmask":"255.255.255.0","netmask6":"/64"},"last_seen":1720595866,"last_trouble":{"code":"07","timestamp":1720039666},"lldp_stat":{"chassis_id":"d0:07:ca:f5:21:00","lldp_med_supported":false,"mgmt_addr":"100.123.105.1","mgmt_addrs":["10.1.5.2","2002:A01:203:0:0:0:0:0"],"port_desc":"ge-0/0/4","port_id":"ge-0/0/4","power_allocated":0,"power_draw":0,"power_request_count":0,"power_requested":0,"system_desc":"Juniper Networks, Inc. ex4300-48t internet router, kernel JUNOS 20.4R3-S7.2, Build date: 2023-04-21 19:47:18 UTC Copyright (c) 1996-2023 Juniper Networks, Inc.","system_name":"Phoenix-Switch"},"lldp_stats":{"eth0":{"chassis_id":"63:68:61:73:73:69","lldp_med_supported":false,"mgmt_addr":"10.1.5.2","mgmt_addrs":["10.1.5.2","2002:A01:203:0:0:0:0:0"],"port_desc":"2/26","port_id":"port1","power_allocated":15500,"power_draw":15000,"power_request_count":3,"power_requested":25500,"system_desc":"HP J9729A 2920-48G-POE+ Switch","system_name":"TC2-OWL-Stack-01"},"eth1":{"chassis_id":"chassis2","lldp_med_supported":true,"mgmt_addr":"10.1.5.3","mgmt_addrs":["10.1.5.3"],"port_desc":"1/12","port_id":"port2","power_allocated":15400,"power_draw":14000,"power_request_count":2,"power_requested":15400,"system_desc":"HP J9728A 2920-24G-POE+ Switch","system_name":"TC2-OWL-Stack-02"}},"mac":"5c5b35d0077b","mem_total_kb":505468,"mem_used_kb":202096,"model":"AP43","modified_time":1718530662,"mount":"faceup","name":"Phoenix","notes":"","num_clients":1,"org_id":"af010e2b-f829-4975-b49e-2e896ed1d627","port_stat":{"eth0":{"full_duplex":true,"rx_bytes":1284143195,"rx_errors":0,"rx_peak_bps":17585,"rx_pkts":5199816,"speed":1000,"tx_bytes":1283744961,"tx_peak_bps":26484,"tx_pkts":3990463,"up":true},"eth1":{"full_duplex":false,"rx_bytes":0,"rx_errors":0,"rx_peak_bps":0,"rx_pkts":0,"speed":0,"tx_bytes":0,"tx_peak_bps":0,"tx_pkts":0,"up":false}},"power_budget":8400,"power_constrained":false,"power_src":"PoE 802.3at","power_srcs":["PoE 802.3at","PoE 802.3af"],"radio_config":{},"radio_stat":{"band_24":{"bandwidth":20,"channel":11,"mac":"5c5b35dea810","noise_floor":-80,"num_clients":0,"power":17,"rx_bytes":12948211,"rx_pkts":65292,"tx_bytes":19071943,"tx_pkts":76926,"usage":"24","util_all":24,"util_non_wifi":2,"util_rx_in_bss":0,"util_rx_other_bss":17,"util_tx":4,"util_undecodable_wifi":0,"util_unknown_wifi":1},"band_5":{"bandwidth":40,"channel":36,"mac":"5c5b35dea7f0","noise_floor":-90,"num_clients":1,"power":17,"rx_bytes":578362619,"rx_pkts":2687577,"tx_bytes":1199571353,"tx_pkts":2479302,"usage":"5","util_all":13,"util_non_wifi":0,"util_rx_in_bss":0,"util_rx_other_bss":10,"util_tx":1,"util_undecodable_wifi":0,"util_unknown_wifi":1}},"rx_bps":9276,"rx_bytes":591310830,"rx_pkts":2752869,"serial":"A0703200709E6","site_id":"46fc665e-9706-4296-8fe2-78f42f2e67e4","status":"connected","switch_redundancy":{"num_redundant_aps":1},"tx_bps":8067,"tx_bytes":1218643296,"tx_pkts":2556228,"type":"ap","uptime":1593120,"version":"0.14.29313"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
