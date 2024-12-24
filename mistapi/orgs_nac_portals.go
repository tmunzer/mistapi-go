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

// OrgsNACPortals represents a controller struct.
type OrgsNACPortals struct {
    baseController
}

// NewOrgsNACPortals creates a new instance of OrgsNACPortals.
// It takes a baseController as a parameter and returns a pointer to the OrgsNACPortals.
func NewOrgsNACPortals(baseController baseController) *OrgsNACPortals {
    orgsNACPortals := OrgsNACPortals{baseController: baseController}
    return &orgsNACPortals
}

// ListOrgNacPortals takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.NacPortal data and
// an error if there was an issue with the request or response.
// List Org NAC Portals
func (o *OrgsNACPortals) ListOrgNacPortals(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.NacPortal],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/nacportals")
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
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result []models.NacPortal
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.NacPortal](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgNacPortal takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.NacPortal data and
// an error if there was an issue with the request or response.
// Create Org NAC Portal
func (o *OrgsNACPortals) CreateOrgNacPortal(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.NacPortal) (
    models.ApiResponse[models.NacPortal],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/nacportals")
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
    
    var result models.NacPortal
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.NacPortal](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgNacPortal takes context, orgId, nacportalId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org NAC Portal
func (o *OrgsNACPortals) DeleteOrgNacPortal(
    ctx context.Context,
    orgId uuid.UUID,
    nacportalId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/nacportals/%v")
    req.AppendTemplateParams(orgId, nacportalId)
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

// GetOrgNacPortal takes context, orgId, nacportalId as parameters and
// returns an models.ApiResponse with models.NacPortal data and
// an error if there was an issue with the request or response.
// Get Org NAC Portal
func (o *OrgsNACPortals) GetOrgNacPortal(
    ctx context.Context,
    orgId uuid.UUID,
    nacportalId uuid.UUID) (
    models.ApiResponse[models.NacPortal],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/nacportals/%v")
    req.AppendTemplateParams(orgId, nacportalId)
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
    
    var result models.NacPortal
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.NacPortal](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgNacPortal takes context, orgId, nacportalId, body as parameters and
// returns an models.ApiResponse with models.NacPortal data and
// an error if there was an issue with the request or response.
// Update Org NAC Portal
func (o *OrgsNACPortals) UpdateOrgNacPortal(
    ctx context.Context,
    orgId uuid.UUID,
    nacportalId uuid.UUID,
    body *models.NacPortal) (
    models.ApiResponse[models.NacPortal],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/nacportals/%v")
    req.AppendTemplateParams(orgId, nacportalId)
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
    
    var result models.NacPortal
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.NacPortal](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListOrgNacPortalSsoLatestFailures takes context, orgId, nacportalId, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseSsoFailureSearch data and
// an error if there was an issue with the request or response.
// Get List of Org NAC Portal SSO Latest Failures
func (o *OrgsNACPortals) ListOrgNacPortalSsoLatestFailures(
    ctx context.Context,
    orgId uuid.UUID,
    nacportalId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseSsoFailureSearch],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/nacportals/%v/failures")
    req.AppendTemplateParams(orgId, nacportalId)
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
    if start != nil {
        req.QueryParam("start", *start)
    }
    if end != nil {
        req.QueryParam("end", *end)
    }
    if duration != nil {
        req.QueryParam("duration", *duration)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result models.ResponseSsoFailureSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseSsoFailureSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgNacPortalImage takes context, orgId, nacportalId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete background image for NAC Portal
// If image is not uploaded or is deleted, NAC Portal will use default image.
func (o *OrgsNACPortals) DeleteOrgNacPortalImage(
    ctx context.Context,
    orgId uuid.UUID,
    nacportalId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "DELETE",
      "/api/v1/orgs/%v/nacportals/%v/portal_image",
    )
    req.AppendTemplateParams(orgId, nacportalId)
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

// UploadOrgNacPortalImage takes context, orgId, nacportalId, file, json as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Upload background image for NAC Portal
func (o *OrgsNACPortals) UploadOrgNacPortalImage(
    ctx context.Context,
    orgId uuid.UUID,
    nacportalId uuid.UUID,
    file *models.FileWrapper,
    json *string) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "POST",
      "/api/v1/orgs/%v/nacportals/%v/portal_image",
    )
    req.AppendTemplateParams(orgId, nacportalId)
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// UpdateOrgNacPortalTempalte takes context, orgId, nacportalId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Update Org NAC Portal Template
func (o *OrgsNACPortals) UpdateOrgNacPortalTempalte(
    ctx context.Context,
    orgId uuid.UUID,
    nacportalId uuid.UUID,
    body *models.NacPortalTemplate) (
    *http.Response,
    error) {
    req := o.prepareRequest(
      ctx,
      "PUT",
      "/api/v1/orgs/%v/nacportals/%v/portal_template",
    )
    req.AppendTemplateParams(orgId, nacportalId)
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

// GetOrgNacPortalSamlMetadata takes context, orgId, nacportalId as parameters and
// returns an models.ApiResponse with models.SamlMetadata data and
// an error if there was an issue with the request or response.
// Get Org NAC Portal SAML Metadata
func (o *OrgsNACPortals) GetOrgNacPortalSamlMetadata(
    ctx context.Context,
    orgId uuid.UUID,
    nacportalId uuid.UUID) (
    models.ApiResponse[models.SamlMetadata],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      "/api/v1/orgs/%v/nacportals/%v/saml_metadata",
    )
    req.AppendTemplateParams(orgId, nacportalId)
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
    
    var result models.SamlMetadata
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SamlMetadata](decoder)
    return models.NewApiResponse(result, resp), err
}

// DownloadOrgNacPortalSamlMetadata takes context, orgId, nacportalId as parameters and
// returns an models.ApiResponse with []byte data and
// an error if there was an issue with the request or response.
// Download Org NAC Portal SAML Metdata
// Example of metadata.xml:
// ```xml
// <?xml version="1.0" encoding="UTF-8"?><md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://api.mist.com/api/v1/saml/5hdF5g/login" validUntil="2027-10-12T21:59:01Z" xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
// <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="true" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
// <md:SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://api.mist.com/api/v1/saml/5hdF5g/logout" />
// <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
// <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://api.mist.com/api/v1/saml/5hdF5g/login" index="0" isDefault="true"/>
// <md:AttributeConsumingService index="0">
// <md:ServiceName xml:lang="en-US">Mist</md:ServiceName>
// <md:RequestedAttribute Name="Role" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic" isRequired="true"/>
// <md:RequestedAttribute Name="FirstName" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic" isRequired="false"/>
// <md:RequestedAttribute Name="LastName" NameFormat="urn:oasis:names:tc:SAML:2.0:attrname-format:basic" isRequired="false"/>
// </md:AttributeConsumingService>
// </md:SPSSODescriptor>
// </md:EntityDescriptor>
// ```
func (o *OrgsNACPortals) DownloadOrgNacPortalSamlMetadata(
    ctx context.Context,
    orgId uuid.UUID,
    nacportalId uuid.UUID) (
    models.ApiResponse[[]byte],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      "/api/v1/orgs/%v/nacportals/%v/saml_metadata.xml",
    )
    req.AppendTemplateParams(orgId, nacportalId)
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
    
    stream, resp, err := req.CallAsStream()
    if err != nil {
        return models.NewApiResponse(stream, resp), err
    }
    return models.NewApiResponse(stream, resp), err
}
