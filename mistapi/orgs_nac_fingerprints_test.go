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

// TestOrgsNACFingerprintsTestCountOrgClientFingerprints tests the behavior of the OrgsNACFingerprints
func TestOrgsNACFingerprintsTestCountOrgClientFingerprints(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.FingerprintsCountDistinctEnum("family")
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsNacFingerprints.CountOrgClientFingerprints(ctx, siteId, &distinct, nil, nil, &duration, &limit)
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

// TestOrgsNACFingerprintsTestCountOrgClientFingerprints1 tests the behavior of the OrgsNACFingerprints
func TestOrgsNACFingerprintsTestCountOrgClientFingerprints1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    distinct := models.FingerprintsCountDistinctEnum("family")
    
    
    duration := "1d"
    limit := int(100)
    apiResponse, err := orgsNacFingerprints.CountOrgClientFingerprints(ctx, siteId, &distinct, nil, nil, &duration, &limit)
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

// TestOrgsNACFingerprintsTestSearchOrgClientFingerprints tests the behavior of the OrgsNACFingerprints
func TestOrgsNACFingerprintsTestSearchOrgClientFingerprints(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    family := "EX Series Switch"
    clientType := models.NacAccessTypeEnum("wired")
    model := "ex4100-f-12p"
    mfg := "Juniper Networks, Inc."
    os := "JUNOS 22.3R1.12"
    osType := "JUNOS"
    mac := "d420b080516d"
    sort := "-family"
    limit := int(100)
    
    
    duration := "1d"
    interval := "10m"
    apiResponse, err := orgsNacFingerprints.SearchOrgClientFingerprints(ctx, siteId, &family, &clientType, &model, &mfg, &os, &osType, &mac, &sort, &limit, nil, nil, &duration, &interval)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1735678700,"limit":10,"results":[{"family":"Apple","mac":"d420b080516e","mfg":"Apple, Inc.","model":"Unknown","org_id":"bb2fb165-0931-49c7-a1b8-9b5814326b7d","os":"iOS 18.1.1","os_type":"iOS","random_mac":true,"site_id":"441a1214-6928-442a-8e92-e1d34b8ec6a6","timestamp":1735678662.58},{"family":"EX Series Switch","mac":"d420b080516d","mfg":"Juniper Networks, Inc.","model":"ex4100-f-12p","org_id":"b6bc08f3-60a3-402b-8f0d-caf9132a1e9a","os":"JUNOS 22.3R1.12","os_type":"JUNOS","random_mac":false,"site_id":"441a1214-6928-442a-8e92-e1d34b8ec6a6","timestamp":1735669092.932}],"start":1735678650,"total":2}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsNACFingerprintsTestSearchOrgClientFingerprints1 tests the behavior of the OrgsNACFingerprints
func TestOrgsNACFingerprintsTestSearchOrgClientFingerprints1(t *testing.T) {
    ctx := context.Background()
    siteId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    family := "EX Series Switch"
    clientType := models.NacAccessTypeEnum("wired")
    model := "ex4100-f-12p"
    mfg := "Juniper Networks, Inc."
    os := "JUNOS 22.3R1.12"
    osType := "JUNOS"
    mac := "d420b080516d"
    sort := "-family"
    limit := int(100)
    
    
    duration := "1d"
    interval := "10m"
    apiResponse, err := orgsNacFingerprints.SearchOrgClientFingerprints(ctx, siteId, &family, &clientType, &model, &mfg, &os, &osType, &mac, &sort, &limit, nil, nil, &duration, &interval)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"end":1735678700,"limit":10,"results":[{"family":"Apple","mac":"d420b080516e","mfg":"Apple, Inc.","model":"Unknown","org_id":"bb2fb165-0931-49c7-a1b8-9b5814326b7d","os":"iOS 18.1.1","os_type":"iOS","random_mac":true,"site_id":"441a1214-6928-442a-8e92-e1d34b8ec6a6","timestamp":1735678662.58},{"family":"EX Series Switch","mac":"d420b080516d","mfg":"Juniper Networks, Inc.","model":"ex4100-f-12p","org_id":"b6bc08f3-60a3-402b-8f0d-caf9132a1e9a","os":"JUNOS 22.3R1.12","os_type":"JUNOS","random_mac":false,"site_id":"441a1214-6928-442a-8e92-e1d34b8ec6a6","timestamp":1735669092.932}],"start":1735678650,"total":2}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
