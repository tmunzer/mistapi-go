package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsSubscriptionsTestUnsubscribeOrgAlarmsReports tests the behavior of the OrgsSubscriptions
func TestOrgsSubscriptionsTestUnsubscribeOrgAlarmsReports(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsSubscriptions.UnsubscribeOrgAlarmsReports(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSubscriptionsTestSubscribeOrgAlarmsReports tests the behavior of the OrgsSubscriptions
func TestOrgsSubscriptionsTestSubscribeOrgAlarmsReports(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsSubscriptions.SubscribeOrgAlarmsReports(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
