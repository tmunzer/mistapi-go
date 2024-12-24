package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsStatsSites represents a controller struct.
type OrgsStatsSites struct {
    baseController
}

// NewOrgsStatsSites creates a new instance of OrgsStatsSites.
// It takes a baseController as a parameter and returns a pointer to the OrgsStatsSites.
func NewOrgsStatsSites(baseController baseController) *OrgsStatsSites {
    orgsStatsSites := OrgsStatsSites{baseController: baseController}
    return &orgsStatsSites
}

// ListOrgSiteStats takes context, orgId, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with []models.StatsSite data and
// an error if there was an issue with the request or response.
// Get List of Org Site Stats
func (o *OrgsStatsSites) ListOrgSiteStats(
    ctx context.Context,
    orgId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.StatsSite],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/stats/sites")
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
    
    var result []models.StatsSite
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.StatsSite](decoder)
    return models.NewApiResponse(result, resp), err
}
