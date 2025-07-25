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

// OrgsAssetFilters represents a controller struct.
type OrgsAssetFilters struct {
    baseController
}

// NewOrgsAssetFilters creates a new instance of OrgsAssetFilters.
// It takes a baseController as a parameter and returns a pointer to the OrgsAssetFilters.
func NewOrgsAssetFilters(baseController baseController) *OrgsAssetFilters {
    orgsAssetFilters := OrgsAssetFilters{baseController: baseController}
    return &orgsAssetFilters
}

// ListOrgAssetFilters takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.AssetFilter data and
// an error if there was an issue with the request or response.
// Get List of Org BLE asset filters. 
// Each asset filter in the list operates independently. For a filter object to match an asset, all of the filter properties must match (logical ‘AND’ of each filter property). For example, the "Visitor Tags" filter below will match an asset when both the "ibeacon\_uuid" and "ibeacon_major" properties match the asset. All non-matching assets are ignored.
func (o *OrgsAssetFilters) ListOrgAssetFilters(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.AssetFilter],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/assetfilters")
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
    
    var result []models.AssetFilter
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.AssetFilter](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgAssetFilter takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.AssetFilter data and
// an error if there was an issue with the request or response.
// Create Asset Filter
// Creates a single BLE asset filter for the given site. Any subset of filter properties can be included in the filter. A matching asset must meet the conditions of all given filter properties (logical ‘AND’).
func (o *OrgsAssetFilters) CreateOrgAssetFilter(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AssetFilter) (
    models.ApiResponse[models.AssetFilter],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/assetfilters")
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
    
    var result models.AssetFilter
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AssetFilter](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgAssetFilter takes context, orgId, assetfilterId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Deletes an existing BLE asset filter for the given site.
func (o *OrgsAssetFilters) DeleteOrgAssetFilter(
    ctx context.Context,
    orgId uuid.UUID,
    assetfilterId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/assetfilters/%v")
    req.AppendTemplateParams(orgId, assetfilterId)
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

// GetOrgAssetFilter takes context, orgId, assetfilterId as parameters and
// returns an models.ApiResponse with models.AssetFilter data and
// an error if there was an issue with the request or response.
// Get Org Asset Filter Details
func (o *OrgsAssetFilters) GetOrgAssetFilter(
    ctx context.Context,
    orgId uuid.UUID,
    assetfilterId uuid.UUID) (
    models.ApiResponse[models.AssetFilter],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/assetfilters/%v")
    req.AppendTemplateParams(orgId, assetfilterId)
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
    
    var result models.AssetFilter
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AssetFilter](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgAssetFilter takes context, orgId, assetfilterId, body as parameters and
// returns an models.ApiResponse with models.AssetFilter data and
// an error if there was an issue with the request or response.
// Updates an existing BLE asset filter for the given site.
func (o *OrgsAssetFilters) UpdateOrgAssetFilter(
    ctx context.Context,
    orgId uuid.UUID,
    assetfilterId uuid.UUID,
    body *models.AssetFilter) (
    models.ApiResponse[models.AssetFilter],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/assetfilters/%v")
    req.AppendTemplateParams(orgId, assetfilterId)
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
    
    var result models.AssetFilter
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AssetFilter](decoder)
    return models.NewApiResponse(result, resp), err
}
