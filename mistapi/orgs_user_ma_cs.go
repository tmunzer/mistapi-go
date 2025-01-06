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

// OrgsUserMACs represents a controller struct.
type OrgsUserMACs struct {
    baseController
}

// NewOrgsUserMACs creates a new instance of OrgsUserMACs.
// It takes a baseController as a parameter and returns a pointer to the OrgsUserMACs.
func NewOrgsUserMACs(baseController baseController) *OrgsUserMACs {
    orgsUserMACs := OrgsUserMACs{baseController: baseController}
    return &orgsUserMACs
}

// CreateOrgUserMacs takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.UserMac data and
// an error if there was an issue with the request or response.
// Create Org User MACs
// ### Usermacs import CSV file format
// mac,labels,vlan,notes 
// 921b638445cd,”bldg1,flor1”,vlan-100 
// 721b638445ef,”bldg2,flor2”,vlan-101,Canon Printers 
// 721b638445ee,”bldg3,flor3”,vlan-102 
// 921b638445ce,”bldg4,flor4”,vlan-103 
// 921b638445cf,”bldg5,flor5”,vlan-104
func (o *OrgsUserMACs) CreateOrgUserMacs(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.UserMac) (
    models.ApiResponse[models.UserMac],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/usermacs")
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
    
    var result models.UserMac
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.UserMac](decoder)
    return models.NewApiResponse(result, resp), err
}

// ImportOrgUserMacs takes context, orgId, file as parameters and
// returns an models.ApiResponse with models.UserMacImport data and
// an error if there was an issue with the request or response.
// Import Org User MACs
// ### CSV Import example
// ```csv 
// mac,labels,vlan,notes,name,radius_group
// 921b638445cd,"bldg1,flor1",vlan-100
// 721b638445ef,"bldg2,flor2",vlan-101,Canon Printers
// 721b638445ee,"bldg3,flor3",vlan-102,Printer2,VIP
// 921b638445ce,"bldg4,flor4",vlan-103
// 921b638445cf,"bldg5,flor5",vlan-104
// ````
func (o *OrgsUserMACs) ImportOrgUserMacs(
    ctx context.Context,
    orgId uuid.UUID,
    file models.FileWrapper) (
    models.ApiResponse[models.UserMacImport],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/usermacs/import")
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
    formFields := []https.FormParam{}
    fileParam := https.FormParam{Key: "file", Value: file, Headers: http.Header{}}
    formFields = append(formFields, fileParam)
    req.FormData(formFields)
    
    var result models.UserMacImport
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.UserMacImport](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgUserMacs takes context, orgId, mac, labels, limit, page as parameters and
// returns an models.ApiResponse with models.UserMacsSearch data and
// an error if there was an issue with the request or response.
// Search Org User MACs
func (o *OrgsUserMACs) SearchOrgUserMacs(
    ctx context.Context,
    orgId uuid.UUID,
    mac *string,
    labels []string,
    limit *int,
    page *int) (
    models.ApiResponse[models.UserMacsSearch],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/usermacs/search")
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
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if labels != nil {
        req.QueryParam("labels", labels)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result models.UserMacsSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.UserMacsSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgUserMac takes context, orgId, usermacId as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete Org User MAC
func (o *OrgsUserMACs) DeleteOrgUserMac(
    ctx context.Context,
    orgId uuid.UUID,
    usermacId uuid.UUID) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/usermacs/%v")
    req.AppendTemplateParams(orgId, usermacId)
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

// GetOrgUserMac takes context, orgId, usermacId as parameters and
// returns an models.ApiResponse with models.UserMac data and
// an error if there was an issue with the request or response.
// Get Org User MAC
func (o *OrgsUserMACs) GetOrgUserMac(
    ctx context.Context,
    orgId uuid.UUID,
    usermacId uuid.UUID) (
    models.ApiResponse[models.UserMac],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/usermacs/%v")
    req.AppendTemplateParams(orgId, usermacId)
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
    
    var result models.UserMac
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.UserMac](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgUserMac takes context, orgId, usermacId, body as parameters and
// returns an models.ApiResponse with models.UserMac data and
// an error if there was an issue with the request or response.
// Update Org User MAC
func (o *OrgsUserMACs) UpdateOrgUserMac(
    ctx context.Context,
    orgId uuid.UUID,
    usermacId uuid.UUID,
    body *models.UserMac) (
    models.ApiResponse[models.UserMac],
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/usermacs/%v")
    req.AppendTemplateParams(orgId, usermacId)
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
    
    var result models.UserMac
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.UserMac](decoder)
    return models.NewApiResponse(result, resp), err
}
