// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package mistapi

import (
    "context"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "github.com/google/uuid"
    "github.com/tmunzer/mistapi-go/mistapi/errors"
    "github.com/tmunzer/mistapi-go/mistapi/models"
)

// OrgsClientsNAC represents a controller struct.
type OrgsClientsNAC struct {
    baseController
}

// NewOrgsClientsNAC creates a new instance of OrgsClientsNAC.
// It takes a baseController as a parameter and returns a pointer to the OrgsClientsNAC.
func NewOrgsClientsNAC(baseController baseController) *OrgsClientsNAC {
    orgsClientsNAC := OrgsClientsNAC{baseController: baseController}
    return &orgsClientsNAC
}

// CountOrgNacClients takes context, orgId, distinct, lastNacruleId, nacruleMatched, authType, lastVlanId, lastNasVendor, idpId, lastSsid, lastUsername, timestamp, siteId, lastAp, mac, lastStatus, mType, mdmComplianceStatus, mdmProvider, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of NAC Clients
func (o *OrgsClientsNAC) CountOrgNacClients(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgNacClientsCountDistinctEnum,
    lastNacruleId *string,
    nacruleMatched *bool,
    authType *string,
    lastVlanId *string,
    lastNasVendor *string,
    idpId *string,
    lastSsid *string,
    lastUsername *string,
    timestamp *float64,
    siteId *string,
    lastAp *string,
    mac *string,
    lastStatus *string,
    mType *string,
    mdmComplianceStatus *string,
    mdmProvider *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/nac_clients/count")
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
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
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
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// CountOrgNacClientEvents takes context, orgId, distinct, mType, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count by Distinct Attributes of NAC Client-Events
func (o *OrgsClientsNAC) CountOrgNacClientEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgNacClientEventsCountDistinctEnum,
    mType *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/nac_clients/events/count")
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
    
    var result models.ResponseCount
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    
    result, err = utilities.DecodeResults[models.ResponseCount](decoder)
    return models.NewApiResponse(result, resp), err
}

// SearchOrgNacClientEvents takes context, orgId, mType, nacruleId, nacruleMatched, dryrunNacruleId, dryrunNacruleMatched, authType, vlan, nasVendor, bssid, idpId, idpRole, idpUsername, respAttrs, ssid, username, siteId, ap, randomMac, mac, timestamp, usermacLabel, text, nasIp, sort, ingressVlan, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseEventsNacClientSearch data and
// an error if there was an issue with the request or response.
// Search NAC Client Events
func (o *OrgsClientsNAC) SearchOrgNacClientEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mType *string,
    nacruleId *uuid.UUID,
    nacruleMatched *bool,
    dryrunNacruleId *string,
    dryrunNacruleMatched *bool,
    authType *string,
    vlan *int,
    nasVendor *string,
    bssid *string,
    idpId *uuid.UUID,
    idpRole *string,
    idpUsername *string,
    respAttrs []string,
    ssid *string,
    username *string,
    siteId *string,
    ap *string,
    randomMac *bool,
    mac *string,
    timestamp *float64,
    usermacLabel *string,
    text *string,
    nasIp *string,
    sort *string,
    ingressVlan *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseEventsNacClientSearch],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/nac_clients/events/search")
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
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
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
    if timestamp != nil {
        req.QueryParam("timestamp", *timestamp)
    }
    if usermacLabel != nil {
        req.QueryParam("usermac_label", *usermacLabel)
    }
    if text != nil {
        req.QueryParam("text", *text)
    }
    if nasIp != nil {
        req.QueryParam("nas_ip", *nasIp)
    }
    if sort != nil {
        req.QueryParam("sort", *sort)
    }
    if ingressVlan != nil {
        req.QueryParam("ingress_vlan", *ingressVlan)
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

// SearchOrgNacClients takes context, orgId, nacruleId, nacruleMatched, authType, vlan, nasVendor, nasIp, idpId, ssid, username, timestamp, siteId, ap, mac, mdmManaged, status, mType, mdmCompliance, family, model, os, hostname, mfg, mdmProvider, sort, usermacLabel, ingressVlan, start, end, duration, limit, page as parameters and
// returns an models.ApiResponse with models.ResponseClientNacSearch data and
// an error if there was an issue with the request or response.
// Search Org NAC Clients
func (o *OrgsClientsNAC) SearchOrgNacClients(
    ctx context.Context,
    orgId uuid.UUID,
    nacruleId *string,
    nacruleMatched *bool,
    authType *string,
    vlan *string,
    nasVendor *string,
    nasIp *string,
    idpId *string,
    ssid *string,
    username *string,
    timestamp *float64,
    siteId *string,
    ap *string,
    mac *string,
    mdmManaged *bool,
    status *models.NacClientLastStatusEnum,
    mType *string,
    mdmCompliance *string,
    family *string,
    model *string,
    os *string,
    hostname *string,
    mfg *string,
    mdmProvider *string,
    sort *string,
    usermacLabel []string,
    ingressVlan *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseClientNacSearch],
    error) {
    req := o.prepareRequest(ctx, "GET", "/api/v1/orgs/%v/nac_clients/search")
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
    if nasIp != nil {
        req.QueryParam("nas_ip", *nasIp)
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
    if siteId != nil {
        req.QueryParam("site_id", *siteId)
    }
    if ap != nil {
        req.QueryParam("ap", *ap)
    }
    if mac != nil {
        req.QueryParam("mac", *mac)
    }
    if mdmManaged != nil {
        req.QueryParam("mdm_managed", *mdmManaged)
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
    if family != nil {
        req.QueryParam("family", *family)
    }
    if model != nil {
        req.QueryParam("model", *model)
    }
    if os != nil {
        req.QueryParam("os", *os)
    }
    if hostname != nil {
        req.QueryParam("hostname", *hostname)
    }
    if mfg != nil {
        req.QueryParam("mfg", *mfg)
    }
    if mdmProvider != nil {
        req.QueryParam("mdm_provider", *mdmProvider)
    }
    if sort != nil {
        req.QueryParam("sort", *sort)
    }
    if usermacLabel != nil {
        req.QueryParam("usermac_label", usermacLabel)
    }
    if ingressVlan != nil {
        req.QueryParam("ingress_vlan", *ingressVlan)
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
