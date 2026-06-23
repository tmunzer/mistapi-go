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

// CountSiteNacClients takes context, siteId, distinct, lastNacruleId, nacruleMatched, authType, lastVlanId, lastNasVendor, idpId, lastSsid, lastUsername, lastAp, mac, lastStatus, mType, mdmComplianceStatus, mdmProvider, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count NAC clients for a site, optionally grouped by the `distinct` field and filtered by authentication, identity, endpoint, network, and time attributes. Use [Count Org NAC Clients]($e/Orgs%20Clients%20-%20NAC/countOrgNacClients) to count NAC clients across the organization.
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
	lastAp *string,
	mac *string,
	lastStatus *string,
	mType *string,
	mdmComplianceStatus *string,
	mdmProvider *string,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/nac_clients/count")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
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

// CountSiteNacClientEvents takes context, siteId, distinct, mType, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count NAC client events for a site, optionally grouped by the `distinct` field and filtered by event type and time range. Use [Count Org NAC Client Events]($e/Orgs%20Clients%20-%20NAC/countOrgNacClientEvents) to count NAC client events across the organization.
func (s *SitesClientsNAC) CountSiteNacClientEvents(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteNacClientEventsCountDistinctEnum,
	mType *string,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/nac_clients/events/count")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
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

// SearchSiteNacClientEvents takes context, siteId, mType, nacruleId, nacruleMatched, dryrunNacruleId, dryrunNacruleMatched, authType, vlan, nasVendor, bssid, idpId, idpRole, idpUsername, respAttrs, ssid, username, ap, randomMac, mac, usermacLabel, text, nasIp, ingressVlan, start, end, duration, limit, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseEventsNacClientSearch data and
// an error if there was an issue with the request or response.
// Search NAC client events for a site with filters for authentication, NAC rule, identity provider, RADIUS, network, endpoint, and time attributes. Use [Search Org NAC Client Events]($e/Orgs%20Clients%20-%20NAC/searchOrgNacClientEvents) to search NAC client events across the organization.
func (s *SitesClientsNAC) SearchSiteNacClientEvents(
	ctx context.Context,
	siteId uuid.UUID,
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
	ap *string,
	randomMac *bool,
	mac *string,
	usermacLabel *string,
	text *string,
	nasIp *string,
	ingressVlan *string,
	start *string,
	end *string,
	duration *string,
	limit *int,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseEventsNacClientSearch],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/nac_clients/events/search")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
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
	if ap != nil {
		req.QueryParam("ap", *ap)
	}
	if randomMac != nil {
		req.QueryParam("random_mac", *randomMac)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
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
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseEventsNacClientSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseEventsNacClientSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteNacClients takes context, siteId, ap, authType, certExpiryDuration, edrManaged, edrProvider, edrStatus, family, hostname, idpId, mac, mdmCompliance, mdmProvider, mdmManaged, mfg, model, nacruleName, nacruleId, nacruleMatched, nasVendor, nasIp, ingressVlan, os, ssid, status, text, mType, usermacLabel, username, vlan, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseClientNacSearch data and
// an error if there was an issue with the request or response.
// Search NAC clients for a site with filters for authentication, endpoint posture, identity, network, NAC rule, and time attributes. Use [Search Org NAC Clients]($e/Orgs%20Clients%20-%20NAC/searchOrgNacClients) to search NAC clients across the organization.
func (s *SitesClientsNAC) SearchSiteNacClients(
	ctx context.Context,
	siteId uuid.UUID,
	ap *string,
	authType *string,
	certExpiryDuration *string,
	edrManaged *bool,
	edrProvider *models.EdrProviderEnum,
	edrStatus *models.EdrStatusEnum,
	family *string,
	hostname *string,
	idpId *string,
	mac *string,
	mdmCompliance *string,
	mdmProvider *string,
	mdmManaged *bool,
	mfg *string,
	model *string,
	nacruleName *string,
	nacruleId *string,
	nacruleMatched *bool,
	nasVendor *string,
	nasIp *string,
	ingressVlan *string,
	os *string,
	ssid *string,
	status *models.NacClientLastStatusEnum,
	text *string,
	mType *string,
	usermacLabel []string,
	username *string,
	vlan *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseClientNacSearch],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/nac_clients/search")
	req.AppendTemplateParams(siteId)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	if ap != nil {
		req.QueryParam("ap", *ap)
	}
	if authType != nil {
		req.QueryParam("auth_type", *authType)
	}
	if certExpiryDuration != nil {
		req.QueryParam("cert_expiry_duration", *certExpiryDuration)
	}
	if edrManaged != nil {
		req.QueryParam("edr_managed", *edrManaged)
	}
	if edrProvider != nil {
		req.QueryParam("edr_provider", *edrProvider)
	}
	if edrStatus != nil {
		req.QueryParam("edr_status", *edrStatus)
	}
	if family != nil {
		req.QueryParam("family", *family)
	}
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if idpId != nil {
		req.QueryParam("idp_id", *idpId)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if mdmCompliance != nil {
		req.QueryParam("mdm_compliance", *mdmCompliance)
	}
	if mdmProvider != nil {
		req.QueryParam("mdm_provider", *mdmProvider)
	}
	if mdmManaged != nil {
		req.QueryParam("mdm_managed", *mdmManaged)
	}
	if mfg != nil {
		req.QueryParam("mfg", *mfg)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if nacruleName != nil {
		req.QueryParam("nacrule_name", *nacruleName)
	}
	if nacruleId != nil {
		req.QueryParam("nacrule_id", *nacruleId)
	}
	if nacruleMatched != nil {
		req.QueryParam("nacrule_matched", *nacruleMatched)
	}
	if nasVendor != nil {
		req.QueryParam("nas_vendor", *nasVendor)
	}
	if nasIp != nil {
		req.QueryParam("nas_ip", *nasIp)
	}
	if ingressVlan != nil {
		req.QueryParam("ingress_vlan", *ingressVlan)
	}
	if os != nil {
		req.QueryParam("os", *os)
	}
	if ssid != nil {
		req.QueryParam("ssid", *ssid)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if text != nil {
		req.QueryParam("text", *text)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if usermacLabel != nil {
		req.QueryParam("usermac_label", usermacLabel)
	}
	if username != nil {
		req.QueryParam("username", *username)
	}
	if vlan != nil {
		req.QueryParam("vlan", *vlan)
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
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseClientNacSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseClientNacSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// SendSiteNacClientCoA takes context, siteId, clientMac, body as parameters and
// returns an models.ApiResponse with models.NacClientCoaResponse data and
// an error if there was an issue with the request or response.
// Sends CoA (Change of Authorization) command to a NAC client.
func (s *SitesClientsNAC) SendSiteNacClientCoA(
	ctx context.Context,
	siteId uuid.UUID,
	clientMac string,
	body *models.NacClientCoa) (
	models.ApiResponse[models.NacClientCoaResponse],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/nac_clients/%v/coa")
	req.AppendTemplateParams(siteId, clientMac)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("csrfToken"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp401},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp403},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’ t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp429},
	})
	req.Header("Content-Type", "application/json")
	if body != nil {
		req.Json(body)
	}

	var result models.NacClientCoaResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.NacClientCoaResponse](decoder)
	return models.NewApiResponse(result, resp), err
}
