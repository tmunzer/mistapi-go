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

// MSPsSSORoles represents a controller struct.
type MSPsSSORoles struct {
    baseController
}

// NewMSPsSSORoles creates a new instance of MSPsSSORoles.
// It takes a baseController as a parameter and returns a pointer to the MSPsSSORoles.
func NewMSPsSSORoles(baseController baseController) *MSPsSSORoles {
    mSPsSSORoles := MSPsSSORoles{baseController: baseController}
    return &mSPsSSORoles
}

// ListMspSsoRoles takes context, mspId as parameters and
// returns an models.ApiResponse with []models.SsoRoleOrg data and
// an error if there was an issue with the request or response.
// Get List of MSP SSO Roles
func (m *MSPsSSORoles) ListMspSsoRoles(
    ctx context.Context,
    mspId uuid.UUID) (
    models.ApiResponse[[]models.SsoRoleOrg],
    error) {
    req := m.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/msps/%v/ssoroles", mspId),
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
    
    var result []models.SsoRoleOrg
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.SsoRoleOrg](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateMspSsoRole takes context, mspId, body as parameters and
// returns an models.ApiResponse with models.SsoRoleMsp data and
// an error if there was an issue with the request or response.
// Create MSP Role
func (m *MSPsSSORoles) CreateMspSsoRole(
    ctx context.Context,
    mspId uuid.UUID,
    body *models.SsoRoleMsp) (
    models.ApiResponse[models.SsoRoleMsp],
    error) {
    req := m.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/msps/%v/ssoroles", mspId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SsoRoleMsp
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SsoRoleMsp](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteMspSsoRole takes context, mspId, ssoroleId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete MSP SSO Roles
func (m *MSPsSSORoles) DeleteMspSsoRole(
    ctx context.Context,
    mspId uuid.UUID,
    ssoroleId uuid.UUID) (
    *http.Response,
    error) {
    req := m.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/msps/%v/ssoroles/%v", mspId, ssoroleId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// UpdateMspSsoRole takes context, mspId, ssoroleId, body as parameters and
// returns an models.ApiResponse with models.SsoRoleMsp data and
// an error if there was an issue with the request or response.
// Update SSO Role
func (m *MSPsSSORoles) UpdateMspSsoRole(
    ctx context.Context,
    mspId uuid.UUID,
    ssoroleId uuid.UUID,
    body *models.SsoRoleMsp) (
    models.ApiResponse[models.SsoRoleMsp],
    error) {
    req := m.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/msps/%v/ssoroles/%v", mspId, ssoroleId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SsoRoleMsp
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SsoRoleMsp](decoder)
    return models.NewApiResponse(result, resp), err
}
