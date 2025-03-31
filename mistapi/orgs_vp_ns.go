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

// OrgsVPNs represents a controller struct.
type OrgsVPNs struct {
    baseController
}

// NewOrgsVPNs creates a new instance of OrgsVPNs.
// It takes a baseController as a parameter and returns a pointer to the OrgsVPNs.
func NewOrgsVPNs(baseController baseController) *OrgsVPNs {
    orgsVPNs := OrgsVPNs{baseController: baseController}
    return &orgsVPNs
}

// ListOrgVpns takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.Vpn data and
// an error if there was an issue with the request or response.
// Get List of Org VPNs
func (o *OrgsVPNs) ListOrgVpns(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Vpn],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/vpns")
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
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result []models.Vpn
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Vpn](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgVpn takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Vpn data and
// an error if there was an issue with the request or response.
// Create Org VPN
func (o *OrgsVPNs) CreateOrgVpn(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Vpn) (
    models.ApiResponse[models.Vpn],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/vpns")
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
    
    var result models.Vpn
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Vpn](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgVpn takes context, orgId, vpnId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org Vpn
func (o *OrgsVPNs) DeleteOrgVpn(
    ctx context.Context,
    orgId uuid.UUID,
    vpnId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/vpns/%v")
    req.AppendTemplateParams(orgId, vpnId)
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

// GetOrgVpn takes context, orgId, vpnId as parameters and
// returns an models.ApiResponse with models.Vpn data and
// an error if there was an issue with the request or response.
// Get Org Vpn
func (o *OrgsVPNs) GetOrgVpn(
    ctx context.Context,
    orgId uuid.UUID,
    vpnId uuid.UUID) (
    models.ApiResponse[models.Vpn],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/vpns/%v")
    req.AppendTemplateParams(orgId, vpnId)
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
    
    var result models.Vpn
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Vpn](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgVpn takes context, orgId, vpnId, body as parameters and
// returns an models.ApiResponse with models.Vpn data and
// an error if there was an issue with the request or response.
// Update Org Vpn
func (o *OrgsVPNs) UpdateOrgVpn(
    ctx context.Context,
    orgId uuid.UUID,
    vpnId uuid.UUID,
    body *models.Vpn) (
    models.ApiResponse[models.Vpn],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/vpns/%v")
    req.AppendTemplateParams(orgId, vpnId)
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
    
    var result models.Vpn
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Vpn](decoder)
    return models.NewApiResponse(result, resp), err
}
