package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestUtilitiesWiFiTestReauthOrgDot1xWirelessClient tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestReauthOrgDot1xWirelessClient(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    clientMac := "0000000000ab"
    resp, err := utilitiesWiFi.ReauthOrgDot1xWirelessClient(ctx, orgId, clientMac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesWiFiTestDisconnectSiteMultipleClients tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestDisconnectSiteMultipleClients(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    resp, err := utilitiesWiFi.DisconnectSiteMultipleClients(ctx, siteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesWiFiTestUnauthorizeSiteMultipleClients tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestUnauthorizeSiteMultipleClients(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    resp, err := utilitiesWiFi.UnauthorizeSiteMultipleClients(ctx, siteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesWiFiTestReauthSiteDot1xWirelessClient tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestReauthSiteDot1xWirelessClient(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    clientMac := "0000000000ab"
    resp, err := utilitiesWiFi.ReauthSiteDot1xWirelessClient(ctx, siteId, clientMac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesWiFiTestDisconnectSiteWirelessClient tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestDisconnectSiteWirelessClient(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    clientMac := "0000000000ab"
    resp, err := utilitiesWiFi.DisconnectSiteWirelessClient(ctx, siteId, clientMac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesWiFiTestUnauthorizeSiteWirelessClient tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestUnauthorizeSiteWirelessClient(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    clientMac := "0000000000ab"
    resp, err := utilitiesWiFi.UnauthorizeSiteWirelessClient(ctx, siteId, clientMac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesWiFiTestReprovisionSiteAllAps tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestReprovisionSiteAllAps(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := utilitiesWiFi.ReprovisionSiteAllAps(ctx, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesWiFiTestResetSiteAllApsToUseRrm tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestResetSiteAllApsToUseRrm(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.UtilsResetRadioConfig
    errBody := json.Unmarshal([]byte(`{"bands":["24","5","6"],"force":false}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := utilitiesWiFi.ResetSiteAllApsToUseRrm(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesWiFiTestZeroizeSiteFipsAllAps tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestZeroizeSiteFipsAllAps(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.UtilsZeroiseFips
    errBody := json.Unmarshal([]byte(`{"password":"NUKETHESITE"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := utilitiesWiFi.ZeroizeSiteFipsAllAps(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesWiFiTestDeauthSiteWirelessClientsConnectedToARogue tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestDeauthSiteWirelessClientsConnectedToARogue(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    rogueBssid := "0000000000ab"
    resp, err := utilitiesWiFi.DeauthSiteWirelessClientsConnectedToARogue(ctx, siteId, rogueBssid)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesWiFiTestOptimizeSiteRrm tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestOptimizeSiteRrm(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.UtilsRrmOptimize
    errBody := json.Unmarshal([]byte(`{"bands":["24","5","6"]}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := utilitiesWiFi.OptimizeSiteRrm(ctx, siteId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesWiFiTestTestSiteWlanTelstraSetup tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestTestSiteWlanTelstraSetup(t *testing.T) {
    ctx := context.Background()
    var body models.TestTelstra
    errBody := json.Unmarshal([]byte(`{"telstra_client_id":"123456","telstra_client_secret":"abcdef","to":"+911122334455"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := utilitiesWiFi.TestSiteWlanTelstraSetup(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestUtilitiesWiFiTestTestSiteWlanTwilioSetup tests the behavior of the UtilitiesWiFi
func TestUtilitiesWiFiTestTestSiteWlanTwilioSetup(t *testing.T) {
    ctx := context.Background()
    var body models.TestTwilio
    errBody := json.Unmarshal([]byte(`{"from":"+185051234567","to":"+19999999999","twilio_auth_token":"2135be04736a1a0a314bce432d61721a","twilio_sid":"AC5f4366878d193fb4865ab151739999eb"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := utilitiesWiFi.TestSiteWlanTwilioSetup(ctx, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
