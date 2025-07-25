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

// OrgsLicenses represents a controller struct.
type OrgsLicenses struct {
    baseController
}

// NewOrgsLicenses creates a new instance of OrgsLicenses.
// It takes a baseController as a parameter and returns a pointer to the OrgsLicenses.
func NewOrgsLicenses(baseController baseController) *OrgsLicenses {
    orgsLicenses := OrgsLicenses{baseController: baseController}
    return &orgsLicenses
}

// ClaimOrgLicense takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponseClaimLicense data and
// an error if there was an issue with the request or response.
// Claim Org licenses / activation codes
func (o *OrgsLicenses) ClaimOrgLicense(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.ClaimActivation) (
    models.ApiResponse[models.ResponseClaimLicense],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/claim")
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
        "400": {Message: "Invalid key (or already used)"},
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401Error},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403Error},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429Error},
    })
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ResponseClaimLicense
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseClaimLicense](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetOrgLicenseAsyncClaimStatus takes context, orgId, detail as parameters and
// returns an models.ApiResponse with models.ResponseAsyncLicense data and
// an error if there was an issue with the request or response.
// Get Processing Status for Async Claim
func (o *OrgsLicenses) GetOrgLicenseAsyncClaimStatus(
    ctx context.Context,
    orgId uuid.UUID,
    detail *bool) (
    models.ApiResponse[models.ResponseAsyncLicense],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/claim/status")
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
    if detail != nil {
        req.QueryParam("detail", *detail)
    }
    
    var result models.ResponseAsyncLicense
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseAsyncLicense](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetOrgLicensesSummary takes context, orgId as parameters and
// returns an models.ApiResponse with models.License data and
// an error if there was an issue with the request or response.
// Get the list of licenses
func (o *OrgsLicenses) GetOrgLicensesSummary(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.License],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/licenses")
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
    
    var result models.License
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.License](decoder)
    return models.NewApiResponse(result, resp), err
}

// MoveOrDeleteOrgLicenseToAnotherOrg takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Move, Undo Move or Delete Org License to Another Org
// If the admin has admin privilege against the `org_id` and `dst_org_id`, he can move some of the licenses to another Org. Given that: 
// 1. the specified license is currently active 
// 2. there’s enough licenses left in the specified license (by subscription_id) 
// 3. there will still be enough entitled licenses for the type of license after the amendment
func (o *OrgsLicenses) MoveOrDeleteOrgLicenseToAnotherOrg(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.OrgLicenseAction) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/licenses")
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// GetOrgLicensesBySite takes context, orgId as parameters and
// returns an models.ApiResponse with []models.LicenseUsageOrg data and
// an error if there was an issue with the request or response.
// Get Licenses Usage by Sites
// This shows license usage (i.e. needed) based on the features enabled for the site.
func (o *OrgsLicenses) GetOrgLicensesBySite(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.LicenseUsageOrg],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/licenses/usages")
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
    
    var result []models.LicenseUsageOrg
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.LicenseUsageOrg](decoder)
    return models.NewApiResponse(result, resp), err
}
