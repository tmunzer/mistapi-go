package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsSites represents a controller struct.
type OrgsSites struct {
    baseController
}

// NewOrgsSites creates a new instance of OrgsSites.
// It takes a baseController as a parameter and returns a pointer to the OrgsSites.
func NewOrgsSites(baseController baseController) *OrgsSites {
    orgsSites := OrgsSites{baseController: baseController}
    return &orgsSites
}

// ListOrgSites takes context, orgId, limit, page as parameters and
// returns an models.ApiResponse with []models.Site data and
// an error if there was an issue with the request or response.
// Get List of Org Sites
func (o *OrgsSites) ListOrgSites(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Site],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sites")
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
    
    var result []models.Site
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[[]models.Site](decoder)
    return models.NewApiResponse(result, resp), err
}

// CreateOrgSite takes context, orgId, body as parameters and
// returns an models.ApiResponse with models.Site data and
// an error if there was an issue with the request or response.
// Create Org Site
func (o *OrgsSites) CreateOrgSite(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Site) (
    models.ApiResponse[models.Site],
    error) {
    req := o.prepareRequest(ctx, "POST", "/api/v1/orgs/%v/sites")
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
    
    var result models.Site
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.Site](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountOrgSites takes context, orgId, distinct, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count Sites
func (o *OrgsSites) CountOrgSites(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgSitesCountDistinctEnum,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sites/count")
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
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgSites takes context, orgId, analyticEnabled, appWaking, assetEnabled, autoUpgradeEnabled, autoUpgradeVersion, countryCode, honeypotEnabled, id, locateUnconnected, meshEnabled, name, rogueEnabled, remoteSyslogEnabled, rtsaEnabled, vnaEnabled, wifiEnabled, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseSiteSearch data and
// an error if there was an issue with the request or response.
// Search Sites
func (o *OrgsSites) SearchOrgSites(
    ctx context.Context,
    orgId uuid.UUID,
    analyticEnabled *bool,
    appWaking *bool,
    assetEnabled *bool,
    autoUpgradeEnabled *bool,
    autoUpgradeVersion *string,
    countryCode *string,
    honeypotEnabled *bool,
    id *string,
    locateUnconnected *bool,
    meshEnabled *bool,
    name *string,
    rogueEnabled *bool,
    remoteSyslogEnabled *bool,
    rtsaEnabled *bool,
    vnaEnabled *bool,
    wifiEnabled *bool,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseSiteSearch],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/sites/search")
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
    if analyticEnabled != nil {
        req.QueryParam("analytic_enabled", *analyticEnabled)
    }
    if appWaking != nil {
        req.QueryParam("app_waking", *appWaking)
    }
    if assetEnabled != nil {
        req.QueryParam("asset_enabled", *assetEnabled)
    }
    if autoUpgradeEnabled != nil {
        req.QueryParam("auto_upgrade_enabled", *autoUpgradeEnabled)
    }
    if autoUpgradeVersion != nil {
        req.QueryParam("auto_upgrade_version", *autoUpgradeVersion)
    }
    if countryCode != nil {
        req.QueryParam("country_code", *countryCode)
    }
    if honeypotEnabled != nil {
        req.QueryParam("honeypot_enabled", *honeypotEnabled)
    }
    if id != nil {
        req.QueryParam("id", *id)
    }
    if locateUnconnected != nil {
        req.QueryParam("locate_unconnected", *locateUnconnected)
    }
    if meshEnabled != nil {
        req.QueryParam("mesh_enabled", *meshEnabled)
    }
    if name != nil {
        req.QueryParam("name", *name)
    }
    if rogueEnabled != nil {
        req.QueryParam("rogue_enabled", *rogueEnabled)
    }
    if remoteSyslogEnabled != nil {
        req.QueryParam("remote_syslog_enabled", *remoteSyslogEnabled)
    }
    if rtsaEnabled != nil {
        req.QueryParam("rtsa_enabled", *rtsaEnabled)
    }
    if vnaEnabled != nil {
        req.QueryParam("vna_enabled", *vnaEnabled)
    }
    if wifiEnabled != nil {
        req.QueryParam("wifi_enabled", *wifiEnabled)
    }
    if limit != nil {
        req.QueryParam("limit", *limit)
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
    
    var result models.ResponseSiteSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseSiteSearch](decoder)
    return models.NewApiResponse(result, resp), err
}
