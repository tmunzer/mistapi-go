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

// OrgsRFTemplates represents a controller struct.
type OrgsRFTemplates struct {
    baseController
}

// NewOrgsRFTemplates creates a new instance of OrgsRFTemplates.
// It takes a baseController as a parameter and returns a pointer to the OrgsRFTemplates.
func NewOrgsRFTemplates(baseController baseController) *OrgsRFTemplates {
    orgsRFTemplates := OrgsRFTemplates{baseController: baseController}
    return &orgsRFTemplates
}

// ListOrgRfTemplates takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.RfTemplate data and
// an error if there was an issue with the request or response.
// Get List of Org RF Template
func (o *OrgsRFTemplates) ListOrgRfTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.RfTemplate],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/rftemplates")
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
    
    var result []models.RfTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.RfTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgRfTemplate takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.RfTemplate data and
// an error if there was an issue with the request or response.
// Create Org RF Template
func (o *OrgsRFTemplates) CreateOrgRfTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.RfTemplate) (
    models.ApiResponse[models.RfTemplate],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/rftemplates")
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
    
    var result models.RfTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RfTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgRfTemplate takes context, orgId, rftemplateId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org RF Template
func (o *OrgsRFTemplates) DeleteOrgRfTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    rftemplateId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/rftemplates/%v")
    req.AppendTemplateParams(orgId, rftemplateId)
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

// GetOrgRfTemplate takes context, orgId, rftemplateId as parameters and
// returns an models.ApiResponse with models.RfTemplate data and
// an error if there was an issue with the request or response.
// Get Org RF Template Details
func (o *OrgsRFTemplates) GetOrgRfTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    rftemplateId uuid.UUID) (
    models.ApiResponse[models.RfTemplate],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/rftemplates/%v")
    req.AppendTemplateParams(orgId, rftemplateId)
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
    
    var result models.RfTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RfTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgRfTemplate takes context, orgId, rftemplateId, body as parameters and
// returns an models.ApiResponse with models.RfTemplate data and
// an error if there was an issue with the request or response.
// Update Org RF Template
func (o *OrgsRFTemplates) UpdateOrgRfTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    rftemplateId uuid.UUID,
    body *models.RfTemplate) (
    models.ApiResponse[models.RfTemplate],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/rftemplates/%v")
    req.AppendTemplateParams(orgId, rftemplateId)
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
    
    var result models.RfTemplate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RfTemplate](decoder)
    return models.NewApiResponse(result, resp), err
}
