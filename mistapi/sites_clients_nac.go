package mistapi

import (
    "context"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// SitesClientsNAC represents a controller struct.
type SitesClientsNAC struct {
    baseController
}

// NewSitesClientsNAC creates a new instance of SitesClientsNAC.
// It takes a baseController as a parameter and returns a pointer to the SitesClientsNAC.
func NewSitesClientsNAC(baseController baseController) *SitesClientsNAC {
    sitesClientsNAC := SitesClientsNAC{baseController: baseController}
    return &sitesClientsNAC
}

// CountSiteNacClients takes context, siteId, distinct, lastNacruleId, nacruleMatched, authType, lastVlanId, lastNasVendor, idpId, lastSsid, lastUsername, timestamp, lastAp, mac, lastStatus, mType, mdmComplianceStatus, mdmProvider, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of NAC Clients
func (s *SitesClientsNAC) CountSiteNacClients(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteNacClientsCountDistinctEnum,
    lastNacruleId *string,
    nacruleMatched *bool,
    authType *string,
    lastVlanId *string,
    lastNasVendor *string,
    idpId *string,
    lastSsid *string,
    lastUsername *string,
    timestamp *float64,
    lastAp *string,
    mac *string,
    lastStatus *string,
    mType *string,
    mdmComplianceStatus *string,
    mdmProvider *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/nac_clients/count", siteId),
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
    if distinct != nil {
        req.QueryParam("distinct", *distinct)
    }
    if lastNacruleId != nil {
        req.QueryParam("last_nacrule_id", *lastNacruleId)
    }
    if nacruleMatched != nil {
        req.QueryParam("nacrule_matched", *nacruleMatched)
    }
    if authType != nil {
        req.QueryParam("auth_type", *authType)
    }
    if lastVlanId != nil {
        req.QueryParam("last_vlan_id", *lastVlanId)
    }
    if lastNasVendor != nil {
        req.QueryParam("last_nas_vendor", *lastNasVendor)
    }
    if idpId != nil {
        req.QueryParam("idp_id", *idpId)
    }
    if lastSsid != nil {
        req.QueryParam("last_ssid", *lastSsid)
    }
    if lastUsername != nil {
        req.QueryParam("last_username", *lastUsername)
    }
    if timestamp != nil {
        req.QueryParam("timestamp", *timestamp)
    }
    if lastAp != nil {
        req.QueryParam("last_ap", *lastAp)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if lastStatus != nil {
        req.QueryParam("last_status", *lastStatus)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if mdmComplianceStatus != nil {
        req.QueryParam("mdm_compliance_status", *mdmComplianceStatus)
    }
    if mdmProvider != nil {
        req.QueryParam("mdm_provider", *mdmProvider)
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
    
    var result models.RepsonseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountSiteNacClientEvents takes context, siteId, distinct, mType, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of NAC Client-Events
func (s *SitesClientsNAC) CountSiteNacClientEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteNacClientEventsCountDistinctEnum,
    mType *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.RepsonseCount],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/nac_clients/events/count", siteId),
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

// SearcsearcSiteNacClientEventsacClientEvents takes context, siteId, mType, nacruleId, nacruleMatched, dryrunNacruleId, dryrunNacruleMatched, authType, vlan, vlanSource, nasVendor, bssid, idpId, idpRole, idpUsername, respAttrs, ssid, username, usermacLabels, ap, randomMac, mac, lookupTimeTaken, timestamp, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseEventsNacClientSearch data and
// an error if there was an issue with the request or response.
// Search NAC Client Events
func (s *SitesClientsNAC) SearcsearcSiteNacClientEventsacClientEvents(
    ctx context.Context,
    siteId uuid.UUID,
    mType *string,
    nacruleId *uuid.UUID,
    nacruleMatched *bool,
    dryrunNacruleId *string,
    dryrunNacruleMatched *bool,
    authType *string,
    vlan *int,
    vlanSource *string,
    nasVendor *string,
    bssid *string,
    idpId *uuid.UUID,
    idpRole *string,
    idpUsername *string,
    respAttrs []string,
    ssid *string,
    username *string,
    usermacLabels []string,
    ap *string,
    randomMac *bool,
    mac *string,
    lookupTimeTaken *float64,
    timestamp *float64,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseEventsNacClientSearch],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/nac_clients/events/search", siteId),
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
    if nacruleId != nil {
        req.QueryParam("nacrule_id", *nacruleId)
    }
    if nacruleMatched != nil {
        req.QueryParam("nacrule_matched", *nacruleMatched)
    }
    if dryrunNacruleId != nil {
        req.QueryParam("dryrun_nacrule_id", *dryrunNacruleId)
    }
    if dryrunNacruleMatched != nil {
        req.QueryParam("dryrun_nacrule_matched", *dryrunNacruleMatched)
    }
    if authType != nil {
        req.QueryParam("auth_type", *authType)
    }
    if vlan != nil {
        req.QueryParam("vlan", *vlan)
    }
    if vlanSource != nil {
        req.QueryParam("vlan_source", *vlanSource)
    }
    if nasVendor != nil {
        req.QueryParam("nas_vendor", *nasVendor)
    }
    if bssid != nil {
        req.QueryParam("bssid", *bssid)
    }
    if idpId != nil {
        req.QueryParam("idp_id", *idpId)
    }
    if idpRole != nil {
        req.QueryParam("idp_role", *idpRole)
    }
    if idpUsername != nil {
        req.QueryParam("idp_username", *idpUsername)
    }
    if respAttrs != nil {
        req.QueryParam("resp_attrs", respAttrs)
    }
    if ssid != nil {
        req.QueryParam("ssid", *ssid)
    }
    if username != nil {
        req.QueryParam("username", *username)
    }
    if usermacLabels != nil {
        req.QueryParam("usermac_labels", usermacLabels)
    }
    if ap != nil {
        req.QueryParam("ap", *ap)
    }
    if randomMac != nil {
        req.QueryParam("random_mac", *randomMac)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if lookupTimeTaken != nil {
        req.QueryParam("lookup_time_taken", *lookupTimeTaken)
    }
    if timestamp != nil {
        req.QueryParam("timestamp", *timestamp)
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
    
    var result models.ResponseEventsNacClientSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseEventsNacClientSearch](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchSiteNacClients takes context, siteId, nacruleId, nacruleMatched, authType, vlan, nasVendor, idpId, ssid, username, timestamp, ap, mac, mxedgeId, nacruleName, status, mType, mdmCompliance, mdmProvider, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseClientNacSearch data and
// an error if there was an issue with the request or response.
// Search Site NAC Clients
func (s *SitesClientsNAC) SearchSiteNacClients(
    ctx context.Context,
    siteId uuid.UUID,
    nacruleId *string,
    nacruleMatched *bool,
    authType *string,
    vlan *string,
    nasVendor *string,
    idpId *string,
    ssid *string,
    username *string,
    timestamp *float64,
    ap *string,
    mac *string,
    mxedgeId *string,
    nacruleName *string,
    status *string,
    mType *string,
    mdmCompliance *string,
    mdmProvider *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseClientNacSearch],
    error) {
    req := s.prepareRequest(
      ctx,
      "GET",
      fmt.Sprintf("/api/v1/sites/%v/nac_clients/search", siteId),
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
    if nacruleId != nil {
        req.QueryParam("nacrule_id", *nacruleId)
    }
    if nacruleMatched != nil {
        req.QueryParam("nacrule_matched", *nacruleMatched)
    }
    if authType != nil {
        req.QueryParam("auth_type", *authType)
    }
    if vlan != nil {
        req.QueryParam("vlan", *vlan)
    }
    if nasVendor != nil {
        req.QueryParam("nas_vendor", *nasVendor)
    }
    if idpId != nil {
        req.QueryParam("idp_id", *idpId)
    }
    if ssid != nil {
        req.QueryParam("ssid", *ssid)
    }
    if username != nil {
        req.QueryParam("username", *username)
    }
    if timestamp != nil {
        req.QueryParam("timestamp", *timestamp)
    }
    if ap != nil {
        req.QueryParam("ap", *ap)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if mxedgeId != nil {
        req.QueryParam("mxedge_id", *mxedgeId)
    }
    if nacruleName != nil {
        req.QueryParam("nacrule_name", *nacruleName)
    }
    if status != nil {
        req.QueryParam("status", *status)
    }
    if mType != nil {
        req.QueryParam("type", *mType)
    }
    if mdmCompliance != nil {
        req.QueryParam("mdm_compliance", *mdmCompliance)
    }
    if mdmProvider != nil {
        req.QueryParam("mdm_provider", *mdmProvider)
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
    
    var result models.ResponseClientNacSearch
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseClientNacSearch](decoder)
    return models.NewApiResponse(result, resp), err
}
