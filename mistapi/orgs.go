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

// Orgs represents a controller struct.
type Orgs struct {
    baseController
}

// NewOrgs creates a new instance of Orgs.
// It takes a baseController as a parameter and returns a pointer to the Orgs.
func NewOrgs(baseController baseController) *Orgs {
    orgs := Orgs{baseController: baseController}
    return &orgs
}

// CreateOrg takes context, body as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Org admin can invite people to manage the org. Furthermore, he can dictate the level of security those accounts are. The check is enforced when the invited admin tries to “accept” the invitation and every time the admin tries to login
func (o *Orgs) CreateOrg(
    ctx context.Context,
    body *models.Org) (
    models.ApiResponse[models.Org],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs")
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

// DeleteOrg takes context, orgId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Org
func (o *Orgs) DeleteOrg(
    ctx context.Context,
    orgId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", fmt.Sprintf("/api/v1/orgs/%v", orgId))
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// GetOrg takes context, orgId as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Get Organization information
func (o *Orgs) GetOrg(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.Org],
    error) {
    req := o.prepareRequest(ctx, "GET", fmt.Sprintf("/api/v1/orgs/%v", orgId))
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
    
    var result models.Org
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Org](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrg takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Update Org
func (o *Orgs) UpdateOrg(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Org) (
    models.ApiResponse[models.Org],
    error) {
    req := o.prepareRequest(ctx, "PUT", fmt.Sprintf("/api/v1/orgs/%v", orgId))
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

// CloneOrg takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Org data and
// an error if there was an issue with the request or response.
// Create an Org by cloning from another one. Org Settings, Templates, Wxlan Tags, Wxlan Tunnels, Wxlan Rules, Org Wlans will be copied. Sites and Site Groups will not be copied, and therefore, the copied template will not be applied to any site/sitegroups.
func (o *Orgs) CloneOrg(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.NameString) (
    models.ApiResponse[models.Org],
    error) {
    req := o.prepareRequest(ctx, "POST", fmt.Sprintf("/api/v1/orgs/%v/clone", orgId))
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

// SearchOrgEvents takes context, orgId, mType, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseEventsOrgsSearch data and
// an error if there was an issue with the request or response.
// Search Org events
// Supported Event Types:
// - CRADLEPOINT_SYNC_FAILED
// - ORG_CA_CERT_ADDED
// - ORG_CA_CERT_REGENERATED
func (o *Orgs) SearchOrgEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mType *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseEventsOrgsSearch],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/events/search", orgId),
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
    if mType != nil {
        req.QueryParam("type", *mType)
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
    
    var result models.ResponseEventsOrgsSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseEventsOrgsSearch](decoder)
    return models.NewApiResponse(result, resp), err
}
