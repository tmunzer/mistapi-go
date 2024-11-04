package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesWlansTestListSiteWlans tests the behavior of the SitesWlans
func TestSitesWlansTestListSiteWlans(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesWlans.ListSiteWlans(ctx, siteId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesWlansTestCreateSiteWlan tests the behavior of the SitesWlans
func TestSitesWlansTestCreateSiteWlan(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Wlan
    errBody := json.Unmarshal([]byte(`{"acct_immediate_update":false,"acct_interim_interval":0,"acct_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1813,"secret":"testing123"}],"airwatch":{"api_key":"aHhlbGxvYXNkZmFzZGZhc2Rmc2RmCg==\"","console_url":"https://hs1.airwatchportals.com","enabled":true,"password":"user1","username":"test123"},"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"ap_ids":["497f6eca-6276-4993-bfeb-53cbbbbb6f08"],"app_limit":{"apps":{"dropbox":300,"netflix":60},"enabled":false,"wxtag_ids":{"f99862d9-2726-931f-7559-3dfdf5d070d3":30}},"app_qos":{"apps":{"skype-business-video":{"dscp":32,"dst_subnet":"10.2.0.0/16","src_subnet":"10.2.0.0/16"}},"enabled":true,"others":[{"dscp":32,"dst_subnet":"10.2.0.0/16","port_ranges":"80,1024-6553","protocol":"udp","src_subnet":"10.2.0.0/16"}]},"apply_to":"site","arp_filter":false,"auth":{"anticlog_threshold":16,"eap_reauth":false,"enable_mac_auth":false,"key_idx":1,"keys":["string"],"multi_psk_only":false,"pairwise":["wpa2-ccmp"],"private_wlan":true,"psk":"foryoureyesonly","type":"psk","wep_as_secondary_auth":true},"auth_server_selection":"ordered","auth_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1812,"secret":"testing123"}],"auth_servers_nas_id":"5c5b350e0101-nas","auth_servers_nas_ip":"15.3.1.5","auth_servers_retries":5,"auth_servers_timeout":5,"band":"string","band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10,20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"cisco_cwa":{"allowed_hostnames":["snapchat.com"],"allowed_subnets":["63.5.3.0/24"],"blocked_subnets":["192.168.0.0/16"],"enabled":false},"client_limit_down":0,"client_limit_down_enabled":false,"client_limit_up":0,"client_limit_up_enabled":false,"coa_servers":[{"disable_event_timestamp_check":false,"enabled":false,"ip":"1.2.3.4","port":3799,"secret":"testing456"}],"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dns_server_rewrite":{"enabled":false,"radius_groups":{"contractor":"172.1.1.1","guest":"8.8.8.8"}},"dtim":2,"dynamic_psk":{"default_psk":"foryoureyesonly","default_vlan_id":999,"enabled":false,"source":"cloud_psks","vlan_ids":[1]},"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"hotspot20":{"domain_name":["mist.com"],"enabled":true,"nai_realms":["string"],"operators":["google","att"],"rcoi":["5A03BA0000"],"venue_name":"some_name"},"interface":"all","isolation":false,"l2_isolation":false,"legacy_overds":false,"limit_bcast":false,"limit_probe_response":true,"max_idletime":1800,"mist_nac":{"enabled":false},"no_static_dns":false,"no_static_ip":false,"ssid":"demo"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesWlans.CreateSiteWlan(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"arp_filter":false,"band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10, 20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"client_limit_down":0,"client_limit_down_enabled":false,"client_limit_up":0,"client_limit_up_enabled":false,"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"ssid":"demo"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWlansTestListSiteWlanDerived tests the behavior of the SitesWlans
func TestSitesWlansTestListSiteWlanDerived(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resolve := bool(false)
    
    apiResponse, err := sitesWlans.ListSiteWlanDerived(ctx, siteId, &resolve, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestSitesWlansTestDeleteSiteWlan tests the behavior of the SitesWlans
func TestSitesWlansTestDeleteSiteWlan(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesWlans.DeleteSiteWlan(ctx, siteId, wlanId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesWlansTestGetSiteWlan tests the behavior of the SitesWlans
func TestSitesWlansTestGetSiteWlan(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesWlans.GetSiteWlan(ctx, siteId, wlanId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"arp_filter":false,"band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10, 20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"client_limit_down":0,"client_limit_down_enabled":false,"client_limit_up":0,"client_limit_up_enabled":false,"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"ssid":"demo"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWlansTestUpdateSiteWlan tests the behavior of the SitesWlans
func TestSitesWlansTestUpdateSiteWlan(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Wlan
    errBody := json.Unmarshal([]byte(`{"acct_immediate_update":false,"acct_interim_interval":0,"acct_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1813,"secret":"testing123"}],"airwatch":{"api_key":"aHhlbGxvYXNkZmFzZGZhc2Rmc2RmCg==\"","console_url":"https://hs1.airwatchportals.com","enabled":true,"password":"user1","username":"test123"},"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"ap_ids":["497f6eca-6276-4993-bfeb-53cbbbbe6f08"],"app_limit":{"apps":{"dropbox":300,"netflix":60},"enabled":false,"wxtag_ids":{"f99862d9-2726-931f-7559-3dfdf5d070d3":30}},"app_qos":{"apps":{"skype-business-video":{"dscp":32,"dst_subnet":"10.2.0.0/16","src_subnet":"10.2.0.0/16"}},"enabled":true,"others":[{"dscp":32,"dst_subnet":"10.2.0.0/16","port_ranges":"80,1024-6553","protocol":"udp","src_subnet":"10.2.0.0/16"}]},"apply_to":"site","arp_filter":false,"auth":{"anticlog_threshold":16,"eap_reauth":false,"enable_mac_auth":false,"key_idx":1,"keys":["string"],"multi_psk_only":false,"pairwise":["wpa2-ccmp"],"private_wlan":true,"psk":"foryoureyesonly","type":"psk","wep_as_secondary_auth":true},"auth_server_selection":"ordered","auth_servers":[{"host":"1.2.3.4","keywrap_enabled":true,"keywrap_format":"hex","keywrap_kek":"1122334455","keywrap_mack":"1122334455","port":1812,"secret":"testing123"}],"auth_servers_nas_id":"5c5b350e0101-nas","auth_servers_nas_ip":"15.3.1.5","auth_servers_retries":5,"auth_servers_timeout":5,"band":"string","band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10,20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"cisco_cwa":{"allowed_hostnames":["snapchat.com"],"allowed_subnets":["63.5.3.0/24"],"blocked_subnets":["192.168.0.0/16"],"enabled":false},"client_limit_down":0,"client_limit_down_enabled":false,"client_limit_up":0,"client_limit_up_enabled":false,"coa_servers":[{"disable_event_timestamp_check":false,"enabled":false,"ip":"1.2.3.4","port":3799,"secret":"testing456"}],"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dns_server_rewrite":{"enabled":false,"radius_groups":{"contractor":"172.1.1.1","guest":"8.8.8.8"}},"dtim":2,"dynamic_psk":{"default_psk":"foryoureyesonly","default_vlan_id":999,"enabled":false,"source":"cloud_psks","vlan_ids":[1]},"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"hotspot20":{"domain_name":["mist.com"],"enabled":true,"nai_realms":["string"],"operators":["google","att"],"rcoi":["5A03BA0000"],"venue_name":"some_name"},"interface":"all","isolation":false,"l2_isolation":false,"legacy_overds":false,"limit_bcast":false,"limit_probe_response":true,"max_idletime":1800,"mist_nac":{"enabled":false},"no_static_dns":false,"no_static_ip":false,"ssid":"demo"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesWlans.UpdateSiteWlan(ctx, siteId, wlanId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"arp_filter":false,"band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10, 20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"client_limit_down":0,"client_limit_down_enabled":false,"client_limit_up":0,"client_limit_up_enabled":false,"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"ssid":"demo"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesWlansTestDeleteSiteWlanPortalImage tests the behavior of the SitesWlans
func TestSitesWlansTestDeleteSiteWlanPortalImage(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesWlans.DeleteSiteWlanPortalImage(ctx, siteId, wlanId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesWlansTestUpdateSiteWlanPortalTemplate tests the behavior of the SitesWlans
func TestSitesWlansTestUpdateSiteWlanPortalTemplate(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.WlanPortalTemplate
    errBody := json.Unmarshal([]byte(`{"accessCodeAlternateEmail":"string","alignment":"string","authButtonAmazon":"string","authButtonAzure":"string","authButtonEmail":"string","authButtonFacebook":"string","authButtonGoogle":"string","authButtonMicrosoft":"string","authButtonPassphrase":"string","authButtonSms":"string","authButtonSponsor":"string","authLabel":"string","backLink":"string","color":"string","colorDark":"string","colorLight":"string","company":true,"companyError":"string","companyLabel":"string","email":true,"emailAccessDomainError":"string","emailCancel":"string","emailCodeError":"string","emailError":"string","emailFieldLabel":"string","emailLabel":"string","emailMessage":"string","emailSubmit":"string","emailTitle":"string","field1":true,"field1Error":"string","field1Label":"string","field1Required":true,"field2":true,"field2Error":"string","field2Label":"string","field2Required":true,"field3":true,"field3Error":"string","field3Label":"string","field3Required":true,"field4":true,"field4Error":"string","field4Label":"string","field4Required":true,"message":"string","name":true,"nameError":"string","nameLabel":"string","optout":true,"optoutLabel":"string","pageTitle":"string","passphraseCancel":"string","passphraseError":"string","passphraseLabel":"string","passphraseMessage":"string","passphraseSubmit":"string","passphraseTitle":"string","poweredBy":true,"requiredFieldLabel":"string","signInLabel":"string","smsCarrierDefault":"string","smsCarrierError":"string","smsCarrierFieldLabel":"string","smsCodeCancel":"string","smsCodeError":"string","smsCodeFieldLabel":"string","smsCodeMessage":"string","smsCodeSubmit":"string","smsCodeTitle":"string","smsCountryFieldLabel":"string","smsCountryFormat":"string","smsHaveAccessCode":"string","smsMessageFormat":"string","smsNumberCancel":"string","smsNumberError":"string","smsNumberFieldLabel":"string","smsNumberFormat":"string","smsNumberMessage":"string","smsNumberSubmit":"string","smsNumberTitle":"string","smsUsernameFormat":"string","smsValidityDuration":0,"sponsorBackLink":"string","sponsorCancel":"string","sponsorEmail":"string","sponsorEmailError":"string","sponsorEmailTemplate":"string","sponsorInfoApproved":"string","sponsorInfoDenied":"string","sponsorInfoPending":"string","sponsorName":"string","sponsorNameError":"string","sponsorNotePending":"string","sponsorStatusApproved":"string","sponsorStatusDenied":"string","sponsorStatusPending":"string","sponsorSubmit":"string","tos":true,"tosAcceptLabel":"string","tosError":"string","tosLink":"string","tosText":"string"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := sitesWlans.UpdateSiteWlanPortalTemplate(ctx, siteId, wlanId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"portal_template":{"accessCodeAlternateEmail":"string","alignment":"left","authButtonAmazon":"string","authButtonAzure":"string","authButtonEmail":"string","authButtonFacebook":"string","authButtonGoogle":"string","authButtonMicrosoft":"string","authButtonPassphrase":"string","authButtonSms":"string","authButtonSponsor":"string","authLabel":"string","backLink":"string","color":"string","colorDark":"string","colorLight":"string","company":true,"companyError":"string","companyLabel":"string","email":true,"emailAccessDomainError":"string","emailCancel":"string","emailCodeError":"string","emailError":"string","emailFieldLabel":"string","emailLabel":"string","emailMessage":"string","emailSubmit":"string","emailTitle":"string","field1":true,"field1Error":"string","field1Label":"string","field1Required":true,"field2":true,"field2Error":"string","field2Label":"string","field2Required":true,"field3":true,"field3Error":"string","field3Label":"string","field3Required":true,"field4":true,"field4Error":"string","field4Label":"string","field4Required":true,"message":"string","name":true,"nameError":"string","nameLabel":"string","optout":true,"optoutLabel":"string","pageTitle":"string","passphraseCancel":"string","passphraseError":"string","passphraseLabel":"string","passphraseMessage":"string","passphraseSubmit":"string","passphraseTitle":"string","poweredBy":true,"requiredFieldLabel":"string","signInLabel":"string","smsCarrierDefault":"string","smsCarrierError":"string","smsCarrierFieldLabel":"string","smsCodeCancel":"string","smsCodeError":"string","smsCodeFieldLabel":"string","smsCodeMessage":"string","smsCodeSubmit":"string","smsCodeTitle":"string","smsCountryFieldLabel":"string","smsCountryFormat":"string","smsHaveAccessCode":"string","smsMessageFormat":"string","smsNumberCancel":"string","smsNumberError":"string","smsNumberFieldLabel":"string","smsNumberFormat":"string","smsNumberMessage":"string","smsNumberSubmit":"string","smsNumberTitle":"string","smsUsernameFormat":"string","smsValidityDuration":5,"sponsorBackLink":"string","sponsorCancel":"string","sponsorEmail":"string","sponsorEmailError":"string","sponsorEmailTemplate":"string","sponsorInfoApproved":"string","sponsorInfoDenied":"string","sponsorInfoPending":"string","sponsorName":"string","sponsorNameError":"string","sponsorNotePending":"string","sponsorStatusApproved":"string","sponsorStatusDenied":"string","sponsorStatusPending":"string","sponsorSubmit":"string","tos":true,"tosAcceptLabel":"string","tosError":"string","tosLink":"string","tosText":"string"}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
