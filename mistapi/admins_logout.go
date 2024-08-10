package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// AdminsLogout represents a controller struct.
type AdminsLogout struct {
    baseController
}

// NewAdminsLogout creates a new instance of AdminsLogout.
// It takes a baseController as a parameter and returns a pointer to the AdminsLogout.
func NewAdminsLogout(baseController baseController) *AdminsLogout {
    adminsLogout := AdminsLogout{baseController: baseController}
    return &adminsLogout
}

// Logout takes context as parameters and
// returns an models.ApiResponse with models.ResponseLogout data and
// an error if there was an issue with the request or response.
// Logout
func (a *AdminsLogout) Logout(ctx context.Context) (
    models.ApiResponse[models.ResponseLogout],
    error) {
    req := a.prepareRequest(ctx, "POST", "/api/v1/logout")
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
    var result models.ResponseLogout
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseLogout](decoder)
    return models.NewApiResponse(result, resp), err
}
