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

// OrgsSDKInvites represents a controller struct.
type OrgsSDKInvites struct {
    baseController
}

// NewOrgsSDKInvites creates a new instance of OrgsSDKInvites.
// It takes a baseController as a parameter and returns a pointer to the OrgsSDKInvites.
func NewOrgsSDKInvites(baseController baseController) *OrgsSDKInvites {
    orgsSDKInvites := OrgsSDKInvites{baseController: baseController}
    return &orgsSDKInvites
}

// ActivateSdkInvite takes context, secret, body as parameters and
// returns an models.ApiResponse with models.ResponseMobileVerifySecret data and
// an error if there was an issue with the request or response.
// Verify secret
func (o *OrgsSDKInvites) ActivateSdkInvite(
    ctx context.Context,
    secret string,
    body *models.DeviceIdString) (
    models.ApiResponse[models.ResponseMobileVerifySecret],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/mobile/verify/%v")
    req.AppendTemplateParams(secret)
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
    
    var result models.ResponseMobileVerifySecret
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseMobileVerifySecret](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSdkInvites takes context, orgId as parameters and
// returns an models.ApiResponse with []models.Sdkinvite data and
// an error if there was an issue with the request or response.
// Get List of Org SDK Invites
func (o *OrgsSDKInvites) ListSdkInvites(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.Sdkinvite],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sdkinvites")
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
    
    var result []models.Sdkinvite
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Sdkinvite](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateSdkInvite takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Sdkinvite data and
// an error if there was an issue with the request or response.
// Create SDK Invite
func (o *OrgsSDKInvites) CreateSdkInvite(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Sdkinvite) (
    models.ApiResponse[models.Sdkinvite],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/sdkinvites")
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
    
    var result models.Sdkinvite
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Sdkinvite](decoder)
    return models.NewApiResponse(result, resp), err
}

// RevokeSdkInvite takes context, orgId, sdkinviteId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Revoke SDK Invite
func (o *OrgsSDKInvites) RevokeSdkInvite(
    ctx context.Context,
    orgId uuid.UUID,
    sdkinviteId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/sdkinvites/%v")
    req.AppendTemplateParams(orgId, sdkinviteId)
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

// GetSdkInvite takes context, orgId, sdkinviteId as parameters and
// returns an models.ApiResponse with models.Sdkinvite data and
// an error if there was an issue with the request or response.
// Get SDK Invite Details
func (o *OrgsSDKInvites) GetSdkInvite(
    ctx context.Context,
    orgId uuid.UUID,
    sdkinviteId uuid.UUID) (
    models.ApiResponse[models.Sdkinvite],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sdkinvites/%v")
    req.AppendTemplateParams(orgId, sdkinviteId)
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
    
    var result models.Sdkinvite
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Sdkinvite](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSdkInvite takes context, orgId, sdkinviteId, body as parameters and
// returns an models.ApiResponse with models.Sdkinvite data and
// an error if there was an issue with the request or response.
// Update SDK Invite
func (o *OrgsSDKInvites) UpdateSdkInvite(
    ctx context.Context,
    orgId uuid.UUID,
    sdkinviteId uuid.UUID,
    body *models.Sdkinvite) (
    models.ApiResponse[models.Sdkinvite],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/sdkinvites/%v")
    req.AppendTemplateParams(orgId, sdkinviteId)
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
    
    var result models.Sdkinvite
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Sdkinvite](decoder)
    return models.NewApiResponse(result, resp), err
}

// SendSdkInviteEmail takes context, orgId, sdkinviteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Send SDK Invite by Email
func (o *OrgsSDKInvites) SendSdkInviteEmail(
    ctx context.Context,
    orgId uuid.UUID,
    sdkinviteId uuid.UUID,
    body *models.EmailString) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/sdkinvites/%v/email")
    req.AppendTemplateParams(orgId, sdkinviteId)
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// GetSdkInviteQrCode takes context, orgId, sdkinviteId as parameters and
// returns an models.ApiResponse with []byte data and
// an error if there was an issue with the request or response.
// Revoke SDK Invite
func (o *OrgsSDKInvites) GetSdkInviteQrCode(
    ctx context.Context,
    orgId uuid.UUID,
    sdkinviteId uuid.UUID) (
    models.ApiResponse[[]byte],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sdkinvites/%v/qrcode")
    req.AppendTemplateParams(orgId, sdkinviteId)
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
    
    stream, resp, err := req.CallAsStream()
    if err != nil {
        return models.NewApiResponse(stream, resp), err
    }
    return models.NewApiResponse(stream, resp), err
}

// SendSdkInviteSms takes context, orgId, sdkinviteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Send SDK Invite by SMS
func (o *OrgsSDKInvites) SendSdkInviteSms(
    ctx context.Context,
    orgId uuid.UUID,
    sdkinviteId uuid.UUID,
    body *models.SdkInviteSms) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/sdkinvites/%v/sms")
    req.AppendTemplateParams(orgId, sdkinviteId)
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}
