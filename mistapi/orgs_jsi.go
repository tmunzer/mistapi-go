package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsJSI represents a controller struct.
type OrgsJSI struct {
    baseController
}

// NewOrgsJSI creates a new instance of OrgsJSI.
// It takes a baseController as a parameter and returns a pointer to the OrgsJSI.
func NewOrgsJSI(baseController baseController) *OrgsJSI {
    orgsJSI := OrgsJSI{baseController: baseController}
    return &orgsJSI
}

// ListOrgJsiDevices takes context, orgId, limit, page, model, serial, mac as parameters and
// returns an models.ApiResponse with []models.JseDevice data and
// an error if there was an issue with the request or response.
// Get List of Org devices that connected to JSI
func (o *OrgsJSI) ListOrgJsiDevices(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int,
    model *string,
    serial *string,
    mac *string) (
    models.ApiResponse[[]models.JseDevice],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/jsi/devices")
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
    if model != nil {
        req.QueryParam("model", *model)
    }
    if serial != nil {
        req.QueryParam("serial", *serial)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    
    var result []models.JseDevice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.JseDevice](decoder)
    return models.NewApiResponse(result, resp), err
}

// AdoptOrgJsiDevice takes context, orgId as parameters and
// returns an models.ApiResponse with models.ResponseDeviceConfigCmd data and
// an error if there was an issue with the request or response.
// Adopt JSI devices
func (o *OrgsJSI) AdoptOrgJsiDevice(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseDeviceConfigCmd],
    error) {
    req := o.prepareRequest(
      ctx,
      "GET",
      "/api/v1/orgs/%v/jsi/devices/outbound_ssh_cmd",
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
    
    var result models.ResponseDeviceConfigCmd
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDeviceConfigCmd](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgJsiDeviceShellSession takes context, orgId, deviceMac as parameters and
// returns an models.ApiResponse with models.WebsocketSessionWithUrl data and
// an error if there was an issue with the request or response.
// Create Shell Session
func (o *OrgsJSI) CreateOrgJsiDeviceShellSession(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    models.ApiResponse[models.WebsocketSessionWithUrl],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/jsi/devices/%v/shell")
    req.AppendTemplateParams(orgId, deviceMac)
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
    
    var result models.WebsocketSessionWithUrl
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.WebsocketSessionWithUrl](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListOrgJsiPastPurchases takes context, orgId, limit, page, model, serial as parameters and
// returns an models.ApiResponse with []models.JsInventoryItem data and
// an error if there was an issue with the request or response.
// This gets all devices purchased from the accounts associated with the Org 
// * Fetch Install base devices for all linked accounts and associated account of the linked accounts. 
// * The primary and the associated account ids will be queries from SFDC by passing the linked account 
// * Returns only the device centric details of the Install base device. No customer specific information will be returned.
func (o *OrgsJSI) ListOrgJsiPastPurchases(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int,
    model *string,
    serial *string) (
    models.ApiResponse[[]models.JsInventoryItem],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/jsi/inventory")
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
        "400": {Message: "Bad Request - no Juniper Account Linked", Unmarshaller: errors.NewResponseDetailString},
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
    if model != nil {
        req.QueryParam("model", *model)
    }
    if serial != nil {
        req.QueryParam("serial", *serial)
    }
    
    var result []models.JsInventoryItem
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.JsInventoryItem](decoder)
    return models.NewApiResponse(result, resp), err
}
