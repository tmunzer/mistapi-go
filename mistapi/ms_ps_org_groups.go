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

// MSPsOrgGroups represents a controller struct.
type MSPsOrgGroups struct {
    baseController
}

// NewMSPsOrgGroups creates a new instance of MSPsOrgGroups.
// It takes a baseController as a parameter and returns a pointer to the MSPsOrgGroups.
func NewMSPsOrgGroups(baseController baseController) *MSPsOrgGroups {
    mSPsOrgGroups := MSPsOrgGroups{baseController: baseController}
    return &mSPsOrgGroups
}

// ListMspOrgGroups takes context, mspId as parameters and
// returns an models.ApiResponse with []models.Orggroup data and
// an error if there was an issue with the request or response.
// Get List of MSP Org Groups
func (m *MSPsOrgGroups) ListMspOrgGroups(
    ctx context.Context,
    mspId uuid.UUID) (
    models.ApiResponse[[]models.Orggroup],
    error) {
    req := m.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/msps/%v/orggroups", mspId),
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
    
    var result []models.Orggroup
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Orggroup](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateMspOrgGroup takes context, mspId, body as parameters and
// returns an models.ApiResponse with models.Orggroup data and
// an error if there was an issue with the request or response.
// Create MSP Org Group
func (m *MSPsOrgGroups) CreateMspOrgGroup(
    ctx context.Context,
    mspId uuid.UUID,
    body *models.Orggroup) (
    models.ApiResponse[models.Orggroup],
    error) {
    req := m.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/msps/%v/orggroups", mspId),
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
    
    var result models.Orggroup
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Orggroup](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteMspOrgGroup takes context, mspId, orggroupId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete MSP Org Group
func (m *MSPsOrgGroups) DeleteMspOrgGroup(
    ctx context.Context,
    mspId uuid.UUID,
    orggroupId uuid.UUID) (
    *http.Response,
    error) {
    req := m.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/msps/%v/orggroups/%v", mspId, orggroupId),
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

// GetMspOrgGroup takes context, mspId, orggroupId as parameters and
// returns an models.ApiResponse with models.Orggroup data and
// an error if there was an issue with the request or response.
// Get MSP Org Group Details
func (m *MSPsOrgGroups) GetMspOrgGroup(
    ctx context.Context,
    mspId uuid.UUID,
    orggroupId uuid.UUID) (
    models.ApiResponse[models.Orggroup],
    error) {
    req := m.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/msps/%v/orggroups/%v", mspId, orggroupId),
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
    
    var result models.Orggroup
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Orggroup](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateMspOrgGroup takes context, mspId, orggroupId, body as parameters and
// returns an models.ApiResponse with models.Orggroup data and
// an error if there was an issue with the request or response.
// Update MSP Org Group
func (m *MSPsOrgGroups) UpdateMspOrgGroup(
    ctx context.Context,
    mspId uuid.UUID,
    orggroupId uuid.UUID,
    body *models.Orggroup) (
    models.ApiResponse[models.Orggroup],
    error) {
    req := m.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/msps/%v/orggroups/%v", mspId, orggroupId),
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
    
    var result models.Orggroup
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Orggroup](decoder)
    return models.NewApiResponse(result, resp), err
}
