package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsStatsTestGetOrgStats tests the behavior of the OrgsStats
func TestOrgsStatsTestGetOrgStats(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	duration := "1d"
	limit := int(100)
	page := int(1)
	apiResponse, err := orgsStats.GetOrgStats(ctx, orgId, nil, nil, &duration, &limit, &page)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"alarmtemplate_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","allow_mist":true,"created_time":0,"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","name":"string","num_devices":0,"num_devices_connected":0,"num_devices_disconnected":0,"num_inventory":0,"num_sites":0,"orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"session_expiry":0,"sle":[{"path":"string","user_minutes":{"ok":0,"total":0}}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
