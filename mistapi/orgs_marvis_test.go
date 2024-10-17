package mistapi

import (
	"context"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/google/uuid"
	"testing"
)

// TestOrgsMarvisTestTroubleshootOrg tests the behavior of the OrgsMarvis
func TestOrgsMarvisTestTroubleshootOrg(t *testing.T) {
	ctx := context.Background()
	orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
	if errUUID != nil {
		t.Error(errUUID)
	}

	apiResponse, err := orgsMarvis.TroubleshootOrg(ctx, orgId, nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"end":1655151856,"results":[{"category":"client","reason":"slow association","recommendation":"Ensure the IP helper-address is configured on the VLAN interface.","text":"Clients of the AP had slow association 8% of the time on Bhavabhi and 5 GHz. ..."}],"start":1655065456}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
