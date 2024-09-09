package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestSitesMapsAutoPlacementTestDeleteSiteApAutoOrientation tests the behavior of the SitesMapsAutoPlacement
func TestSitesMapsAutoPlacementTestDeleteSiteApAutoOrientation(t *testing.T) {
    ctx := context.Background()
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesMapsAutoPlacement.DeleteSiteApAutoOrientation(ctx, mapId, siteId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesMapsAutoPlacementTestStartSiteApAutoOrientation tests the behavior of the SitesMapsAutoPlacement
func TestSitesMapsAutoPlacementTestStartSiteApAutoOrientation(t *testing.T) {
    ctx := context.Background()
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := sitesMapsAutoPlacement.StartSiteApAutoOrientation(ctx, mapId, siteId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"state":"Not Started","time_queued":1675414259}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMapsAutoPlacementTestDeleteSiteApAutoplacement tests the behavior of the SitesMapsAutoPlacement
func TestSitesMapsAutoPlacementTestDeleteSiteApAutoplacement(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := sitesMapsAutoPlacement.DeleteSiteApAutoplacement(ctx, siteId, mapId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesMapsAutoPlacementTestGetSiteApAutoPlacement tests the behavior of the SitesMapsAutoPlacement
func TestSitesMapsAutoPlacementTestGetSiteApAutoPlacement(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := sitesMapsAutoPlacement.GetSiteApAutoPlacement(ctx, siteId, mapId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end_time":1678900362,"start_time":1678900062,"status":"done"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSitesMapsAutoPlacementTestRunSiteApAutoplacement tests the behavior of the SitesMapsAutoPlacement
func TestSitesMapsAutoPlacementTestRunSiteApAutoplacement(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    resp, err := sitesMapsAutoPlacement.RunSiteApAutoplacement(ctx, siteId, mapId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesMapsAutoPlacementTestClearSiteApAutoOrient tests the behavior of the SitesMapsAutoPlacement
func TestSitesMapsAutoPlacementTestClearSiteApAutoOrient(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    resp, err := sitesMapsAutoPlacement.ClearSiteApAutoOrient(ctx, siteId, mapId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesMapsAutoPlacementTestClearSiteApAutoplacement tests the behavior of the SitesMapsAutoPlacement
func TestSitesMapsAutoPlacementTestClearSiteApAutoplacement(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    resp, err := sitesMapsAutoPlacement.ClearSiteApAutoplacement(ctx, siteId, mapId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSitesMapsAutoPlacementTestConfirmSiteApLocalizationData tests the behavior of the SitesMapsAutoPlacement
func TestSitesMapsAutoPlacementTestConfirmSiteApLocalizationData(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    mapId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.UseAutoApValues
    errBody := json.Unmarshal([]byte(`{"accept":false,"device_macs":["string"],"for":"placement"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := sitesMapsAutoPlacement.ConfirmSiteApLocalizationData(ctx, siteId, mapId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
