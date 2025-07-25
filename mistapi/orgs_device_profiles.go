// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// OrgsDeviceProfiles represents a controller struct.
type OrgsDeviceProfiles struct {
    baseController
}

// NewOrgsDeviceProfiles creates a new instance of OrgsDeviceProfiles.
// It takes a baseController as a parameter and returns a pointer to the OrgsDeviceProfiles.
func NewOrgsDeviceProfiles(baseController baseController) *OrgsDeviceProfiles {
    orgsDeviceProfiles := OrgsDeviceProfiles{baseController: baseController}
    return &orgsDeviceProfiles
}

// ListOrgDeviceProfiles takes context, orgId, mType, limit, page as parameters and
// returns an models.ApiResponse with []models.Deviceprofile data and
// an error if there was an issue with the request or response.
// Get List of Org Device Profiles
func (o *OrgsDeviceProfiles) ListOrgDeviceProfiles(
    ctx context.Context,
    orgId uuid.UUID,
    mType *models.DeviceTypeDefaultApEnum,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Deviceprofile],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/deviceprofiles")
    req.AppendTemplateParams(orgId)
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result []models.Deviceprofile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Deviceprofile](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgDeviceProfile takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Deviceprofile data and
// an error if there was an issue with the request or response.
// Create Device Profile
func (o *OrgsDeviceProfiles) CreateOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Deviceprofile) (
    models.ApiResponse[models.Deviceprofile],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/deviceprofiles")
    req.AppendTemplateParams(orgId)
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.Deviceprofile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Deviceprofile](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgDeviceProfile takes context, orgId, deviceprofileId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org Device Profile
func (o *OrgsDeviceProfiles) DeleteOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/deviceprofiles/%v")
    req.AppendTemplateParams(orgId, deviceprofileId)
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// GetOrgDeviceProfile takes context, orgId, deviceprofileId as parameters and
// returns an models.ApiResponse with models.Deviceprofile data and
// an error if there was an issue with the request or response.
// Get Org device Profile Details
func (o *OrgsDeviceProfiles) GetOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID) (
    models.ApiResponse[models.Deviceprofile],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/deviceprofiles/%v")
    req.AppendTemplateParams(orgId, deviceprofileId)
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    
    var result models.Deviceprofile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Deviceprofile](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgDeviceProfile takes context, orgId, deviceprofileId, body as parameters and
// returns an models.ApiResponse with models.Deviceprofile data and
// an error if there was an issue with the request or response.
// Update Org Device Profile
func (o *OrgsDeviceProfiles) UpdateOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID,
    body *models.Deviceprofile) (
    models.ApiResponse[models.Deviceprofile],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/deviceprofiles/%v")
    req.AppendTemplateParams(orgId, deviceprofileId)
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.Deviceprofile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Deviceprofile](decoder)
    return models.NewApiResponse(result, resp), err
}

// AssignOrgDeviceProfile takes context, orgId, deviceprofileId, body as parameters and
// returns an models.ApiResponse with models.ResponseAssignSuccess data and
// an error if there was an issue with the request or response.
// Assign Org Device Profile to Devices
func (o *OrgsDeviceProfiles) AssignOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.ResponseAssignSuccess],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/deviceprofiles/%v/assign")
    req.AppendTemplateParams(orgId, deviceprofileId)
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ResponseAssignSuccess
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseAssignSuccess](decoder)
    return models.NewApiResponse(result, resp), err
}

// UnassignOrgDeviceProfile takes context, orgId, deviceprofileId, body as parameters and
// returns an models.ApiResponse with models.ResponseAssignSuccess data and
// an error if there was an issue with the request or response.
// Unassign Org Device Profile from Devices
func (o *OrgsDeviceProfiles) UnassignOrgDeviceProfile(
    ctx context.Context,
    orgId uuid.UUID,
    deviceprofileId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.ResponseAssignSuccess],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      "/api/v1/orgs/%v/deviceprofiles/%v/unassign",
    )
    req.AppendTemplateParams(orgId, deviceprofileId)
    req.Authenticate(
        NewOrAuth(
            NewAuth("apiToken"),
            NewAuth("basicAuth"),
            NewAndAuth(
                NewAuth("basicAuth"),
                NewAuth("csrfToken"),
            ),

        ),
    )
    req.AppendErrors(map[string]https.ErrorBuilder[error]{
        "400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ResponseAssignSuccess
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseAssignSuccess](decoder)
    return models.NewApiResponse(result, resp), err
}
