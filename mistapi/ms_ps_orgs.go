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

// MSPsOrgs represents a controller struct.
type MSPsOrgs struct {
    baseController
}

// NewMSPsOrgs creates a new instance of MSPsOrgs.
// It takes a baseController as a parameter and returns a pointer to the MSPsOrgs.
func NewMSPsOrgs(baseController baseController) *MSPsOrgs {
    mSPsOrgs := MSPsOrgs{baseController: baseController}
    return &mSPsOrgs
}

// ListMspOrgs takes context, mspId as parameters and
// returns an models.ApiResponse with []models.Org data and
// an error if there was an issue with the request or response.
// Get List of MSP Orgs
func (m *MSPsOrgs) ListMspOrgs(
    ctx context.Context,
    mspId uuid.UUID) (
    models.ApiResponse[[]models.Org],
    error) {
    req := m.prepareRequest(ctx, "GET", fmt.Sprintf("/api/v1/msps/%v/orgs", mspId))
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
    
    var result []models.Org
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Org](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateMspOrg takes context, mspId, body as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Create an Org under MSP
func (m *MSPsOrgs) CreateMspOrg(
    ctx context.Context,
    mspId uuid.UUID,
    body *models.Org) (
    models.ApiResponse[models.Org],
    error) {
    req := m.prepareRequest(ctx, "POST", fmt.Sprintf("/api/v1/msps/%v/orgs", mspId))
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
    
    var result models.Org
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Org](decoder)
    return models.NewApiResponse(result, resp), err
}

// ManageMspOrgs takes context, mspId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Assign or Unassign Orgs to an MSP account
func (m *MSPsOrgs) ManageMspOrgs(
    ctx context.Context,
    mspId uuid.UUID,
    body *models.MspOrgChange) (
    *http.Response,
    error) {
    req := m.prepareRequest(ctx, "PUT", fmt.Sprintf("/api/v1/msps/%v/orgs", mspId))
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// SearchMspOrgs takes context, mspId, name, orgId, subInsufficient, trialEnabled, usageTypes, limit as parameters and
// returns an models.ApiResponse with models.ResponseOrgSearch data and
// an error if there was an issue with the request or response.
// Search Org in MSP
func (m *MSPsOrgs) SearchMspOrgs(
    ctx context.Context,
    mspId uuid.UUID,
    name *string,
    orgId *uuid.UUID,
    subInsufficient *bool,
    trialEnabled *bool,
    usageTypes []string,
    limit *int) (
    models.ApiResponse[models.ResponseOrgSearch],
    error) {
    req := m.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/msps/%v/orgs/search", mspId),
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
    if name != nil {
        req.QueryParam("name", *name)
    }
    if orgId != nil {
        req.QueryParam("org_id", *orgId)
    }
    if subInsufficient != nil {
        req.QueryParam("sub_insufficient", *subInsufficient)
    }
    if trialEnabled != nil {
        req.QueryParam("trial_enabled", *trialEnabled)
    }
    if usageTypes != nil {
        req.QueryParam("usage_types", usageTypes)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    
    var result models.ResponseOrgSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseOrgSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteMspOrg takes context, mspId, orgId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// delete MSP Org
func (m *MSPsOrgs) DeleteMspOrg(
    ctx context.Context,
    mspId uuid.UUID,
    orgId uuid.UUID) (
    *http.Response,
    error) {
    req := m.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/msps/%v/orgs/%v", mspId, orgId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// GetMspOrg takes context, mspId, orgId as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Get MSP Org Details
func (m *MSPsOrgs) GetMspOrg(
    ctx context.Context,
    mspId uuid.UUID,
    orgId uuid.UUID) (
    models.ApiResponse[models.Org],
    error) {
    req := m.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/msps/%v/orgs/%v", mspId, orgId),
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
    
    var result models.Org
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Org](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateMspOrg takes context, mspId, orgId, body as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Update MSP Org
func (m *MSPsOrgs) UpdateMspOrg(
    ctx context.Context,
    mspId uuid.UUID,
    orgId uuid.UUID,
    body *models.Org) (
    models.ApiResponse[models.Org],
    error) {
    req := m.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/msps/%v/orgs/%v", mspId, orgId),
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
    
    var result models.Org
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Org](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListMspOrgStats takes context, mspId, limit, page as parameters and
// returns an models.ApiResponse with []models.StatsOrg data and
// an error if there was an issue with the request or response.
// Get List of MSP Orgs Stats
func (m *MSPsOrgs) ListMspOrgStats(
    ctx context.Context,
    mspId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.StatsOrg],
    error) {
    req := m.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/msps/%v/stats/orgs", mspId),
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
    
    var result []models.StatsOrg
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.StatsOrg](decoder)
    return models.NewApiResponse(result, resp), err
}
