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

// OrgsCert represents a controller struct.
type OrgsCert struct {
    baseController
}

// NewOrgsCert creates a new instance of OrgsCert.
// It takes a baseController as a parameter and returns a pointer to the OrgsCert.
func NewOrgsCert(baseController baseController) *OrgsCert {
    orgsCert := OrgsCert{baseController: baseController}
    return &orgsCert
}

// GetOrgCertificates takes context, orgId as parameters and
// returns an models.ApiResponse with models.ResponseCertificate data and
// an error if there was an issue with the request or response.
// Get Org Certificates
func (o *OrgsCert) GetOrgCertificates(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseCertificate],
    error) {
    req := o.prepareRequest(ctx, "GET", fmt.Sprintf("/api/v1/orgs/%v/cert", orgId))
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
    
    var result models.ResponseCertificate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCertificate](decoder)
    return models.NewApiResponse(result, resp), err
}

// ClearOrgCertificates takes context, orgId as parameters and
// returns an models.ApiResponse with models.ResponseCertificate data and
// an error if there was an issue with the request or response.
// Clear Org Certificates
func (o *OrgsCert) ClearOrgCertificates(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseCertificate],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/cert/regenerate", orgId),
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
    
    var result models.ResponseCertificate
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCertificate](decoder)
    return models.NewApiResponse(result, resp), err
}

// TruncateOrgCrlFile takes context, orgId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// By default, all certs used by recently unclaimed devices within 9 month will be included in CRL. If the list grows too big, you can truncate it
func (o *OrgsCert) TruncateOrgCrlFile(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.DaysNumber) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/crl/truncate", orgId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}
