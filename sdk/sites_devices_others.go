package mistapi

import (
	"context"
	"fmt"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapi/sdk/errors"
	"github.com/tmunzer/mistapi/sdk/models"
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

// ListSiteOtherDevices takes context, siteId, vendor, mac, serial, model, name, page, limit as parameters and
// returns an models.ApiResponse with []models.DeviceOther data and
// an error if there was an issue with the request or response.
// Get List of Site other devices (3rd party devices)
func (s *SitesDevicesOthers) ListSiteOtherDevices(
	ctx context.Context,
	siteId uuid.UUID,
	vendor *string,
	mac *string,
	serial *string,
	model *string,
	name *string,
	page *int,
	limit *int) (
	models.ApiResponse[[]models.DeviceOther],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/otherdevices", siteId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
		),
	)
	req.AppendErrors(map[string]https.ErrorBuilder[error]{
		"400": {Message: "Bad Syntax", Unmarshaller: errors.NewResponseHttp400},
		"401": {Message: "Unauthorized", Unmarshaller: errors.NewResponseHttp400},
		"403": {Message: "Permission Denied", Unmarshaller: errors.NewResponseHttp400},
		"404": {Message: "Not found. The API endpoint doesn’t exist or resource doesn’t exist", Unmarshaller: errors.NewResponseHttp404},
		"429": {Message: "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold", Unmarshaller: errors.NewResponseHttp400},
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
	if page != nil {
		req.QueryParam("page", *page)
	}
	if limit != nil {
		req.QueryParam("limit", *limit)
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
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Site OtherDevices Events
func (s *SitesDevicesOthers) CountSiteOtherDeviceEvents(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteOtherDeviceEventsCountDistinctEnum,
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
		fmt.Sprintf("/api/v1/sites/%v/otherdevices/events/count", siteId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
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

// SearchSiteOtherDeviceEvents takes context, siteId, mac, deviceMac, vendor, mType, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseEventsOtherDevicesSearch data and
// an error if there was an issue with the request or response.
// Search Site OtherDevices Events
func (s *SitesDevicesOthers) SearchSiteOtherDeviceEvents(
	ctx context.Context,
	siteId uuid.UUID,
	mac *string,
	deviceMac *string,
	vendor *string,
	mType *string,
	start *int,
	end *int,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseEventsOtherDevicesSearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/otherdevices/events/search", siteId),
	)
	req.Authenticate(
		NewOrAuth(
			NewAuth("apiToken"),
			NewAuth("basicAuth"),
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
	if deviceMac != nil {
		req.QueryParam("device_mac", *deviceMac)
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