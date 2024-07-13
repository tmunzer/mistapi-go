package mistapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tmunzer/mistapi-go/mistapi/errors"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
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

// ListSiteDevices takes context, siteId, mType, name, page, limit as parameters and
// returns an models.ApiResponse with []models.ListSiteDevicesResponse data and
// an error if there was an issue with the request or response.
// Get list of devices on the site.
func (s *SitesDevices) ListSiteDevices(
	ctx context.Context,
	siteId uuid.UUID,
	mType *models.DeviceTypeWithAllEnum,
	name *string,
	page *int,
	limit *int) (
	models.ApiResponse[[]models.ListSiteDevicesResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/devices", siteId),
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
	if name != nil {
		req.QueryParam("name", *name)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
	}

	var result []models.ListSiteDevicesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ListSiteDevicesResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteDeviceConfigHistory takes context, siteId, distinct, mac, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Counts the number of entries in device config history for distinct field with given filters
func (s *SitesDevices) CountSiteDeviceConfigHistory(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *string,
	mac *string,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/devices/config_history/count", siteId),
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
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if page != nil {
		req.QueryParam("page", *page)
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

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteDeviceConfigHistory takes context, siteId, mType, mac, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseConfigHistorySearch data and
// an error if there was an issue with the request or response.
// Search for entries in device config history
func (s *SitesDevices) SearchSiteDeviceConfigHistory(
	ctx context.Context,
	siteId uuid.UUID,
	mType *models.DeviceTypeEnum,
	mac *string,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseConfigHistorySearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/devices/config_history/search", siteId),
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

	var result models.ResponseConfigHistorySearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseConfigHistorySearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteDevices takes context, siteId, distinct, hostname, model, mac, version, mxtunnelStatus, mxedgeId, lldpSystemName, lldpSystemDesc, lldpPortId, lldpMgmtAddr, mapId, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
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
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/devices/count", siteId),
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
	if page != nil {
		req.QueryParam("page", *page)
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

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteDeviceEvents takes context, siteId, distinct, model, mType, typeCode, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Counts the number of entries in ap events history for distinct field with given filters
func (s *SitesDevices) CountSiteDeviceEvents(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteDeviceEventsCountDistinctEnum,
	model *string,
	mType *string,
	typeCode *string,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/devices/events/count", siteId),
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
	if model != nil {
		req.QueryParam("model", *model)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if typeCode != nil {
		req.QueryParam("type_code", *typeCode)
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

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteDeviceEvents takes context, siteId, mac, model, text, timestamp, mType, lastBy, limit, start, end, duration as parameters and
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
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseEventsDevices],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/devices/events/search", siteId),
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

	var result models.ResponseEventsDevices
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseEventsDevices](decoder)
	return models.NewApiResponse(result, resp), err
}

// ExportSiteDevices takes context, siteId as parameters and
// returns an models.ApiResponse with []byte data and
// an error if there was an issue with the request or response.
// To download the exported device information
func (s *SitesDevices) ExportSiteDevices(
	ctx context.Context,
	siteId uuid.UUID) (
	models.ApiResponse[[]byte],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/devices/export", siteId),
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

	stream, resp, err := req.CallAsStream()
	if err != nil {
		return models.NewApiResponse(stream, resp), err
	}
	return models.NewApiResponse(stream, resp), err
}

// ImportSiteDevices takes context, siteId, body as parameters and
// returns an models.ApiResponse with []models.ImportSiteDevicesResponse data and
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
	body []models.DeviceAp) (
	models.ApiResponse[[]models.ImportSiteDevicesResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/devices/import", siteId),
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

	var result []models.ImportSiteDevicesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.ImportSiteDevicesResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteDeviceLastConfig takes context, siteId, distinct, page, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Counts the number of entries in device config history for distinct field with given filters
func (s *SitesDevices) CountSiteDeviceLastConfig(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteDeviceLastConfigCountDistinctEnum,
	page *int,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/devices/last_config/count", siteId),
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
	if page != nil {
		req.QueryParam("page", *page)
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

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteDeviceLastConfigs takes context, siteId, mType, mac, version, name, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseConfigHistorySearch data and
// an error if there was an issue with the request or response.
// Search Device Last Configs
func (s *SitesDevices) SearchSiteDeviceLastConfigs(
	ctx context.Context,
	siteId uuid.UUID,
	mType *models.DeviceTypeEnum,
	mac *string,
	version *string,
	name *string,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseConfigHistorySearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/devices/last_config/search", siteId),
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

	var result models.ResponseConfigHistorySearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseConfigHistorySearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// RestartSiteMultipleDevices takes context, siteId, body as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Note that only the devices that are connected will be restarted.
func (s *SitesDevices) RestartSiteMultipleDevices(
	ctx context.Context,
	siteId uuid.UUID,
	body *models.UtilsDevicesRestartMulti) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/devices/restart", siteId),
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

// SearchSiteDevices takes context, siteId, hostname, mType, model, mac, version, powerConstrained, ipAddress, mxtunnelStatus, mxedgeId, lldpSystemName, lldpSystemDesc, lldpPortId, lldpMgmtAddr, band24Channel, band5Channel, band6Channel, band24Bandwith, band5Bandwith, band6Bandwith, eth0PortSpeed, sort, descSort, stats, limit, start, end, duration as parameters and
// returns an models.ApiResponse with models.ResponseDeviceSearch data and
// an error if there was an issue with the request or response.
// Search Device
func (s *SitesDevices) SearchSiteDevices(
	ctx context.Context,
	siteId uuid.UUID,
	hostname *string,
	mType *models.DeviceTypeEnum,
	model *string,
	mac *string,
	version *string,
	powerConstrained *bool,
	ipAddress *string,
	mxtunnelStatus *models.SearchSiteDevicesMxtunnelStatusEnum,
	mxedgeId *uuid.UUID,
	lldpSystemName *string,
	lldpSystemDesc *string,
	lldpPortId *string,
	lldpMgmtAddr *string,
	band24Channel *int,
	band5Channel *int,
	band6Channel *int,
	band24Bandwith *int,
	band5Bandwith *int,
	band6Bandwith *int,
	eth0PortSpeed *int,
	sort *models.SearchSiteDevicesSortEnum,
	descSort *models.SearchSiteDevicesDescSortEnum,
	stats *bool,
	limit *int,
	start *int,
	end *int,
	duration *string) (
	models.ApiResponse[models.ResponseDeviceSearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/devices/search", siteId),
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
	if version != nil {
		req.QueryParam("version", *version)
	}
	if powerConstrained != nil {
		req.QueryParam("power_constrained", *powerConstrained)
	}
	if ipAddress != nil {
		req.QueryParam("ip_address", *ipAddress)
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
	if band24Channel != nil {
		req.QueryParam("band_24_channel", *band24Channel)
	}
	if band5Channel != nil {
		req.QueryParam("band_5_channel", *band5Channel)
	}
	if band6Channel != nil {
		req.QueryParam("band_6_channel", *band6Channel)
	}
	if band24Bandwith != nil {
		req.QueryParam("band_24_bandwith", *band24Bandwith)
	}
	if band5Bandwith != nil {
		req.QueryParam("band_5_bandwith", *band5Bandwith)
	}
	if band6Bandwith != nil {
		req.QueryParam("band_6_bandwith", *band6Bandwith)
	}
	if eth0PortSpeed != nil {
		req.QueryParam("eth0_port_speed", *eth0PortSpeed)
	}
	if sort != nil {
		req.QueryParam("sort", *sort)
	}
	if descSort != nil {
		req.QueryParam("desc_sort", *descSort)
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

	var result models.ResponseDeviceSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseDeviceSearch](decoder)
	return models.NewApiResponse(result, resp), err
}

// GetSiteDevice takes context, siteId, deviceId as parameters and
// returns an models.ApiResponse with models.GetSiteDeviceResponse data and
// an error if there was an issue with the request or response.
// Get Device Configuration
func (s *SitesDevices) GetSiteDevice(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID) (
	models.ApiResponse[models.GetSiteDeviceResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/devices/%v", siteId, deviceId),
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

	var result models.GetSiteDeviceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSiteDeviceResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// UpdateSiteDevice takes context, siteId, deviceId, body as parameters and
// returns an models.ApiResponse with models.UpdateSiteDeviceResponse data and
// an error if there was an issue with the request or response.
// Update Device Configuration
func (s *SitesDevices) UpdateSiteDevice(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	body *models.MistDevice) (
	models.ApiResponse[models.UpdateSiteDeviceResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/api/v1/sites/%v/devices/%v", siteId, deviceId),
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

	var result models.UpdateSiteDeviceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.UpdateSiteDeviceResponse](decoder)
	return models.NewApiResponse(result, resp), err
}

// DeleteSiteDeviceImage takes context, siteId, deviceId, imageNumber as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Delete image from a device
func (s *SitesDevices) DeleteSiteDeviceImage(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	imageNumber int) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/api/v1/sites/%v/devices/%v/image%v", siteId, deviceId, imageNumber),
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

// AddSiteDeviceImage takes context, siteId, deviceId, imageNumber, file, json as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
// Attach up to 3 images to a device
func (s *SitesDevices) AddSiteDeviceImage(
	ctx context.Context,
	siteId uuid.UUID,
	deviceId uuid.UUID,
	imageNumber int,
	file models.FileWrapper,
	json *string) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/api/v1/sites/%v/devices/%v/image%v", siteId, deviceId, imageNumber),
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
	req.Header("Content-Type", "multipart/form_data")
	formFields := []https.FormParam{}
	fileParam := https.FormParam{Key: "file", Value: file, Headers: http.Header{}}
	formFields = append(formFields, fileParam)
	if json != nil {
		jsonParam := https.FormParam{Key: "json", Value: *json, Headers: http.Header{}}
		formFields = append(formFields, jsonParam)
	}
	req.FormData(formFields)

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}
