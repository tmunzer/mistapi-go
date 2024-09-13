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

// OrgsSecIntelProfiles represents a controller struct.
type OrgsSecIntelProfiles struct {
    baseController
}

// NewOrgsSecIntelProfiles creates a new instance of OrgsSecIntelProfiles.
// It takes a baseController as a parameter and returns a pointer to the OrgsSecIntelProfiles.
func NewOrgsSecIntelProfiles(baseController baseController) *OrgsSecIntelProfiles {
    orgsSecIntelProfiles := OrgsSecIntelProfiles{baseController: baseController}
    return &orgsSecIntelProfiles
}

// ListOrgSecIntelProfiles takes context, orgId as parameters and
// returns an models.ApiResponse with []models.SecintelProfile data and
// an error if there was an issue with the request or response.
// Get List of Sec Intel Profiles
func (o *OrgsSecIntelProfiles) ListOrgSecIntelProfiles(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.SecintelProfile],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/secintelprofiles", orgId),
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
    
    var result []models.SecintelProfile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.SecintelProfile](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgSecIntelProfile takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.SecintelProfile data and
// an error if there was an issue with the request or response.
// Create Sec Intel Profiles
func (o *OrgsSecIntelProfiles) CreateOrgSecIntelProfile(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.SecintelProfile) (
    models.ApiResponse[models.SecintelProfile],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/secintelprofiles", orgId),
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
    
    var result models.SecintelProfile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SecintelProfile](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgSecIntelProfile takes context, orgId, secintelprofileId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete Sec Intel Profile
func (o *OrgsSecIntelProfiles) DeleteOrgSecIntelProfile(
    ctx context.Context,
    orgId uuid.UUID,
    secintelprofileId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/orgs/%v/secintelprofiles/%v", orgId, secintelprofileId),
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

// GetOrgSecIntelProfile takes context, orgId, secintelprofileId as parameters and
// returns an models.ApiResponse with models.SecintelProfile data and
// an error if there was an issue with the request or response.
// Get Sec Intel Profile
func (o *OrgsSecIntelProfiles) GetOrgSecIntelProfile(
    ctx context.Context,
    orgId uuid.UUID,
    secintelprofileId uuid.UUID) (
    models.ApiResponse[models.SecintelProfile],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/secintelprofiles/%v", orgId, secintelprofileId),
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
    
    var result models.SecintelProfile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SecintelProfile](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgSecIntelProfile takes context, orgId, secintelprofileId, body as parameters and
// returns an models.ApiResponse with models.SecintelProfile data and
// an error if there was an issue with the request or response.
// Update Sec Intel Profile
func (o *OrgsSecIntelProfiles) UpdateOrgSecIntelProfile(
    ctx context.Context,
    orgId uuid.UUID,
    secintelprofileId uuid.UUID,
    body *models.SecintelProfile) (
    models.ApiResponse[models.SecintelProfile],
    error) {
    req := o.prepareRequest(
      ctx,
      "PUT",
      fmt.Sprintf("/api/v1/orgs/%v/secintelprofiles/%v", orgId, secintelprofileId),
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
    
    var result models.SecintelProfile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SecintelProfile](decoder)
    return models.NewApiResponse(result, resp), err
}
