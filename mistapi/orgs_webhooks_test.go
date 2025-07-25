// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "testing"
)

// TestOrgsWebhooksTestListOrgWebhooks tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestListOrgWebhooks(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsWebhooks.ListOrgWebhooks(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"enabled":true,"headers":{},"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","secret":"string","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","splunk_token":"string","topics":["location"],"type":"http-post","url":"string","verify_cert":true}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWebhooksTestListOrgWebhooks1 tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestListOrgWebhooks1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    limit := int(100)
    page := int(1)
    apiResponse, err := orgsWebhooks.ListOrgWebhooks(ctx, orgId, &limit, &page)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"created_time":0,"enabled":true,"headers":{},"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","secret":"string","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","splunk_token":"string","topics":["location"],"type":"http-post","url":"string","verify_cert":true}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWebhooksTestCreateOrgWebhook tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestCreateOrgWebhook(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsWebhooks.CreateOrgWebhook(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"enabled":true,"headers":{},"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","secret":"string","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","splunk_token":"string","topics":["location"],"type":"http-post","url":"string","verify_cert":true}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWebhooksTestCreateOrgWebhook1 tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestCreateOrgWebhook1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsWebhooks.CreateOrgWebhook(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"enabled":true,"headers":{},"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","secret":"string","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","splunk_token":"string","topics":["location"],"type":"http-post","url":"string","verify_cert":true}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWebhooksTestDeleteOrgWebhook tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestDeleteOrgWebhook(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    webhookId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsWebhooks.DeleteOrgWebhook(ctx, orgId, webhookId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsWebhooksTestGetOrgWebhook tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestGetOrgWebhook(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    webhookId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsWebhooks.GetOrgWebhook(ctx, orgId, webhookId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"enabled":true,"headers":{},"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","secret":"string","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","splunk_token":"string","topics":["location"],"type":"http-post","url":"string","verify_cert":true}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWebhooksTestGetOrgWebhook1 tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestGetOrgWebhook1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    webhookId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsWebhooks.GetOrgWebhook(ctx, orgId, webhookId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"enabled":true,"headers":{},"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","secret":"string","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","splunk_token":"string","topics":["location"],"type":"http-post","url":"string","verify_cert":true}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWebhooksTestUpdateOrgWebhook tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestUpdateOrgWebhook(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    webhookId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsWebhooks.UpdateOrgWebhook(ctx, orgId, webhookId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"enabled":true,"headers":{},"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","secret":"string","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","splunk_token":"string","topics":["location"],"type":"http-post","url":"string","verify_cert":true}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWebhooksTestUpdateOrgWebhook1 tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestUpdateOrgWebhook1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    webhookId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsWebhooks.UpdateOrgWebhook(ctx, orgId, webhookId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"created_time":0,"enabled":true,"headers":{},"id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","modified_time":0,"name":"string","org_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","secret":"string","site_id":"b069b358-4c97-5319-1f8c-7c5ca64d6ab1","splunk_token":"string","topics":["location"],"type":"http-post","url":"string","verify_cert":true}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWebhooksTestCountOrgWebhooksDeliveries tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestCountOrgWebhooksDeliveries(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    webhookId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    status := models.WebhookDeliveryStatusEnum("failure")
    topic := models.WebhookDeliveryTopicEnum("audits")
    distinct := models.WebhookDeliveryDistinctEnum("webhook_id")
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsWebhooks.CountOrgWebhooksDeliveries(ctx, orgId, webhookId, nil, nil, &status, &topic, &distinct, nil, nil, &duration, &limit)
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

// TestOrgsWebhooksTestCountOrgWebhooksDeliveries1 tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestCountOrgWebhooksDeliveries1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    webhookId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    status := models.WebhookDeliveryStatusEnum("failure")
    topic := models.WebhookDeliveryTopicEnum("audits")
    distinct := models.WebhookDeliveryDistinctEnum("webhook_id")
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsWebhooks.CountOrgWebhooksDeliveries(ctx, orgId, webhookId, nil, nil, &status, &topic, &distinct, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"distinct":"string","end":0,"limit":0,"results":[{"count":0,"property":"string"}],"start":0,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWebhooksTestSearchOrgWebhooksDeliveries tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestSearchOrgWebhooksDeliveries(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    webhookId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    status := models.WebhookDeliveryStatusEnum("failure")
    topic := models.WebhookDeliveryTopicEnum("audits")
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsWebhooks.SearchOrgWebhooksDeliveries(ctx, orgId, webhookId, nil, nil, &status, &topic, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1688035193,"limit":10,"results":[{"error":"string","id":"55b0f02f-ebf6-4ad2-8b10-200508a97581","org_id":"fc7e2967-e7ef-41e6-b007-1217713de05a","req_headers":"{\\\"Content-Type\\\":[\\\"application/json\\\"],\\\"User-Agent\\\":[\\\"Mist-webhook\\\"]}","req_payload":"{\\\"topic\\\":\\\"audits\\\",\\\"events\\\":[{\\\"admin_name\\\":\\\"John Doe john.doe@juniper.net\\\",\\\"after\\\":\\\"{\\\\\"radio_config\\\\\": {\\\\\"band_24\\\\\": {\\\\\"disabled\\\\\": false, \\\\\"allow_rrm_disable\\\\\": false, \\\\\"power_min\\\\\": null, \\\\\"power_max\\\\\": null, \\\\\"power\\\\\": 10, \\\\\"preamble\\\\\": \\\\\"short\\\\\", \\\\\"channels\\\\\": [1, 10], \\\\\"bandwidth\\\\\": 20}}}\\\",\\\"before\\\":\\\"{\\\\\"radio_config\\\\\": {\\\\\"band_24\\\\\": {\\\\\"disabled\\\\\": false, \\\\\"allow_rrm_disable\\\\\": false, \\\\\"power_min\\\\\": 8, \\\\\"power_max\\\\\": 18, \\\\\"power\\\\\": null, \\\\\"preamble\\\\\": \\\\\"long\\\\\", \\\\\"channels\\\\\": [1, 10], \\\\\"bandwidth\\\\\": 20}}}\\\",\\\"id\\\":\\\"737909a2-04ff-4aeb-b9da-cc924e74a4dd\\\",\\\"message\\\":\\\"Update Site Settings\\\",\\\"org_id\\\":\\\"fc7e2967-e7ef-41e6-b007-1217713de05a\\\",\\\"site_id\\\":\\\"256c3a35-9cb7-436e-bc6d-314972645d95\\\",\\\"site_name\\\":\\\"Test Site\\\",\\\"src_ip\\\":\\\"1.2.3.4\\\",\\\"timestamp\\\":1685956576.923601,\\\"user_agent\\\":\\\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36\\\"}]}","req_url":"https://example.com","resp_body":"Ok","resp_headers":"string","site_id":"256c3a35-9cb7-436e-bc6d-314972645d95","status":"success","status_code":200,"timestamp":1687962508.583656,"topic":"audits","webhook_id":"7a11b901-f719-4c91-8aef-deb8699a6364"}],"start":1687948793,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWebhooksTestSearchOrgWebhooksDeliveries1 tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestSearchOrgWebhooksDeliveries1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    webhookId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    
    status := models.WebhookDeliveryStatusEnum("failure")
    topic := models.WebhookDeliveryTopicEnum("audits")
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsWebhooks.SearchOrgWebhooksDeliveries(ctx, orgId, webhookId, nil, nil, &status, &topic, nil, nil, &duration, &limit)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1688035193,"limit":10,"results":[{"error":"string","id":"55b0f02f-ebf6-4ad2-8b10-200508a97581","org_id":"fc7e2967-e7ef-41e6-b007-1217713de05a","req_headers":"{\\\"Content-Type\\\":[\\\"application/json\\\"],\\\"User-Agent\\\":[\\\"Mist-webhook\\\"]}","req_payload":"{\\\"topic\\\":\\\"audits\\\",\\\"events\\\":[{\\\"admin_name\\\":\\\"John Doe john.doe@juniper.net\\\",\\\"after\\\":\\\"{\\\\\"radio_config\\\\\": {\\\\\"band_24\\\\\": {\\\\\"disabled\\\\\": false, \\\\\"allow_rrm_disable\\\\\": false, \\\\\"power_min\\\\\": null, \\\\\"power_max\\\\\": null, \\\\\"power\\\\\": 10, \\\\\"preamble\\\\\": \\\\\"short\\\\\", \\\\\"channels\\\\\": [1, 10], \\\\\"bandwidth\\\\\": 20}}}\\\",\\\"before\\\":\\\"{\\\\\"radio_config\\\\\": {\\\\\"band_24\\\\\": {\\\\\"disabled\\\\\": false, \\\\\"allow_rrm_disable\\\\\": false, \\\\\"power_min\\\\\": 8, \\\\\"power_max\\\\\": 18, \\\\\"power\\\\\": null, \\\\\"preamble\\\\\": \\\\\"long\\\\\", \\\\\"channels\\\\\": [1, 10], \\\\\"bandwidth\\\\\": 20}}}\\\",\\\"id\\\":\\\"737909a2-04ff-4aeb-b9da-cc924e74a4dd\\\",\\\"message\\\":\\\"Update Site Settings\\\",\\\"org_id\\\":\\\"fc7e2967-e7ef-41e6-b007-1217713de05a\\\",\\\"site_id\\\":\\\"256c3a35-9cb7-436e-bc6d-314972645d95\\\",\\\"site_name\\\":\\\"Test Site\\\",\\\"src_ip\\\":\\\"1.2.3.4\\\",\\\"timestamp\\\":1685956576.923601,\\\"user_agent\\\":\\\"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36\\\"}]}","req_url":"https://example.com","resp_body":"Ok","resp_headers":"string","site_id":"256c3a35-9cb7-436e-bc6d-314972645d95","status":"success","status_code":200,"timestamp":1687962508.583656,"topic":"audits","webhook_id":"7a11b901-f719-4c91-8aef-deb8699a6364"}],"start":1687948793,"total":0}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsWebhooksTestPingOrgWebhook tests the behavior of the OrgsWebhooks
func TestOrgsWebhooksTestPingOrgWebhook(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    webhookId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsWebhooks.PingOrgWebhook(ctx, orgId, webhookId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}
