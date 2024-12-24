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

// OrgsDevicesOthers represents a controller struct.
type OrgsDevicesOthers struct {
    baseController
}

// NewOrgsDevicesOthers creates a new instance of OrgsDevicesOthers.
// It takes a baseController as a parameter and returns a pointer to the OrgsDevicesOthers.
func NewOrgsDevicesOthers(baseController baseController) *OrgsDevicesOthers {
    orgsDevicesOthers := OrgsDevicesOthers{baseController: baseController}
    return &orgsDevicesOthers
}

// ListOrgOtherDevices takes context, orgId, vendor, mac, serial, model, name, limit, page as parameters and
// returns an models.ApiResponse with []models.DeviceOther data and
// an error if there was an issue with the request or response.
// Get List of Org other devices (3rd party devices)
func (o *OrgsDevicesOthers) ListOrgOtherDevices(
    ctx context.Context,
    orgId uuid.UUID,
    vendor *string,
    mac *string,
    serial *string,
    model *string,
    name *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.DeviceOther],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/otherdevices")
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
    if vendor != nil {
        req.QueryParam("vendor", *vendor)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if serial != nil {
        req.QueryParam("serial", *serial)
    }
    if model != nil {
        req.QueryParam("model", *model)
    }
    if name != nil {
        req.QueryParam("name", *name)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
    }
    if page != nil {
        req.QueryParam("page", *page)
    }
    
    var result []models.DeviceOther
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.DeviceOther](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgOtherDevices takes context, orgId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Assign or unassign OtherDevices to and from a site.
func (o *OrgsDevicesOthers) UpdateOrgOtherDevices(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.OtherDeviceUpdateMulti) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/otherdevices")
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

// CountOrgOtherDeviceEvents takes context, orgId, distinct, mType, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Org OtherDevices Events
func (o *OrgsDevicesOthers) CountOrgOtherDeviceEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgOtherdevicesEventsCountDistinctEnum,
    mType *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/otherdevices/events/count")
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
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
    }
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
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgOtherDeviceEvents takes context, orgId, siteId, mac, deviceMac, model, vendor, mType, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseEventsOtherDevicesSearch data and
// an error if there was an issue with the request or response.
// Search Org OtherDevices Events
func (o *OrgsDevicesOthers) SearchOrgOtherDeviceEvents(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string,
    mac *string,
    deviceMac *string,
    model *string,
    vendor *string,
    mType *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseEventsOtherDevicesSearch],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/otherdevices/events/search")
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
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if deviceMac != nil {
        req.QueryParam("device_mac", *deviceMac)
    }
    if model != nil {
        req.QueryParam("model", *model)
    }
    if vendor != nil {
        req.QueryParam("vendor", *vendor)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
    }
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
    
    var result models.ResponseEventsOtherDevicesSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseEventsOtherDevicesSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// DeleteOrgOtherDevice takes context, orgId, deviceMac as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete OtherDevice
func (o *OrgsDevicesOthers) DeleteOrgOtherDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "DELETE", "/api/v1/orgs/%v/otherdevices/%v")
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}

// GetOrgOtherDevice takes context, orgId, deviceMac as parameters and
// returns an models.ApiResponse with models.DeviceOther data and
// an error if there was an issue with the request or response.
// Get Org other device (3rd party device)
func (o *OrgsDevicesOthers) GetOrgOtherDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    models.ApiResponse[models.DeviceOther],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/otherdevices/%v")
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
    
    var result models.DeviceOther
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.DeviceOther](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpdateOrgOtherDevice takes context, orgId, deviceMac, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// If the Site / Device cannot be identified, a manual association can be made
func (o *OrgsDevicesOthers) UpdateOrgOtherDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string,
    body *models.OtherDeviceUpdate) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "PUT", "/api/v1/orgs/%v/otherdevices/%v")
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

// RebootOrgOtherDevice takes context, orgId, deviceMac as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Reboot OtherDevice
func (o *OrgsDevicesOthers) RebootOrgOtherDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string) (
    *http.Response,
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/otherdevices/%v/reboot")
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
    
    httpCtx, err := req.Call()
    if err != nil {
        return httpCtx.Response, err
    }
    return httpCtx.Response, err
}
