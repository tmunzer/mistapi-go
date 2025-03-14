package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
    "net/http"
)

// OrgsSCEP represents a controller struct.
type OrgsSCEP struct {
    baseController
}

// NewOrgsSCEP creates a new instance of OrgsSCEP.
// It takes a baseController as a parameter and returns a pointer to the OrgsSCEP.
func NewOrgsSCEP(baseController baseController) *OrgsSCEP {
    orgsSCEP := OrgsSCEP{baseController: baseController}
    return &orgsSCEP
}

// DisableOrgMistScep takes context, orgId as parameters and
// returns an models.ApiResponse with models.OrgSettingScep data and
// an error if there was an issue with the request or response.
// Disable Mist SCEP Org setting
func (o *OrgsSCEP) DisableOrgMistScep(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.OrgSettingScep],
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/setting/mist_scep")
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
    
    var result models.OrgSettingScep
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.OrgSettingScep](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetOrgMistScep takes context, orgId as parameters and
// returns an models.ApiResponse with models.OrgSettingScep data and
// an error if there was an issue with the request or response.
// Get Mist SCEP Org setting
func (o *OrgsSCEP) GetOrgMistScep(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.OrgSettingScep],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/setting/mist_scep")
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
    
    var result models.OrgSettingScep
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.OrgSettingScep](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgMistScep takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.OrgSettingScep data and
// an error if there was an issue with the request or response.
// Update Mist SCEP Org setting
func (o *OrgsSCEP) UpdateOrgMistScep(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.OrgSettingScep) (
    models.ApiResponse[models.OrgSettingScep],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/setting/mist_scep")
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.OrgSettingScep
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.OrgSettingScep](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListOrgIssuedClientCertificates takes context, orgId, ssoNameId, serialNumber, deviceId as parameters and
// returns an models.ApiResponse with models.IssuedClientCertificatesResults data and
// an error if there was an issue with the request or response.
// Get Issued Client Certificates
func (o *OrgsSCEP) ListOrgIssuedClientCertificates(
    ctx context.Context,
    orgId uuid.UUID,
    ssoNameId *string,
    serialNumber *string,
    deviceId *string) (
    models.ApiResponse[models.IssuedClientCertificatesResults],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      "/api/v1/orgs/%v/setting/mist_scep/client_certs",
    )
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
    if ssoNameId != nil {
        req.QueryParam("sso_name_id", *ssoNameId)
    }
    if serialNumber != nil {
        req.QueryParam("serial_number", *serialNumber)
    }
    if deviceId != nil {
        req.QueryParam("device_id", *deviceId)
    }
    
    var result models.IssuedClientCertificatesResults
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.IssuedClientCertificatesResults](decoder)
    return models.NewApiResponse(result, resp), err
}

// RevokeOrgIssuedClientCertificates takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Revoke Issued Client Certificates
func (o *OrgsSCEP) RevokeOrgIssuedClientCertificates(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.ClientCertSerialNumbers) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      "/api/v1/orgs/%v/setting/mist_scep/client_certs/revoke",
    )
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}
