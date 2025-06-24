package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsWlansTestListOrgWlans tests the behavior of the OrgsWlans
func TestOrgsWlansTestListOrgWlans(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsWlans.ListOrgWlans(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsWlansTestListOrgWlans1 tests the behavior of the OrgsWlans
func TestOrgsWlansTestListOrgWlans1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsWlans.ListOrgWlans(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

// TestOrgsWlansTestCreateOrgWlan tests the behavior of the OrgsWlans
func TestOrgsWlansTestCreateOrgWlan(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsWlans.CreateOrgWlan(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"arp_filter":false,"band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10, 20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"client_limit_down":1000,"client_limit_down_enabled":false,"client_limit_up":512,"client_limit_up_enabled":false,"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"ssid":"demo"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWlansTestCreateOrgWlan1 tests the behavior of the OrgsWlans
func TestOrgsWlansTestCreateOrgWlan1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsWlans.CreateOrgWlan(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"arp_filter":false,"band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10, 20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"client_limit_down":1000,"client_limit_down_enabled":false,"client_limit_up":512,"client_limit_up_enabled":false,"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"ssid":"demo"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWlansTestDeleteOrgWlan tests the behavior of the OrgsWlans
func TestOrgsWlansTestDeleteOrgWlan(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsWlans.DeleteOrgWlan(ctx, orgId, wlanId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsWlansTestGetOrgWLAN tests the behavior of the OrgsWlans
func TestOrgsWlansTestGetOrgWLAN(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsWlans.GetOrgWLAN(ctx, orgId, wlanId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"arp_filter":false,"band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10, 20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"client_limit_down":1000,"client_limit_down_enabled":false,"client_limit_up":512,"client_limit_up_enabled":false,"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"ssid":"demo"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWlansTestGetOrgWLAN1 tests the behavior of the OrgsWlans
func TestOrgsWlansTestGetOrgWLAN1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsWlans.GetOrgWLAN(ctx, orgId, wlanId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"arp_filter":false,"band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10, 20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"client_limit_down":1000,"client_limit_down_enabled":false,"client_limit_up":512,"client_limit_up_enabled":false,"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"ssid":"demo"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWlansTestUpdateOrgWlan tests the behavior of the OrgsWlans
func TestOrgsWlansTestUpdateOrgWlan(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Wlan
    errBody := json.Unmarshal([]byte(`{"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"arp_filter":false,"band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10,20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"client_limit_down":1000,"client_limit_down_enabled":false,"client_limit_up":512,"client_limit_up_enabled":false,"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"ssid":"demo"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsWlans.UpdateOrgWlan(ctx, orgId, wlanId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"arp_filter":false,"band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10, 20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"client_limit_down":1000,"client_limit_down_enabled":false,"client_limit_up":512,"client_limit_up_enabled":false,"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"ssid":"demo"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWlansTestUpdateOrgWlan1 tests the behavior of the OrgsWlans
func TestOrgsWlansTestUpdateOrgWlan1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.Wlan
    errBody := json.Unmarshal([]byte(`{"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"arp_filter":false,"band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10,20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"client_limit_down":1000,"client_limit_down_enabled":false,"client_limit_up":512,"client_limit_up_enabled":false,"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"ssid":"demo"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    apiResponse, err := orgsWlans.UpdateOrgWlan(ctx, orgId, wlanId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"allow_ipv6_ndp":true,"allow_mdns":false,"allow_ssdp":false,"arp_filter":false,"band_steer":false,"band_steer_force_band5":false,"bands":["24","5"],"block_blacklist_clients":false,"bonjour":{"additional_vlan_ids":"10, 20","enabled":false,"services":{"airplay":{"radius_groups":["teachers"],"scope":"same_ap"}}},"client_limit_down":1000,"client_limit_down_enabled":false,"client_limit_up":512,"client_limit_up_enabled":false,"disable_11ax":false,"disable_ht_vht_rates":false,"disable_uapsd":false,"disable_v1_roam_notify":false,"disable_v2_roam_notify":false,"disable_wmm":false,"dynamic_vlan":{"default_vlan_id":999,"enabled":false,"local_vlan_ids":[1],"type":"airespace-interface-name","vlans":{"131":"default","322":"fast,video"}},"enable_local_keycaching":false,"enable_wireless_bridging":false,"enabled":true,"fast_dot1x_timers":false,"hide_ssid":false,"hostname_ie":false,"ssid":"demo"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWlansTestDeleteOrgWlanPortalImage tests the behavior of the OrgsWlans
func TestOrgsWlansTestDeleteOrgWlanPortalImage(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsWlans.DeleteOrgWlanPortalImage(ctx, orgId, wlanId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsWlansTestUpdateOrgWlanPortalTemplate tests the behavior of the OrgsWlans
func TestOrgsWlansTestUpdateOrgWlanPortalTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsWlans.UpdateOrgWlanPortalTemplate(ctx, orgId, wlanId, nil)
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

// TestOrgsWlansTestUpdateOrgWlanPortalTemplate1 tests the behavior of the OrgsWlans
func TestOrgsWlansTestUpdateOrgWlanPortalTemplate1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsWlans.UpdateOrgWlanPortalTemplate(ctx, orgId, wlanId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"portal_template":{"accessCodeAlternateEmail":"string","alignment":"left","authButtonAmazon":"string","authButtonAzure":"string","authButtonEmail":"string","authButtonFacebook":"string","authButtonGoogle":"string","authButtonMicrosoft":"string","authButtonPassphrase":"string","authButtonSms":"string","authButtonSponsor":"string","authLabel":"string","backLink":"string","color":"string","colorDark":"string","colorLight":"string","company":true,"companyError":"string","companyLabel":"string","email":true,"emailAccessDomainError":"string","emailCancel":"string","emailCodeError":"string","emailError":"string","emailFieldLabel":"string","emailLabel":"string","emailMessage":"string","emailSubmit":"string","emailTitle":"string","field1":true,"field1Error":"string","field1Label":"string","field1Required":true,"field2":true,"field2Error":"string","field2Label":"string","field2Required":true,"field3":true,"field3Error":"string","field3Label":"string","field3Required":true,"field4":true,"field4Error":"string","field4Label":"string","field4Required":true,"message":"string","name":true,"nameError":"string","nameLabel":"string","optout":true,"optoutLabel":"string","pageTitle":"string","passphraseCancel":"string","passphraseError":"string","passphraseLabel":"string","passphraseMessage":"string","passphraseSubmit":"string","passphraseTitle":"string","poweredBy":true,"requiredFieldLabel":"string","signInLabel":"string","smsCarrierDefault":"string","smsCarrierError":"string","smsCarrierFieldLabel":"string","smsCodeCancel":"string","smsCodeError":"string","smsCodeFieldLabel":"string","smsCodeMessage":"string","smsCodeSubmit":"string","smsCodeTitle":"string","smsCountryFieldLabel":"string","smsCountryFormat":"string","smsHaveAccessCode":"string","smsMessageFormat":"string","smsNumberCancel":"string","smsNumberError":"string","smsNumberFieldLabel":"string","smsNumberFormat":"string","smsNumberMessage":"string","smsNumberSubmit":"string","smsNumberTitle":"string","smsUsernameFormat":"string","smsValidityDuration":5,"sponsorBackLink":"string","sponsorCancel":"string","sponsorEmail":"string","sponsorEmailError":"string","sponsorEmailTemplate":"string","sponsorInfoApproved":"string","sponsorInfoDenied":"string","sponsorInfoPending":"string","sponsorName":"string","sponsorNameError":"string","sponsorNotePending":"string","sponsorStatusApproved":"string","sponsorStatusDenied":"string","sponsorStatusPending":"string","sponsorSubmit":"string","tos":true,"tosAcceptLabel":"string","tosError":"string","tosLink":"string","tosText":"string"}}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
