package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "mistapi/errors"
    "mistapi/models"
)

// OrgsVars represents a controller struct.
type OrgsVars struct {
    baseController
}

// NewOrgsVars creates a new instance of OrgsVars.
// It takes a baseController as a parameter and returns a pointer to the OrgsVars.
func NewOrgsVars(baseController baseController) *OrgsVars {
    orgsVars := OrgsVars{baseController: baseController}
    return &orgsVars
}

// SearchOrgVars takes context, orgId, siteId, vars, src, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseSearchVar data and
// an error if there was an issue with the request or response.
// Search vars
// Example: /api/v1/orgs/:org_id/vars/search?vars=*
func (o *OrgsVars) SearchOrgVars(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string,
    vars *string,
    src *models.VarSourceEnum,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseSearchVar],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/vars/search", orgId),
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
        "404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
        "429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
    })
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if vars != nil {
        req.QueryParam("vars", *vars)
    }
    if src != nil {
        req.QueryParam("src", *src)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result models.ResponseSearchVar
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseSearchVar](decoder)
    return models.NewApiResponse(result, resp), err
}
