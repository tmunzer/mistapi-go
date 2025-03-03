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

// SitesAssets represents a controller struct.
type SitesAssets struct {
    baseController
}

// NewSitesAssets creates a new instance of SitesAssets.
// It takes a baseController as a parameter and returns a pointer to the SitesAssets.
func NewSitesAssets(baseController baseController) *SitesAssets {
    sitesAssets := SitesAssets{baseController: baseController}
    return &sitesAssets
}

// ListSiteAssets takes context, siteId, limit, page as parameters and
// returns an models.ApiResponse with []models.Asset data and
// an error if there was an issue with the request or response.
// Get List of Site Assets
func (s *SitesAssets) ListSiteAssets(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Asset],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/assets")
    req.AppendTemplateParams(siteId)
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
    
    var result []models.Asset
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Asset](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateSiteAsset takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.Asset data and
// an error if there was an issue with the request or response.
// Create Site Asset
func (s *SitesAssets) CreateSiteAsset(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Asset) (
    models.ApiResponse[models.Asset],
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/assets")
    req.AppendTemplateParams(siteId)
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
    
    var result models.Asset
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Asset](decoder)
    return models.NewApiResponse(result, resp), err
}

// ImportSiteAssets takes context, siteId, upsert, file as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Import Site Assets. 
// It can be done via a CSV file or a JSON payload.
// ## CSV File Format
// ```csv
// name,mac
// "asset_name",5c5b53010101
// ```
func (s *SitesAssets) ImportSiteAssets(
    ctx context.Context,
    siteId uuid.UUID,
    upsert *models.ImportSiteAssetsUpsertEnum,
    file *models.FileWrapper) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/assets/import")
    req.AppendTemplateParams(siteId)
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
    if upsert != nil {
        req.QueryParam("upsert", *upsert)
    }
    formFields := []https.FormParam{}
    if file != nil {
        fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
        formFields = append(formFields, fileParam)
    }
    req.FormData(formFields)
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// DeleteSiteAsset takes context, siteId, assetId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Site Asset
func (s *SitesAssets) DeleteSiteAsset(
    ctx context.Context,
    siteId uuid.UUID,
    assetId uuid.UUID) (
    *http.Response,
    error) {
    req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/assets/%v")
    req.AppendTemplateParams(siteId, assetId)
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

// GetSiteAsset takes context, siteId, assetId as parameters and
// returns an models.ApiResponse with models.Asset data and
// an error if there was an issue with the request or response.
// Get Site Asset Details
func (s *SitesAssets) GetSiteAsset(
    ctx context.Context,
    siteId uuid.UUID,
    assetId uuid.UUID) (
    models.ApiResponse[models.Asset],
    error) {
    req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/assets/%v")
    req.AppendTemplateParams(siteId, assetId)
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
    
    var result models.Asset
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Asset](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateSiteAsset takes context, siteId, assetId, body as parameters and
// returns an models.ApiResponse with models.Asset data and
// an error if there was an issue with the request or response.
// Update Site Asset
func (s *SitesAssets) UpdateSiteAsset(
    ctx context.Context,
    siteId uuid.UUID,
    assetId uuid.UUID,
    body *models.Asset) (
    models.ApiResponse[models.Asset],
    error) {
    req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/assets/%v")
    req.AppendTemplateParams(siteId, assetId)
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
    
    var result models.Asset
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Asset](decoder)
    return models.NewApiResponse(result, resp), err
}
