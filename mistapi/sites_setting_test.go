package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSitesSettingTestGetSiteSetting tests the behavior of the SitesSetting
func TestSitesSettingTestGetSiteSetting(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	apiResponse, err := sitesSetting.GetSiteSetting(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"additional_config_cmds":["set snmp community public"],"analytic":{"enabled":false},"ap_matching":{"enabled":true,"rules":[{"eth1,eth2":{"port_vlan_id":1,"vlan_ids":[1,10,50]}}]},"ap_port_config":{"model_specific":{"AP32":{"eth1,eth2":{"port_vlan_id":1,"vlan_ids":[1,10,50]}}}},"auto_placement":{"orientation":45,"x":30,"y":60},"auto_upgrade":{"custom_versions":{"AP21":"stable","AP41":"0.1.5135","AP61":"0.1.7215"},"day_of_week":"sun","enabled":false,"time_of_day":"12:00","version":"beta"},"blacklist_url":"https://papi.s3.amazonaws.com/blacklist/xxx...","ble_config":{"beacon_enabled":false,"beacon_rate":3,"beacon_rate_mode":"custom","beam_disabled":[1,3,6],"custom_ble_packet_enabled":false,"custom_ble_packet_frame":"0x........","custom_ble_packet_freq_msec":300,"eddystone_uid_adv_power":-65,"eddystone_uid_beams":"2-4,7","eddystone_uid_enabled":false,"eddystone_uid_freq_msec":200,"eddystone_uid_instance":"5c5b35000001","eddystone_uid_namespace":"2818e3868dec25629ede","eddystone_url_adv_power":-65,"eddystone_url_beams":"2-4,7","eddystone_url_enabled":true,"eddystone_url_freq_msec":1000,"eddystone_url_url":"https://www.abc.com","ibeacon_adv_power":-65,"ibeacon_beams":"2-4,7","ibeacon_enabled":false,"ibeacon_freq_msec":0,"ibeacon_major":13,"ibeacon_minor":138,"ibeacon_uuid":"f3f17139-704a-f03a-2786-0400279e37c3","power":10,"power_mode":"custom"},"config_auto_revert":false,"created_time":0,"device_updown_threshold":0,"dns_servers":["string"],"dns_suffix":["string"],"engagement":{"dwell_tag_names":{"bounce":"Bounce","engaged":"Engaged","passerby":"Passer By","stationed":"Stationed"},"dwell_tags":{"bounce":null,"engaged":"300-14400","passerby":null,"stationed":"14400-43200"},"hours":{"fri":"09:00-17:00","mon":"09:00-17:00","sat":"09:00-12:00","sun":"09:00-12:00","thu":"09:00-17:00","tue":"09:00-17:00","wed":"09:00-17:00"},"max_dwell":43200,"min_dwell":0},"evpn_options":{"auto_loopback_subnet":"100.101.0.0/16","auto_router_id_subnet":"100.100.0.0/24","core_as_border":false,"overlay":{"as":65000},"per_vlan_vga_v4_mac":false,"routed_at":"edge","underlay":{"as_base":65001,"routed_id_prefix":"/24","subnet":"10.255.240.0/20"}},"flags":{"property1":"string","property2":"string"},"for_site":true,"gateway_additional_config_cmds":["set snmp community public"],"gateway_mgmt":{"admin_sshkeys":["string"],"app_probing":{"apps":["string"],"custom_apps":[{"app_type":"string","hostname":["string"],"name":"string","protocol":"http"}],"enabled":true},"app_usage":true,"auto_signature_update":{"day_of_week":"mon","enable":true,"time_of_day":"string"},"config_revert_timer":10,"probe_hosts":["string"],"root_password":"string","security_log_source_address":"192.168.1.1","security_log_source_interface":"string"},"id":"497f6eca-6276-4993-bfeb-53cbbbba6f09","led":{"brightness":255,"enabled":true},"modified_time":0,"mxedge":{"mist_das":{"coa_servers":[{"disable_event_timestamp_check":false,"enabled":true,"host":"string","port":3799,"secret":"string"}],"enabled":false},"radsec":{"acct_servers":[{"host":"string","port":1813,"secret":"string","ssids":["string"]}],"auth_servers":[{"host":"string","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"string","keywrap_mack":"string","port":1812,"secret":"string","ssids":["string"]}],"enabled":true,"match_ssid":true,"proxy_hosts":["string"],"server_selection":"ordered","source":"any"}},"mxedge_mgmt":{"mist_password":"MIST_PASSWORD","root_password":"ROOT_PASSWORD"},"ntp_servers":["pool.ntp.org"],"occupancy":{"assets_enabled":false,"clients_enabled":true,"min_duration":3000,"sdkclients_enabled":false,"unconnected_clients_enabled":false},"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","ospf_areas":{"property1":{"include_loopback":false,"networks":{"corp":{"auth_keys":{"1":"auth-key-1"},"auth_type":"md5","bfd_minimum_interval":500,"dead_interval":40,"hello_interval":10,"interface_type":"nbma","metric":10000},"guest":{"passive":true}},"type":"default"},"property2":{"include_loopback":false,"networks":{"corp":{"auth_keys":{"1":"auth-key-1"},"auth_type":"md5","bfd_minimum_interval":500,"dead_interval":40,"hello_interval":10,"interface_type":"nbma","metric":10000},"guest":{"passive":true}},"type":"default"}},"persist_config_on_device":false,"port_mirroring":{"property1":{"input_networks_ingress":["corp"],"input_port_ids_egress":["ge-0/0/3"],"input_port_ids_ingress":["ge-0/0/3"],"output_network":"analyze","output_port_id":"ge-0/0/5"},"property2":{"input_networks_ingress":["corp"],"input_port_ids_egress":["ge-0/0/3"],"input_port_ids_ingress":["ge-0/0/3"],"output_network":"analyze","output_port_id":"ge-0/0/5"}},"port_usages":{"dynamic":{"mode":"dynamic","reset_default_when":"link_down","rules":[{"equals":"string","equals_any":["string"],"expression":"string","src":"lldp_chassis_id","usage":"string"}]},"property1":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","mtu":0,"networks":["string"],"persist_mac":false,"poe_disabled":false,"port_auth":"dot1x","port_network":"string","rejected_network":null,"speed":"string","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"},"property2":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","mtu":0,"networks":["string"],"persist_mac":false,"poe_disabled":false,"port_network":"string","rejected_network":null,"speed":"string","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"}},"proxy":{"url":"http://proxy.internal:8080/"},"radius_config":{"acct_interim_interval":0,"acct_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1813,"secret":"testing123"}],"auth_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1812,"secret":"testing123"}],"auth_servers_retries":3,"auth_servers_timeout":5,"coa_enabled":false,"coa_port":3799,"network":"string","source_ip":"string"},"remote_syslog":{"archive":{"files":20,"size":"5m"},"console":{"contents":[{"facility":"config","severity":"warning"}]},"enabled":false,"files":[{"archive":{"files":10,"size":"5m"},"contents":[{"facility":"config","severity":"warning"}],"explicit_priority":true,"file":"file-name","match":"!alarm|ntp|errors.crc_error[chan]","structured_data":true}],"network":"default","send_to_all_servers":false,"servers":[{"facility":"config","host":"syslogd.internal","port":514,"protocol":"udp","severity":"info","tag":""}],"time_format":"millisecond","users":[{"contents":[{"facility":"config","severity":"warning"}],"match":"\"!alarm|ntp|errors.crc_error[chan]\"","user":"*"}]},"report_gatt":false,"rogue":{"enabled":false,"honeypot_enabled":false,"min_duration":10,"min_rssi":-80,"whitelisted_bssids":["NeighborSSID"],"whitelisted_ssids":["cc:8e:6f:d4:bf:16","cc-8e-6f-d4-bf-16","cc-73-*","cc:82:*"]},"rtsa":{"app_waking":false,"disable_dead_reckoning":true,"disable_pressure_sensor":false,"enabled":true,"track_asset":false},"simple_alert":{"arp_failure":{"client_count":10,"duration":20,"incident_count":10},"dhcp_failure":{"client_count":10,"duration":10,"incident_count":20},"dns_failure":{"client_count":20,"duration":10,"incident_count":30}},"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","skyatp":{"enabled":true,"send_ip_mac_mapping":true},"srx_app":{"enabled":false},"ssh_keys":["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAA...Wxa6p6UW0ZbcP john@host"],"ssr":{"conductor_hosts":["\"1.1.1.1\", \"2.2.2.2\""],"disable_stats":true},"status_portal":{"enabled":false,"hostnames":["my.misty.com"]},"switch_mgmt":{"ap_affinity_threshold":10,"config_revert_timer":10,"dhcp_option_fqdn":false,"mxedge_proxy_host":"string","mxedge_proxy_port":2222,"root_password":"string","tacacs":{"acct_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}],"enabled":true,"network":"string","tacplus_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}]},"use_mxedge_proxy":true},"vars":{"RADIUS_IP1":"172.31.2.5","RADIUS_SECRET":"11s64632d"},"vna":{"enabled":false},"vrf_instances":{"guest":{"extra_routes":{"0.0.0.0/0":{"via":"192.168.31.1"}},"networks":["guest"]}},"vrrp_groups":{"property1":{"auth_key":"auth-key-1","auth_password":"string","auth_type":"md5","networks":{"data":{"ip":"10.182.96.1"},"mgmt":{"ip":"10.182.104.1"},"v10":{"ip":"10.182.104.129"},"wap":{"ip":"10.182.102.1"}}},"property2":{"auth_key":"auth-key-1","auth_password":"string","auth_type":"md5","networks":{"data":{"ip":"10.182.96.1"},"mgmt":{"ip":"10.182.104.1"},"v10":{"ip":"10.182.104.129"},"wap":{"ip":"10.182.102.1"}}}},"wan_vna":{"enabled":false},"watched_station_url":"https://papi.s3.amazonaws.com/watched_station/xxx...","whitelist_url":"https://papi.s3.amazonaws.com/whitelist/xxx...","wids":{"repeated_auth_failures":{"duration":60,"threshold":0}},"wifi":{"cisco_enabled":true,"disable_11k":false,"disable_radios_when_power_constrained":false,"enable_arp_spoof_check":false,"enable_shared_radio_scanning":true,"enabled":true,"locate_connected":true,"locate_unconnected":false,"mesh_allow_dfs":false,"mesh_enable_crm":false,"mesh_enabled":false,"mesh_psk":"string","mesh_ssid":"string","proxy_arp":"default"},"wired_vna":{"enabled":false},"zone_occupancy_alert":{"email_notifiers":["foo@juniper.net","bar@juniper.net"],"enabled":false,"threshold":5}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesSettingTestUpdateSiteSettings tests the behavior of the SitesSetting
func TestSitesSettingTestUpdateSiteSettings(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.SiteSetting
	errBody := json.Unmarshal([]byte(`{"additional_config_cmds":["set snmp community public"],"analytic":{"enabled":false},"ap_matching":{"enabled":true,"rules":[{"eth1,eth2":{"port_vlan_id":1,"vlan_ids":[1,10,50]}}]},"ap_port_config":{"model_specific":{"AP32":{"eth1,eth2":{"port_vlan_id":1,"vlan_ids":[1,10,50]}}}},"auto_upgrade":{"custom_versions":{"AP21":"stable","AP41":"0.1.5135","AP61":"0.1.7215"},"day_of_week":"sun","enabled":false,"time_of_day":"12:00","version":"beta"},"config_auto_revert":false,"device_updown_threshold":0,"dns_servers":["string"],"dns_suffix":["string"],"engagement":{"dwell_tag_names":{"bounce":"Bounce","engaged":"Engaged","passerby":"Passer By","stationed":"Stationed"},"dwell_tags":{"bounce":null,"engaged":"300-14400","passerby":null,"stationed":"14400-43200"},"hours":{"fri":"09:00-17:00","mon":"09:00-17:00","sat":"09:00-12:00","sun":"09:00-12:00","thu":"09:00-17:00","tue":"09:00-17:00","wed":"09:00-17:00"},"max_dwell":43200,"min_dwell":0},"evpn_options":{"auto_loopback_subnet":"100.101.0.0/16","auto_router_id_subnet":"100.100.0.0/24","core_as_border":false,"overlay":{"as":65000},"per_vlan_vga_v4_mac":false,"routed_at":"edge","underlay":{"as_base":65001,"routed_id_prefix":"/24","subnet":"10.255.240.0/20"}},"gateway_additional_config_cmds":["set snmp community public"],"gateway_mgmt":{"admin_sshkeys":["string"],"app_probing":{"apps":["string"],"custom_apps":[{"app_type":"string","hostname":["string"],"name":"string","protocol":"http"}],"enabled":true},"app_usage":true,"auto_signature_update":{"day_of_week":"any","enable":true,"time_of_day":"string"},"config_revert_timer":10,"probe_hosts":["string"],"root_password":"string","security_log_source_address":"192.168.1.1","security_log_source_interface":"string"},"led":{"brightness":255,"enabled":true},"mxedge_mgmt":{"mist_password":"MIST_PASSWORD","root_password":"ROOT_PASSWORD"},"networks":{"property1":{"dns":["string"],"dns_suffix":["string"],"gateway":"string","ospf_interface_type":"string","subnet":"string","vlan_id":10,"zone":"string"},"property2":{"dns":["string"],"dns_suffix":["string"],"gateway":"string","ospf_interface_type":"string","subnet":"string","vlan_id":10,"zone":"string"}},"ntp_servers":["string"],"occupancy":{"assets_enabled":false,"clients_enabled":true,"min_duration":3000,"sdkclients_enabled":false,"unconnected_clients_enabled":false},"ospf_areas":{"property1":{"include_loopback":false,"networks":{"corp":{"auth_keys":{"1":"auth-key-1"},"auth_type":"md5","bfd_minimum_interval":500,"dead_interval":40,"hello_interval":10,"interface_type":"nbma","metric":10000},"guest":{"passive":true}},"type":"default"},"property2":{"include_loopback":false,"networks":{"corp":{"auth_keys":{"1":"auth-key-1"},"auth_type":"md5","bfd_minimum_interval":500,"dead_interval":40,"hello_interval":10,"interface_type":"nbma","metric":10000},"guest":{"passive":true}},"type":"default"}},"persist_config_on_device":false,"port_mirroring":{"property1":{"input_networks_ingress":["corp"],"input_port_ids_egress":["ge-0/0/3"],"input_port_ids_ingress":["ge-0/0/3"],"output_network":"analyze","output_port_id":"ge-0/0/5"},"property2":{"input_networks_ingress":["corp"],"input_port_ids_egress":["ge-0/0/3"],"input_port_ids_ingress":["ge-0/0/3"],"output_network":"analyze","output_port_id":"ge-0/0/5"}},"port_usages":{"dynamic":{"mode":"dynamic","reset_default_when":"link_down","rules":[{"equals":"string","equals_any":["string"],"expression":"string","src":"lldp_chassis_id","usage":"string"}]},"property1":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","mtu":0,"networks":["string"],"persist_mac":false,"poe_disabled":false,"port_auth":"dot1x","port_network":"string","rejected_network":null,"speed":"string","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"},"property2":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","mtu":0,"networks":["string"],"persist_mac":false,"poe_disabled":false,"port_network":"string","rejected_network":null,"speed":"string","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"}},"proxy":{"url":"http://proxy.internal:8080/*"},"rogue":{"enabled":false,"honeypot_enabled":false,"min_duration":10,"min_rssi":-80,"whitelisted_bssids":["NeighborSSID"],"whitelisted_ssids":["cc:8e:6f:d4:bf:16","cc-8e-6f-d4-bf-16","cc-73-*","cc:82:*"]},"simple_alert":{"arp_failure":{"client_count":10,"duration":20,"incident_count":10},"dhcp_failure":{"client_count":10,"duration":10,"incident_count":20},"dns_failure":{"client_count":20,"duration":10,"incident_count":30}},"skyatp":{"enabled":true,"send_ip_mac_mapping":true},"srx_app":{"enabled":false},"ssh_keys":["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAA...Wxa6p6UW0ZbcP john@host"],"ssr":{"conductor_hosts":["\"1.1.1.1\", \"2.2.2.2\""],"disable_stats":true},"status_portal":{"enabled":false,"hostnames":["my.misty.com"]},"vars":{"RADIUS_IP1":"172.31.2.5","RADIUS_SECRET":"11s64632d"},"vna":{"enabled":false},"wan_vna":{"enabled":false},"wids":{"repeated_auth_failures":{"duration":60,"threshold":0}},"wifi":{"cisco_enabled":true,"disable_11k":false,"disable_radios_when_power_constrained":false,"enable_arp_spoof_check":false,"enable_shared_radio_scanning":true,"enabled":true,"locate_connected":true,"locate_unconnected":false,"mesh_allow_dfs":false,"mesh_enable_crm":false,"mesh_enabled":false,"mesh_psk":"string","mesh_ssid":"string","proxy_arp":"default"},"wired_vna":{"enabled":false},"zone_occupancy_alert":{"email_notifiers":["foo@juniper.net","bar@juniper.net"],"enabled":false,"threshold":5}}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesSetting.UpdateSiteSettings(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"additional_config_cmds":["set snmp community public"],"analytic":{"enabled":false},"ap_matching":{"enabled":true,"rules":[{"eth1,eth2":{"port_vlan_id":1,"vlan_ids":[1,10,50]}}]},"ap_port_config":{"model_specific":{"AP32":{"eth1,eth2":{"port_vlan_id":1,"vlan_ids":[1,10,50]}}}},"auto_placement":{"orientation":45,"x":30,"y":60},"auto_upgrade":{"custom_versions":{"AP21":"stable","AP41":"0.1.5135","AP61":"0.1.7215"},"day_of_week":"sun","enabled":false,"time_of_day":"12:00","version":"beta"},"blacklist_url":"https://papi.s3.amazonaws.com/blacklist/xxx...","ble_config":{"beacon_enabled":false,"beacon_rate":3,"beacon_rate_mode":"custom","beam_disabled":[1,3,6],"custom_ble_packet_enabled":false,"custom_ble_packet_frame":"0x........","custom_ble_packet_freq_msec":300,"eddystone_uid_adv_power":-65,"eddystone_uid_beams":"2-4,7","eddystone_uid_enabled":false,"eddystone_uid_freq_msec":200,"eddystone_uid_instance":"5c5b35000001","eddystone_uid_namespace":"2818e3868dec25629ede","eddystone_url_adv_power":-65,"eddystone_url_beams":"2-4,7","eddystone_url_enabled":true,"eddystone_url_freq_msec":1000,"eddystone_url_url":"https://www.abc.com","ibeacon_adv_power":-65,"ibeacon_beams":"2-4,7","ibeacon_enabled":false,"ibeacon_freq_msec":0,"ibeacon_major":13,"ibeacon_minor":138,"ibeacon_uuid":"f3f17139-704a-f03a-2786-0400279e37c3","power":10,"power_mode":"custom"},"config_auto_revert":false,"created_time":0,"device_updown_threshold":0,"dns_servers":["string"],"dns_suffix":["string"],"engagement":{"dwell_tag_names":{"bounce":"Bounce","engaged":"Engaged","passerby":"Passer By","stationed":"Stationed"},"dwell_tags":{"bounce":null,"engaged":"300-14400","passerby":null,"stationed":"14400-43200"},"hours":{"fri":"09:00-17:00","mon":"09:00-17:00","sat":"09:00-12:00","sun":"09:00-12:00","thu":"09:00-17:00","tue":"09:00-17:00","wed":"09:00-17:00"},"max_dwell":43200,"min_dwell":0},"evpn_options":{"auto_loopback_subnet":"100.101.0.0/16","auto_router_id_subnet":"100.100.0.0/24","core_as_border":false,"overlay":{"as":65000},"per_vlan_vga_v4_mac":false,"routed_at":"edge","underlay":{"as_base":65001,"routed_id_prefix":"/24","subnet":"10.255.240.0/20"}},"flags":{"property1":"string","property2":"string"},"for_site":true,"gateway_additional_config_cmds":["set snmp community public"],"gateway_mgmt":{"admin_sshkeys":["string"],"app_probing":{"apps":["string"],"custom_apps":[{"app_type":"string","hostname":["string"],"name":"string","protocol":"http"}],"enabled":true},"app_usage":true,"auto_signature_update":{"day_of_week":"mon","enable":true,"time_of_day":"string"},"config_revert_timer":10,"probe_hosts":["string"],"root_password":"string","security_log_source_address":"192.168.1.1","security_log_source_interface":"string"},"id":"497f6eca-6276-4993-bfeb-53cbbbba6f09","led":{"brightness":255,"enabled":true},"modified_time":0,"mxedge":{"mist_das":{"coa_servers":[{"disable_event_timestamp_check":false,"enabled":true,"host":"string","port":3799,"secret":"string"}],"enabled":false},"radsec":{"acct_servers":[{"host":"string","port":1813,"secret":"string","ssids":["string"]}],"auth_servers":[{"host":"string","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"string","keywrap_mack":"string","port":1812,"secret":"string","ssids":["string"]}],"enabled":true,"match_ssid":true,"proxy_hosts":["string"],"server_selection":"ordered","source":"any"}},"mxedge_mgmt":{"mist_password":"MIST_PASSWORD","root_password":"ROOT_PASSWORD"},"ntp_servers":["pool.ntp.org"],"occupancy":{"assets_enabled":false,"clients_enabled":true,"min_duration":3000,"sdkclients_enabled":false,"unconnected_clients_enabled":false},"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","ospf_areas":{"property1":{"include_loopback":false,"networks":{"corp":{"auth_keys":{"1":"auth-key-1"},"auth_type":"md5","bfd_minimum_interval":500,"dead_interval":40,"hello_interval":10,"interface_type":"nbma","metric":10000},"guest":{"passive":true}},"type":"default"},"property2":{"include_loopback":false,"networks":{"corp":{"auth_keys":{"1":"auth-key-1"},"auth_type":"md5","bfd_minimum_interval":500,"dead_interval":40,"hello_interval":10,"interface_type":"nbma","metric":10000},"guest":{"passive":true}},"type":"default"}},"persist_config_on_device":false,"port_mirroring":{"property1":{"input_networks_ingress":["corp"],"input_port_ids_egress":["ge-0/0/3"],"input_port_ids_ingress":["ge-0/0/3"],"output_network":"analyze","output_port_id":"ge-0/0/5"},"property2":{"input_networks_ingress":["corp"],"input_port_ids_egress":["ge-0/0/3"],"input_port_ids_ingress":["ge-0/0/3"],"output_network":"analyze","output_port_id":"ge-0/0/5"}},"port_usages":{"dynamic":{"mode":"dynamic","reset_default_when":"link_down","rules":[{"equals":"string","equals_any":["string"],"expression":"string","src":"lldp_chassis_id","usage":"string"}]},"property1":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","mtu":0,"networks":["string"],"persist_mac":false,"poe_disabled":false,"port_auth":"dot1x","port_network":"string","rejected_network":null,"speed":"string","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"},"property2":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","mtu":0,"networks":["string"],"persist_mac":false,"poe_disabled":false,"port_network":"string","rejected_network":null,"speed":"string","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"}},"proxy":{"url":"http://proxy.internal:8080/"},"radius_config":{"acct_interim_interval":0,"acct_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1813,"secret":"testing123"}],"auth_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1812,"secret":"testing123"}],"auth_servers_retries":3,"auth_servers_timeout":5,"coa_enabled":false,"coa_port":3799,"network":"string","source_ip":"string"},"remote_syslog":{"archive":{"files":20,"size":"5m"},"console":{"contents":[{"facility":"config","severity":"warning"}]},"enabled":false,"files":[{"archive":{"files":10,"size":"5m"},"contents":[{"facility":"config","severity":"warning"}],"explicit_priority":true,"file":"file-name","match":"!alarm|ntp|errors.crc_error[chan]","structured_data":true}],"network":"default","send_to_all_servers":false,"servers":[{"facility":"config","host":"syslogd.internal","port":514,"protocol":"udp","severity":"info","tag":""}],"time_format":"millisecond","users":[{"contents":[{"facility":"config","severity":"warning"}],"match":"\"!alarm|ntp|errors.crc_error[chan]\"","user":"*"}]},"report_gatt":false,"rogue":{"enabled":false,"honeypot_enabled":false,"min_duration":10,"min_rssi":-80,"whitelisted_bssids":["NeighborSSID"],"whitelisted_ssids":["cc:8e:6f:d4:bf:16","cc-8e-6f-d4-bf-16","cc-73-*","cc:82:*"]},"rtsa":{"app_waking":false,"disable_dead_reckoning":true,"disable_pressure_sensor":false,"enabled":true,"track_asset":false},"simple_alert":{"arp_failure":{"client_count":10,"duration":20,"incident_count":10},"dhcp_failure":{"client_count":10,"duration":10,"incident_count":20},"dns_failure":{"client_count":20,"duration":10,"incident_count":30}},"site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","skyatp":{"enabled":true,"send_ip_mac_mapping":true},"srx_app":{"enabled":false},"ssh_keys":["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAA...Wxa6p6UW0ZbcP john@host"],"ssr":{"conductor_hosts":["\"1.1.1.1\", \"2.2.2.2\""],"disable_stats":true},"status_portal":{"enabled":false,"hostnames":["my.misty.com"]},"switch_mgmt":{"ap_affinity_threshold":10,"config_revert_timer":10,"dhcp_option_fqdn":false,"mxedge_proxy_host":"string","mxedge_proxy_port":2222,"root_password":"string","tacacs":{"acct_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}],"enabled":true,"network":"string","tacplus_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}]},"use_mxedge_proxy":true},"vars":{"RADIUS_IP1":"172.31.2.5","RADIUS_SECRET":"11s64632d"},"vna":{"enabled":false},"vrf_instances":{"guest":{"extra_routes":{"0.0.0.0/0":{"via":"192.168.31.1"}},"networks":["guest"]}},"vrrp_groups":{"property1":{"auth_key":"auth-key-1","auth_password":"string","auth_type":"md5","networks":{"data":{"ip":"10.182.96.1"},"mgmt":{"ip":"10.182.104.1"},"v10":{"ip":"10.182.104.129"},"wap":{"ip":"10.182.102.1"}}},"property2":{"auth_key":"auth-key-1","auth_password":"string","auth_type":"md5","networks":{"data":{"ip":"10.182.96.1"},"mgmt":{"ip":"10.182.104.1"},"v10":{"ip":"10.182.104.129"},"wap":{"ip":"10.182.102.1"}}}},"wan_vna":{"enabled":false},"watched_station_url":"https://papi.s3.amazonaws.com/watched_station/xxx...","whitelist_url":"https://papi.s3.amazonaws.com/whitelist/xxx...","wids":{"repeated_auth_failures":{"duration":60,"threshold":0}},"wifi":{"cisco_enabled":true,"disable_11k":false,"disable_radios_when_power_constrained":false,"enable_arp_spoof_check":false,"enable_shared_radio_scanning":true,"enabled":true,"locate_connected":true,"locate_unconnected":false,"mesh_allow_dfs":false,"mesh_enable_crm":false,"mesh_enabled":false,"mesh_psk":"string","mesh_ssid":"string","proxy_arp":"default"},"wired_vna":{"enabled":false},"zone_occupancy_alert":{"email_notifiers":["foo@juniper.net","bar@juniper.net"],"enabled":false,"threshold":5}}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesSettingTestDeleteSiteWirelessClientsBlocklist tests the behavior of the SitesSetting
func TestSitesSettingTestDeleteSiteWirelessClientsBlocklist(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesSetting.DeleteSiteWirelessClientsBlocklist(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesSettingTestCreateSiteWirelessClientsBlocklist tests the behavior of the SitesSetting
func TestSitesSettingTestCreateSiteWirelessClientsBlocklist(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.MacAddresses
	errBody := json.Unmarshal([]byte(`{"macs":["18-65-90-de-f4-c6","84-89-ad-5d-69-0d"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesSetting.CreateSiteWirelessClientsBlocklist(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"macs":["18-65-90-de-f4-c6","84-89-ad-5d-69-0d"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesSettingTestDeleteSiteWatchedStations tests the behavior of the SitesSetting
func TestSitesSettingTestDeleteSiteWatchedStations(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesSetting.DeleteSiteWatchedStations(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesSettingTestCreateSiteWatchedStations tests the behavior of the SitesSetting
func TestSitesSettingTestCreateSiteWatchedStations(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	var body models.MacAddresses
	errBody := json.Unmarshal([]byte(`{"macs":["18-65-90-de-f4-c6","84-89-ad-5d-69-0d"]}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := sitesSetting.CreateSiteWatchedStations(ctx, siteId, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"macs":["18-65-90-de-f4-c6","84-89-ad-5d-69-0d"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesSettingTestDeleteSiteWirelessClientsAllowlist tests the behavior of the SitesSetting
func TestSitesSettingTestDeleteSiteWirelessClientsAllowlist(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}
	resp, err := sitesSetting.DeleteSiteWirelessClientsAllowlist(ctx, siteId)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesSettingTestCreateSiteWirelessClientsAllowlist tests the behavior of the SitesSetting
func TestSitesSettingTestCreateSiteWirelessClientsAllowlist(t *testing.T) {
	ctx := context.Background()
	siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := sitesSetting.CreateSiteWirelessClientsAllowlist(ctx, siteId, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"macs":["18-65-90-de-f4-c6","84-89-ad-5d-69-0d"]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
