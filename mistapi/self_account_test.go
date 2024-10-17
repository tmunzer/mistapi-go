package mistapi

import (
	"context"
	"encoding/json"
	"github.com/apimatic/go-core-runtime/testHelper"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"testing"
)

// TestSelfAccountTestDeleteSelf tests the behavior of the SelfAccount
func TestSelfAccountTestDeleteSelf(t *testing.T) {
	ctx := context.Background()
	resp, err := selfAccount.DeleteSelf(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestSelfAccountTestGetSelf tests the behavior of the SelfAccount
func TestSelfAccountTestGetSelf(t *testing.T) {
	ctx := context.Background()
	apiResponse, err := selfAccount.GetSelf(ctx)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"admin_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","email":"user@example.com","first_name":"string","last_name":"string","password_modified_time":1656353525,"privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"]}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSelfAccountTestUpdateSelf tests the behavior of the SelfAccount
func TestSelfAccountTestUpdateSelf(t *testing.T) {
	ctx := context.Background()
	var body models.Admin
	errBody := json.Unmarshal([]byte(`{"email":"john.smith@mycorp.net","first_name":"John","last_name":"Smith","persona":"security","phone":"14081112222","phone2":"14083334444"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	apiResponse, err := selfAccount.UpdateSelf(ctx, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
	expectedHeaders := []testHelper.TestHeader{
		testHelper.NewTestHeader(true, "Content-Type", "application/json"),
	}
	testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
	expected := `{"admin_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","email":"user@example.com","first_name":"string","last_name":"string","password_modified_time":1656353525,"privileges":[{"msp_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","msp_name":"string","name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","org_name":"string","orggroup_ids":["b069b358-4c97-5319-1f8c-7c5ca64d6ab1"],"role":"admin","scope":"org","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","sitegroup_ids":["6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"]}]}`
	testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestSelfAccountTestUpdateSelfEmail tests the behavior of the SelfAccount
func TestSelfAccountTestUpdateSelfEmail(t *testing.T) {
	ctx := context.Background()
	var body models.EmailString
	errBody := json.Unmarshal([]byte(`{"email":"new@mistsys.com"}`), &body)
	if errBody != nil {
		t.Errorf("Cannot parse the model object.")
	}
	resp, err := selfAccount.UpdateSelfEmail(ctx, &body)
	if err != nil {
		t.Errorf("Endpoint call failed: %v", err)
	}
	testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
