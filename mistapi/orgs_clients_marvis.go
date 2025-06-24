package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "net/http"
)

// OrgsClientsMarvis represents a controller struct.
type OrgsClientsMarvis struct {
    baseController
}

// NewOrgsClientsMarvis creates a new instance of OrgsClientsMarvis.
// It takes a baseController as a parameter and returns a pointer to the OrgsClientsMarvis.
func NewOrgsClientsMarvis(baseController baseController) *OrgsClientsMarvis {
    orgsClientsMarvis := OrgsClientsMarvis{baseController: baseController}
    return &orgsClientsMarvis
}

// DeleteOrgMarvisClient takes context, orgId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Marvis Client
func (o *OrgsClientsMarvis) DeleteOrgMarvisClient(
    ctx context.Context,
    orgId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/stats/marvisclients")
    req.AppendTemplateParams(orgId)
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
