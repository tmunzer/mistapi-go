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

// SitesDevicesOthers represents a controller struct.
type SitesDevicesOthers struct {
	baseController
}

// NewSitesDevicesOthers creates a new instance of SitesDevicesOthers.
// It takes a baseController as a parameter and returns a pointer to the SitesDevicesOthers.
func NewSitesDevicesOthers(baseController baseController) *SitesDevicesOthers {
	sitesDevicesOthers := SitesDevicesOthers{baseController: baseController}
	return &sitesDevicesOthers
}

// ListSiteOtherDevices takes context, siteId, vendor, mac, serial, model, name, limit, page as parameters and
// returns an models.ApiResponse with []models.DeviceOther data and
// an error if there was an issue with the request or response.
// List third-party devices in a site, such as devices discovered or tracked outside the managed Mist device inventory. Use [List Org Other Devices]($e/Orgs%20Devices%20-%20Others/listOrgOtherDevices) to retrieve third-party devices across the organization.
func (s *SitesDevicesOthers) ListSiteOtherDevices(
	ctx context.Context,
	siteId uuid.UUID,
	vendor *string,
	mac *string,
	serial *string,
	model *string,
	name *string,
	limit *int,
	page *int) (
	models.ApiResponse[[]models.DeviceOther],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/otherdevices")
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

// CountSiteOtherDeviceEvents takes context, siteId, distinct, mType, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseCount data and
// an error if there was an issue with the request or response.
// Count third-party device events for a site, optionally grouped by the `distinct` field and filtered by event type and time range. Use [Count Org Other Device Events]($e/Orgs%20Devices%20-%20Others/countOrgOtherDeviceEvents) to count third-party device events across the organization.
func (s *SitesDevicesOthers) CountSiteOtherDeviceEvents(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteOtherDeviceEventsCountDistinctEnum,
	mType *string,
	start *string,
	end *string,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseCount],
	error) {
	req := s.prepareRequest(ctx, "GET", "/api/v1/sites/%v/otherdevices/events/count")
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

// SearchSiteOtherDeviceEvents takes context, siteId, mac, deviceMac, vendor, mType, limit, start, end, duration, sort, searchAfter as parameters and
// returns an models.ApiResponse with models.ResponseEventsOtherDevicesSearch data and
// an error if there was an issue with the request or response.
// Search third-party device events for a site with filters for device identifiers, model, vendor, event type, and time range. Use [Search Org Other Device Events]($e/Orgs%20Devices%20-%20Others/searchOrgOtherDeviceEvents) to search third-party device events across the organization.
func (s *SitesDevicesOthers) SearchSiteOtherDeviceEvents(
	ctx context.Context,
	siteId uuid.UUID,
	mac *string,
	deviceMac *string,
	vendor *string,
	mType *string,
	limit *int,
	start *string,
	end *string,
	duration *string,
	sort *string,
	searchAfter *string) (
	models.ApiResponse[models.ResponseEventsOtherDevicesSearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		"/api/v1/sites/%v/otherdevices/events/search",
	)
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
	if mac != nil {
		req.QueryParam("mac", *mac)
	}
	if deviceMac != nil {
		req.QueryParam("device_mac", *deviceMac)
	}
	if vendor != nil {
		req.QueryParam("vendor", *vendor)
	}
	if mType != nil {
		req.QueryParam("type", *mType)
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

	var result models.ResponseEventsOtherDevicesSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseEventsOtherDevicesSearch](decoder)
	return models.NewApiResponse(result, resp), err
}
