package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// OrgsWxTunnels represents a controller struct.
type OrgsWxTunnels struct {
    baseController
}

// NewOrgsWxTunnels creates a new instance of OrgsWxTunnels.
// It takes a baseController as a parameter and returns a pointer to the OrgsWxTunnels.
func NewOrgsWxTunnels(baseController baseController) *OrgsWxTunnels {
    orgsWxTunnels := OrgsWxTunnels{baseController: baseController}
    return &orgsWxTunnels
}

// ListOrgWxTunnels takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.WxlanTunnel data and
// an error if there was an issue with the request or response.
// Get List of Org WxLAN Tunnels
func (o *OrgsWxTunnels) ListOrgWxTunnels(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.WxlanTunnel],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/wxtunnels", orgId),
    )
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
    
    var result []models.WxlanTunnel
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.WxlanTunnel](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgWxTunnel takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.WxlanTunnel data and
// an error if there was an issue with the request or response.
// Create Org WxAN Tunnel
func (o *OrgsWxTunnels) CreateOrgWxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.WxlanTunnel) (
    models.ApiResponse[models.WxlanTunnel],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/wxtunnels", orgId),
    )
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
    
    var result models.WxlanTunnel
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WxlanTunnel](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgWxTunnel takes context, orgId, wxtunnelId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org WxLAN Tunnel
func (o *OrgsWxTunnels) DeleteOrgWxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    wxtunnelId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/orgs/%v/wxtunnels/%v", orgId, wxtunnelId),
    )
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

// GetOrgWxTunnel takes context, orgId, wxtunnelId as parameters and
// returns an models.ApiResponse with models.WxlanTunnel data and
// an error if there was an issue with the request or response.
// Get Org WxLAN Tunnel Details
func (o *OrgsWxTunnels) GetOrgWxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    wxtunnelId uuid.UUID) (
    models.ApiResponse[models.WxlanTunnel],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/wxtunnels/%v", orgId, wxtunnelId),
    )
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
    
    var result models.WxlanTunnel
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WxlanTunnel](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgWxTunnel takes context, orgId, wxtunnelId, body as parameters and
// returns an models.ApiResponse with models.WxlanTunnel data and
// an error if there was an issue with the request or response.
// Update Org WxLAN Tunnel
func (o *OrgsWxTunnels) UpdateOrgWxTunnel(
    ctx context.Context,
    orgId uuid.UUID,
    wxtunnelId uuid.UUID,
    body *models.WxlanTunnel) (
    models.ApiResponse[models.WxlanTunnel],
    error) {
    req := o.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/orgs/%v/wxtunnels/%v", orgId, wxtunnelId),
    )
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
    
    var result models.WxlanTunnel
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WxlanTunnel](decoder)
    return models.NewApiResponse(result, resp), err
}
