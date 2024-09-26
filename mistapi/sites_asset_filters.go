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

// SitesAssetFilters represents a controller struct.
type SitesAssetFilters struct {
    baseController
}

// NewSitesAssetFilters creates a new instance of SitesAssetFilters.
// It takes a baseController as a parameter and returns a pointer to the SitesAssetFilters.
func NewSitesAssetFilters(baseController baseController) *SitesAssetFilters {
    sitesAssetFilters := SitesAssetFilters{baseController: baseController}
    return &sitesAssetFilters
}

// ListSiteAssetFilters takes context, siteId, limit, page as parameters and
// returns an models.ApiResponse with []models.AssetFilter data and
// an error if there was an issue with the request or response.
// Get List of Site Asset Filters
func (s *SitesAssetFilters) ListSiteAssetFilters(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.AssetFilter],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/assetfilters", siteId),
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
    
    var result []models.AssetFilter
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.AssetFilter](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateSiteAssetFilters takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.AssetFilter data and
// an error if there was an issue with the request or response.
// Create Site Asset Filter
func (s *SitesAssetFilters) CreateSiteAssetFilters(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.AssetFilter) (
    models.ApiResponse[models.AssetFilter],
    error) {
    req := s.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/assetfilters", siteId),
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
    
    var result models.AssetFilter
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AssetFilter](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteSiteAssetFilter takes context, siteId, assetfilterId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Deletes an existing BLE asset filter for the given site.
func (s *SitesAssetFilters) DeleteSiteAssetFilter(
    ctx context.Context,
    siteId uuid.UUID,
    assetfilterId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/sites/%v/assetfilters/%v", siteId, assetfilterId),
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

// GetSiteAssetFilter takes context, siteId, assetfilterId as parameters and
// returns an models.ApiResponse with models.AssetFilter data and
// an error if there was an issue with the request or response.
// Get Site Asset Filter Details
func (s *SitesAssetFilters) GetSiteAssetFilter(
    ctx context.Context,
    siteId uuid.UUID,
    assetfilterId uuid.UUID) (
    models.ApiResponse[models.AssetFilter],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/assetfilters/%v", siteId, assetfilterId),
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
    
    var result models.AssetFilter
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AssetFilter](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteAssetFilter takes context, siteId, assetfilterId, body as parameters and
// returns an models.ApiResponse with models.AssetFilter data and
// an error if there was an issue with the request or response.
// Updates an existing BLE asset filter for the given site.
func (s *SitesAssetFilters) UpdateSiteAssetFilter(
    ctx context.Context,
    siteId uuid.UUID,
    assetfilterId uuid.UUID,
    body *models.AssetFilter) (
    models.ApiResponse[models.AssetFilter],
    error) {
    req := s.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/sites/%v/assetfilters/%v", siteId, assetfilterId),
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
    
    var result models.AssetFilter
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.AssetFilter](decoder)
    return models.NewApiResponse(result, resp), err
}
