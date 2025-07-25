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

// TestOrgsSecurityPoliciesTestListOrgSecPolicies tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestListOrgSecPolicies(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsSecurityPolicies.ListOrgSecPolicies(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"name":"corporate only","wlans":[{"auth":{"pairwise":["wpa1-tkip","wpa2-tkip"],"type":"psk"},"band":"both","ssid":"office"},{"auth":{"type":"open"},"band":"5","ssid":"office-guest"}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityPoliciesTestListOrgSecPolicies1 tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestListOrgSecPolicies1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsSecurityPolicies.ListOrgSecPolicies(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"name":"corporate only","wlans":[{"auth":{"pairwise":["wpa1-tkip","wpa2-tkip"],"type":"psk"},"band":"both","ssid":"office"},{"auth":{"type":"open"},"band":"5","ssid":"office-guest"}]}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityPoliciesTestCreateOrgSecPolicy tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestCreateOrgSecPolicy(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Secpolicy
    errBody := json.Unmarshal([]byte(`{"name":"string","wlans":[{"acct_immediate_update":false,"acct_interim_interval":0,"acct_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1813,"secret":"testing123"}],"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"app_limit":{"apps":{"dropbox":300,"netflix":60},"enabled":false,"wxtag_ids":{"f99862d9-2726-931f-7559-3dfdf5d070d3":30}},"app_qos":{"apps":{"skype-business-video":{"dscp":32,"dst_subnet":"10.2.0.0/16","src_subnet":"10.2.0.0/16"}},"enabled":true,"others":[{"dscp":32,"dst_subnet":"10.2.0.0/16","port_ranges":"80,1024-6553","protocol":"udp","src_subnet":"10.2.0.0/16"}]},"arp_filter":false,"auth":{"anticlog_threshold":16,"eap_reauth":false,"enable_mac_auth":false,"key_idx":1,"keys":["string"],"multi_psk_only":false,"pairwise":["wpa2-ccmp"],"private_wlan":true,"psk":"foryoureyesonly","type":"psk","wep_as_secondary_auth":true},"auth_server_selection":"ordered","auth_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1812,"secret":"testing123"}],"auth_servers_nas_id":"5c5b350e0101-nas","auth_servers_nas_ip":"15.3.1.5","auth_servers_retries":5,"auth_servers_timeout":5,"band":"string","band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10,20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"cisco_cwa":{"allowed_hostnames":["snapchat.com"],"allowed_subnets":["63.5.3.0/24"],"blocked_subnets":["192.168.0.0/16"],"enabled":false},"client_limit_down":1000,"client_limit_down_enabled":false,"client_limit_up":512,"client_limit_up_enabled":false,"coa_servers":[{"disable_event_timestamp_check":false,"enabled":false,"ip":"1.2.3.4","port":3799,"secret":"testing456"}],"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dns_server_rewrite":{"enabled":false,"radius_groups":{"contractor":"172.1.1.1","guest":"8.8.8.8"}},"dtim":2,"dynamic_psk":{"default_psk":"foryoureyesonly","default_vlan_id":999,"enabled":false,"source":"cloud_psks"},"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"hotspot20":{"domain_name":["mist.com"],"enabled":true,"nai_realms":["string"],"operators":["google","att"],"rcoi":["5A03BA0000"],"venue_name":"some_name"},"interface":"all","isolation":false,"l2_isolation":false,"legacy_overds":false,"limit_bcast":false,"limit_probe_response":true,"max_idletime":1800,"mist_nac":{"enabled":false},"no_static_dns":false,"no_static_ip":false,"portal":{"amazon_client_id":"string","amazon_client_secret":"string","amazon_email_domains":["string"],"amazon_enabled":false,"auth":"none","azure_client_id":"string","azure_client_secret":"string","azure_enabled":false,"azure_tenant_id":"string","broadnet_password":"password","broadnet_sid":"MIST","broadnet_user_id":"juniper","bypass_when_cloud_down":false,"clickatell_api_key":"string","cross_site":false,"email_enabled":true,"enabled":false,"expire":1440,"external_portal_url":"string","facebook_client_id":"string","facebook_client_secret":"string","facebook_email_domains":["string"],"facebook_enabled":false,"forward":false,"forward_url":"https://abc.com/promotions","google_client_id":"string","google_client_secret":"string","google_email_domains":["mydomain.edu","mydomain.org"],"google_enabled":false,"gupshup_password":"string","gupshup_userid":"string","microsoft_client_id":"string","microsoft_client_secret":"string","microsoft_email_domains":["string"],"microsoft_enabled":false,"passphrase_enabled":false,"password":"let me in","portal_api_secret":"string","portal_image":"https://url/to/image.png","predefined_sponsors_enabled":true,"privacy":true,"puzzel_password":"string","puzzel_service_id":"string","puzzel_username":"string","smsMessageFormat":"string","sms_enabled":false,"sms_provider":"twilio","sponsor_auto_approve":false,"sponsor_email_domains":["reserved.net","reserved.org"],"sponsor_enabled":false,"sponsor_link_validity_duration":"30","sponsor_notify_all":false,"sponsor_status_notify":false,"sponsors":{"sponsor1@company.com":"FirstName1 LastName1","sponsor2@company.com":"FirstName2 LastName2"},"sso_default_role":"string","sso_forced_role":"string","sso_idp_cert":"string","sso_idp_sign_algo":"sha256","sso_idp_sso_url":"string","sso_issuer":"string","sso_nameid_format":"email","telstra_client_id":"string","telstra_client_secret":"string","thumbnail":"string","twilio_auth_token":"af9dac44c344a875ab5d31cb7abcdefg","twilio_phone_number":"+18548888888","twilio_sid":"AC72ec6ba0ec5af30e6731c5e47abcdefgh"},"portal_allowed_hostnames":["snapchat.com","ibm.com"],"portal_allowed_subnets":["63.5.3.0/24"],"portal_api_secret":"EIfPMOykI3lMlDdNPub2WcbqT6dNOtWwmYHAd6bY","portal_denied_hostnames":["msg.snapchat.com"],"portal_image":"https://example.com","portal_template_url":"string","qos":{"class":"best_effort","overwrite":false},"radsec":{"enabled":true,"idle_timeout":60,"mxcluster_ids":["572586b7-f97b-a22b-526c-8b97a3f609c4"],"proxy_hosts":["mxedge1.local"],"server_name":"radsec.abc.com","servers":[{"host":"1.1.1.1","port":1812}],"use_mxedge":true,"use_site_mxedge":false},"rateset":{"24":{"ht":"00ff00ff00ff","legacy":["6","9","12","18","24b","36","48","54"],"min_rssi":-70,"template":"custom","vht":"03ff03ff03ff01ff"},"5":{"ht":"00ff00ff00ff","legacy":["6","9","12","18","24b","36","48","54"],"min_rssi":-70,"template":"custom","vht":"03ff03ff03ff01ff"}},"roam_mode":"NONE","schedule":{"enabled":false,"hours":{"fri":"09:00-17:00","mon":"09:00-17:00"}},"sle_excluded":false,"ssid":"corporate","template_id":"c6d67e98-83ea-49f0-8812-e4abae2b68bc","thumbnail":"https://example.com","use_eapol_v1":false,"vlan_enabled":false,"vlan_ids":[3,4,5],"vlan_pooling":false,"wxtag_ids":["497f6eca-6276-4993-bfeb-53e4bbba6f08"],"wxtunnel_id":"string","wxtunnel_remote_id":"string"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsSecurityPolicies.CreateOrgSecPolicy(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"name":"corporate only","wlans":[{"auth":{"pairwise":["wpa1-tkip","wpa2-tkip"],"type":"psk"},"band":"both","ssid":"office"},{"auth":{"type":"open"},"band":"5","ssid":"office-guest"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityPoliciesTestCreateOrgSecPolicy1 tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestCreateOrgSecPolicy1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Secpolicy
    errBody := json.Unmarshal([]byte(`{"name":"string","wlans":[{"acct_immediate_update":false,"acct_interim_interval":0,"acct_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1813,"secret":"testing123"}],"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"app_limit":{"apps":{"dropbox":300,"netflix":60},"enabled":false,"wxtag_ids":{"f99862d9-2726-931f-7559-3dfdf5d070d3":30}},"app_qos":{"apps":{"skype-business-video":{"dscp":32,"dst_subnet":"10.2.0.0/16","src_subnet":"10.2.0.0/16"}},"enabled":true,"others":[{"dscp":32,"dst_subnet":"10.2.0.0/16","port_ranges":"80,1024-6553","protocol":"udp","src_subnet":"10.2.0.0/16"}]},"arp_filter":false,"auth":{"anticlog_threshold":16,"eap_reauth":false,"enable_mac_auth":false,"key_idx":1,"keys":["string"],"multi_psk_only":false,"pairwise":["wpa2-ccmp"],"private_wlan":true,"psk":"foryoureyesonly","type":"psk","wep_as_secondary_auth":true},"auth_server_selection":"ordered","auth_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1812,"secret":"testing123"}],"auth_servers_nas_id":"5c5b350e0101-nas","auth_servers_nas_ip":"15.3.1.5","auth_servers_retries":5,"auth_servers_timeout":5,"band":"string","band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10,20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"cisco_cwa":{"allowed_hostnames":["snapchat.com"],"allowed_subnets":["63.5.3.0/24"],"blocked_subnets":["192.168.0.0/16"],"enabled":false},"client_limit_down":1000,"client_limit_down_enabled":false,"client_limit_up":512,"client_limit_up_enabled":false,"coa_servers":[{"disable_event_timestamp_check":false,"enabled":false,"ip":"1.2.3.4","port":3799,"secret":"testing456"}],"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dns_server_rewrite":{"enabled":false,"radius_groups":{"contractor":"172.1.1.1","guest":"8.8.8.8"}},"dtim":2,"dynamic_psk":{"default_psk":"foryoureyesonly","default_vlan_id":999,"enabled":false,"source":"cloud_psks"},"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"hotspot20":{"domain_name":["mist.com"],"enabled":true,"nai_realms":["string"],"operators":["google","att"],"rcoi":["5A03BA0000"],"venue_name":"some_name"},"interface":"all","isolation":false,"l2_isolation":false,"legacy_overds":false,"limit_bcast":false,"limit_probe_response":true,"max_idletime":1800,"mist_nac":{"enabled":false},"no_static_dns":false,"no_static_ip":false,"portal":{"amazon_client_id":"string","amazon_client_secret":"string","amazon_email_domains":["string"],"amazon_enabled":false,"auth":"none","azure_client_id":"string","azure_client_secret":"string","azure_enabled":false,"azure_tenant_id":"string","broadnet_password":"password","broadnet_sid":"MIST","broadnet_user_id":"juniper","bypass_when_cloud_down":false,"clickatell_api_key":"string","cross_site":false,"email_enabled":true,"enabled":false,"expire":1440,"external_portal_url":"string","facebook_client_id":"string","facebook_client_secret":"string","facebook_email_domains":["string"],"facebook_enabled":false,"forward":false,"forward_url":"https://abc.com/promotions","google_client_id":"string","google_client_secret":"string","google_email_domains":["mydomain.edu","mydomain.org"],"google_enabled":false,"gupshup_password":"string","gupshup_userid":"string","microsoft_client_id":"string","microsoft_client_secret":"string","microsoft_email_domains":["string"],"microsoft_enabled":false,"passphrase_enabled":false,"password":"let me in","portal_api_secret":"string","portal_image":"https://url/to/image.png","predefined_sponsors_enabled":true,"privacy":true,"puzzel_password":"string","puzzel_service_id":"string","puzzel_username":"string","smsMessageFormat":"string","sms_enabled":false,"sms_provider":"twilio","sponsor_auto_approve":false,"sponsor_email_domains":["reserved.net","reserved.org"],"sponsor_enabled":false,"sponsor_link_validity_duration":"30","sponsor_notify_all":false,"sponsor_status_notify":false,"sponsors":{"sponsor1@company.com":"FirstName1 LastName1","sponsor2@company.com":"FirstName2 LastName2"},"sso_default_role":"string","sso_forced_role":"string","sso_idp_cert":"string","sso_idp_sign_algo":"sha256","sso_idp_sso_url":"string","sso_issuer":"string","sso_nameid_format":"email","telstra_client_id":"string","telstra_client_secret":"string","thumbnail":"string","twilio_auth_token":"af9dac44c344a875ab5d31cb7abcdefg","twilio_phone_number":"+18548888888","twilio_sid":"AC72ec6ba0ec5af30e6731c5e47abcdefgh"},"portal_allowed_hostnames":["snapchat.com","ibm.com"],"portal_allowed_subnets":["63.5.3.0/24"],"portal_api_secret":"EIfPMOykI3lMlDdNPub2WcbqT6dNOtWwmYHAd6bY","portal_denied_hostnames":["msg.snapchat.com"],"portal_image":"https://example.com","portal_template_url":"string","qos":{"class":"best_effort","overwrite":false},"radsec":{"enabled":true,"idle_timeout":60,"mxcluster_ids":["572586b7-f97b-a22b-526c-8b97a3f609c4"],"proxy_hosts":["mxedge1.local"],"server_name":"radsec.abc.com","servers":[{"host":"1.1.1.1","port":1812}],"use_mxedge":true,"use_site_mxedge":false},"rateset":{"24":{"ht":"00ff00ff00ff","legacy":["6","9","12","18","24b","36","48","54"],"min_rssi":-70,"template":"custom","vht":"03ff03ff03ff01ff"},"5":{"ht":"00ff00ff00ff","legacy":["6","9","12","18","24b","36","48","54"],"min_rssi":-70,"template":"custom","vht":"03ff03ff03ff01ff"}},"roam_mode":"NONE","schedule":{"enabled":false,"hours":{"fri":"09:00-17:00","mon":"09:00-17:00"}},"sle_excluded":false,"ssid":"corporate","template_id":"c6d67e98-83ea-49f0-8812-e4abae2b68bc","thumbnail":"https://example.com","use_eapol_v1":false,"vlan_enabled":false,"vlan_ids":[3,4,5],"vlan_pooling":false,"wxtag_ids":["497f6eca-6276-4993-bfeb-53e4bbba6f08"],"wxtunnel_id":"string","wxtunnel_remote_id":"string"}]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsSecurityPolicies.CreateOrgSecPolicy(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"name":"corporate only","wlans":[{"auth":{"pairwise":["wpa1-tkip","wpa2-tkip"],"type":"psk"},"band":"both","ssid":"office"},{"auth":{"type":"open"},"band":"5","ssid":"office-guest"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityPoliciesTestDeleteOrgSecPolicy tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestDeleteOrgSecPolicy(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    secpolicyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsSecurityPolicies.DeleteOrgSecPolicy(ctx, orgId, secpolicyId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSecurityPoliciesTestGetOrgSecPolicy tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestGetOrgSecPolicy(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    secpolicyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSecurityPolicies.GetOrgSecPolicy(ctx, orgId, secpolicyId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"name":"corporate only","wlans":[{"auth":{"pairwise":["wpa1-tkip","wpa2-tkip"],"type":"psk"},"band":"both","ssid":"office"},{"auth":{"type":"open"},"band":"5","ssid":"office-guest"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityPoliciesTestGetOrgSecPolicy1 tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestGetOrgSecPolicy1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    secpolicyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSecurityPolicies.GetOrgSecPolicy(ctx, orgId, secpolicyId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"name":"corporate only","wlans":[{"auth":{"pairwise":["wpa1-tkip","wpa2-tkip"],"type":"psk"},"band":"both","ssid":"office"},{"auth":{"type":"open"},"band":"5","ssid":"office-guest"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityPoliciesTestUpdateOrgSecPolicy tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestUpdateOrgSecPolicy(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    secpolicyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSecurityPolicies.UpdateOrgSecPolicy(ctx, orgId, secpolicyId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"name":"corporate only","wlans":[{"auth":{"pairwise":["wpa1-tkip","wpa2-tkip"],"type":"psk"},"band":"both","ssid":"office"},{"auth":{"type":"open"},"band":"5","ssid":"office-guest"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSecurityPoliciesTestUpdateOrgSecPolicy1 tests the behavior of the OrgsSecurityPolicies
func TestOrgsSecurityPoliciesTestUpdateOrgSecPolicy1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    secpolicyId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSecurityPolicies.UpdateOrgSecPolicy(ctx, orgId, secpolicyId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"name":"corporate only","wlans":[{"auth":{"pairwise":["wpa1-tkip","wpa2-tkip"],"type":"psk"},"band":"both","ssid":"office"},{"auth":{"type":"open"},"band":"5","ssid":"office-guest"}]}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
