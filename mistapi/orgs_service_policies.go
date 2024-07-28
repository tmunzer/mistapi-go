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

// OrgsServicePolicies represents a controller struct.
type OrgsServicePolicies struct {
    baseController
}

// NewOrgsServicePolicies creates a new instance of OrgsServicePolicies.
// It takes a baseController as a parameter and returns a pointer to the OrgsServicePolicies.
func NewOrgsServicePolicies(baseController baseController) *OrgsServicePolicies {
    orgsServicePolicies := OrgsServicePolicies{baseController: baseController}
    return &orgsServicePolicies
}

// ListOrgServicePolicies takes context, orgId, page, limit as parameters and
// returns an models.ApiResponse with []models.OrgServicePolicy data and
// an error if there was an issue with the request or response.
// Get List of Org Service Policies
func (o *OrgsServicePolicies) ListOrgServicePolicies(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.OrgServicePolicy],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/servicepolicies", orgId),
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
    
    var result []models.OrgServicePolicy
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.OrgServicePolicy](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgServicePolicy takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.OrgServicePolicy data and
// an error if there was an issue with the request or response.
// Create Org Serrvice Policy
func (o *OrgsServicePolicies) CreateOrgServicePolicy(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.OrgServicePolicy) (
    models.ApiResponse[models.OrgServicePolicy],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/servicepolicies", orgId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.OrgServicePolicy
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.OrgServicePolicy](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgServicePolicy takes context, orgId, servicepolicyId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Org Service Policuy
func (o *OrgsServicePolicies) DeleteOrgServicePolicy(
    ctx context.Context,
    orgId uuid.UUID,
    servicepolicyId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/orgs/%v/servicepolicies/%v", orgId, servicepolicyId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// GetOrgServicePolicy takes context, orgId, servicepolicyId as parameters and
// returns an models.ApiResponse with models.OrgServicePolicy data and
// an error if there was an issue with the request or response.
// Get Org Service Policy Details
func (o *OrgsServicePolicies) GetOrgServicePolicy(
    ctx context.Context,
    orgId uuid.UUID,
    servicepolicyId uuid.UUID) (
    models.ApiResponse[models.OrgServicePolicy],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/servicepolicies/%v", orgId, servicepolicyId),
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
    
    var result models.OrgServicePolicy
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.OrgServicePolicy](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgServicePolicy takes context, orgId, servicepolicyId, body as parameters and
// returns an models.ApiResponse with models.OrgServicePolicy data and
// an error if there was an issue with the request or response.
// Update Org Serrvice Policy
func (o *OrgsServicePolicies) UpdateOrgServicePolicy(
    ctx context.Context,
    orgId uuid.UUID,
    servicepolicyId uuid.UUID,
    body *models.OrgServicePolicy) (
    models.ApiResponse[models.OrgServicePolicy],
    error) {
    req := o.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/orgs/%v/servicepolicies/%v", orgId, servicepolicyId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.OrgServicePolicy
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.OrgServicePolicy](decoder)
    return models.NewApiResponse(result, resp), err
}
