package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsNACFingerprints represents a controller struct.
type OrgsNACFingerprints struct {
    baseController
}

// NewOrgsNACFingerprints creates a new instance of OrgsNACFingerprints.
// It takes a baseController as a parameter and returns a pointer to the OrgsNACFingerprints.
func NewOrgsNACFingerprints(baseController baseController) *OrgsNACFingerprints {
    orgsNACFingerprints := OrgsNACFingerprints{baseController: baseController}
    return &orgsNACFingerprints
}

// CountOrgClientFingerprints takes context, siteId, distinct, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count Client Fingerprints
func (o *OrgsNACFingerprints) CountOrgClientFingerprints(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.FingerprintsCountDistinctEnum,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      "/api/v1/sites/%v/insights/fingerprints/count",
    )
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
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgClientFingerprints takes context, siteId, family, model, mfg, os, osType, mac, sort, limit, start, end, duration, interval as parameters and
// returns an models.ApiResponse with models.FingerprintSearchResult data and
// an error if there was an issue with the request or response.
// Search Client Fingerprints
func (o *OrgsNACFingerprints) SearchOrgClientFingerprints(
    ctx context.Context,
    siteId uuid.UUID,
    family *string,
    model *string,
    mfg *string,
    os *string,
    osType *string,
    mac *string,
    sort *string,
    limit *int,
    start *int,
    end *int,
    duration *string,
    interval *string) (
    models.ApiResponse[models.FingerprintSearchResult],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      "/api/v1/sites/%v/insights/fingerprints/search",
    )
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
    if family != nil {
        req.QueryParam("family", *family)
    }
    if model != nil {
        req.QueryParam("model", *model)
    }
    if mfg != nil {
        req.QueryParam("mfg", *mfg)
    }
    if os != nil {
        req.QueryParam("os", *os)
    }
    if osType != nil {
        req.QueryParam("os_type", *osType)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if sort != nil {
        req.QueryParam("sort", *sort)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if interval != nil {
        req.QueryParam("interval", *interval)
    }
    
    var result models.FingerprintSearchResult
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.FingerprintSearchResult](decoder)
    return models.NewApiResponse(result, resp), err
}
