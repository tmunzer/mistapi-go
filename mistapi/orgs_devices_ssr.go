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

// OrgsDevicesSSR represents a controller struct.
type OrgsDevicesSSR struct {
    baseController
}

// NewOrgsDevicesSSR creates a new instance of OrgsDevicesSSR.
// It takes a baseController as a parameter and returns a pointer to the OrgsDevicesSSR.
func NewOrgsDevicesSSR(baseController baseController) *OrgsDevicesSSR {
    orgsDevicesSSR := OrgsDevicesSSR{baseController: baseController}
    return &orgsDevicesSSR
}

// GetOrg128TRegistrationCommands takes context, orgId as parameters and
// returns an models.ApiResponse with models.ReponseRouter128TRegisterCmd data and
// an error if there was an issue with the request or response.
// 128T devices can be managed/adopted by Mist.
func (o *OrgsDevicesSSR) GetOrg128TRegistrationCommands(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ReponseRouter128TRegisterCmd],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/128routers/register_cmd", orgId),
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
    
    var result models.ReponseRouter128TRegisterCmd
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ReponseRouter128TRegisterCmd](decoder)
    return models.NewApiResponse(result, resp), err
}
