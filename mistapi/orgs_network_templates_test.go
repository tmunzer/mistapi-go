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

// TestOrgsNetworkTemplatesTestListOrgNetworkTemplates tests the behavior of the OrgsNetworkTemplates
func TestOrgsNetworkTemplatesTestListOrgNetworkTemplates(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsNetworkTemplates.ListOrgNetworkTemplates(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"additional_config_cmds":["set snmp community public"],"created_time":0,"dhcp_snooping":{"all_networks":true,"enable_arp_spoof_check":true,"enable_ip_source_guard":true,"enabled":true,"networks":["string"]},"dns_servers":["8.8.8.8","4.4.4.4"],"dns_suffix":[".mist.local",".mist.com"],"extra_routes":{"0.0.0.0/0":{"via":"1.2.3.4"}},"group_tags":{},"id":"497f6eca-6276-4993-bfeb-53cbbbba6808","import_org_networks":["ap"],"mist_nac":{"enabled":true,"network":"default"},"modified_time":0,"name":"template_name","networks":{"corp":{"vlan_id":600},"default":{"subnet":"192.168.1.0/24","vlan_id":1},"guest":{"vlan_id":700},"mgmt":{"vlan_id":500}},"ntp_servers":["192.168.1.10"],"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","port_usages":{"ap":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"WAP","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":false,"enable_qos":true,"mac_auth_only":false,"mac_limit":0,"mode":"trunk","networks":["guest","corp"],"persist_mac":false,"poe_disabled":false,"port_network":"default","rejected_network":null,"storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true},"dynamic":{"mode":"dynamic","reset_default_when":"link_down","rules":[{"equals":"string","equals_any":["string"],"expression":"string","src":"lldp_chassis_id","usage":"string"}]},"iot":{"allow_dhcpd":true,"mode":"access","port_network":"default","stp_edge":true},"uplink":{"all_networks":true,"enable_qos":false,"mode":"trunk","port_network":"default","stp_edge":false}},"radius_config":{"acct_interim_interval":0,"acct_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1813,"secret":"testing123"}],"auth_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1812,"secret":"testing123"}],"auth_servers_retries":3,"auth_servers_timeout":5,"coa_enabled":false,"coa_port":3799,"network":"default"},"remote_syslog":{"archive":{"files":20,"size":"5m"},"console":{"contents":[{"facility":"config","severity":"warning"}]},"enabled":false,"files":[{"archive":{"files":10,"size":"5m"},"contents":[{"facility":"config","severity":"warning"}],"explicit_priority":true,"file":"file-name","match":"!alarm|ntp|errors.crc_error[chan]","structured_data":true}],"network":"default","send_to_all_servers":false,"servers":[{"facility":"config","host":"syslogd.internal","port":514,"protocol":"udp","severity":"info","tag":""}],"time_format":"millisecond","users":[{"contents":[{"facility":"config","severity":"warning"}],"match":"\"!alarm|ntp|errors.crc_error[chan]\"","user":"*"}]},"switch_matching":{"enable":true,"rules":[{"additional_config_cmds":["set snmp community public"],"match_model":"EX4300","match_name[0:3]":"abc","name":"match by name","port_config":{"ge-0/0/0":{"usage":"uplink"},"ge-0/0/8-16ge-1/0/0-47":{"usage":"ap"}}},{"additional_config_cmds":["set snmp community public2"],"match_role":"access","name":"match by role","port_config":{"ge-0/0/0":{"usage":"uplink"}}}]},"switch_mgmt":{"config_revert":10,"protect_re":{"enabled":false},"root_password":"string","tacacs":{"acct_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}],"enabled":true,"network":"string","tacplus_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}]}},"vrf_config":{"enabled":false},"vrf_instances":{"property1":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.10"}},"networks":["string"]},"property2":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.10"}},"networks":["string"]}}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNetworkTemplatesTestListOrgNetworkTemplates1 tests the behavior of the OrgsNetworkTemplates
func TestOrgsNetworkTemplatesTestListOrgNetworkTemplates1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsNetworkTemplates.ListOrgNetworkTemplates(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"additional_config_cmds":["set snmp community public"],"created_time":0,"dhcp_snooping":{"all_networks":true,"enable_arp_spoof_check":true,"enable_ip_source_guard":true,"enabled":true,"networks":["string"]},"dns_servers":["8.8.8.8","4.4.4.4"],"dns_suffix":[".mist.local",".mist.com"],"extra_routes":{"0.0.0.0/0":{"via":"1.2.3.4"}},"group_tags":{},"id":"497f6eca-6276-4993-bfeb-53cbbbba6808","import_org_networks":["ap"],"mist_nac":{"enabled":true,"network":"default"},"modified_time":0,"name":"template_name","networks":{"corp":{"vlan_id":600},"default":{"subnet":"192.168.1.0/24","vlan_id":1},"guest":{"vlan_id":700},"mgmt":{"vlan_id":500}},"ntp_servers":["192.168.1.10"],"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","port_usages":{"ap":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"WAP","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":false,"enable_qos":true,"mac_auth_only":false,"mac_limit":0,"mode":"trunk","networks":["guest","corp"],"persist_mac":false,"poe_disabled":false,"port_network":"default","rejected_network":null,"storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true},"dynamic":{"mode":"dynamic","reset_default_when":"link_down","rules":[{"equals":"string","equals_any":["string"],"expression":"string","src":"lldp_chassis_id","usage":"string"}]},"iot":{"allow_dhcpd":true,"mode":"access","port_network":"default","stp_edge":true},"uplink":{"all_networks":true,"enable_qos":false,"mode":"trunk","port_network":"default","stp_edge":false}},"radius_config":{"acct_interim_interval":0,"acct_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1813,"secret":"testing123"}],"auth_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1812,"secret":"testing123"}],"auth_servers_retries":3,"auth_servers_timeout":5,"coa_enabled":false,"coa_port":3799,"network":"default"},"remote_syslog":{"archive":{"files":20,"size":"5m"},"console":{"contents":[{"facility":"config","severity":"warning"}]},"enabled":false,"files":[{"archive":{"files":10,"size":"5m"},"contents":[{"facility":"config","severity":"warning"}],"explicit_priority":true,"file":"file-name","match":"!alarm|ntp|errors.crc_error[chan]","structured_data":true}],"network":"default","send_to_all_servers":false,"servers":[{"facility":"config","host":"syslogd.internal","port":514,"protocol":"udp","severity":"info","tag":""}],"time_format":"millisecond","users":[{"contents":[{"facility":"config","severity":"warning"}],"match":"\"!alarm|ntp|errors.crc_error[chan]\"","user":"*"}]},"switch_matching":{"enable":true,"rules":[{"additional_config_cmds":["set snmp community public"],"match_model":"EX4300","match_name[0:3]":"abc","name":"match by name","port_config":{"ge-0/0/0":{"usage":"uplink"},"ge-0/0/8-16ge-1/0/0-47":{"usage":"ap"}}},{"additional_config_cmds":["set snmp community public2"],"match_role":"access","name":"match by role","port_config":{"ge-0/0/0":{"usage":"uplink"}}}]},"switch_mgmt":{"config_revert":10,"protect_re":{"enabled":false},"root_password":"string","tacacs":{"acct_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}],"enabled":true,"network":"string","tacplus_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}]}},"vrf_config":{"enabled":false},"vrf_instances":{"property1":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.10"}},"networks":["string"]},"property2":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.10"}},"networks":["string"]}}}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNetworkTemplatesTestCreateOrgNetworkTemplate tests the behavior of the OrgsNetworkTemplates
func TestOrgsNetworkTemplatesTestCreateOrgNetworkTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.NetworkTemplate
    errBody := json.Unmarshal([]byte(`{"additional_config_cmds":["set snmp community public"],"dhcp_snooping":{"all_networks":true,"enable_arp_spoof_check":true,"enable_ip_source_guard":true,"enabled":true,"networks":["string"]},"dns_servers":["string"],"dns_suffix":["string"],"import_org_networks":["ap"],"mist_nac":{"enabled":true,"network":"string"},"name":"string","networks":{"property1":{"subnet":"192.168.1.0/24","vlan_id":10},"property2":{"subnet":"192.168.1.0/24","vlan_id":10}},"ntp_servers":["string"],"port_usages":{"dynamic":{"mode":"dynamic","reset_default_when":"link_down","rules":[{"equals":"string","equals_any":["string"],"expression":"string","src":"lldp_chassis_id","usage":"string"}]},"property1":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","networks":["string"],"persist_mac":false,"poe_disabled":false,"port_auth":"dot1x","port_network":"string","rejected_network":null,"speed":"auto","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"},"property2":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","networks":["string"],"persist_mac":false,"poe_disabled":false,"port_network":"string","rejected_network":null,"speed":"auto","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"}},"switch_mgmt":{"config_revert":10,"protect_re":{"enabled":false},"root_password":"string","tacacs":{"acct_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}],"enabled":true,"network":"string","tacplus_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}]}},"vrf_config":{"enabled":false},"vrf_instances":{"property1":{"extra_routes":{"property1":{"via":"192.0.2.10"},"property2":{"via":"198.51.100.1"}},"networks":["string"]},"property2":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.10"}},"networks":["string"]}}}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsNetworkTemplates.CreateOrgNetworkTemplate(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"additional_config_cmds":["set snmp community public"],"created_time":0,"dhcp_snooping":{"all_networks":true,"enable_arp_spoof_check":true,"enable_ip_source_guard":true,"enabled":true,"networks":["string"]},"dns_servers":["string"],"dns_suffix":["string"],"extra_routes":{"property1":{"via":"string"},"property2":{"via":"string"}},"group_tags":{},"id":"497f6eca-6276-4993-bfeb-53cbbbba6708","import_org_networks":["ap"],"mist_nac":{"enabled":true,"network":"string"},"modified_time":0,"name":"string","networks":{"property1":{"subnet":"192.168.1.0/24","vlan_id":10},"property2":{"subnet":"192.168.1.0/24","vlan_id":10}},"ntp_servers":["string"],"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","port_usages":{"dynamic":{"mode":"dynamic","reset_default_when":"link_down","rules":[{"equals":"string","equals_any":["string"],"expression":"string","src":"lldp_chassis_id","usage":"string"}]},"property1":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"networks":["string"],"persist_mac":false,"poe_disabled":false,"port_auth":"dot1x","port_network":"string","rejected_network":null,"speed":"auto","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"},"property2":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","networks":["string"],"persist_mac":false,"poe_disabled":false,"port_network":"string","rejected_network":null,"speed":"auto","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"}},"switch_mgmt":{"config_revert":10,"protect_re":{"enabled":false},"root_password":"string","tacacs":{"acct_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}],"enabled":true,"network":"string","tacplus_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}]}},"vrf_config":{"enabled":false},"vrf_instances":{"property1":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.10"}},"networks":["string"]},"property2":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.10"}},"networks":["string"]}}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNetworkTemplatesTestCreateOrgNetworkTemplate1 tests the behavior of the OrgsNetworkTemplates
func TestOrgsNetworkTemplatesTestCreateOrgNetworkTemplate1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.NetworkTemplate
    errBody := json.Unmarshal([]byte(`{"additional_config_cmds":["set snmp community public"],"dhcp_snooping":{"all_networks":true,"enable_arp_spoof_check":true,"enable_ip_source_guard":true,"enabled":true,"networks":["string"]},"dns_servers":["string"],"dns_suffix":["string"],"import_org_networks":["ap"],"mist_nac":{"enabled":true,"network":"string"},"name":"string","networks":{"property1":{"subnet":"192.168.1.0/24","vlan_id":10},"property2":{"subnet":"192.168.1.0/24","vlan_id":10}},"ntp_servers":["string"],"port_usages":{"dynamic":{"mode":"dynamic","reset_default_when":"link_down","rules":[{"equals":"string","equals_any":["string"],"expression":"string","src":"lldp_chassis_id","usage":"string"}]},"property1":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","networks":["string"],"persist_mac":false,"poe_disabled":false,"port_auth":"dot1x","port_network":"string","rejected_network":null,"speed":"auto","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"},"property2":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","networks":["string"],"persist_mac":false,"poe_disabled":false,"port_network":"string","rejected_network":null,"speed":"auto","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"}},"switch_mgmt":{"config_revert":10,"protect_re":{"enabled":false},"root_password":"string","tacacs":{"acct_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}],"enabled":true,"network":"string","tacplus_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}]}},"vrf_config":{"enabled":false},"vrf_instances":{"property1":{"extra_routes":{"property1":{"via":"192.0.2.10"},"property2":{"via":"198.51.100.1"}},"networks":["string"]},"property2":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.10"}},"networks":["string"]}}}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsNetworkTemplates.CreateOrgNetworkTemplate(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"additional_config_cmds":["set snmp community public"],"created_time":0,"dhcp_snooping":{"all_networks":true,"enable_arp_spoof_check":true,"enable_ip_source_guard":true,"enabled":true,"networks":["string"]},"dns_servers":["string"],"dns_suffix":["string"],"extra_routes":{"property1":{"via":"string"},"property2":{"via":"string"}},"group_tags":{},"id":"497f6eca-6276-4993-bfeb-53cbbbba6708","import_org_networks":["ap"],"mist_nac":{"enabled":true,"network":"string"},"modified_time":0,"name":"string","networks":{"property1":{"subnet":"192.168.1.0/24","vlan_id":10},"property2":{"subnet":"192.168.1.0/24","vlan_id":10}},"ntp_servers":["string"],"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","port_usages":{"dynamic":{"mode":"dynamic","reset_default_when":"link_down","rules":[{"equals":"string","equals_any":["string"],"expression":"string","src":"lldp_chassis_id","usage":"string"}]},"property1":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"networks":["string"],"persist_mac":false,"poe_disabled":false,"port_auth":"dot1x","port_network":"string","rejected_network":null,"speed":"auto","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"},"property2":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","networks":["string"],"persist_mac":false,"poe_disabled":false,"port_network":"string","rejected_network":null,"speed":"auto","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"}},"switch_mgmt":{"config_revert":10,"protect_re":{"enabled":false},"root_password":"string","tacacs":{"acct_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}],"enabled":true,"network":"string","tacplus_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}]}},"vrf_config":{"enabled":false},"vrf_instances":{"property1":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.10"}},"networks":["string"]},"property2":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.10"}},"networks":["string"]}}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNetworkTemplatesTestDeleteOrgNetworkTemplate tests the behavior of the OrgsNetworkTemplates
func TestOrgsNetworkTemplatesTestDeleteOrgNetworkTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    networktemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsNetworkTemplates.DeleteOrgNetworkTemplate(ctx, orgId, networktemplateId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsNetworkTemplatesTestGetOrgNetworkTemplate tests the behavior of the OrgsNetworkTemplates
func TestOrgsNetworkTemplatesTestGetOrgNetworkTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    networktemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsNetworkTemplates.GetOrgNetworkTemplate(ctx, orgId, networktemplateId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsNetworkTemplatesTestUpdateOrgNetworkTemplate tests the behavior of the OrgsNetworkTemplates
func TestOrgsNetworkTemplatesTestUpdateOrgNetworkTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    networktemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.NetworkTemplate
    errBody := json.Unmarshal([]byte(`{"additional_config_cmds":["set snmp community public"],"dhcp_snooping":{"all_networks":true,"enable_arp_spoof_check":true,"enable_ip_source_guard":true,"enabled":true,"networks":["string"]},"dns_servers":["string"],"dns_suffix":["string"],"extra_routes":{"property1":{"via":"string"},"property2":{"via":"string"}},"group_tags":{},"import_org_networks":["ap"],"mist_nac":{"enabled":true,"network":"string"},"name":"string","networks":{"property1":{"subnet":"192.168.1.0/24","vlan_id":10},"property2":{"subnet":"192.168.1.0/24","vlan_id":10}},"ntp_servers":["string"],"port_usages":{"dynamic":{"mode":"dynamic","reset_default_when":"link_down","rules":[{"equals":"string","equals_any":["string"],"expression":"string","src":"lldp_chassis_id","usage":"string"}]},"property1":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","networks":["string"],"persist_mac":false,"poe_disabled":false,"port_auth":"dot1x","port_network":"string","rejected_network":null,"speed":"auto","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"},"property2":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","networks":["string"],"persist_mac":false,"poe_disabled":false,"port_network":"string","rejected_network":null,"speed":"auto","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"}},"switch_mgmt":{"config_revert":10,"protect_re":{"enabled":false},"root_password":"string","tacacs":{"acct_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}],"enabled":true,"network":"string","tacplus_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}]}},"vrf_config":{"enabled":false},"vrf_instances":{"property1":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.2"}},"networks":["string"]},"property2":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.2"}},"networks":["string"]}}}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsNetworkTemplates.UpdateOrgNetworkTemplate(ctx, orgId, networktemplateId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"additional_config_cmds":["set snmp community public"],"dhcp_snooping":{"all_networks":true,"enable_arp_spoof_check":true,"enable_ip_source_guard":true,"enabled":true,"networks":["string"]},"dns_servers":["string"],"dns_suffix":["string"],"extra_routes":{"property1":{"via":"string"},"property2":{"via":"string"}},"import_org_networks":["ap"],"mist_nac":{"enabled":true,"network":"string"},"name":"string","networks":{"property1":{"subnet":"192.168.1.0/24","vlan_id":10},"property2":{"subnet":"192.168.1.0/24","vlan_id":10}},"ntp_servers":["string"],"port_usages":{"dynamic":{"mode":"dynamic","reset_default_when":"link_down","rules":[{"equals":"string","equals_any":["string"],"expression":"string","src":"lldp_chassis_id","usage":"string"}]},"property1":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","networks":["string"],"persist_mac":false,"poe_disabled":false,"port_auth":"dot1x","port_network":"string","rejected_network":null,"speed":"auto","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"},"property2":{"all_networks":false,"allow_dhcpd":true,"authentication_protocol":"pap","bypass_auth_when_server_down":true,"description":"string","disable_autoneg":false,"disabled":false,"duplex":"auto","enable_mac_auth":true,"enable_qos":true,"guest_network":"string","mac_auth_only":true,"mac_limit":0,"mode":"access","networks":["string"],"persist_mac":false,"poe_disabled":false,"port_network":"string","rejected_network":null,"speed":"auto","storm_control":{"no_broadcast":false,"no_multicast":false,"no_registered_multicast":false,"no_unknown_unicast":false,"percentage":80},"stp_edge":true,"voip_network":"string"}},"switch_mgmt":{"config_revert":10,"protect_re":{"enabled":false},"root_password":"string","tacacs":{"acct_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}],"enabled":true,"network":"string","tacplus_servers":[{"host":"198.51.100.1","port":"49","secret":"string","timeout":10}]}},"vrf_config":{"enabled":false},"vrf_instances":{"property1":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.2"}},"networks":["string"]},"property2":{"extra_routes":{"property1":{"via":"198.51.100.1"},"property2":{"via":"198.51.100.2"}},"networks":["string"]}}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
