package mistapigo

import (
	"context"
	"fmt"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"github.com/google/uuid"
	"github.com/tmunzer/mistapigo/sdk/errors"
	"github.com/tmunzer/mistapigo/sdk/models"
)

// SitesServices represents a controller struct.
type SitesServices struct {
	baseController
}

// NewSitesServices creates a new instance of SitesServices.
// It takes a baseController as a parameter and returns a pointer to the SitesServices.
func NewSitesServices(baseController baseController) *SitesServices {
	sitesServices := SitesServices{baseController: baseController}
	return &sitesServices
}

// ListSiteServicesDerived takes context, siteId, resolve as parameters and
// returns an models.ApiResponse with []models.Service data and
// an error if there was an issue with the request or response.
// Retrieves the list of Services available for the Site
func (s *SitesServices) ListSiteServicesDerived(
	ctx context.Context,
	siteId uuid.UUID,
	resolve *bool) (
	models.ApiResponse[[]models.Service],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/services/derived", siteId),
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
	if resolve != nil {
		req.QueryParam("resolve", *resolve)
	}

	var result []models.Service
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[[]models.Service](decoder)
	return models.NewApiResponse(result, resp), err
}

// CountSiteServicePathEvents takes context, siteId, distinct, mType, text, vpnName, vpnPath, policy, portId, model, version, timestamp, mac, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.RepsonseCount data and
// an error if there was an issue with the request or response.
// Count Service Path Events
func (s *SitesServices) CountSiteServicePathEvents(
	ctx context.Context,
	siteId uuid.UUID,
	distinct *models.SiteServiceEventsCountDistinctEnum,
	mType *string,
	text *string,
	vpnName *string,
	vpnPath *string,
	policy *string,
	portId *string,
	model *string,
	version *string,
	timestamp *float64,
	mac *string,
	start *int,
	end *int,
	duration *string,
	limit *int) (
	models.ApiResponse[models.RepsonseCount],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/services/events/count", siteId),
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
	if text != nil {
		req.QueryParam("text", *text)
	}
	if vpnName != nil {
		req.QueryParam("vpn_name", *vpnName)
	}
	if vpnPath != nil {
		req.QueryParam("vpn_path", *vpnPath)
	}
	if policy != nil {
		req.QueryParam("policy", *policy)
	}
	if portId != nil {
		req.QueryParam("port_id", *portId)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if version != nil {
		req.QueryParam("version", *version)
	}
	if timestamp != nil {
		req.QueryParam("timestamp", *timestamp)
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

	var result models.RepsonseCount
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.RepsonseCount](decoder)
	return models.NewApiResponse(result, resp), err
}

// SearchSiteServicePathEvents takes context, siteId, mType, text, vpnName, vpnPath, policy, portId, model, version, timestamp, mac, start, end, duration, limit as parameters and
// returns an models.ApiResponse with models.ResponseEventsPathSearch data and
// an error if there was an issue with the request or response.
// Search Service Path Events
func (s *SitesServices) SearchSiteServicePathEvents(
	ctx context.Context,
	siteId uuid.UUID,
	mType *string,
	text *string,
	vpnName *string,
	vpnPath *string,
	policy *string,
	portId *string,
	model *string,
	version *string,
	timestamp *float64,
	mac *string,
	start *int,
	end *int,
	duration *string,
	limit *int) (
	models.ApiResponse[models.ResponseEventsPathSearch],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/api/v1/sites/%v/services/events/search", siteId),
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
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if text != nil {
		req.QueryParam("text", *text)
	}
	if vpnName != nil {
		req.QueryParam("vpn_name", *vpnName)
	}
	if vpnPath != nil {
		req.QueryParam("vpn_path", *vpnPath)
	}
	if policy != nil {
		req.QueryParam("policy", *policy)
	}
	if portId != nil {
		req.QueryParam("port_id", *portId)
	}
	if model != nil {
		req.QueryParam("model", *model)
	}
	if version != nil {
		req.QueryParam("version", *version)
	}
	if timestamp != nil {
		req.QueryParam("timestamp", *timestamp)
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

	var result models.ResponseEventsPathSearch
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ResponseEventsPathSearch](decoder)
	return models.NewApiResponse(result, resp), err
}
