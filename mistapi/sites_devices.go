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
	"net/http"
)

// SitesDevices represents a controller struct.
type SitesDevices struct {
	baseController
}

// NewSitesDevices creates a new instance of SitesDevices.
// It takes a baseController as a parameter and returns a pointer to the SitesDevices.
func NewSitesDevices(baseController baseController) *SitesDevices {
	sitesDevices := SitesDevices{baseController: baseController}
	return &sitesDevices
}

// ListSiteDevices takes context, siteId, mType, name, limit, page as parameters and
// returns an models.ApiResponse with []models.ConfigDevice data and
// an error if there was an issue with the request or response.
// Get list of devices on the site.
func (s *SitesDevices) ListSiteDevices(
	ctx context.Context,
	siteId uuid.UUID,
	mType *models.DeviceTypeWithAllEnum,
	name *string,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.ConfigDevice],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices")
	req.AppendTemplateParams(siteId)
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
	if name != nil {
		req.QueryParam("name", *name)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}

	var result []models.ConfigDevice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConfigDevice](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteDeviceConfigHistory takes context, siteId, distinct, mac, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Counts the number of entries in device config history for distinct field with given filters
func (s *SitesDevices) CountSiteDeviceConfigHistory(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *string,
	mac *string,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		"/api/v1/sites/%v/devices/config_history/count",
	)
	req.AppendTemplateParams(siteId)
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
	if mac != nil {
		req.QueryParam("mac", *mac)
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

// SearchSiteDeviceConfigHistory takes context, siteId, mType, mac, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseConfigHistorySearch data and
// an error if there was an issue with the request or response.
// Search for entries in device config history
func (s *SitesDevices) SearchSiteDeviceConfigHistory(
	ctx context.Context,
	siteId uuid.UUID,
	mType *models.DeviceTypeDefaultApEnum,
	mac *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseConfigHistorySearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		"/api/v1/sites/%v/devices/config_history/search",
	)
	req.AppendTemplateParams(siteId)
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
	if mac != nil {
		req.QueryParam("mac", *mac)
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

	var result models.ResponseConfigHistorySearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseConfigHistorySearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteDevices takes context, siteId, distinct, hostname, model, mac, version, mxtunnelStatus, mxedgeId, lldpSystemName, lldpSystemDesc, lldpPortId, lldpMgmtAddr, mapId, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Counts the number of entries in ap events history for distinct field with given filters
func (s *SitesDevices) CountSiteDevices(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteDevicesCountDistinctEnum,
	hostname *string,
	model *string,
	mac *string,
	version *string,
	mxtunnelStatus *string,
	mxedgeId *string,
	lldpSystemName *string,
	lldpSystemDesc *string,
	lldpPortId *string,
	lldpMgmtAddr *string,
	mapId *string,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/count")
	req.AppendTemplateParams(siteId)
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
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if version != nil {
		req.QueryParam("version", *version)
	}
	if mxtunnelStatus != nil {
		req.QueryParam("mxtunnel_status", *mxtunnelStatus)
	}
	if mxedgeId != nil {
		req.QueryParam("mxedge_id", *mxedgeId)
	}
	if lldpSystemName != nil {
		req.QueryParam("lldp_system_name", *lldpSystemName)
	}
	if lldpSystemDesc != nil {
		req.QueryParam("lldp_system_desc", *lldpSystemDesc)
	}
	if lldpPortId != nil {
		req.QueryParam("lldp_port_id", *lldpPortId)
	}
	if lldpMgmtAddr != nil {
		req.QueryParam("lldp_mgmt_addr", *lldpMgmtAddr)
	}
	if mapId != nil {
		req.QueryParam("map_id", *mapId)
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

// CountSiteDeviceEvents takes context, siteId, distinct, model, mType, typeCode, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Counts the number of entries in ap events history for distinct field with given filters
func (s *SitesDevices) CountSiteDeviceEvents(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteDeviceEventsCountDistinctEnum,
	model *string,
	mType *string,
	typeCode *string,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/events/count")
	req.AppendTemplateParams(siteId)
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
	if model != nil {
		req.QueryParam("model", *model)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if typeCode != nil {
		req.QueryParam("type_code", *typeCode)
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

// SearchSiteDeviceEvents takes context, siteId, mac, model, text, timestamp, mType, lastBy, includes, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseEventsDevices data and
// an error if there was an issue with the request or response.
// Search Devices Events
func (s *SitesDevices) SearchSiteDeviceEvents(
	ctx context.Context,
	siteId uuid.UUID,
	mac *string,
	model *string,
	text *string,
	timestamp *string,
	mType *string,
	lastBy *string,
	includes *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseEventsDevices],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/events/search")
	req.AppendTemplateParams(siteId)
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
	if model != nil {
		req.QueryParam("model", *model)
	}
	if text != nil {
		req.QueryParam("text", *text)
	}
	if timestamp != nil {
		req.QueryParam("timestamp", *timestamp)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if lastBy != nil {
		req.QueryParam("last_by", *lastBy)
	}
	if includes != nil {
		req.QueryParam("includes", *includes)
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

	var result models.ResponseEventsDevices
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseEventsDevices](decoder)
	return models.NewApiResponse(result, resp), err
}

// ExportSiteDevices takes context, siteId as parameters and
// returns an models.ApiResponse with string data and
// an error if there was an issue with the request or response.
// To download the exported device information
func (s *SitesDevices) ExportSiteDevices(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[string],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/export")
	req.AppendTemplateParams(siteId)
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

	str, resp, err := req.CallAsText()
	var result string = str

	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	return models.NewApiResponse(result, resp), err
}

// SetSiteDevicesGbpTag takes context, siteId, body as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Set GBP Tag for multiple devices
func (s *SitesDevices) SetSiteDevicesGbpTag(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.DevicesGbpTag) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/gbp_tag")
	req.AppendTemplateParams(siteId)
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

// ImportSiteDevices takes context, siteId, file as parameters and
// returns an models.ApiResponse with []models.ConfigDevice data and
// an error if there was an issue with the request or response.
// Import Information for Multiple Devices
// CSV format:
// ```csv
// mac,name,map_id,x,y,height,orientation,labels,band_24.power,band_24.bandwidth,band_24.channel,band_24.disabled,band_5.power,band_5.bandwidth,band_5.channel,band_5.disabled,band_6.power,band_6.bandwidth,band_6.channel,band_6.disabled
// 5c5b53010101,"AP 1",845a23bf-bed9-e43c-4c86-6fa474be7ae5,30,10,2.3,45,"guest, campus, vip",1,20,0,false,0,40,0,false,17,80,0,false
// ```
func (s *SitesDevices) ImportSiteDevices(
	ctx context.Context,
	siteId uuid.UUID,
	file string) (
	models.ApiResponse[[]models.ConfigDevice],
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/import")
	req.AppendTemplateParams(siteId)
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
	req.FormParam("file", file)

	var result []models.ConfigDevice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ConfigDevice](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteDeviceLastConfig takes context, siteId, distinct, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Counts the number of entries in device config history for distinct field with given filters
func (s *SitesDevices) CountSiteDeviceLastConfig(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteDeviceLastConfigCountDistinctEnum,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/last_config/count")
	req.AppendTemplateParams(siteId)
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

	var result models.ResponseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteDeviceLastConfigs takes context, siteId, certExpiryDuration, deviceType, mac, version, name, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseConfigHistorySearch data and
// an error if there was an issue with the request or response.
// Search Device Last Configs
func (s *SitesDevices) SearchSiteDeviceLastConfigs(
	ctx context.Context,
	siteId uuid.UUID,
	certExpiryDuration *string,
	deviceType *models.LastConfigDeviceTypeEnum,
	mac *string,
	version *string,
	name *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseConfigHistorySearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		"/api/v1/sites/%v/devices/last_config/search",
	)
	req.AppendTemplateParams(siteId)
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
	if certExpiryDuration != nil {
		req.QueryParam("cert_expiry_duration", *certExpiryDuration)
	}
	if deviceType != nil {
		req.QueryParam("device_type", *deviceType)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if version != nil {
		req.QueryParam("version", *version)
	}
	if name != nil {
		req.QueryParam("name", *name)
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

	var result models.ResponseConfigHistorySearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseConfigHistorySearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteDevices takes context, siteId, hostname, mType, model, mac, extIp, version, powerConstrained, ip, mxtunnelStatus, mxedgeId, mxedgeIds, lastHostname, lastConfigStatus, radiusStats, cpu, node0Mac, clustered, t128agentVersion, node1Mac, node, evpntopoId, lldpSystemName, lldpSystemDesc, lldpPortId, lldpMgmtAddr, band24Channel, band5Channel, band6Channel, band24Bandwidth, band5Bandwidth, band6Bandwidth, eth0PortSpeed, stats, limit, start, end, duration, sort, descSort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseDeviceSearch data and
// an error if there was an issue with the request or response.
// Search Device
func (s *SitesDevices) SearchSiteDevices(
	ctx context.Context,
	siteId uuid.UUID,
	hostname *string,
	mType *models.DeviceTypeDefaultApEnum,
	model *string,
	mac *string,
	extIp *string,
	version *string,
	powerConstrained *bool,
	ip *string,
	mxtunnelStatus *models.SearchSiteDevicesMxtunnelStatusEnum,
	mxedgeId *uuid.UUID,
	mxedgeIds []uuid.UUID,
	lastHostname *string,
	lastConfigStatus *string,
	radiusStats *string,
	cpu *string,
	node0Mac *string,
	clustered *bool,
	t128agentVersion *string,
	node1Mac *string,
	node *models.HaClusterNodeEnum,
	evpntopoId *string,
	lldpSystemName *string,
	lldpSystemDesc *string,
	lldpPortId *string,
	lldpMgmtAddr *string,
	band24Channel *int,
	band5Channel *int,
	band6Channel *int,
	band24Bandwidth *int,
	band5Bandwidth *int,
	band6Bandwidth *int,
	eth0PortSpeed *int,
	stats *bool,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *models.SearchSiteDevicesSortEnum,
	descSort *models.SearchSiteDevicesDescSortEnum,
	searchAfter *string) (
	models.ApiResponse[models.ResponseDeviceSearch],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/search")
	req.AppendTemplateParams(siteId)
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
	if hostname != nil {
		req.QueryParam("hostname", *hostname)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if extIp != nil {
		req.QueryParam("ext_ip", *extIp)
	}
	if version != nil {
		req.QueryParam("version", *version)
	}
	if powerConstrained != nil {
		req.QueryParam("power_constrained", *powerConstrained)
	}
	if ip != nil {
		req.QueryParam("ip", *ip)
	}
	if mxtunnelStatus != nil {
		req.QueryParam("mxtunnel_status", *mxtunnelStatus)
	}
	if mxedgeId != nil {
		req.QueryParam("mxedge_id", *mxedgeId)
	}
	if mxedgeIds != nil {
		req.QueryParam("mxedge_ids", mxedgeIds)
	}
	if lastHostname != nil {
		req.QueryParam("last_hostname", *lastHostname)
	}
	if lastConfigStatus != nil {
		req.QueryParam("last_config_status", *lastConfigStatus)
	}
	if radiusStats != nil {
		req.QueryParam("radius_stats", *radiusStats)
	}
	if cpu != nil {
		req.QueryParam("cpu", *cpu)
	}
	if node0Mac != nil {
		req.QueryParam("node0_mac", *node0Mac)
	}
	if clustered != nil {
		req.QueryParam("clustered", *clustered)
	}
	if t128agentVersion != nil {
		req.QueryParam("t128agent_version", *t128agentVersion)
	}
	if node1Mac != nil {
		req.QueryParam("node1_mac", *node1Mac)
	}
	if node != nil {
		req.QueryParam("node", *node)
	}
	if evpntopoId != nil {
		req.QueryParam("evpntopo_id", *evpntopoId)
	}
	if lldpSystemName != nil {
		req.QueryParam("lldp_system_name", *lldpSystemName)
	}
	if lldpSystemDesc != nil {
		req.QueryParam("lldp_system_desc", *lldpSystemDesc)
	}
	if lldpPortId != nil {
		req.QueryParam("lldp_port_id", *lldpPortId)
	}
	if lldpMgmtAddr != nil {
		req.QueryParam("lldp_mgmt_addr", *lldpMgmtAddr)
	}
	if band24Channel != nil {
		req.QueryParam("band_24_channel", *band24Channel)
	}
	if band5Channel != nil {
		req.QueryParam("band_5_channel", *band5Channel)
	}
	if band6Channel != nil {
		req.QueryParam("band_6_channel", *band6Channel)
	}
	if band24Bandwidth != nil {
		req.QueryParam("band_24_bandwidth", *band24Bandwidth)
	}
	if band5Bandwidth != nil {
		req.QueryParam("band_5_bandwidth", *band5Bandwidth)
	}
	if band6Bandwidth != nil {
		req.QueryParam("band_6_bandwidth", *band6Bandwidth)
	}
	if eth0PortSpeed != nil {
		req.QueryParam("eth0_port_speed", *eth0PortSpeed)
	}
	if stats != nil {
		req.QueryParam("stats", *stats)
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
	if descSort != nil {
		req.QueryParam("desc_sort", *descSort)
	}
	if searchAfter != nil {
		req.QueryParam("search_after", *searchAfter)
	}

	var result models.ResponseDeviceSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseDeviceSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteDevice takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with models.MistDevice data and
// an error if there was an issue with the request or response.
// Get Device Configuration
func (s *SitesDevices) GetSiteDevice(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID) (
	models.ApiResponse[models.MistDevice],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/devices/%v")
	req.AppendTemplateParams(siteId, deviceId)
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

	var result models.MistDevice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MistDevice](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSiteDevice takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.MistDevice data and
// an error if there was an issue with the request or response.
// Update Device Configuration
func (s *SitesDevices) UpdateSiteDevice(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	body *models.MistDevice) (
	models.ApiResponse[models.MistDevice],
	error) {
	req := s.prepareRequest(ctx, "PUT", "/api/v1/sites/%v/devices/%v")
	req.AppendTemplateParams(siteId, deviceId)
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

	var result models.MistDevice
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.MistDevice](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteSiteDeviceImage takes context, siteId, deviceId, imageNumber as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Delete image from a device
func (s *SitesDevices) DeleteSiteDeviceImage(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	imageNumber int) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "DELETE", "/api/v1/sites/%v/devices/%v/image%v")
	req.AppendTemplateParams(siteId, deviceId, imageNumber)
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

// AddSiteDeviceImage takes context, siteId, deviceId, imageNumber, file, json as parameters and
// returns an *Response and
// an error if there was an issue with the request or response.
// Attach up to 3 images to a device
func (s *SitesDevices) AddSiteDeviceImage(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	imageNumber int,
	file string,
	json *string) (
	*http.Response,
	error) {
	req := s.prepareRequest(ctx, "POST", "/api/v1/sites/%v/devices/%v/image%v")
	req.AppendTemplateParams(siteId, deviceId, imageNumber)
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
	req.FormParam("file", file)
	if json != nil {
		req.FormParam("json", *json)
	}

	httpCtx, err := req.Call()
	if err != nil {
		return httpCtx.Response, err
	}
	return httpCtx.Response, err
}
