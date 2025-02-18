package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesGuestsTestListSiteAllGuestAuthorizations tests the behavior of the SitesGuests
func TestSitesGuestsTestListSiteAllGuestAuthorizations(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := sitesGuests.ListSiteAllGuestAuthorizations(ctx, siteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"authorized":true,"authorized_expiring_time":0,"authorized_time":0,"company":"string","email":"user@example.com","field1":"string","field2":"string","field3":"string","field4":"string","mac":"string","minutes":0,"name":"string"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesGuestsTestCountSiteGuestAuthorizations tests the behavior of the SitesGuests
func TestSitesGuestsTestCountSiteGuestAuthorizations(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.SiteGuestsCountDistinctEnum("auth_method")
    
    
    duration := "1d"
    limit := int(100)
    page := int(1)
    apiResponse, err := sitesGuests.CountSiteGuestAuthorizations(ctx, siteId, &distinct, nil, nil, &duration, &limit, &page)
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

// TestSitesGuestsTestListSiteAllGuestAuthorizationsDerived tests the behavior of the SitesGuests
func TestSitesGuestsTestListSiteAllGuestAuthorizationsDerived(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    crossSite := bool(false)
    apiResponse, err := sitesGuests.ListSiteAllGuestAuthorizationsDerived(ctx, siteId, nil, &crossSite)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"authorized":true,"authorized_expiring_time":0,"authorized_time":0,"company":"string","email":"user@example.com","field1":"string","field2":"string","field3":"string","field4":"string","mac":"string","minutes":0,"name":"string"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesGuestsTestSearchSiteGuestAuthorization tests the behavior of the SitesGuests
func TestSitesGuestsTestSearchSiteGuestAuthorization(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    wlanId := "00000000-0000-0000-0000-000000000000"
    
    
    limit := int(100)
    
    
    duration := "1d"
    apiResponse, err := sitesGuests.SearchSiteGuestAuthorization(ctx, siteId, &wlanId, nil, nil, &limit, nil, nil, &duration)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1531862583,"limit":2,"next":"/api/v1/sites/8aaba0aa-09cc-44bd-9709-33b98040550c/guests/search?wlan_id=88ffe630-95b8-11e8-b294-346895ed1b7d&end=1531855849.000&limit=2&start=1531776183.0","results":[{"ap":"5c5b350e0001","auth_method":"passphrase","authorized_expiring_time":1531810258.186273,"authorized_time":1531782218,"company":"mistsystems","email":"user@mistsys.com","name":"john","ssid":"openNet","timestamp":1531782218},{"ap":"5c5b350e0001","auth_method":"facebook","authorized_expiring_time":1531810821.145,"authorized_time":1531782632,"company":"xyz inc.","email":"cool_user@yahoo.com","name":"John White","ssid":"openNet","timestamp":1531782632}],"start":1531776183,"total":14}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesGuestsTestDeleteSiteGuestAuthorization tests the behavior of the SitesGuests
func TestSitesGuestsTestDeleteSiteGuestAuthorization(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    guestMac := "0000000000ab"
    resp, err := sitesGuests.DeleteSiteGuestAuthorization(ctx, siteId, guestMac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesGuestsTestGetSiteGuestAuthorization tests the behavior of the SitesGuests
func TestSitesGuestsTestGetSiteGuestAuthorization(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    guestMac := "0000000000ab"
    apiResponse, err := sitesGuests.GetSiteGuestAuthorization(ctx, siteId, guestMac)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"authorized":true,"authorized_expiring_time":0,"authorized_time":0,"company":"string","email":"user@example.com","field1":"string","field2":"string","field3":"string","field4":"string","mac":"string","minutes":0,"name":"string"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesGuestsTestUpdateSiteGuestAuthorization tests the behavior of the SitesGuests
func TestSitesGuestsTestUpdateSiteGuestAuthorization(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    guestMac := "0000000000ab"
    
    apiResponse, err := sitesGuests.UpdateSiteGuestAuthorization(ctx, siteId, guestMac, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"authorized":true,"authorized_expiring_time":0,"authorized_time":0,"company":"string","email":"user@example.com","field1":"string","field2":"string","field3":"string","field4":"string","mac":"string","minutes":0,"name":"string"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
