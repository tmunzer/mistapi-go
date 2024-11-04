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

// OrgsNACCRL represents a controller struct.
type OrgsNACCRL struct {
    baseController
}

// NewOrgsNACCRL creates a new instance of OrgsNACCRL.
// It takes a baseController as a parameter and returns a pointer to the OrgsNACCRL.
func NewOrgsNACCRL(baseController baseController) *OrgsNACCRL {
    orgsNACCRL := OrgsNACCRL{baseController: baseController}
    return &orgsNACCRL
}

// GetOrgNacCrl takes context, orgId as parameters and
// returns an models.ApiResponse with models.ResponseNacCrlFiles data and
// an error if there was an issue with the request or response.
// Returns all uploaded CRL file IDs with names for the orgI
func (o *OrgsNACCRL) GetOrgNacCrl(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseNacCrlFiles],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/setting/mist_nac_crls", orgId),
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
    
    var result models.ResponseNacCrlFiles
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseNacCrlFiles](decoder)
    return models.NewApiResponse(result, resp), err
}

// ImportOrgNacCrl takes context, orgId, file, json as parameters and
// returns an models.ApiResponse with models.NacCrlFile data and
// an error if there was an issue with the request or response.
// The Import NAC Org CRL File endpoint allows users to manually upload a Certificate Revocation List (CRL) file in either PEM or DER format. This is a multipart POST request. We support one file upload per issuer, and re-uploads for the same issuer will overwrite the existing file.
func (o *OrgsNACCRL) ImportOrgNacCrl(
    ctx context.Context,
    orgId uuid.UUID,
    file *models.FileWrapper,
    json *string) (
    models.ApiResponse[models.NacCrlFile],
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/setting/mist_nac_crls", orgId),
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
    formFields := []https.FormParam{}
    if file != nil {
        fileParam := https.FormParam{Key: "file", Value: *file, Headers: http.Header{}}
        formFields = append(formFields, fileParam)
    }
    if json != nil {
        jsonParam := https.FormParam{Key: "json", Value: *json, Headers: http.Header{}}
        formFields = append(formFields, jsonParam)
    }
    req.FormData(formFields)
    
    var result models.NacCrlFile
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.NacCrlFile](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgNacCrl takes context, orgId, naccrlId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete NAC Org CRL file is a DELETE request to delete CRL file identified by its ID (ID assigned on file upload/creation)
func (o *OrgsNACCRL) DeleteOrgNacCrl(
    ctx context.Context,
    orgId uuid.UUID,
    naccrlId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      fmt.Sprintf("/api/v1/orgs/%v/setting/mist_nac_crls/%v", orgId, naccrlId),
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
