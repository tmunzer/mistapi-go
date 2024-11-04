package mistapi

import (
    "context"
    "encoding/json"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsCradlepointTestDeleteOrgCradlepointConnection tests the behavior of the OrgsCradlepoint
func TestOrgsCradlepointTestDeleteOrgCradlepointConnection(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsCradlepoint.DeleteOrgCradlepointConnection(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsCradlepointTestSetupOrgCradlepointConnectionToMist tests the behavior of the OrgsCradlepoint
func TestOrgsCradlepointTestSetupOrgCradlepointConnectionToMist(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    resp, err := orgsCradlepoint.SetupOrgCradlepointConnectionToMist(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsCradlepointTestUpdateOrgCradlepointConnectionToMist tests the behavior of the OrgsCradlepoint
func TestOrgsCradlepointTestUpdateOrgCradlepointConnectionToMist(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    var body models.AccountCradlepointConfig
    errBody := json.Unmarshal([]byte(`{"cp_api_id":"84446d61-2206-4ea5-855a-0043f980be54","cp_api_key":"79c329da9893e34099c7d8ad5cb9c941","ecm_api_id":"73446d61-2206-4ea5-855a-0043f980be62","ecm_api_key":"68b329da9893e34099c7d8ad5cb9c9405"}`), &body)
    if errBody != nil {
        t.Errorf("Cannot parse the model object.")
    }
    resp, err := orgsCradlepoint.UpdateOrgCradlepointConnectionToMist(ctx, orgId, &body)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsCradlepointTestSyncOrgCradlepointRouters tests the behavior of the OrgsCradlepoint
func TestOrgsCradlepointTestSyncOrgCradlepointRouters(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsCradlepoint.SyncOrgCradlepointRouters(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
