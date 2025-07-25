// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "github.com/google/uuid"
    "testing"
)

// TestOrgsSDKTemplatesTestListSdkTemplates tests the behavior of the OrgsSDKTemplates
func TestOrgsSDKTemplatesTestListSdkTemplates(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSdkTemplates.ListSdkTemplates(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"bg_image":"https://p.imgci.com/db/PICTURES/CMS/227700/227791.4.jpg","btn_flr_bgcolor":"#282828","default":true,"header_txt":"Mist","name":"default","search_txtcolor":"#282828","welcome_msg":"Welcome to Mist"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKTemplatesTestListSdkTemplates1 tests the behavior of the OrgsSDKTemplates
func TestOrgsSDKTemplatesTestListSdkTemplates1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSdkTemplates.ListSdkTemplates(ctx, orgId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `[{"bg_image":"https://p.imgci.com/db/PICTURES/CMS/227700/227791.4.jpg","btn_flr_bgcolor":"#282828","default":true,"header_txt":"Mist","name":"default","search_txtcolor":"#282828","welcome_msg":"Welcome to Mist"}]`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKTemplatesTestCreateSdkTemplate tests the behavior of the OrgsSDKTemplates
func TestOrgsSDKTemplatesTestCreateSdkTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSdkTemplates.CreateSdkTemplate(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"bg_image":"https://p.imgci.com/db/PICTURES/CMS/227700/227791.4.jpg","btn_flr_bgcolor":"#282828","default":true,"header_txt":"Mist","name":"default","search_txtcolor":"#282828","welcome_msg":"Welcome to Mist"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKTemplatesTestCreateSdkTemplate1 tests the behavior of the OrgsSDKTemplates
func TestOrgsSDKTemplatesTestCreateSdkTemplate1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSdkTemplates.CreateSdkTemplate(ctx, orgId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"bg_image":"https://p.imgci.com/db/PICTURES/CMS/227700/227791.4.jpg","btn_flr_bgcolor":"#282828","default":true,"header_txt":"Mist","name":"default","search_txtcolor":"#282828","welcome_msg":"Welcome to Mist"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKTemplatesTestDeleteSdkTemplate tests the behavior of the OrgsSDKTemplates
func TestOrgsSDKTemplatesTestDeleteSdkTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdktemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    resp, err := orgsSdkTemplates.DeleteSdkTemplate(ctx, orgId, sdktemplateId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, resp.StatusCode, 200)
}

// TestOrgsSDKTemplatesTestGetSdkTemplate tests the behavior of the OrgsSDKTemplates
func TestOrgsSDKTemplatesTestGetSdkTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdktemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSdkTemplates.GetSdkTemplate(ctx, orgId, sdktemplateId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"bg_image":"https://p.imgci.com/db/PICTURES/CMS/227700/227791.4.jpg","btn_flr_bgcolor":"#282828","default":true,"header_txt":"Mist","name":"default","search_txtcolor":"#282828","welcome_msg":"Welcome to Mist"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKTemplatesTestGetSdkTemplate1 tests the behavior of the OrgsSDKTemplates
func TestOrgsSDKTemplatesTestGetSdkTemplate1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdktemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    apiResponse, err := orgsSdkTemplates.GetSdkTemplate(ctx, orgId, sdktemplateId)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"bg_image":"https://p.imgci.com/db/PICTURES/CMS/227700/227791.4.jpg","btn_flr_bgcolor":"#282828","default":true,"header_txt":"Mist","name":"default","search_txtcolor":"#282828","welcome_msg":"Welcome to Mist"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKTemplatesTestUpdateSdkTemplate tests the behavior of the OrgsSDKTemplates
func TestOrgsSDKTemplatesTestUpdateSdkTemplate(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdktemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSdkTemplates.UpdateSdkTemplate(ctx, orgId, sdktemplateId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"bg_image":"https://p.imgci.com/db/PICTURES/CMS/227700/227791.4.jpg","btn_flr_bgcolor":"#282828","default":true,"header_txt":"Mist","name":"default","search_txtcolor":"#282828","welcome_msg":"Welcome to Mist"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}

// TestOrgsSDKTemplatesTestUpdateSdkTemplate1 tests the behavior of the OrgsSDKTemplates
func TestOrgsSDKTemplatesTestUpdateSdkTemplate1(t *testing.T) {
    ctx := context.Background()
    orgId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    sdktemplateId, errUUID := uuid.Parse("000000ab-00ab-00ab-00ab-0000000000ab")
    if errUUID != nil {
        t.Error(errUUID)
    }
    
    apiResponse, err := orgsSdkTemplates.UpdateSdkTemplate(ctx, orgId, sdktemplateId, nil)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/vnd.api+json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
    expected := `{"bg_image":"https://p.imgci.com/db/PICTURES/CMS/227700/227791.4.jpg","btn_flr_bgcolor":"#282828","default":true,"header_txt":"Mist","name":"default","search_txtcolor":"#282828","welcome_msg":"Welcome to Mist"}`
    testHelper.KeysBodyMatcher(t, expected, apiResponse.Response.Body, false, false)
}
