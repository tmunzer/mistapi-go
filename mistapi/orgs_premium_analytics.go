package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsPremiumAnalytics represents a controller struct.
type OrgsPremiumAnalytics struct {
    baseController
}

// NewOrgsPremiumAnalytics creates a new instance of OrgsPremiumAnalytics.
// It takes a baseController as a parameter and returns a pointer to the OrgsPremiumAnalytics.
func NewOrgsPremiumAnalytics(baseController baseController) *OrgsPremiumAnalytics {
    orgsPremiumAnalytics := OrgsPremiumAnalytics{baseController: baseController}
    return &orgsPremiumAnalytics
}

// ListOrgPmaDashboards takes context, orgId, page, limit as parameters and
// returns an models.ApiResponse with []models.PmaDashboard data and
// an error if there was an issue with the request or response.
// Get List of premium analytics dashboards for this Org
func (o *OrgsPremiumAnalytics) ListOrgPmaDashboards(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.PmaDashboard],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/pma/dashboards", orgId),
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
        "401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
        "403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if page != nil {
        req.QueryParam("page", *page)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    
    var result []models.PmaDashboard
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.PmaDashboard](decoder)
    return models.NewApiResponse(result, resp), err
}
