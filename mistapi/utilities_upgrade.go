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

// UtilitiesUpgrade represents a controller struct.
type UtilitiesUpgrade struct {
    baseController
}

// NewUtilitiesUpgrade creates a new instance of UtilitiesUpgrade.
// It takes a baseController as a parameter and returns a pointer to the UtilitiesUpgrade.
func NewUtilitiesUpgrade(baseController baseController) *UtilitiesUpgrade {
    utilitiesUpgrade := UtilitiesUpgrade{baseController: baseController}
    return &utilitiesUpgrade
}

// ListOrgDeviceUpgrades takes context, orgId as parameters and
// returns an models.ApiResponse with []models.OrgDeviceUpgrade data and
// an error if there was an issue with the request or response.
// Get List of Org multiple devces upgrades
func (u *UtilitiesUpgrade) ListOrgDeviceUpgrades(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.OrgDeviceUpgrade],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/devices/upgrade", orgId),
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
    
    var result []models.OrgDeviceUpgrade
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.OrgDeviceUpgrade](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpgradeOrgDevices takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.ResponseUpgradeOrgDevices data and
// an error if there was an issue with the request or response.
// Upgrade Multiple Sites (Only supported for Access Points ugprades)
func (u *UtilitiesUpgrade) UpgradeOrgDevices(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.UpgradeOrgDevices) (
    models.ApiResponse[models.ResponseUpgradeOrgDevices],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/devices/upgrade", orgId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ResponseUpgradeOrgDevices
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseUpgradeOrgDevices](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetOrgDeviceUpgrade takes context, orgId, upgradeId as parameters and
// returns an models.ApiResponse with models.ResponseUpgradeOrgDevices data and
// an error if there was an issue with the request or response.
// Get Multiple Devices Upgrade
func (u *UtilitiesUpgrade) GetOrgDeviceUpgrade(
    ctx context.Context,
    orgId uuid.UUID,
    upgradeId uuid.UUID) (
    models.ApiResponse[models.ResponseUpgradeOrgDevices],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/devices/upgrade/%v", orgId, upgradeId),
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
    
    var result models.ResponseUpgradeOrgDevices
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseUpgradeOrgDevices](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpgradeOrgJsiDevice takes context, orgId, deviceMac, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Upgrade
func (u *UtilitiesUpgrade) UpgradeOrgJsiDevice(
    ctx context.Context,
    orgId uuid.UUID,
    deviceMac string,
    body *models.VersionString) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/jsi/devices/%v/upgrade", orgId, deviceMac),
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

// ListOrgMxEdgeUpgrades takes context, orgId as parameters and
// returns an models.ApiResponse with []models.ResponseMxedgeUpgrade data and
// an error if there was an issue with the request or response.
// Get List of Org Mist Edge Upgrades
func (u *UtilitiesUpgrade) ListOrgMxEdgeUpgrades(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.ResponseMxedgeUpgrade],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/mxedges/upgrade", orgId),
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
    
    var result []models.ResponseMxedgeUpgrade
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ResponseMxedgeUpgrade](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpgradeOrgMxEdges takes context, orgId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Upgrade Mist Edges
func (u *UtilitiesUpgrade) UpgradeOrgMxEdges(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.MxedgeUpgradeMulti) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/mxedges/upgrade", orgId),
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

// GetOrgMxEdgeUpgrade takes context, orgId, upgradeId as parameters and
// returns an models.ApiResponse with models.ResponseMxedgeUpgrade data and
// an error if there was an issue with the request or response.
// Get Mist Edge Upgrade
func (u *UtilitiesUpgrade) GetOrgMxEdgeUpgrade(
    ctx context.Context,
    orgId uuid.UUID,
    upgradeId uuid.UUID) (
    models.ApiResponse[models.ResponseMxedgeUpgrade],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/mxedges/upgrade/%v", orgId, upgradeId),
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
    
    var result models.ResponseMxedgeUpgrade
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseMxedgeUpgrade](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListOrgSsrUpgrades takes context, orgId as parameters and
// returns an models.ApiResponse with []models.SsrUpgradeResponse data and
// an error if there was an issue with the request or response.
// Get List of Org SSR Upgrades
func (u *UtilitiesUpgrade) ListOrgSsrUpgrades(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.SsrUpgradeResponse],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/ssr/upgrade", orgId),
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
    
    var result []models.SsrUpgradeResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.SsrUpgradeResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpgradeOrgSsrs takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.SsrUpgradeResponse data and
// an error if there was an issue with the request or response.
// Upgrade Org SSRs
func (u *UtilitiesUpgrade) UpgradeOrgSsrs(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.SsrUpgradeMulti) (
    models.ApiResponse[models.SsrUpgradeResponse],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/ssr/upgrade", orgId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SsrUpgradeResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SsrUpgradeResponse](decoder)
    return models.NewApiResponse(result, resp), err
}

// CancelOrgSsrUpgrade takes context, orgId, upgradeId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Best effort to cancel an upgrade. Devices which are already upgraded wont be touched↵
func (u *UtilitiesUpgrade) CancelOrgSsrUpgrade(
    ctx context.Context,
    orgId uuid.UUID,
    upgradeId uuid.UUID) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/orgs/%v/ssr/upgrade/%v/cancel", orgId, upgradeId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// ListOrgAvailableSsrVersions takes context, orgId, channel as parameters and
// returns an models.ApiResponse with []models.SsrVersion data and
// an error if there was an issue with the request or response.
// Get available version for SSR
func (u *UtilitiesUpgrade) ListOrgAvailableSsrVersions(
    ctx context.Context,
    orgId uuid.UUID,
    channel *string) (
    models.ApiResponse[[]models.SsrVersion],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/orgs/%v/ssr/versions", orgId),
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
    if channel != nil {
        req.QueryParam("channel", *channel)
    }
    
    var result []models.SsrVersion
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.SsrVersion](decoder)
    return models.NewApiResponse(result, resp), err
}

// ListSiteDeviceUpgrades takes context, siteId, status as parameters and
// returns an models.ApiResponse with []models.ResponseSiteDeviceUpgrade data and
// an error if there was an issue with the request or response.
// Get all upgrades for site
func (u *UtilitiesUpgrade) ListSiteDeviceUpgrades(
    ctx context.Context,
    siteId uuid.UUID,
    status *models.DeviceUpgradeStatusEnum) (
    models.ApiResponse[[]models.ResponseSiteDeviceUpgrade],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/devices/upgrade", siteId),
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
    if status != nil {
        req.QueryParam("status", *status)
    }
    
    var result []models.ResponseSiteDeviceUpgrade
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.ResponseSiteDeviceUpgrade](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpgradeSiteDevices takes context, siteId, body as parameters and
// returns an models.ApiResponse with models.ResponseUpgradeSiteDevices data and
// an error if there was an issue with the request or response.
// Upgrade Site Device
// **Note**: this call doesn’t guarantee the devices to be upgraded right away (they may be offline)
func (u *UtilitiesUpgrade) UpgradeSiteDevices(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.UpgradeSiteDevices) (
    models.ApiResponse[models.ResponseUpgradeSiteDevices],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/upgrade", siteId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ResponseUpgradeSiteDevices
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseUpgradeSiteDevices](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteDeviceUpgrade takes context, siteId, upgradeId as parameters and
// returns an models.ApiResponse with models.ResponseDeviceUpgrade data and
// an error if there was an issue with the request or response.
// Get Site Device Upgrade
func (u *UtilitiesUpgrade) GetSiteDeviceUpgrade(
    ctx context.Context,
    siteId uuid.UUID,
    upgradeId uuid.UUID) (
    models.ApiResponse[models.ResponseDeviceUpgrade],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/devices/upgrade/%v", siteId, upgradeId),
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
    
    var result models.ResponseDeviceUpgrade
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseDeviceUpgrade](decoder)
    return models.NewApiResponse(result, resp), err
}

// CancelSiteDeviceUpgrade takes context, siteId, upgradeId as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Best effort to cancel an upgrade. Devices which are already upgraded wont be touched
func (u *UtilitiesUpgrade) CancelSiteDeviceUpgrade(
    ctx context.Context,
    siteId uuid.UUID,
    upgradeId uuid.UUID) (
    *http.Response,
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/upgrade/%v/cancel", siteId, upgradeId),
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
    
    context, err := req.Call()
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// ListSiteAvailableDeviceVersions takes context, siteId, mType, model as parameters and
// returns an models.ApiResponse with []models.DeviceVersionItem data and
// an error if there was an issue with the request or response.
// Get List of Available Device Versions
func (u *UtilitiesUpgrade) ListSiteAvailableDeviceVersions(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.DeviceTypeEnum,
    model *string) (
    models.ApiResponse[[]models.DeviceVersionItem],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/devices/versions", siteId),
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
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if model != nil {
        req.QueryParam("model", *model)
    }
    
    var result []models.DeviceVersionItem
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.DeviceVersionItem](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpgradeDevice takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.ResponseUpgradeDevice data and
// an error if there was an issue with the request or response.
// Device Upgrade
func (u *UtilitiesUpgrade) UpgradeDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.DeviceUpgrade) (
    models.ApiResponse[models.ResponseUpgradeDevice],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/devices/%v/upgrade", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.ResponseUpgradeDevice
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseUpgradeDevice](decoder)
    return models.NewApiResponse(result, resp), err
}

// GetSiteSsrUpgrade takes context, siteId, upgradeId as parameters and
// returns an models.ApiResponse with models.ResponseSsrUpgradeStatus data and
// an error if there was an issue with the request or response.
// Get Specific Site SSR Upgrade
func (u *UtilitiesUpgrade) GetSiteSsrUpgrade(
    ctx context.Context,
    siteId uuid.UUID,
    upgradeId uuid.UUID) (
    models.ApiResponse[models.ResponseSsrUpgradeStatus],
    error) {
    req := u.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/ssr/upgrade/%v", siteId, upgradeId),
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
    
    var result models.ResponseSsrUpgradeStatus
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseSsrUpgradeStatus](decoder)
    return models.NewApiResponse(result, resp), err
}

// UpgradeSsr takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.SsrUpgradeResponse data and
// an error if there was an issue with the request or response.
// Upgrade Site SSR device
func (u *UtilitiesUpgrade) UpgradeSsr(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.SsrUpgrade) (
    models.ApiResponse[models.SsrUpgradeResponse],
    error) {
    req := u.prepareRequest(
      ctx,
      "POST",
      fmt.Sprintf("/api/v1/sites/%v/ssr/%v/upgrade", siteId, deviceId),
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
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(body)
    }
    
    var result models.SsrUpgradeResponse
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.SsrUpgradeResponse](decoder)
    return models.NewApiResponse(result, resp), err
}
